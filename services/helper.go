package services

import "time"

func ParseTanggalLahir(tanggal string) time.Time {
	if tanggal == "" {
		return time.Time{}
	}

	formats := []string{
		"2006-01-02",
		"02/01/2006",
		"02-01-2006",
	}

	for _, format := range formats {
		t, err := time.Parse(format, tanggal)
		if err == nil {
			return t
		}
	}

	return time.Time{}
}
