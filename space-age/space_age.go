//Package space give you different function to determine cool stuff in space.
package space

// Planet is a planer
type Planet string

// Age determines the age of someone in seconds on different planets.
func Age(f float64, p Planet) float64 {
	earthYear := (f / 31557600)
	switch p {
	case "Earth":
		return earthYear
	case "Mercury":
		return earthYear / 0.2408467
	case "Venus":
		return earthYear / 0.61519726
	case "Mars":
		return earthYear / 1.8808158
	case "Jupiter":
		return earthYear / 11.862615
	case "Saturn":
		return earthYear / 29.447498
	case "Uranus":
		return earthYear / 84.016846
	case "Neptune":
		return earthYear / 164.79132
	default:
		return 0
	}

}
