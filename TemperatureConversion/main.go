package main

import (
	"fmt"
	"math"
)

type (
	// Celsius is a scale and unit of measurement for temperature.
	Celsius float64
	// Fahrenheit is a scale and unit of measurement for temperature.
	Fahrenheit float64
	// Kelvin is a scale and unit of measurement for temperature.
	Kelvin float64
)

const (
	absoluteZeroC = 273.15
	absoluteZeroF = 459.67
)

func main() {

	var continueToConversion bool
	continueToConversion = true

	var usersDegree float64
	var usersRequest string
	var tempType string
	var convertTo string

	for continueToConversion == true {
		fmt.Println(" Welcome to the temperature type conversion program ^-^ !! ")

		fmt.Print(" What is the type of your temperature ? (f/k/c) ")
		fmt.Scanln(&tempType)

		fmt.Print(" What is your temperature in degrees? ")
		fmt.Scanln(&usersDegree)

		fmt.Print(" Which type of temperature do you want to convert to? (f/k/c) ")
		fmt.Scanln(&convertTo)

		if tempType == "f" {
			if convertTo == "k" {
				fmt.Println(FahrenheitToKelvin(Fahrenheit(usersDegree)))
			} else if convertTo == "c" {
				fmt.Println(FahrenheitToCelsius(Fahrenheit(usersDegree)))
			} else {
				fmt.Println("Please enter a valid type to convert, program is terminated")
			}

		} else if tempType == "k" {
			if convertTo == "f" {
				fmt.Println(KelvinToFahrenheit(Kelvin(usersDegree)))
			} else if convertTo == "c" {
				fmt.Println(KelvinToCelcius(Kelvin(usersDegree)))
			} else {
				fmt.Println("Please enter a valid type to convert, program is terminated")
			}

		} else if tempType == "c" {
			if convertTo == "f" {
				fmt.Println(CelsiusToFahrenheit(Celsius(usersDegree)))
			} else if convertTo == "k" {
				fmt.Println(CelsiusToKelvin(Celsius(usersDegree)))
			} else {
				fmt.Println("Please enter a valid type to convert, program is terminated")
			}

		} else {
			fmt.Println("Please enter a valid temperature type !")
		}

		fmt.Print("Do you want to continue to conversion? (yes/no) ")
		fmt.Scanln(&usersRequest)
		if usersRequest == "no" {
			continueToConversion = false
		}
	}

}

// CelsiusToFahrenheit converts celsius to fahrenheit.
func CelsiusToFahrenheit(c Celsius) Fahrenheit {
	value := round(float64((c * 9 / 5) + 32))
	return Fahrenheit(value)
}

// CelsiusToKelvin converts celsius to kelvin.
func CelsiusToKelvin(c Celsius) Kelvin {
	value := round(float64(c + absoluteZeroC))
	return Kelvin(value)
}

// FahrenheitToCelsius converts fahrenheit to celsius.
func FahrenheitToCelsius(f Fahrenheit) Celsius {
	value := round(float64((f - 32) * 5 / 9))
	return Celsius(value)
}

// FahrenheitToKelvin converts fahrenheit to kelvin.
func FahrenheitToKelvin(f Fahrenheit) Kelvin {
	value := round(float64((f + absoluteZeroF) * 5 / 9))
	return Kelvin(value)
}

// KelvinToCelcius converts kelvin to celsius.
func KelvinToCelcius(k Kelvin) Celsius {
	value := round(float64(k - absoluteZeroC))
	return Celsius(value)
}

// KelvinToFahrenheit converts kelvin to fahrenheit.
func KelvinToFahrenheit(k Kelvin) Fahrenheit {
	value := round(float64((k * 9 / 5) - absoluteZeroF))
	return Fahrenheit(value)
}

// Round a value to 1 decimal place.
func round(val float64) float64 {
	return math.Round(val*10) / 10
}
