// Package weather is used to forecast the weather.
package weather

// CurrentCondition is used to store the current condition.
var CurrentCondition string

// CurrentLocation is used to store the current location.
var CurrentLocation string

// Forecast is used to forecast the by accepting the location and the condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
