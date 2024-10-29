package controller

import (
	"AI-Recruitment-backend/internal/global"
	"AI-Recruitment-backend/internal/global/param"
	"AI-Recruitment-backend/internal/global/response"
	"AI-Recruitment-backend/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetJobList(c *gin.Context) {
	// get all jobs
	location := c.Query("location")
	company := c.Query("company")
	salary := c.Query("salary")

	job := &model.Job{
		Location: location,
		Company:  company,
		Salary:   salary,
	}
	jobs, err := job.GetAll(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "get all jobs failed", err.Error())
		return
	}
	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"jobs": jobs}, "get all jobs success")
}

func GetJobByID(c *gin.Context) {
	// get job by id
	id := c.Param("id")
	job := &model.Job{}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid id", err.Error())
		return
	}
	job.ID = uint(idInt)
	job, err = job.Get(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "get job failed", err.Error())
		return
	}
	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"job": job}, "get job success")
}

func CreateJob(c *gin.Context) {
	var req param.ReqCreateJob
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid params", err.Error())
		return
	}

	job := &model.Job{
		Title:       req.Title,
		Description: req.Description,
		Location:    req.Location,
		Company:     req.Company,
		Salary:      req.Salary,
	}

	id, err := job.Create(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "create job failed", err.Error())
		return
	}
	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"id": id}, "create job success")
}
