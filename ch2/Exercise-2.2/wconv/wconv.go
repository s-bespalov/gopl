// package wconv converts weight in different units
package wconv

import "fmt"

type Kilogram float64
type Pound float64

func (p Pound) String() string    { return fmt.Sprintf("%glbs", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
func KToP(k Kilogram) Pound       { return Pound(k / 0.45359237) }
func PToK(p Pound) Kilogram       { return Kilogram(p * 0.45359237) }
