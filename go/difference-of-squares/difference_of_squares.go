package diffsquares

// SquareOfSum sums numbers from 1 to n and then it squares the result
func SquareOfSum(num int) (res int) {
	res = (num * (num + 1)) / 2
	return res * res
}

// SumOfSquares squares numbers from 1 to n and then sums them
func SumOfSquares(num int) int {
	return (num * (num + 1) * (2*num + 1)) / 6
}

// Difference subtracts the SumOfSquares to the SquareOfSum
func Difference(num int) int {
	return (num*num - 1) * (3*num + 2) * num / 12
}
