// Package diffsquares gives you functions to work with squares.
package diffsquares

// SumOfSquares find the sum of squares through n
func SumOfSquares(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// SquareOfSum find the square of the sum through n
func SquareOfSum(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// Difference finds the difference between the square of the sum and the sum of the square.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
