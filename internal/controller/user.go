package controller

import (
	"AI-Recruitment-backend/internal/global"
	"AI-Recruitment-backend/internal/global/param"
	"AI-Recruitment-backend/internal/global/response"
	"AI-Recruitment-backend/internal/model"
	"AI-Recruitment-backend/pkg/jwt"
	"AI-Recruitment-backend/pkg/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {
	var req param.ReqRegister
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid params", err.Error())
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "hash password failed", err.Error())
		return
	}
	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.Role,
		Age:      req.Age,
		Degree:   req.Degree,
	}

	id, err := user.Create(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "create user failed", err.Error())
		return
	}
	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"id": id}, "register success")
}

func Login(c *gin.Context) {
	var req param.ReqLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid params", err.Error())
		return
	}

	user := &model.User{
		Username: req.Username,
	}
	user, err := user.GetByUsername(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "get user failed", err.Error())
		return
	}

	if ok := util.CheckHashedPassword(req.Password, user.Password); !ok {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "password incorrect", "")
		return
	}

	tokenClaims := jwt.GenerateJwtToken(strconv.Itoa(int(user.ID)), user.Role, global.Config.Jwt.Expire, global.Config.Jwt.Issuer)
	token, err := jwt.GenerateJwtTokenString(tokenClaims, []byte(global.Config.Jwt.Key))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "generate jwt token failed", err.Error())
		return
	}
	response.Success(
		c, http.StatusOK,
		response.CodeSuccess,
		response.Data{
			"id":     user.ID,
			"role":   user.Role,
			"token":  token,
			"expire": global.Config.Jwt.Expire,
		},
		"login success")
}

func GetProfile(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid id", err.Error())
		return
	}
	user := &model.User{
		Model: &gorm.Model{ID: uint(idUint)},
	}
	user, err = user.Get(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "get user failed", err.Error())
		return
	}

	userData := response.UserData{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Age:      user.Age,
		Degree:   user.Degree,
	}
	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"profile": userData}, "get user profile success")
}
