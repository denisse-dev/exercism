package grains

import (
	"errors"
)

// Square calculates number of grains of wheat given that the number on each
// square doubles using geometric series.
func Square(grains int) (uint64, error) {
	if grains <= 0 || grains >= 65 {
		return 0, errors.New("invalid number of grains")
	}

	return 1 << (uint64(grains) - 1), nil
}

// Total calculates the number of grains of wheat on a chessboard
func Total() uint64 {
	return 1<<64 - 1
}
