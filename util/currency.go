package util

// Constants for all supported currencies
const (
	BDT = "BDT"
	CAD = "CAD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case BDT, CAD:
		return true
	}
	return false
}
