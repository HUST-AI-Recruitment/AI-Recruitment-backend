package param

type ReqCreateApplication struct {
	JobID uint `json:"job_id" binding:"required"`
}

type ReqUpdateApplication struct {
	ID       uint `json:"id" binding:"required"`
	Accepted bool `json:"accepted"`
}
