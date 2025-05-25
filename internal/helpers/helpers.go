package helpers

import (
	"slices"
	"time"
)

type IHelpers interface {
	GetNowDatePointer() *time.Time
}

func GetNowDatePointer() *time.Time {
	t := time.Now()
	return &t
}

func DoesInclude[T comparable](val T, slice []T) bool {
	return slices.Contains(slice, val)
}
