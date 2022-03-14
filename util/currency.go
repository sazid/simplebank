package util

// Constants for all supported currencies
const (
	BDT = "BDT"
	CAD = "CAD"
	DHS = "DHS"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case BDT, CAD, DHS:
		return true
	}
	return false
}
