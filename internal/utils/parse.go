package utils

import (
	"strconv"
	"strings"
	"time"
)

func ParseFloat(v string) *float64 {
	if v == "" {
		return nil
	}

	v = strings.TrimSpace(v)
	v = strings.ReplaceAll(v, " ", "")
	v = strings.ReplaceAll(v, ",", ".")

	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return nil
	}
	return &f
}

func ParseDate(v string) *time.Time {
	if v == "" {
		return nil
	}

	layouts := []string{
		"02.01.2006",
		"02.01.2006 15:04",
		"2006-01-02",
		"2006-01-02 15:04:05",
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, v); err == nil {
			return &t
		}
	}
	return nil
}
