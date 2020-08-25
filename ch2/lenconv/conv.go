package lenconv

// MToF converts length in meters to feets
func MToF(m Meter) Feet { return Feet(m * 3.281) }

// FToM converts length in feets to meters
func FToM(f Feet) Meter { return Meter(f / 3.281) }
