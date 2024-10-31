package common

type Progress int

const (
	_ Progress = iota
	delivered
	accepted
	rejected
)
