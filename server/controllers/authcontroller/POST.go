package authcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"net/http"
	"outstagram/server/dtos/authdtos"
	"outstagram/server/models"
	"outstagram/server/utils"
	"time"
)

func (ac *Controller) Login(c *gin.Context) {
	var reqBody authdtos.LoginRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Some required fields missing", nil)
		return
	}

	user, err := ac.userService.FindByUsername(reqBody.Username)
	if gorm.IsRecordNotFoundError(err) {
		utils.ResponseWithError(c, http.StatusNotFound, "Username not found", nil)
		return
	}

	if user.Password != reqBody.Password {
		utils.ResponseWithError(c, http.StatusConflict, "Username or password incorrect", nil)
		return
	}

	token, err := utils.SignToken(user)
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Signing token failed", err.Error())
		return
	}

	user.LastLogin = utils.NewTimePointer(time.Now())
	if err = ac.userService.Save(user); err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Saving user failed", err.Error())
		return
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Login successfully", token)
}

func (ac *Controller) Register(c *gin.Context) {
	var reqBody authdtos.RegisterRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Some required fields missing", err.Error())
		return
	}

	if ac.userService.CheckExistsByUsername(reqBody.Username) {
		utils.ResponseWithError(c, http.StatusConflict, "Username already used", nil)
		return
	}

	if ac.userService.CheckExistsByEmail(reqBody.Username) {
		utils.ResponseWithError(c, http.StatusConflict, "Email already used", nil)
		return
	}

	newUser := models.User{}
	if err := copier.Copy(&newUser, &reqBody); err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while copying from request body to model", err.Error())
		return
	}

	if err := ac.userService.Save(&newUser); err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Creating user failed", err.Error())
		return
	}

	utils.ResponseWithSuccess(c, http.StatusCreated, "Create user successfully", nil)
}
