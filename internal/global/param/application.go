package param

type ReqCreateApplication struct {
	JobID uint `json:"job_id" binding:"required"`
}
