package main

import (
	"fmt"
	"os"
	"strconv"
)

func celciusToFarenheit(temp float64) float64 {
	return (temp * (9.0 / 5)) + 32
}

func farenheitToCelcius(temp float64) float64 {
	return (temp - 32) * (5.0 / 9)
}

// Converts the temperature from either Celcius to Farenheit (opposite of what is in used as command line arg)
func main() {
	args := os.Args

	// CHECK ARGS LENGTH
	if len(args) != 3 {
		fmt.Println("invalid number of command line args")
		fmt.Printf("Usage: %s <temp> <unit as either 'C' or 'F'>\n", args[0])
		os.Exit(1)
	}

	// CHECK UNIT
	if args[2] != "C" && args[2] != "F" {
		fmt.Println(`temperature unit must be specified by either "C" or "F"`)
		os.Exit(1)
	}

	// CHECK TEMP
	s_temp := args[1]
	if f_temp, err := strconv.ParseFloat(s_temp, 64); err == nil {
		fmt.Printf("Converting temp %.2f %s\n\nRESULT:\n", f_temp, args[2])
		switch args[2] {
		case "C":
			ans := celciusToFarenheit(f_temp)
			fmt.Printf("%.2f F\n", ans)
		case "F":
			ans := farenheitToCelcius(f_temp)
			fmt.Printf("%.2f C\n", ans)
		}

	} else {
		fmt.Printf("temperature value invalid: %s\n", s_temp)
		os.Exit(1)
	}

}
