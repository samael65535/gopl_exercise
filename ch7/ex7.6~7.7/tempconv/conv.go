package tempconv

func convCToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func convFToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func convCToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

func convKToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}


