package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/LaQuannT/credit-card-validator/model"
	"github.com/LaQuannT/credit-card-validator/utils"
)

const (
	binLength       = 6
	jsonContentType = "application/json"
	usageMessage    = "to validate a card number send a 'GET' request containing a json body with 'card_number' as the key for the card number to be validated"
)

type CardServer struct {
	store model.CardIssuerStore
	http.Server
}

func NewCardServer(addr string, store model.CardIssuerStore) *CardServer {
	srv := CardServer{store: store}

	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(srv.homeHandler))
	router.HandleFunc("/cards", http.HandlerFunc(srv.cardHandler))

	srv.Server = http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		Handler:      router,
	}

	return &srv
}

func (s *CardServer) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", jsonContentType)

	msg := struct {
		Message string `json:"message"`
	}{
		Message: usageMessage,
	}
	err := json.NewEncoder(w).Encode(msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error encoding response msg: %v", err)
	}
}

func (s *CardServer) cardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	s.getIssuer(w, r.Body)
}

func (s *CardServer) getIssuer(w http.ResponseWriter, payload io.ReadCloser) {
	var data model.Payload
	w.Header().Set("Content-Type", jsonContentType)

	err := json.NewDecoder(payload).Decode(&data)
	if err != nil {
		if errors.Is(err, io.EOF) {
			http.Error(w, "request body is empty, please provide a card number", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error decoding json payload: %v", err)
		return
	}

	isvalid := utils.IsValidLuhn(data.CardNumber)
	if !isvalid {
		http.Error(w, fmt.Sprintf("%q is not valid credit/debit card number", data.CardNumber), http.StatusBadRequest)
		return
	}

	issuer, ok := s.store.GetCardIssuer(data.CardNumber[:binLength])
	if !ok {
		var uknownIssuer model.UnknownCardIssuer

		uknownIssuer.PaymentSystem = utils.CheckCardType(data.CardNumber)
		err := json.NewEncoder(w).Encode(uknownIssuer)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("error encoding json data: %v", err)
			return
		}

		return
	}
	err = json.NewEncoder(w).Encode(issuer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error encoding json data: %v", err)
		return
	}
}
