# Credit Card Validator

a simple JSON API for validating credit card numbers.

## EndPoints

Server will be ran on localhost:8080. Only GET request are allowed.

- "/" 
- "/cards"

## Usage

Send an GET request to endpoint "/cards" with a json body containing card number.
```
{
  "card_number":"[some test card number]"
}
```

if card number is valid the response will be that off the card payment system or a struct containing details about the card provider if the card's BIN is known.
```
{
  "payment_system" : ""
}
or
{
    "bin": "",
    "payment_system": "",
    "bank": "",
    "card_type": "",
    "country": "",
    "country_code": "",
    "website": ""
}
```
