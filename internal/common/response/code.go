package response

type ErrorCode int32

const (
	CodeSuccess      ErrorCode = 0
	CodeInvalidToken ErrorCode = 10001
	CodeUnauthorized ErrorCode = 10002
)
