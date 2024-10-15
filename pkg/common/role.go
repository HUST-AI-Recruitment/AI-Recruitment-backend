package common

type Role int

const (
	Admin Role = iota
	Recruiter
	Candidate
)

func (r Role) String() string {
	switch r {
	case Admin:
		return "admin"
	case Recruiter:
		return "recruiter"
	case Candidate:
		return "candidate"
	default:
		return "unknown"
	}
}
