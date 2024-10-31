package common

type Progress int

const (
	_ Progress = iota
	delivered
	accepted
	rejected
)

func (p Progress) String() string {
	switch p {
	case delivered:
		return "delivered"
	case accepted:
		return "accepted"
	case rejected:
		return "rejected"
	default:
		return "unknown"
	}
}
