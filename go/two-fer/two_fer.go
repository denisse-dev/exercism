// Package twofer implements the Two-fer game in go
package twofer

import (
	"fmt"
)

// ShareWith prints the name of the person (if it receives one)
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
