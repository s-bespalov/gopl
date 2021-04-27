// Package lenconv converts length in different units
package lenconv

import "fmt"

type Feet float64
type Meter float64

func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

// FToM converts length in Feet to length in Meters
func FToM(f Feet) Meter { return Meter(f / 3.2808) }

// MToF converts length in Meters to length in Feets
func MToF(m Meter) Feet { return Feet(m * 3.2808) }
