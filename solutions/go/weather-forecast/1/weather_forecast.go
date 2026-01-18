/* Package weather provides information about the weather conditions of different cities. */
package weather





var (
    // CurrentCondition defines the condition of a city.
	CurrentCondition string

	// CurrentLocation defines the location of a city.
	CurrentLocation  string
)

/* Forecast returns the CurrentLocation and its assigned CurrentCondition. */
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
