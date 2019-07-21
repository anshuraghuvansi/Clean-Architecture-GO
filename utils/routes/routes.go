package routes

import (
	"user/utils/middlewares"

	"github.com/gin-gonic/gin"
)

//Routes :
type Routes struct {
	Private *gin.RouterGroup
	Public  *gin.RouterGroup
}

//CreatePublicPrivateRoute :
func CreatePublicPrivateRoute(name string, apiGroup *gin.RouterGroup) Routes {
	privateAPI := createPrivateRoute(name, apiGroup)
	publicAPI := CreatePublicRoute(name, apiGroup)
	return Routes{Public: publicAPI, Private: privateAPI}
}

func createPrivateRoute(name string, g *gin.RouterGroup) *gin.RouterGroup {
	privateGroup := g.Group(name)
	privateGroup.Use(middlewares.AuthRequired())
	return privateGroup
}

//CreatePublicRoute :
func CreatePublicRoute(name string, g *gin.RouterGroup) *gin.RouterGroup {
	return g.Group(name)
}
