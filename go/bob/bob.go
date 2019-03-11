// Package bob let’s you talk with the lackadaisical teen named Bob.
package bob

import (
	"regexp"
	"strings"
)

var ask = regexp.MustCompile(`\?$`)
var yell = regexp.MustCompile(`^[^a-z]*$`)
var yellAsk = regexp.MustCompile(`^[^a-z]*\?$`)
var upCase = regexp.MustCompile(`[A-Z]`)

// Hey lets you say something to Bob.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	switch {
	case remark == "":
		return "Fine. Be that way!"
	case yellAsk.MatchString(remark) && upCase.MatchString(remark):
		return "Calm down, I know what I'm doing!"
	case ask.MatchString(remark):
		return "Sure."
	case yell.MatchString(remark) && upCase.MatchString(remark):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
