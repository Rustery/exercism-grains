// Package grains is Exercism.io exercise
package grains

import (
	"errors"
)

// Square — get how many grains were on a given square of chessboard
func Square(input int) (uint64, error) {
	if input < 1 || input > 64 {
		return 0, errors.New("square should be greater than 0 and less or equal 64")
	}
	return 1 << (input - 1), nil
}

// Total — get the total number of grains on the chessboard
func Total() uint64 {
	return 1<<64 - 1
}
