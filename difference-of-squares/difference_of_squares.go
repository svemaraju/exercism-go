package diffsquares

// SquareOfSum calculates square of sum of first n natural numbers.
func SquareOfSum(n int) int {
	sum := (n*(n+1))/2
	return sum*sum
}

// SumOfSquares calculates sum of squares of first n natural numbers.
func SumOfSquares(n int) int {
     return (n*(n+1)*(2*n+1))/6
}

// Difference returns difference between square of the sum and the sum of the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}