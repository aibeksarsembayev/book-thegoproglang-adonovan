package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	converter "main/converter"
)

func main() {
	if len(os.Args) == 1 {

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			t, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "mycf: %v\n", err)
				os.Exit(1)
			}
			convAndPrint(t)
		}

	} else {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "mycf: %v\n", err)
				os.Exit(1)
			}
			convAndPrint(t)
		}
	}
}

// Converts and Prints
func convAndPrint(input float64) {
	// degC and degF
	f := converter.Fahrenheit(input)
	c := converter.Celsius(input)
	fmt.Printf("%s = %s, %s = %s\n", f, converter.FToC(f), c, converter.CToF(c))
	// Meter and Foot
	m := converter.Meter(input)
	ft := converter.Foot(input)
	fmt.Printf("%s = %s, %s = %s\n", m, converter.MToF(m), ft, converter.FToM(ft))
	// Kilo and Pound
	k := converter.Kilogramm(input)
	p := converter.Pound(input)
	fmt.Printf("%s = %s, %s = %s\n", k, converter.KToP(k), p, converter.PToK(p))
}
