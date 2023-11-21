package genlib

//Fahrenheit to Celsius
func FahToCel(f float64) float64 {
	return (f - 32) / 1.8
}

// Celsius to Fahrenheit
func CelToFah(c float64) float64 {
	return (c * 1.8) + 32
}

/// ULA ULA NİYAZİİİİİİİİ

// Feet to Meter
func FeetToMeter(feet float64) float64 {
	return feet / 3.2808
}

// Kg to Pound
func KgToPound(kg float64) float64 {
	return kg * 2.2046
}

// Kilometers to Mile
func KmToMiles(km float64) float64 {
	return km * 0.621371192
}
