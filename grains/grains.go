package grains

import (
	"errors"
	"math"
)

// Square returns the number of grains on a particular square.
func Square(num int) (uint64, error) {
	if num > 0 && num < 65 {
		return uint64(math.Pow(2, float64(num-1))), nil
	} else {
		return 0, errors.New("Invalid input.")
	}
}

// Total returns the total number of grains on the chessboard.
func Total() uint64 {
	var sum uint64
	for i := 1; i < 65; i++ {
		s, _ := Square(i)
		sum += s
	}
	return sum
}
