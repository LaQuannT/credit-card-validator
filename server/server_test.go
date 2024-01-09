package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LaQuannT/credit-card-validator/model"
)

type StubCardIssuerStore struct {
	issuers map[string]model.CardIssuer
}

func (s *StubCardIssuerStore) GetCardIssuer(BIN string) (model.CardIssuer, bool) {
	issuer, ok := s.issuers[BIN]
	return issuer, ok
}

func TestGetIssuer(t *testing.T) {
	store := StubCardIssuerStore{
		issuers: map[string]model.CardIssuer{
			"542523": {Bank: "Allied Irish Banks"},
		},
	}

	server := NewCardServer("", &store)

	t.Run("returns Visa for unknown Visa issuer", func(t *testing.T) {
		req, err := newGetIssuerRequest([]byte(`{"card_number": "4347699988887777"}`))
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()

		server.cardHandler(res, req)
		var issuer model.UnknownCardIssuer

		json.NewDecoder(res.Body).Decode(&issuer)
		want := "Visa"

		if issuer.PaymentSystem != want {
			t.Errorf("expected payment system %q, got %q", want, issuer.PaymentSystem)
		}
	})

	t.Run("returns Allied Irish Bank for known card issuer", func(t *testing.T) {
		req, err := newGetIssuerRequest([]byte(`{"card_number": "5425233430109903"}`))
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()

		server.cardHandler(res, req)

		var issuer model.CardIssuer

		json.NewDecoder(res.Body).Decode(&issuer)
		want := "Allied Irish Banks"

		if issuer.Bank != want {
			t.Errorf("expected bank %q, got %q", want, issuer.Bank)
		}
	})

	t.Run("returns 400 http status for invalid card card_number", func(t *testing.T) {
		req, err := newGetIssuerRequest([]byte(`{"card_number": "5425233430109905""}`))
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()

		server.cardHandler(res, req)

		if res.Code != http.StatusBadRequest {
			t.Errorf("expected http status %d got %d", http.StatusBadRequest, res.Code)
		}
	})
}

func newGetIssuerRequest(body []byte) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, "/cards", bytes.NewReader(body))
	return req, err
}
