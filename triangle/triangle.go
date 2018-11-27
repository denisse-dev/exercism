// Package triangle determines if a triangle is equilateral, isocelees,
// degenerate or escalene.
package triangle

import (
	"math"
	"sort"
)

// Kind is returned by KindFromSides() returns the kind of triangle.
type Kind string

const (
	// NaT referes to not a triangle
	NaT = Kind("NaT")
	// Equ equilateral
	Equ = Kind("Equ")
	// Iso isosceles
	Iso = Kind("Iso")
	// Sca scalene
	Sca = Kind("Sca")
	// Deg degenerate triangle
	Deg = Kind("Deg")
)

// KindFromSides returns the type of triangle given 3 float64 values.
func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	sort.Float64s(sides)

	if sides[0] <= 0 || sides[0]+sides[1] < sides[2] || invalidSide(sides) {
		return NaT
	}
	if sides[0]+sides[1] == sides[2] {
		return Deg
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || c == a {
		return Iso
	}
	return Sca
}

// invalidSide determines if the given side is invalid for a triangule
func invalidSide(sides []float64) bool {
	for _, i := range sides {
		if math.IsNaN(i) || math.IsInf(i, 0) {
			return true
		}
	}
	return false
}
