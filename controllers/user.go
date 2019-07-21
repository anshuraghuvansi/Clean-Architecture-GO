package controllers

/*
Interface Adapters is a set of adapters that convert data from the format most convenient for
the use cases and entities, to the format most convenient for some external agency
such as the Database or the Web. It is this layer, for example, that will wholly
contain the MVC architecture of a GUI. The Presenters, Views,
and Controllers all belong in here.
*/

import (
	"user/constants"
	"user/models"
	"user/services"
	"user/utils/formatter"
	"user/utils/middlewares"
	"user/utils/routes"

	"github.com/gin-gonic/gin"
)

type userController struct {
	service services.UserService
}

//NewUserController :
func NewUserController(routes *routes.Routes, service services.UserService) {

	handler := userController{service: service}

	routes.Public.POST("signup", handler.signup)
	routes.Public.POST("signin", handler.signin)
	routes.Private.GET("profile", handler.getProfile)
	routes.Private.PUT("profile/update", handler.updateProfile)
}

func (c *userController) signup(context *gin.Context) {

	var params models.SignupBodyParams

	err := context.ShouldBindJSON(&params)
	if err != nil {
		formatter.HandleResponse(context, nil, constants.ErrorInvalidParams)
		return
	}

	err = params.Validate()
	if err != nil {
		formatter.HandleResponse(context, nil, err)
		return
	}

	token, err := c.service.SignupUser(params)
	formatter.HandleResponse(context, token, err)
}

func (c *userController) signin(context *gin.Context) {

	var params models.SigninBodyParams

	err := context.ShouldBindJSON(&params)
	if err != nil {
		formatter.HandleResponse(context, nil, constants.ErrorInvalidParams)
		return
	}

	err = params.Validate()
	if err != nil {
		formatter.HandleResponse(context, nil, err)
		return
	}

	token, err := c.service.SigninUser(params)
	formatter.HandleResponse(context, token, err)
}

func (c *userController) getProfile(context *gin.Context) {

	userID := context.GetInt64(middlewares.UserID)
	profile, err := c.service.GetProfile(userID)
	formatter.HandleResponse(context, profile, err)
}

func (c *userController) updateProfile(context *gin.Context) {

	var params models.UpdateProfileBodyParams

	err := context.ShouldBindJSON(&params)
	if err != nil {
		formatter.HandleResponse(context, nil, constants.ErrorInvalidParams)
		return
	}

	err = params.Validate()
	if err != nil {
		formatter.HandleResponse(context, nil, err)
		return
	}

	userID := context.GetInt64(middlewares.UserID)
	profile, err := c.service.UpdateProfile(userID, params)
	formatter.HandleResponse(context, profile, err)
}
