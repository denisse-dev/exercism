package summultiples

// SumMultiples finds the sum of all the unique multiples of particular numbers
// up to but not including that number.
func SumMultiples(num int, divisors ...int) (sum int) {
	multiples := make(map[int]struct{})

	for _, div := range divisors {
		if div == 0 {
			continue
		}
		for i := 0; i < num; i += div {
			multiples[i] = struct{}{}
		}
	}

	for key := range multiples {
		sum += key
	}

	return sum
}
