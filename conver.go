package main

import (
	"fmt"
	"strings"
)

// Function to convert Celsius to Fahrenheit
func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

// Function to convert Celsius to Kelvin
func celsiusToKelvin(c float64) float64 {
	return c + 273.15
}

// Function to convert Fahrenheit to Celsius
func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

// Function to convert Fahrenheit to Kelvin
func fahrenheitToKelvin(f float64) float64 {
	return fahrenheitToCelsius(f) + 273.15
}

// Function to convert Kelvin to Celsius
func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

// Function to convert Kelvin to Fahrenheit
func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

func main() {
	fmt.Println("Temperature Converter")
	fmt.Println("=======================")
	fmt.Println("Enter a temperature and its unit (e.g., '32 F' or '0 C').")
	fmt.Println("Valid units: C (Celsius), F (Fahrenheit), K (Kelvin).")
	fmt.Println("Type 'exit' to quit.\n")

	for {
		fmt.Print("Enter temperature and unit: ")
		var input string
		fmt.Scanln(&input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		var value float64
		var unit string
		_, err := fmt.Sscanf(input, "%f %s", &value, &unit)
		if err != nil || (unit != "C" && unit != "F" && unit != "K") {
			fmt.Println("Invalid input. Please enter a temperature and unit (e.g., '32 F').\n")
			continue
		}

		fmt.Println("Conversions:")
		switch unit {
		case "C":
			fmt.Printf("%.2f C = %.2f F\n", value, celsiusToFahrenheit(value))
			fmt.Printf("%.2f C = %.2f K\n", value, celsiusToKelvin(value))
		case "F":
			fmt.Printf("%.2f F = %.2f C\n", value, fahrenheitToCelsius(value))
			fmt.Printf("%.2f F = %.2f K\n", value, fahrenheitToKelvin(value))
		case "K":
			fmt.Printf("%.2f K = %.2f C\n", value, kelvinToCelsius(value))
			fmt.Printf("%.2f K = %.2f F\n", value, kelvinToFahrenheit(value))
		}
		fmt.Println()
	}
}
