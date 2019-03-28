package strain

// Ints defines a collection of ints in a slice
type Ints []int

// Lists defines a collection of ints in a slice inside a slice
type Lists [][]int

// Strings defines a collection of strings in a slice
type Strings []string

// Keep returns a new collection containing those elements where the predicate
// is true.
func (input Ints) Keep(evaluate func(int) bool) (output Ints) {
	if input == nil {
		return nil
	}

	for _, num := range input {
		if evaluate(num) {
			output = append(output, num)
		}
	}

	return output
}

// Discard returns a new collection containing those elements where the
// predicate is false.
func (input Ints) Discard(evaluate func(int) bool) (output Ints) {
	if input == nil {
		return nil
	}

	for _, num := range input {
		if !evaluate(num) {
			output = append(output, num)
		}
	}

	return output
}

// Keep returns a new collection containing those elements where the predicate
// is true.
func (input Lists) Keep(evaluate func([]int) bool) (output Lists) {
	if input == nil {
		return nil
	}

	for _, num := range input {
		if evaluate(num) {
			output = append(output, num)
		}
	}

	return output
}

// Keep returns a new collection containing those elements where the predicate
// is true.
func (input Strings) Keep(evaluate func(string) bool) (output Strings) {
	if input == nil {
		return nil
	}

	for _, str := range input {
		if evaluate(str) {
			output = append(output, str)
		}
	}

	return output
}
