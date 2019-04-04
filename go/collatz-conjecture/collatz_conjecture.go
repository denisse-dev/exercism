// package collatzconjecture solves the Collatz Conjecture or 3x+1 problem
package collatzconjecture

import (
	"errors"
)

// CollatzConjecture returns the number of steps needed to solve the Collatz
// Conjecture problem
func CollatzConjecture(num int) (steps int, err error) {
	if num < 1 {
		return 0, errors.New("can't divide by zero or negative numbers")
	}

	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else {
			num *= 3
			num += 1
		}
		steps += 1
	}

	return steps, err
}
