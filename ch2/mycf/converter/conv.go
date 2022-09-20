package converter

import "fmt"

type Celsius float64
type Fahrenheit float64
type Meter float64
type Foot float64
type Kilogramm float64
type Pound float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (m Meter) String() string      { return fmt.Sprintf("%gm", m) }
func (ft Foot) String() string      { return fmt.Sprintf("%gft", ft) }
func (k Kilogramm) String() string  { return fmt.Sprintf("%gkg", k) }
func (p Pound) String() string      { return fmt.Sprintf("%glbs", p) }

// FToC converts a Fahrenheit to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToF converts a Celsius to Fahreinheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// MToF converts a Meter to Foot.
func MToF(m Meter) Foot { return Foot(m / 0.3048) }

// FToM converts Foot to Meter.
func FToM(ft Foot) Meter { return Meter(ft * 0.3048) }

// KToP converts Kilo to Pound
func KToP(k Kilogramm) Pound { return Pound(k / 0.453592) }

// PToK converts Pound to Kilo.
func PToK(p Pound) Kilogramm { return Kilogramm(p * .0453592) }
