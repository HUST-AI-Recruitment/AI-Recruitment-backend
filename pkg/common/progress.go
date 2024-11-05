package common

type Progress int

const (
	_ Progress = iota
	CandidateApplied
	RecruiterReviewed
	RecruiterAccepted
	RecruiterRejected
	CandidateAccepted
	CandidateRejected
)

func (p Progress) String() string {
	switch p {
	case CandidateApplied:
		return "candidate applied"
	case RecruiterReviewed:
		return "recruiter reviewed"
	case RecruiterAccepted:
		return "recruiter accepted"
	case RecruiterRejected:
		return "recruiter rejected"
	case CandidateAccepted:
		return "candidate accepted"
	case CandidateRejected:
		return "candidate rejected"
	default:
		return "unknown"
	}
}
