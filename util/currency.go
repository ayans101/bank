package util

// Constants for all supported currencies
const (
	INR = "INR"
	EUR = "EUR"
	RUB = "RUB"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case INR, EUR, RUB:
		return true
	}
	return false
}
