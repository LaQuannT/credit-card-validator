package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/LaQuannT/credit-card-validator/model"
)

const (
	minCreditCardLen = 15
	maxCreditCardLen = 19
)

var (
	amexRegex     = regexp.MustCompile(`^3[47][0-9]{0,}$`)
	visaRegex     = regexp.MustCompile(`^4[0-9]{0,}$`)
	masterCdRegex = regexp.MustCompile(`^(5[1-5]|222[1-9]|22[3-9]|2[3-6]|27[01]|2720)[0-9]{0,}$`)
	discoverRegex = regexp.MustCompile(`^65[4-9][0-9]{13}|64[4-9][0-9]{13}|6011[0-9]{12}|(622(?:12[6-9]|1[3-9][0-9]|[2-8][0-9][0-9]|9[01][0-9]|92[0-5])[0-9]{10})$`)
)

func IsValidLuhn(cardNumber string) bool {
	NumberLength := len(cardNumber)
	if NumberLength > maxCreditCardLen || NumberLength < minCreditCardLen {
		return false
	}
	var sum int
	isDigitToBeDoubled := false

	for i := NumberLength - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(cardNumber[i]))
		if err != nil {
			return false
		}

		if isDigitToBeDoubled {
			digit *= 2

			if digit > 9 {
				digit -= 9
			}

			sum += digit
		} else {
			sum += digit
		}
		isDigitToBeDoubled = !isDigitToBeDoubled
	}

	return sum%10 == 0
}

func CheckCardType(cardNumber string) string {
	switch {
	case amexRegex.MatchString(cardNumber):
		return "American Express"
	case masterCdRegex.MatchString(cardNumber):
		return "MasterCard"
	case visaRegex.MatchString(cardNumber):
		return "Visa"
	case discoverRegex.MatchString(cardNumber):
		return "Discover"
	default:
		return "UNKNOWN"
	}
}

func PopulateStore(filePath string, store map[string]model.CardIssuer) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("error opening file %s: %v", filePath, err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("error reading contents of file %s: %v", filePath, err)
	}

	var cardIssuers model.CardIssuers
	rdr := bytes.NewReader(content)

	err = json.NewDecoder(rdr).Decode(&cardIssuers)
	if err != nil {
		log.Fatalf("error decoding file contents: %v", err)
	}

	for _, issuer := range cardIssuers.Issuers {
		store[issuer.Bin] = issuer
	}
}
