package utils

import (
	"strconv"
	"strings"
)

func BuildBusinessKey(
	vin string,
	contract string,
	subject string,
	brand string,
	year int,
) string {

	vin = strings.TrimSpace(vin)
	contract = strings.TrimSpace(contract)

	if vin != "" && contract != "" {
		return strings.ToUpper(vin) + "|" + contract
	}

	if contract == "" {
		return ""
	}

	parts := []string{
		contract,
		strings.TrimSpace(subject),
		strings.TrimSpace(brand),
		strconv.Itoa(year),
	}

	return strings.Join(parts, "|")
}
