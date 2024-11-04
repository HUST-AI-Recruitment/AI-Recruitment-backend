package param

import (
	"AI-Recruitment-backend/pkg/common"
	"time"
)

type ReqUpsertResume struct {
	Name        string                      `json:"name" binding:"required"`
	Gender      int                         `json:"gender" binding:"required"`
	Phone       string                      `json:"phone"`
	Email       string                      `json:"email"`
	Wechat      string                      `json:"wechat"`
	State       common.State                `json:"state" binding:"required"`
	Description string                      `json:"description"`
	Education   []ReqUpsertResumeEducation  `json:"education"`
	Experience  []ReqUpsertResumeExperience `json:"experience"`
	Project     []ReqUpsertResumeProject    `json:"project"`
}

type ReqUpsertResumeEducation struct {
	School    string        `json:"school" binding:"required"`
	Major     string        `json:"major" binding:"required"`
	Degree    common.Degree `json:"degree" binding:"required"`
	StartTime time.Time     `json:"start_time" binding:"required"`
	EndTime   time.Time     `json:"end_time" binding:"required"`
}

type ReqUpsertResumeExperience struct {
	Company   string    `json:"company" binding:"required"`
	Position  string    `json:"position" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
}

type ReqUpsertResumeProject struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
}
