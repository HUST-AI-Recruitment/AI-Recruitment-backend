package response

import (
	"AI-Recruitment-backend/internal/global"
	"AI-Recruitment-backend/pkg/common"
	"github.com/gin-gonic/gin"
	"time"
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
	OwnerID     uint   `json:"owner_id"`
}

type ResumeData struct {
	ID          uint               `json:"id"`
	UserID      uint               `json:"user_id"`
	Name        string             `json:"name"`
	Gender      int                `json:"gender"`
	Phone       string             `json:"phone"`
	Email       string             `json:"email"`
	Wechat      string             `json:"wechat"`
	State       common.State       `json:"state"`
	Description string             `json:"description"`
	Education   []ResumeEducation  `json:"education"`
	Experience  []ResumeExperience `json:"experience"`
	Project     []ResumeProject    `json:"project"`
}

type ResumeEducation struct {
	School    string        `json:"school"`
	Major     string        `json:"major"`
	Degree    common.Degree `json:"degree"`
	StartTime time.Time     `json:"start_time"`
	EndTime   time.Time     `json:"end_time"`
}

type ResumeExperience struct {
	Company   string    `json:"company"`
	Position  string    `json:"position"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type ResumeProject struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}

type ApplicationData struct {
	ID       uint            `json:"id"`
	UserID   uint            `json:"user_id"`
	JobID    uint            `json:"job_id"`
	Progress common.Progress `json:"progress"`
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
