package lenconv

import "fmt"

type Meter float64
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%gMeter", m) }
func (f Feet) String() string  { return fmt.Sprintf("%gFeet", f) }
