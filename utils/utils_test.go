package utils

import "testing"

func TestIsValidluhn(t *testing.T) {
	testCases := []struct {
		name         string
		cardNumber   string
		expectedBool bool
	}{
		{
			"returns true for valid card number",
			"5425233430109903",
			true,
		},
		{
			"returns false for invalid card number",
			"5425233430109605",
			false,
		},
		{
			"returns false for input containing letters",
			"jdfeti0000000000",
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isvalid := IsValidLuhn(tc.cardNumber)

			if isvalid != tc.expectedBool {
				t.Errorf("expected is card valid to be %v, got %v", tc.expectedBool, isvalid)
			}
		})
	}
}

func TestCheckCardType(t *testing.T) {
	testCases := []struct {
		name         string
		cardNumber   string
		expectedType string
	}{
		{
			"returns card type of american express",
			"378282246310005",
			"American Express",
		},
		{
			"returns card type of visa",
			"4007702835532454",
			"Visa",
		},
		{
			"returns card type of mastercard",
			"2222420000001113",
			"MasterCard",
		},
		{
			"returns card type of discover",
			"6011000180331112",
			"Discover",
		},
		{
			"returns card type of unknown",
			"3530111333300000",
			"UNKNOWN",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cardType := CheckCardType(tc.cardNumber)

			if cardType != tc.expectedType {
				t.Errorf("expected card type of %q, got %q", tc.expectedType, cardType)
			}
		})
	}
}
