package convert

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KToP(k Kg) Pound {
	return Pound(k * 2.20462)
}

func PToK(p Pound) Kg {
	return Kg(p / 2.20462)
}

func MToF(m Meters) Feet {
	return Feet(m / 0.3048)
}

func FToM(t Feet) Meters {
	return Meters(t * 0.3048)
}
