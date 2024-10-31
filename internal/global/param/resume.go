package param

type ReqCreateResume struct {
	OwnerID uint `json:"owner_id" binding:"required"`
}
