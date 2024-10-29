package response

type ErrorCode int32

const (
	CodeSuccess       ErrorCode = 0
	CodeInvalidToken  ErrorCode = 40001
	CodeUnauthorized  ErrorCode = 40002
	CodeInvalidParams ErrorCode = 40003
	CodeForbidden     ErrorCode = 40004
	CodeServerBusy    ErrorCode = 50001
)
