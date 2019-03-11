package accumulate

// Accumulate perfomrs the desired operation in a given collection.
func Accumulate(collection []string, operation func(string) string) (res []string) {
	for _, str := range collection {
		res = append(res, operation(str))
	}

	return res
}
