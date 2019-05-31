package prime

import "math"

// Nth returns the nth prime number
func Nth(nth int) (int, bool) {
	if nth < 1 {
		return 0, false
	}
	if nth <= 2 {
		return nth + 1, true
	}

	prime := 0
	count := 2
	for num := 5; count < nth; num += 2 {
		if isPrime(num) {
			prime = num
			count++
		}
	}

	return prime, true
}

func isPrime(num int) bool {
	if num%2 == 0 {
		return false
	}

	lim := int(math.Sqrt(float64(num))) + 1
	for i := 3; i < lim; i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}
