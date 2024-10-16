package common

type Degree int

const (
	HighSchool Degree = iota
	Bachelor
	Master
	Doctor
)

func (d Degree) String() string {
	switch d {
	case HighSchool:
		return "high school"
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
