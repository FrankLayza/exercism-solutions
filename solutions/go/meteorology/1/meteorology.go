package meteorology

import "fmt"

type Stringer interface{
	String() string
}

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (sc TemperatureUnit) String() string{
	units := []string{"°C", "°F"}
	return units[sc]
}
func (t Temperature) String() string{
	return fmt.Sprintf("%d %s",t.degree, t.unit)
}
// Add a String method to the Temperature type

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (speed SpeedUnit) String() string {
    switch speed {
    case KmPerHour:
        return "km/h"
    case MilesPerHour:
        return "mph"
    default:
        // Returns a safe fallback instead of crashing
        return fmt.Sprintf("unknown unit (%d)", speed)
    }
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed)String() string{
	return fmt.Sprintf("%d %s", s.magnitude, s.unit)
}
type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (m MeteorologyData)String()string{
	return fmt.Sprintf("%v: %v, Wind %s at %s, %d%% Humidity", m.location, m.temperature, m.windDirection,m.windSpeed,m.humidity)
}