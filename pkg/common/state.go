package common

type State int

const (
	_ State = iota
	Graduated
	UndergraduateNotReady
	UndergraduateThinking
	UndergraduateReady
)

func (s State) String() string {
	switch s {
	case Graduated:
		return "graduated"
	case UndergraduateNotReady:
		return "undergraduate not ready"
	case UndergraduateThinking:
		return "undergraduate thinking"
	case UndergraduateReady:
		return "undergraduate ready"
	default:
		return "unknown"
	}
}
