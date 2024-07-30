package util

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

func GetUUID() string {
	return uuid.New().String()
}

func ParseDate(date string) time.Time {
	layout := "2006-01-02"
	res, error := time.Parse(layout, date)
	if error != nil {
		panic(error)
	}
	return res
}

func ToUint(s string) (uint, error) {
	res, error := strconv.ParseUint(s, 10, 32)
	if error != nil {
		return 0, error
	}
	return uint(res), nil
}
