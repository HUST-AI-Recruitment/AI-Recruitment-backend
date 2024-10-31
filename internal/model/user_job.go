package model

import (
	"AI-Recruitment-backend/pkg/common"
	"gorm.io/gorm"
)

type UserJob struct {
	*gorm.Model
	UserID   uint            `gorm:"type:int;primary_key;not null" json:"user_id"`
	JobID    uint            `gorm:"type:int;primary_key;not null" json:"job_id"`
	Progress common.Progress `gorm:"type:int;not null" json:"progress"`
}

func (uj UserJob) TableName() string {
	return "user_job"
}

func (uj UserJob) Create(db *gorm.DB) error {
	return db.Create(&uj).Error
}

func (uj UserJob) Delete(db *gorm.DB) error {
	return db.Delete(&uj).Error
}

func (uj UserJob) Get(db *gorm.DB) (*UserJob, error) {
	var userJob UserJob
	err := db.Where("user_id = ? AND job_id = ?", uj.UserID, uj.JobID).First(&userJob).Error
	if err != nil {
		return &userJob, err
	}
	return &userJob, nil
}

func (uj UserJob) GetByUserID(db *gorm.DB) (*[]UserJob, error) {
	var userJobs []UserJob
	err := db.Where("user_id = ?", uj.UserID).Find(&userJobs).Error
	if err != nil {
		return &userJobs, err
	}
	return &userJobs, nil
}

func (uj UserJob) GetByJobID(db *gorm.DB) (*[]UserJob, error) {
	var userJobs []UserJob
	err := db.Where("job_id = ?", uj.JobID).Find(&userJobs).Error
	if err != nil {
		return &userJobs, err
	}
	return &userJobs, nil
}
