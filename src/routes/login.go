package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jhonnydsl/clinify-backend/src/controllers"
	"github.com/jhonnydsl/clinify-backend/src/repository"
	"github.com/jhonnydsl/clinify-backend/src/services"
)

func SetupLoginRoutes(app *gin.RouterGroup) {
	loginService := &services.LoginService{Repo: &repository.LoginRepository{}}
	loginController := &controllers.LoginController{Service: loginService}

	app.POST("/login", loginController.LoginUser)
}