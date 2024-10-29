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

func GetJobList(c *gin.Context) {
	// get all jobs
	location := c.Query("location")
	company := c.Query("company")
	salary := c.Query("salary")
	jobType := c.Query("job_type")

	job := &model.Job{
		Location: location,
		Company:  company,
		Salary:   salary,
		JobType:  jobType,
	}
	jobs, err := job.GetAll(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "get all jobs failed", err.Error())
		return
	}

	var jobsData []response.JobData
	for _, j := range *jobs {
		jobsData = append(jobsData, response.JobData{
			ID:          j.ID,
			Title:       j.Title,
			Description: j.Description,
			Demand:      j.Demand,
			Location:    j.Location,
			Company:     j.Company,
			Salary:      j.Salary,
			JobType:     j.JobType,
		})
	}

	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"jobs": jobsData}, "get all jobs success")
}

func GetJobByID(c *gin.Context) {
	// get job by id
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid id", err.Error())
		return
	}
	job := &model.Job{
		Model: &gorm.Model{
			ID: uint(idUint),
		},
	}
	job, err = job.Get(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "get job failed", err.Error())
		return
	}
	jobData := response.JobData{
		ID:          job.ID,
		Title:       job.Title,
		Description: job.Description,
		Demand:      job.Demand,
		Location:    job.Location,
		Company:     job.Company,
		Salary:      job.Salary,
		JobType:     job.JobType,
	}
	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"job": jobData}, "get job success")
}

func CreateJob(c *gin.Context) {
	// check role of user
	userData, _ := c.Get("user")
	role := userData.(map[string]string)["role"]
	if role != common.Role(1).String() {
		response.Error(c, http.StatusForbidden, response.CodeForbidden, "permission denied", "")
		return
	}

	var req param.ReqCreateJob
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid params", err.Error())
		return
	}

	job := &model.Job{
		Title:       req.Title,
		Description: req.Description,
		Demand:      req.Demand,
		Location:    req.Location,
		Company:     req.Company,
		Salary:      req.Salary,
		JobType:     req.JobType,
	}

	id, err := job.Create(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "create job failed", err.Error())
		return
	}
	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"id": id}, "create job success")
}
