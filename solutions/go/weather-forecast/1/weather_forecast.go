// Package weather has an important function in 
//various cities in Goblinocus. It has two variables //CurrentCondition and Currentlocation.The Forecast //func has the logic to forecast the weathercondition.
package weather

var (
// CurrentCondition stores the most updated //information about the condition in each city.
	CurrentCondition string

// CurrentLocation describes where(city, region) the //forecasts was taken.
	CurrentLocation  string
)

// Forecast func takes the city and the condition and //forecasts how is the weather there.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
