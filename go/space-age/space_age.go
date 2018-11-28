// Package space calculates how old someone would be on different planets given
// an age in seconds
package space

// Planet type represents a planet name
type Planet = string

// Age calculates the age of someone on different planets
func Age(seconds float64, planet string) float64 {
	var orbitalPeriod float64
	earthAge := seconds / 31557600

	switch planet {
	case "Mercury":
		orbitalPeriod = 0.2408467
	case "Venus":
		orbitalPeriod = 0.61519726
	case "Mars":
		orbitalPeriod = 1.8808158
	case "Jupiter":
		orbitalPeriod = 11.862615
	case "Saturn":
		orbitalPeriod = 29.447498
	case "Uranus":
		orbitalPeriod = 84.016846
	case "Neptune":
		orbitalPeriod = 164.79132
	default:
		return earthAge
	}
	return earthAge / orbitalPeriod
}
