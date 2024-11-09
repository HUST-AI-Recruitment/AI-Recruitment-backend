package controller

import (
	"AI-Recruitment-backend/internal/global"
	"AI-Recruitment-backend/internal/global/param"
	"AI-Recruitment-backend/internal/global/response"
	"AI-Recruitment-backend/internal/model"
	"AI-Recruitment-backend/pkg/common"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
)

func RecommendJobs(c *gin.Context) {
	userData, _ := c.Get("user")
	role := userData.(map[string]string)["role"]
	uid := userData.(map[string]string)["id"]
	uidInt, _ := strconv.Atoi(uid)
	if role != common.Candidate.String() {
		response.Error(c, http.StatusForbidden, response.CodeForbidden, "permission denied", "")
		return
	}

	// get resume
	resume := &model.Resume{
		UserID: uint(uidInt),
	}
	resume, err := resume.GetByUserID(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "get resume list failed", err.Error())
		return
	}

	resumeData := &response.ResumeData{
		ID:          resume.ID,
		UserID:      resume.UserID,
		Name:        resume.Name,
		Gender:      resume.Gender,
		Phone:       resume.Phone,
		Email:       resume.Email,
		Wechat:      resume.Wechat,
		State:       resume.State,
		Description: resume.Description,
	}

	edu := &model.ResumeEducation{
		ResumeID: resume.ID,
	}
	eduList, err := edu.GetByResumeID(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "get resume education list failed", err.Error())
		return
	}
	var eduDataList []response.ResumeEducation
	for _, v := range *eduList {
		eduDataList = append(eduDataList, response.ResumeEducation{
			School:    v.School,
			Major:     v.Major,
			Degree:    v.Degree,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}

	exp := &model.ResumeExperience{
		ResumeID: resume.ID,
	}
	expList, err := exp.GetByResumeID(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "get resume experience list failed", err.Error())
		return
	}
	var expDataList []response.ResumeExperience
	for _, v := range *expList {
		expDataList = append(expDataList, response.ResumeExperience{
			Company:   v.Company,
			Position:  v.Position,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}

	project := &model.ResumeProject{
		ResumeID: resume.ID,
	}
	projectList, err := project.GetByResumeID(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "get resume project list failed", err.Error())
		return
	}
	var projectDataList []response.ResumeProject
	for _, v := range *projectList {
		projectDataList = append(projectDataList, response.ResumeProject{
			Name:        v.Name,
			Description: v.Description,
			StartTime:   v.StartTime,
			EndTime:     v.EndTime,
		})
	}

	resumeData.Education = eduDataList
	resumeData.Experience = expDataList
	resumeData.Project = projectDataList

	// get all jobs
	job := &model.Job{}

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
			OwnerID:     j.OwnerID,
		})
	}

	// recommend jobs
	requestBody := struct {
		Resume *response.ResumeData `json:"resume"`
		Jobs   []response.JobData   `json:"jobs"`
	}{
		Resume: resumeData,
		Jobs:   jobsData,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "json marshal failed", err.Error())
		return
	}
	url := fmt.Sprintf("%s/recommend_jobs/resume", global.Config.AI.Addr)

	recommendReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "create request failed", err.Error())
		return
	}
	recommendReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(recommendReq)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "send request failed", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "read response body failed", err.Error())
		return
	}
	var recommendJobs struct {
		Job []int `json:"job"`
	}
	if err := json.Unmarshal(body, &recommendJobs); err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "json unmarshal failed", err.Error())
		return
	}

	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"jobs": recommendJobs.Job}, "recommend jobs success")
}

func RecommendJobsByDescription(c *gin.Context) {
	userData, _ := c.Get("user")
	role := userData.(map[string]string)["role"]
	if role != common.Candidate.String() {
		response.Error(c, http.StatusForbidden, response.CodeForbidden, "permission denied", "")
		return
	}

	var req param.ReqRecommendJobsDescription
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid param", err.Error())
		return
	}
	description := req.Description

	// get all jobs
	job := &model.Job{}

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
			OwnerID:     j.OwnerID,
		})
	}

	// recommend jobs
	requestBody := struct {
		Description string             `json:"description"`
		Jobs        []response.JobData `json:"jobs"`
	}{
		Description: description,
		Jobs:        jobsData,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "json marshal failed", err.Error())
		return
	}
	url := fmt.Sprintf("%s/recommend_jobs/description", global.Config.AI.Addr)

	recommendReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "create request failed", err.Error())
		return
	}
	recommendReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(recommendReq)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "send request failed", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "read response body failed", err.Error())
		return
	}
	var recommendJobs struct {
		Job []int `json:"job"`
	}
	if err := json.Unmarshal(body, &recommendJobs); err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "json unmarshal failed", err.Error())
		return
	}

	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"jobs": recommendJobs.Job}, "recommend jobs success")
}

