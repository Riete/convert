package size

import (
	"fmt"
)

const (
	KiB float64 = 1 << 10
	MiB float64 = 1 << 20
	GiB float64 = 1 << 30
	TiB float64 = 1 << 40
)

type Num interface {
	~int64 | ~uint64 | ~int32 | ~uint32 | ~int16 | ~uint16 | ~int8 | ~uint8 | ~int | ~uint | ~float64 | ~float32
}

func ToHumanSize[T Num](size T) string {
	s := float64(size)
	switch {
	case s >= TiB:
		return fmt.Sprintf("%.2fTiB", s/TiB)
	case s >= GiB:
		return fmt.Sprintf("%.2fGiB", s/GiB)
	case s >= MiB:
		return fmt.Sprintf("%.2fMiB", s/MiB)
	case s >= KiB:
		return fmt.Sprintf("%.2fKiB", s/KiB)
	default:
		return fmt.Sprintf("%.0fB", s)
	}
}
