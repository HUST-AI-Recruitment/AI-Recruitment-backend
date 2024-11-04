package controller

import (
	"AI-Recruitment-backend/internal/global"
	"AI-Recruitment-backend/internal/global/param"
	"AI-Recruitment-backend/internal/global/response"
	"AI-Recruitment-backend/internal/model"
	"AI-Recruitment-backend/pkg/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateResume(c *gin.Context) {
	userData, _ := c.Get("user")
	role := userData.(map[string]string)["role"]
	uid := userData.(map[string]string)["id"]
	uidInt, _ := strconv.Atoi(uid)
	if role != common.Role(2).String() {
		response.Error(c, http.StatusForbidden, response.CodeForbidden, "permission denied", "")
		return
	}

	var req param.ReqUpsertResume
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid params", err.Error())
		return
	}

	resume := &model.Resume{
		UserID:      uint(uidInt),
		Name:        req.Name,
		Gender:      req.Gender,
		Phone:       req.Phone,
		Email:       req.Email,
		Wechat:      req.Wechat,
		State:       req.State,
		Description: req.Description,
	}
	var resumeEducation []model.ResumeEducation
	for _, v := range req.Education {
		resumeEducation = append(resumeEducation, model.ResumeEducation{
			School:    v.School,
			Major:     v.Major,
			Degree:    v.Degree,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}
	var resumeExperience []model.ResumeExperience
	for _, v := range req.Experience {
		resumeExperience = append(resumeExperience, model.ResumeExperience{
			Company:   v.Company,
			Position:  v.Position,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}
	var resumeProject []model.ResumeProject
	for _, v := range req.Project {
		resumeProject = append(resumeProject, model.ResumeProject{
			Name:        v.Name,
			Description: v.Description,
			StartTime:   v.StartTime,
			EndTime:     v.EndTime,
		})
	}

	id, err := resume.CreateResume(global.DBEngine, &resumeEducation, &resumeExperience, &resumeProject)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "create resume failed", err.Error())
		return
	}

	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"id": id}, "create resume success")
}

func GetResumeByUserId(c *gin.Context) {
	userId := c.Param("user_id")
	uidInt, err := strconv.Atoi(userId)
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid user id", err.Error())
		return
	}

	resume := &model.Resume{
		UserID: uint(uidInt),
	}
	resume, err = resume.GetByUserID(global.DBEngine)
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

	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"resume": resumeData}, "get resume success")
}

func UpdateResumeByUserId(c *gin.Context) {
	userData, _ := c.Get("user")
	role := userData.(map[string]string)["role"]
	uid := userData.(map[string]string)["id"]
	uidInt, _ := strconv.Atoi(uid)
	if role != common.Role(2).String() {
		response.Error(c, http.StatusForbidden, response.CodeForbidden, "permission denied", "")
		return
	}

	var req param.ReqUpsertResume
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid params", err.Error())
		return
	}

	resume := &model.Resume{
		UserID:      uint(uidInt),
		Name:        req.Name,
		Gender:      req.Gender,
		Phone:       req.Phone,
		Email:       req.Email,
		Wechat:      req.Wechat,
		State:       req.State,
		Description: req.Description,
	}
	var resumeEducation []model.ResumeEducation
	for _, v := range req.Education {
		resumeEducation = append(resumeEducation, model.ResumeEducation{
			School:    v.School,
			Major:     v.Major,
			Degree:    v.Degree,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}
	var resumeExperience []model.ResumeExperience
	for _, v := range req.Experience {
		resumeExperience = append(resumeExperience, model.ResumeExperience{
			Company:   v.Company,
			Position:  v.Position,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}
	var resumeProject []model.ResumeProject
	for _, v := range req.Project {
		resumeProject = append(resumeProject, model.ResumeProject{
			Name:        v.Name,
			Description: v.Description,
			StartTime:   v.StartTime,
			EndTime:     v.EndTime,
		})
	}

	resume, err := resume.UpdateResume(global.DBEngine, &resumeEducation, &resumeExperience, &resumeProject)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "update resume failed", err.Error())
		return
	}
	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"id": resume.ID}, "update resume success")
}
