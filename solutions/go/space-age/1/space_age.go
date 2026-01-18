package space

type Planet string

var planets = map[string]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func ConvertYears(planet string, seconds int) float64 {
	const earth_year_in_seconds = 31557600
	value := planets[planet]
	equivalentEarthyYear := earth_year_in_seconds * value
	return float64(seconds) / equivalentEarthyYear
}

func Age(seconds float64, planet Planet) float64 {
	const earth_year_in_seconds = 31557600
	planets := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	orbitalYears, exists := planets[Planet(planet)]

	if !exists {
		return -1
	}
	planetYearInSeconds := earth_year_in_seconds * orbitalYears
	return float64(seconds) / planetYearInSeconds
}
