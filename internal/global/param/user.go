package param

import "AI-Recruitment-backend/pkg/common"

type ReqRegister struct {
	Username string        `json:"username" binding:"required,min=2,max=255"`
	Email    string        `json:"email" binding:"required,max=255,email"`
	Password string        `json:"password" binding:"required,min=6,max=255"`
	Role     common.Role   `json:"role" binding:"required,gt=0"`
	Age      int           `json:"age" binding:"required,min=0,max=150"`
	Degree   common.Degree `json:"degree" binding:"required"`
}

type ReqLogin struct {
	Username string `json:"username" binding:"required,min=2,max=255"`
	Password string `json:"password" binding:"required,min=6,max=255"`
}
