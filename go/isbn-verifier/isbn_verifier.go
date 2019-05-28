package isbn

import "unicode"

// IsValidISBN checks if the provided string is a valid ISBN-10.
func IsValidISBN(isbn string) bool {
	sum := 0
	count := 10

	for _, num := range isbn {
		if num == '-' {
			continue
		}

		if !unicode.IsNumber(num) && num != 'X' {
			return false
		}

		if num == 'X' {
			sum += 10
			return sum%11 == 0 && count == 1
		}

		sum += (int(num) - '0') * count
		count--
	}

	return sum%11 == 0 && count == 0
}
