package repository

import (
	"github.com/LaQuannT/credit-card-validator/model"
	"github.com/LaQuannT/credit-card-validator/utils"
)

type InMemoryCardIssuerRepo struct {
	database map[string]model.CardIssuer
}

func NewInMemoryCardIssuerStore(filepath string) *InMemoryCardIssuerRepo {
	db := make(map[string]model.CardIssuer)
	utils.PopulateStore(filepath, db)
	return &InMemoryCardIssuerRepo{
		database: db,
	}
}

func (i *InMemoryCardIssuerRepo) GetCardIssuer(BIN string) (model.CardIssuer, bool) {
	issuer, ok := i.database[BIN]
	return issuer, ok
}
