package lenconv

func FTToM(f Foot) Meter {
	return Meter(0.3048 * f)
}

func MToFT(m Meter) Foot {
	return Foot(m / 0.3048)
}
