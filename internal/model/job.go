package model

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Job struct {
	*gorm.Model
	Title       string `gorm:"type:varchar(255);not null" json:"title"`
	Description string `gorm:"type:text;not null" json:"description"`
	Demand      string `gorm:"type:text;not null" json:"demand"`
	Location    string `gorm:"type:varchar(255);not null" json:"location"`
	Company     string `gorm:"type:varchar(255);not null" json:"company"`
	Salary      string `gorm:"type:varchar(255);not null" json:"salary"`
	JobType     string `gorm:"type:varchar(255);not null" json:"job_type"`
	OwnerID     uint   `gorm:"type:int;not null" json:"owner_id"`
}

func (j Job) TableName() string {
	return "job"
}

func (j Job) Create(db *gorm.DB) (uint, error) {
	res := db.Create(&j)
	if res.Error != nil {
		return 0, res.Error
	}
	return j.ID, nil
}

func (j Job) Update(db *gorm.DB) (*Job, error) {
	var jobs []Job
	if err := db.Model(&Job{}).Clauses(clause.Returning{}).Where("id = ?", j.ID).Updates(j).Scan(&jobs).Error; err != nil {
		return nil, err
	}
	return &jobs[0], nil
}

func (j Job) Delete(db *gorm.DB) error {
	return db.Delete(&Job{}, j.Model.ID).Error
}

func (j Job) Get(db *gorm.DB) (*Job, error) {
	var job Job
	err := db.Where("id = ?", j.Model.ID).First(&job).Error
	if err != nil {
		return &job, err
	}
	return &job, nil
}

func (j Job) GetAll(db *gorm.DB) (*[]Job, error) {
	var jobs []Job
	query := db.Model(&Job{})

	if j.Location != "" {
		query = query.Where("location = ?", j.Location)
	}
	if j.Company != "" {
		query = query.Where("company = ?", j.Company)
	}
	if j.Salary != "" {
		query = query.Where("salary = ?", j.Salary)
	}
	if j.OwnerID != 0 {
		query = query.Where("owner_id = ?", j.OwnerID)
	}

	err := query.Find(&jobs).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &jobs, err
	}
	return &jobs, nil
}
