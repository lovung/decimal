package decimal

type Type int

const (
	Recurring Type = iota
	Irrational
	PositiveInfinity
	NegativeInfinity
	PositiveZero
	NegativeZero
)
