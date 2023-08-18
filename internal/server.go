package app

import (
	"campfire/internal/app/auth"
	"campfire/internal/app/organization"
	"os"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func NewApp() *App {
	app := &App{}
	app.router = gin.Default()
	app.router.SetTrustedProxies([]string{"http://localhost:3000"})

	auth.AuthRoutes(app.router)
	organization.OrganizationRoutes(app.router)

	return app
}

func (app App) Start() {
	app.router.Run(":" + os.Getenv("SERVER_PORT"))
}
