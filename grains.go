// Package grains is Exercism.io exercise
package grains

import (
	"errors"
	"math"
)

// Square — get how many grains were on a given square of chessboard
func Square(input int) (uint64, error) {
	if input < 1 || input > 64 {
		return 0, errors.New("square should be greater than 0 and less or equal 64")
	}
	return uint64(math.Pow(2, float64(input-1))), nil
}

// Total — get the total number of grains on the chessboard
func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		square, _ := Square(i)
		total += square
	}
	return total
}
