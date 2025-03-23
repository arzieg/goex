package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"umrechner/convert"
	myflags "umrechner/flag"
)

func main() {
	flag.Parse()

	if myflags.Verbose {

		fmt.Println("Celsius:", myflags.Celsius)
		fmt.Println("Fahrenheit:", myflags.Fahrenheit)
		fmt.Println("KG:", myflags.Kg)
		fmt.Println("Pound:", myflags.Pound)
	}

	switch {
	case myflags.Celsius != "":
		z, err := strconv.ParseFloat(myflags.Celsius, 64)
		c := convert.Celsius(z)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s = %s\n", c, convert.CToF(c))
	case myflags.Fahrenheit != "":
		z, err := strconv.ParseFloat(myflags.Fahrenheit, 64)
		f := convert.Fahrenheit(z)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s = %s\n", f, convert.FToC(f))
	case myflags.Kg != "":
		z, err := strconv.ParseFloat(myflags.Kg, 64)
		k := convert.Kg(z)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s = %s\n", k, convert.KToP(k))
	case myflags.Pound != "":
		z, err := strconv.ParseFloat(myflags.Pound, 64)
		p := convert.Pound(z)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s = %s\n", p, convert.PToK(p))
	case myflags.Meters != "":
		z, err := strconv.ParseFloat(myflags.Meters, 64)
		m := convert.Meters(z)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s = %s\n", m, convert.MToF(m))
	case myflags.Feet != "":
		z, err := strconv.ParseFloat(myflags.Feet, 64)
		t := convert.Feet(z)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s = %s\n", t, convert.FToM(t))
	default:
		fmt.Fprintf(os.Stderr, "Wert konnte nicht konvertiert werden")
		os.Exit(1)
	}
}
