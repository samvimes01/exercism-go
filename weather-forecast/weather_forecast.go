// Package weather provides tools for getting a weather forecast for a certin location.
package weather

// CurrentCondition represents the weather condition.
var CurrentCondition string

// CurrentLocation represents the weather location.
var CurrentLocation string

// Forecast returns a string representation of weather condition of certain location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
