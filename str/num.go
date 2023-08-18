package str

import (
	"strconv"
)

func ToInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64](s string, r *T, bitSize int) error {
	n, err := strconv.ParseInt(s, 10, bitSize)
	*r = T(n)
	return err
}

func ToUint[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](s string, r *T, bitSize int) error {
	n, err := strconv.ParseUint(s, 10, bitSize)
	*r = T(n)
	return err
}

func ToFloat[T ~float32 | ~float64](s string, r *T, bitSize int) error {
	f, err := strconv.ParseFloat(s, bitSize)
	*r = T(f)
	return err
}
