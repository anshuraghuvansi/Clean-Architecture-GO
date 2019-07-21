package routes

import (
	"user/controllers"
	"user/repositories"
	"user/services"
	"user/utils/database"
	route "user/utils/routes"

	"github.com/gin-gonic/gin"
)

//Init :
func Init(app *gin.Engine, db *database.DataBase) {
	apiGroup := app.Group("/api")
	initV1Routes(apiGroup, db)
}

func initV1Routes(apiGroup *gin.RouterGroup, db *database.DataBase) {
	v1Group := route.CreatePublicRoute("/v1", apiGroup)
	initUserRoutes(v1Group, db)
}

func initUserRoutes(v1Group *gin.RouterGroup, db *database.DataBase) {
	userRoute := route.CreatePublicPrivateRoute("/user", v1Group)
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	controllers.NewUserController(&userRoute, userService)
}
