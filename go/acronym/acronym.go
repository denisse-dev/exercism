// Package acronym helps generate some jargon by converting a long name like
// Portable Network Graphics to its acronym (PNG).
package acronym

import (
	"regexp"
	"strings"
)

var match = regexp.MustCompile(`(\b\w)`)

// Abbreviate converts a phrase to its acronym.
func Abbreviate(s string) string {
	s = strings.Replace(s, "'", "", -1)
	letters := match.FindAllString(s, -1)

	return strings.ToUpper(strings.Join(letters, ""))
}
