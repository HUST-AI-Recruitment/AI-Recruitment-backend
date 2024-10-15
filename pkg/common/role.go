package common

type Role int

const (
	Admin Role = iota
	Recruiter
	Candidate
)
