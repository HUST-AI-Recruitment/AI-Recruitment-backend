package model

import (
	"AI-Recruitment-backend/pkg/common"
	"gorm.io/gorm"
)

type Application struct {
	*gorm.Model
	UserID   uint            `gorm:"type:int;primary_key;not null" json:"user_id"`
	JobID    uint            `gorm:"type:int;primary_key;not null" json:"job_id"`
	Progress common.Progress `gorm:"type:int;not null" json:"progress"`
}

func (a Application) TableName() string {
	return "application"
}

func (a Application) Create(db *gorm.DB) (uint, error) {
	res := db.Create(&a)
	if res.Error != nil {
		return 0, res.Error
	}
	return a.ID, nil
}

func (a Application) Delete(db *gorm.DB) error {
	return db.Delete(&a).Error
}

func (a Application) Get(db *gorm.DB) (*Application, error) {
	var userJob Application
	err := db.Where("user_id = ? AND job_id = ?", a.UserID, a.JobID).First(&userJob).Error
	if err != nil {
		return &userJob, err
	}
	return &userJob, nil
}

func (a Application) GetByUserID(db *gorm.DB) (*[]Application, error) {
	var userJobs []Application
	err := db.Where("user_id = ?", a.UserID).Find(&userJobs).Error
	if err != nil {
		return &userJobs, err
	}
	return &userJobs, nil
}

func (a Application) GetByJobID(db *gorm.DB) (*[]Application, error) {
	var userJobs []Application
	err := db.Where("job_id = ?", a.JobID).Find(&userJobs).Error
	if err != nil {
		return &userJobs, err
	}
	return &userJobs, nil
}
