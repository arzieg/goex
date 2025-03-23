package myflags

import (
	"flag"
)

var (
	// commandline flags
	Celsius    string
	Fahrenheit string
	Kg         string
	Pound      string
	Meters     string
	Feet       string
	Verbose    bool
)

func init() {
	flag.BoolVar(&Verbose, "v", false, "verbose output")
	flag.StringVar(&Celsius, "c", "", "Celsius")
	flag.StringVar(&Fahrenheit, "f", "", "Fahrenheit")
	flag.StringVar(&Kg, "k", "", "KG")
	flag.StringVar(&Pound, "p", "", "Pound")
	flag.StringVar(&Meters, "m", "", "Meters")
	flag.StringVar(&Feet, "t", "", "Feet")

}
