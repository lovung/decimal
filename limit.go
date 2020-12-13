package decimal

// Type of the decimal number
type Type int

// List of types
const (
	Recurring Type = iota
	Irrational
	PositiveInfinity
	NegativeInfinity
	PositiveZero
	NegativeZero
)
