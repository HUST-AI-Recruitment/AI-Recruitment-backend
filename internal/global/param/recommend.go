package param

type ReqRecommendJobsDescription struct {
	Description string `json:"description" binding:"required"`
}
