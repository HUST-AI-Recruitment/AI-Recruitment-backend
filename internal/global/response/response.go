package response

import (
	"AI-Recruitment-backend/internal/global"
	"AI-Recruitment-backend/pkg/common"
	"github.com/gin-gonic/gin"
)

type response struct {
	Code  ErrorCode `json:"code"`
	Data  Data      `json:"data,omitempty"`
	Msg   string    `json:"msg,omitempty"`
	Error string    `json:"error,omitempty"` // only available in debug mode
}

type Data map[string]any

type UserData struct {
	ID       uint          `json:"id"`
	Username string        `json:"username"`
	Email    string        `json:"email"`
	Role     common.Role   `json:"role"`
	Age      int           `json:"age"`
	Degree   common.Degree `json:"degree"`
}

type JobData struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Demand      string `json:"demand"`
	Location    string `json:"location"`
	Company     string `json:"company"`
	Salary      string `json:"salary"`
	JobType     string `json:"job_type"`
}

func Success(c *gin.Context, status int, code ErrorCode, data Data, msg string) {
	c.JSON(status, response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Error(c *gin.Context, status int, code ErrorCode, msg string, err ...string) {
	if global.Config.App.Debug {
		c.JSON(status, response{
			Code:  code,
			Msg:   msg,
			Error: err[0],
		})
	} else {
		c.JSON(status, response{
			Code: code,
			Msg:  msg,
		})
	}
}
