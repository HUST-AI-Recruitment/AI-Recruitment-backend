package model

import (
	"AI-Recruitment-backend/pkg/common"
	"gorm.io/gorm"
	"time"
)

type Resume struct {
	*gorm.Model
	UserID      uint         `gorm:"type:int;unique;not null" json:"user_id"`
	Name        string       `gorm:"type:varchar(255);not null" json:"name"`
	Gender      int          `gorm:"type:int;not null" json:"gender"`
	Phone       string       `gorm:"type:varchar(255)" json:"phone"`
	Email       string       `gorm:"type:varchar(255)" json:"email"`
	Wechat      string       `gorm:"type:varchar(255)" json:"wechat"`
	State       common.State `gorm:"type:int;not null" json:"state"`
	Description string       `gorm:"type:text" json:"description"`
}

type ResumeEducation struct {
	*gorm.Model
	ResumeID  uint          `gorm:"type:int;not null" json:"resume_id"`
	School    string        `gorm:"type:varchar(255);not null" json:"school"`
	Major     string        `gorm:"type:varchar(255);not null" json:"major"`
	Degree    common.Degree `gorm:"type:int;not null" json:"degree"`
	StartTime time.Time     `gorm:"type:date;not null" json:"start_time"`
	EndTime   time.Time     `gorm:"type:date;not null" json:"end_time"`
}

type ResumeExperience struct {
	*gorm.Model
	ResumeID  uint      `gorm:"type:int;not null" json:"resume_id"`
	Company   string    `gorm:"type:varchar(255);not null" json:"company"`
	Position  string    `gorm:"type:varchar(255);not null" json:"position"`
	StartTime time.Time `gorm:"type:date;not null" json:"start_time"`
	EndTime   time.Time `gorm:"type:date;not null" json:"end_time"`
}

type ResumeProject struct {
	*gorm.Model
	ResumeID    uint      `gorm:"type:int;not null" json:"resume_id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	StartTime   time.Time `gorm:"type:date;not null" json:"start_time"`
	EndTime     time.Time `gorm:"type:date;not null" json:"end_time"`
}

func (r Resume) TableName() string {
	return "resume"
}

func (re ResumeEducation) TableName() string {
	return "resume_education"
}

func (re ResumeExperience) TableName() string {
	return "resume_experience"
}

func (rp ResumeProject) TableName() string {
	return "resume_project"
}

func (r Resume) CreateResume(db *gorm.DB, edu *[]ResumeEducation, exp *[]ResumeExperience, project *[]ResumeProject) (uint, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&r).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	for _, e := range *edu {
		e.ResumeID = r.ID
		if err := tx.Create(&e).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	for _, e := range *exp {
		e.ResumeID = r.ID
		if err := tx.Create(&e).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	for _, e := range *project {
		e.ResumeID = r.ID
		if err := tx.Create(&e).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	return r.ID, nil
}

func (r Resume) Create(db *gorm.DB) (uint, error) {
	res := db.Create(&r)
	if res.Error != nil {
		return 0, res.Error
	}
	return r.ID, nil
}

func (r Resume) Update(db *gorm.DB, values map[string]interface{}) error {
	return db.Model(&Resume{}).Where("id = ?", r.Model.ID).Updates(values).Error
}

func (r Resume) Delete(db *gorm.DB) error {
	return db.Delete(&Resume{}, r.Model.ID).Error
}

func (r Resume) Get(db *gorm.DB) (*Resume, error) {
	var resume Resume
	err := db.Where("id = ?", r.Model.ID).First(&resume).Error
	if err != nil {
		return &resume, err
	}
	return &resume, nil
}

func (r Resume) GetByUserID(db *gorm.DB) (*Resume, error) {
	var resume Resume
	err := db.Where("user_id = ?", r.UserID).First(&resume).Error
	if err != nil {
		return &resume, err
	}
	return &resume, nil
}

func (re ResumeEducation) Create(db *gorm.DB) (uint, error) {
	res := db.Create(&re)
	if res.Error != nil {
		return 0, res.Error
	}
	return re.ID, nil
}

func (re ResumeEducation) Update(db *gorm.DB, values map[string]interface{}) error {
	return db.Model(&ResumeEducation{}).Where("id = ?", re.Model.ID).Updates(values).Error
}

func (re ResumeEducation) Delete(db *gorm.DB) error {
	return db.Delete(&ResumeEducation{}, re.Model.ID).Error
}

func (re ResumeEducation) Get(db *gorm.DB) (*ResumeEducation, error) {
	var resumeEducation ResumeEducation
	err := db.Where("id = ?", re.Model.ID).First(&resumeEducation).Error
	if err != nil {
		return &resumeEducation, err
	}
	return &resumeEducation, nil
}

func (re ResumeEducation) GetByResumeID(db *gorm.DB) (*[]ResumeEducation, error) {
	var resumeEducations []ResumeEducation
	err := db.Where("resume_id = ?", re.ResumeID).Find(&resumeEducations).Error
	if err != nil {
		return &resumeEducations, err
	}
	return &resumeEducations, nil
}

func (re ResumeExperience) Create(db *gorm.DB) (uint, error) {
	res := db.Create(&re)
	if res.Error != nil {
		return 0, res.Error
	}
	return re.ID, nil
}

func (re ResumeExperience) Update(db *gorm.DB, values map[string]interface{}) error {
	return db.Model(&ResumeExperience{}).Where("id = ?", re.Model.ID).Updates(values).Error
}

func (re ResumeExperience) Delete(db *gorm.DB) error {
	return db.Delete(&ResumeExperience{}, re.Model.ID).Error
}

func (re ResumeExperience) Get(db *gorm.DB) (*ResumeExperience, error) {
	var resumeExperience ResumeExperience
	err := db.Where("id = ?", re.Model.ID).First(&resumeExperience).Error
	if err != nil {
		return &resumeExperience, err
	}
	return &resumeExperience, nil
}

func (re ResumeExperience) GetByResumeID(db *gorm.DB) (*[]ResumeExperience, error) {
	var resumeExperiences []ResumeExperience
	err := db.Where("resume_id = ?", re.ResumeID).Find(&resumeExperiences).Error
	if err != nil {
		return &resumeExperiences, err
	}
	return &resumeExperiences, nil
}

func (rp ResumeProject) Create(db *gorm.DB) (uint, error) {
	res := db.Create(&rp)
	if res.Error != nil {
		return 0, res.Error
	}
	return rp.ID, nil
}

func (rp ResumeProject) Update(db *gorm.DB, values map[string]interface{}) error {
	return db.Model(&ResumeProject{}).Where("id = ?", rp.Model.ID).Updates(values).Error
}

func (rp ResumeProject) Delete(db *gorm.DB) error {
	return db.Delete(&ResumeProject{}, rp.Model.ID).Error
}

func (rp ResumeProject) Get(db *gorm.DB) (*ResumeProject, error) {
	var resumeProject ResumeProject
	err := db.Where("id = ?", rp.Model.ID).First(&resumeProject).Error
	if err != nil {
		return &resumeProject, err
	}
	return &resumeProject, nil
}

func (rp ResumeProject) GetByResumeID(db *gorm.DB) (*[]ResumeProject, error) {
	var resumeProjects []ResumeProject
	err := db.Where("resume_id = ?", rp.ResumeID).Find(&resumeProjects).Error
	if err != nil {
		return &resumeProjects, err
	}
	return &resumeProjects, nil
}
