package encode

import (
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode compresses data using the Run-length encoding algorithm.
func RunLengthEncode(decoded string) string {
	var encoded strings.Builder

	for _, char := range decoded {
		start := len(decoded)
		decoded = strings.TrimLeft(decoded, string(char))

		count := start - len(decoded)
		if count == 0 {
			continue
		}
		if count > 1 {
			encoded.WriteString(strconv.Itoa(count))
		}
		encoded.WriteRune(char)
	}

	return encoded.String()
}

// RunLengthDecode reconstructs data using the Run-length decoding algorithm.
func RunLengthDecode(encoded string) string {
	var decoded strings.Builder
	wasNumber := false
	num := ""

	for _, char := range encoded {
		if unicode.IsNumber(char) {
			num += string(char)
			wasNumber = true
			continue
		}
		if wasNumber {
			times, _ := strconv.Atoi(num)
			decoded.WriteString(strings.Repeat(string(char), times))
			wasNumber = false
			num = ""
			continue
		}
		decoded.WriteRune(char)
	}

	return decoded.String()
}
