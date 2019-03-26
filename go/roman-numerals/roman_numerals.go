package romannumerals

import "errors"

var romanNumerals = []struct {
	number int
	roman  string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRomanNumeral attempts to convert an integer to Roman Numerals returning the
// number as a string or an error if it can't be converted.
func ToRomanNumeral(number int) (result string, err error) {
	if number <= 0 || number > 3000 {
		return result, errors.New("Number is not in range")
	}

	for _, transform := range romanNumerals {
		for number >= transform.number {
			result += transform.roman
			number -= transform.number
		}
	}

	return result, nil
}
