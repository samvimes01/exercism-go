package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	period := planet.earthPeriod()
	if period < 0 {
		return period
	}
	return seconds / 31557600 / period
}

func (p Planet) earthPeriod() float64 {
	switch p {
	case "Mercury":
		return 0.2408467
	case "Venus":
		return 0.61519726
	case "Earth":
		return 1.0
	case "Mars":
		return 1.8808158
	case "Jupiter":
		return 11.862615
	case "Saturn":
		return 29.447498
	case "Uranus":
		return 84.016846
	case "Neptune":
		return 164.79132
	}
	return -1.0
}