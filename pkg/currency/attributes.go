package currency

import "fmt"

// Brazilian Real	BRL	986

var ISO4217 = []attrs{
	{code: "BRL", number: "986", name: "Brazilian Real"},
	{code: "AFN", number: "971", name: "Afghani"},
	{code: "ALL", number: "008", name: "Lek"},
	{code: "DZD", number: "012", name: "Algerian Dinar"},
	{code: "USD", number: "840", name: "US Dollar"},
	{code: "EUR", number: "978", name: "Euro"},
	{code: "AOA", number: "973", name: "Kwanza"},
	{code: "XCD", number: "951", name: "East Caribbean Dollar"},
	{code: "ARS", number: "032", name: "Argentine Peso"},
	{code: "AMD", number: "051", name: "Armenian Dram"},
	{code: "AWG", number: "533", name: "Aruban Florin"},
	{code: "AUD", number: "036", name: "Australian Dollar"},
	{code: "EUR", number: "978", name: "Euro"},
	{code: "AZN", number: "944", name: "Azerbaijanian Manat"},
}

type attrs struct {
	code   string
	name   string
	number string
}

func newAttributes(identifier string) (*attrs, error) {
	for _, attr := range ISO4217 {
		if identifier == attr.name || identifier == attr.number || identifier == attr.code {
			return &attr, nil
		}
	}
	return nil, fmt.Errorf("unable to identify currency")
}
