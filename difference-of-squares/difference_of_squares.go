package diffsquares

func SquareOfSum(n int) int {
	// Gauss summ (I knew about this formula so used it from the start
	sum := 0
	if n% 2 == 0 {
		sum = (n * (n + 1))/2
	} else {
		nm := n - 1
		sum = (nm * (nm + 1))/2 + n
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	// straightforward - slow
	// sum := 0
	// for i := 1; i <= n; i++ {
	// 	sum += i * i
	// }
	// return sum
	// https://en.wikipedia.org/wiki/Square_pyramidal_number
	return n * (n + 1) * (2*n + 1)/6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