func RankCandidates(c *gin.Context) {
	userData, _ := c.Get("user")
	role := userData.(map[string]string)["role"]
	uid := userData.(map[string]string)["id"]
	uidInt, _ := strconv.Atoi(uid)
	if role != common.Recruiter.String() {
		response.Error(c, http.StatusForbidden, response.CodeForbidden, "permission denied", "")
		return
	}

	jobIDStr := c.Query("job_id")
	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid job id", err.Error())
		return
	}

	// get job
	job := &model.Job{
		Model: &gorm.Model{
			ID: uint(jobID),
		},
	}
	job, err = job.Get(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "get job failed", err.Error())
		return
	}
	// verify if the job belongs to the recruiter
	if job.OwnerID != uint(uidInt) {
		response.Error(c, http.StatusForbidden, response.CodeForbidden, "permission denied", "")
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
		OwnerID:     job.OwnerID,
	}

	// get resumes
	application := &model.Application{
		JobID: uint(jobID),
	}
	applications, err := application.GetByJobID(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "database error", err.Error())
		return
	}

	var resumesData []response.ResumeData

	for _, a := range *applications {
		resume := &model.Resume{
			UserID: a.UserID,
		}
		resume, err := resume.GetByUserID(global.DBEngine)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "get resume list failed", err.Error())
			return
		}

		resumeData := &response.ResumeData{
			ID:          resume.ID,
			UserID:      resume.UserID,
			Name:        resume.Name,
			Gender:      resume.Gender,
			Phone:       resume.Phone,
			Email:       resume.Email,
			Wechat:      resume.Wechat,
			State:       resume.State,
			Description: resume.Description,
		}

		edu := &model.ResumeEducation{
			ResumeID: resume.ID,
		}
		eduList, err := edu.GetByResumeID(global.DBEngine)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "get resume education list failed", err.Error())
			return
		}
		var eduDataList []response.ResumeEducation
		for _, v := range *eduList {
			eduDataList = append(eduDataList, response.ResumeEducation{
				School:    v.School,
				Major:     v.Major,
				Degree:    v.Degree,
				StartTime: v.StartTime,
				EndTime:   v.EndTime,
			})
		}

		exp := &model.ResumeExperience{
			ResumeID: resume.ID,
		}
		expList, err := exp.GetByResumeID(global.DBEngine)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "get resume experience list failed", err.Error())
			return
		}
		var expDataList []response.ResumeExperience
		for _, v := range *expList {
			expDataList = append(expDataList, response.ResumeExperience{
				Company:   v.Company,
				Position:  v.Position,
				StartTime: v.StartTime,
				EndTime:   v.EndTime,
			})
		}

		project := &model.ResumeProject{
			ResumeID: resume.ID,
		}
		projectList, err := project.GetByResumeID(global.DBEngine)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "get resume project list failed", err.Error())
			return
		}
		var projectDataList []response.ResumeProject
		for _, v := range *projectList {
			projectDataList = append(projectDataList, response.ResumeProject{
				Name:        v.Name,
				Description: v.Description,
				StartTime:   v.StartTime,
				EndTime:     v.EndTime,
			})
		}

		resumeData.Education = eduDataList
		resumeData.Experience = expDataList
		resumeData.Project = projectDataList

		resumesData = append(resumesData, *resumeData)
	}

	// recommend resumes
	requestBody := struct {
		Job     response.JobData      `json:"job"`
		Resumes []response.ResumeData `json:"resumes"`
	}{
		Job:     jobData,
		Resumes: resumesData,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "json marshal failed", err.Error())
		return
	}
	url := fmt.Sprintf("%s/rank_candidates", global.Config.AI.Addr)

	recommendReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "create request failed", err.Error())
		return
	}
	recommendReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(recommendReq)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "send request failed", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "read response body failed", err.Error())
		return
	}
	var recommendResumes []struct {
		ID    uint `json:"id"`
		Score int  `json:"score"`
	}
	if err := json.Unmarshal(body, &recommendResumes); err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "json unmarshal failed, body: "+string(body), err.Error())
		return
	}

	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"score": recommendResumes}, "rank candidates success")
}
