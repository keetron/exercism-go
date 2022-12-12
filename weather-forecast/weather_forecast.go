// Package weather allows us to get a forecast by city.
package weather

// CurrentCondition contains the current condition for the given location.
var CurrentCondition string

// CurrentLocation contains the current location for the weather condition.
var CurrentLocation string

// Forecast returns a string with the current weather conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
