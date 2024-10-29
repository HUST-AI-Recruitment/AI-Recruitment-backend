package param

type ReqCreateJob struct {
	Title       string `json:"title" binding:"required,min=2,max=255"`
	Description string `json:"description" binding:"required,min=2"`
	Demand      string `json:"demand" binding:"required,min=2"`
	Location    string `json:"location" binding:"required,min=2,max=255"`
	Company     string `json:"company" binding:"required,min=2,max=255"`
	Salary      string `json:"salary" binding:"required,min=2,max=255"`
	JobType     string `json:"job_type" binding:"required,min=2,max=255"`
}
