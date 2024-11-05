package controller

import (
	"AI-Recruitment-backend/internal/global"
	"AI-Recruitment-backend/internal/global/param"
	"AI-Recruitment-backend/internal/global/response"
	"AI-Recruitment-backend/internal/model"
	"AI-Recruitment-backend/pkg/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func CreateApplication(c *gin.Context) {
	userData, _ := c.Get("user")
	role := userData.(map[string]string)["role"]
	uid := userData.(map[string]string)["id"]
	uidInt, _ := strconv.Atoi(uid)
	if role != common.Candidate.String() {
		response.Error(c, http.StatusForbidden, response.CodeForbidden, "permission denied", "")
		return
	}

	var req param.ReqCreateApplication
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid params", err.Error())
		return
	}

	// verify if the job exists
	job := &model.Job{
		Model: &gorm.Model{ID: req.JobID},
	}
	if _, err := job.Get(global.DBEngine); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "job does not exist", err.Error())
		return
	}

	application := &model.Application{
		UserID:   uint(uidInt),
		JobID:    req.JobID,
		Progress: common.CandidateApplied,
	}

	id, err := application.Create(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "database error", err.Error())
		return
	}

	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"id": id}, "application created")
}

func GetApplicationsByUserId(c *gin.Context) {
	userData, _ := c.Get("user")
	role := userData.(map[string]string)["role"]
	uid := userData.(map[string]string)["id"]
	uidInt, _ := strconv.Atoi(uid)
	if role != common.Candidate.String() {
		response.Error(c, http.StatusForbidden, response.CodeForbidden, "permission denied", "")
		return
	}

	application := &model.Application{
		UserID: uint(uidInt),
	}

	applications, err := application.GetByUserID(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "database error", err.Error())
		return
	}

	var applicationsData []response.ApplicationData
	for _, application := range *applications {
		applicationsData = append(applicationsData, response.ApplicationData{
			ID:       application.ID,
			UserID:   application.UserID,
			JobID:    application.JobID,
			Progress: application.Progress,
		})
	}

	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"applications": applicationsData}, "applications retrieved")
}

func GetApplicationsByJobId(c *gin.Context) {
	userData, _ := c.Get("user")
	role := userData.(map[string]string)["role"]
	uid := userData.(map[string]string)["id"]
	uidInt, _ := strconv.Atoi(uid)
	if role != common.Recruiter.String() {
		response.Error(c, http.StatusForbidden, response.CodeForbidden, "permission denied", "")
		return
	}

	jobID := c.Param("job_id")
	jobIDInt, err := strconv.Atoi(jobID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid job_id", err.Error())
		return
	}

	// verify if the job exists and belongs to the recruiter
	job := &model.Job{
		Model: &gorm.Model{ID: uint(jobIDInt)},
	}

	job, err = job.Get(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "job does not exist", err.Error())
		return
	}
	if job.OwnerID != uint(uidInt) {
		response.Error(c, http.StatusForbidden, response.CodeForbidden, "permission denied", "")
		return
	}

	application := &model.Application{
		JobID: uint(jobIDInt),
	}

	applications, err := application.GetByJobID(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "database error", err.Error())
		return
	}

	var applicationsData []response.ApplicationData
	for _, application := range *applications {
		applicationsData = append(applicationsData, response.ApplicationData{
			ID:       application.ID,
			UserID:   application.UserID,
			JobID:    application.JobID,
			Progress: application.Progress,
		})
	}

	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"applications": applicationsData}, "applications retrieved")
}

func UpdateApplicationProgress(c *gin.Context) {
	userData, _ := c.Get("user")
	role := userData.(map[string]string)["role"]
	uid := userData.(map[string]string)["id"]
	uidInt, _ := strconv.Atoi(uid)

	var req param.ReqUpdateApplication
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid params", err.Error())
		return
	}
	// get application
	application := &model.Application{
		Model: &gorm.Model{ID: req.ID},
	}
	application, err := application.Get(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "application does not exist", err.Error())
		return
	}

	// check permission
	if role == common.Candidate.String() {
		if application.UserID != uint(uidInt) {
			response.Error(c, http.StatusForbidden, response.CodeForbidden, "permission denied", "")
			return
		}
		if application.Progress != common.RecruiterAccepted {
			response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid progress", "")
			return
		}
	} else if role == common.Recruiter.String() {
		job := &model.Job{
			Model: &gorm.Model{ID: application.JobID},
		}
		job, err := job.Get(global.DBEngine)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "database error", err.Error())
			return
		}
		if job.OwnerID != uint(uidInt) {
			response.Error(c, http.StatusForbidden, response.CodeForbidden, "permission denied", "")
			return
		}
		if application.Progress != common.CandidateApplied && application.Progress != common.RecruiterReviewed {
			response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid progress", "")
			return
		}
	} else {
		response.Error(c, http.StatusForbidden, response.CodeForbidden, "permission denied", "")
		return
	}

	// update progress
	switch application.Progress {
	case common.CandidateApplied:
		if req.Accepted {
			application.Progress = common.RecruiterReviewed
		} else {
			application.Progress = common.CandidateRejected
		}
	case common.RecruiterReviewed:
		if req.Accepted {
			application.Progress = common.RecruiterAccepted
		} else {
			application.Progress = common.RecruiterRejected
		}
	case common.RecruiterAccepted:
		if req.Accepted {
			application.Progress = common.CandidateAccepted
		} else {
			application.Progress = common.CandidateRejected
		}
	default:
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid progress", "")
	}

	application, err = application.Update(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "database error", err.Error())
		return
	}
	response.Success(c, http.StatusOK, response.CodeSuccess,
		response.Data{"id": application.ID, "user_id": application.UserID, "job_id": application.JobID, "progress": application.Progress},
		"progress updated")
}
