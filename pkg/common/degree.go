package common

type Degree int

const (
	_ Degree = iota
	Bachelor
	Master
	Doctor
)

func (d Degree) String() string {
	switch d {
	case Bachelor:
		return "bachelor"
	case Master:
		return "master"
	case Doctor:
		return "doctor"
	default:
		return "unknown"
	}
}
