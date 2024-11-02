package param

import (
	"AI-Recruitment-backend/pkg/common"
	"time"
)

type ReqCreateResume struct {
	Name        string                      `json:"name" binding:"required"`
	Gender      int                         `json:"gender" binding:"required"`
	Phone       string                      `json:"phone"`
	Email       string                      `json:"email"`
	Wechat      string                      `json:"wechat"`
	State       common.State                `json:"state" binding:"required"`
	Description string                      `json:"description"`
	Education   []ReqCreateResumeEducation  `json:"education"`
	Experience  []ReqCreateResumeExperience `json:"experience"`
	Project     []ReqCreateResumeProject    `json:"project"`
}

type ReqCreateResumeEducation struct {
	School    string        `json:"school" binding:"required"`
	Major     string        `json:"major" binding:"required"`
	Degree    common.Degree `json:"degree" binding:"required"`
	StartTime time.Time     `json:"start_time" binding:"required"`
	EndTime   time.Time     `json:"end_time" binding:"required"`
}

type ReqCreateResumeExperience struct {
	Company   string    `json:"company" binding:"required"`
	Position  string    `json:"position" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
}

type ReqCreateResumeProject struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
}
