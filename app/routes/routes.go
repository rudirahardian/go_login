package routes

import (
	"github.com/rudirahardian/go_env/app/controller"
	"github.com/rudirahardian/go_env/app/middleware"
	"github.com/gin-gonic/gin"
)

func RouteInit(port string, mode string){
	gin.SetMode(mode)
	router := gin.Default()
	router.Use(middleware.PanicHandler)

	v1User := router.Group("/api/v1/user")
	{
		v1User.POST("/register", controller.V1UserRegister)
		v1User.POST("/login", controller.V1UserLogin)
	}
	{
		v1UserNoAuth := v1User.Group("/")
		v1UserNoAuth.Use(middleware.AuthMiddleware)
		v1UserNoAuth.GET("/get-user", controller.V1UserGet)
	}

	router.Run(":"+port)
}