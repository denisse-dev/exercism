package luhn

import "strings"

// Valid determines if a number has a valid checksum using the Luhn algorithm.
func Valid(numbers string) bool {
	var sum int
	numbers = strings.Replace(numbers, " ", "", -1)

	if len(numbers) <= 1 {
		return false
	}

	double := len(numbers)%2 == 0
	for _, r := range numbers {
		num := int(r - '0')
		if num < 0 || num > 9 {
			return false
		}
		if double {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}

		sum += num
		double = !double
	}

	return sum%10 == 0
}
