// Package routes contains all routes of the application
package routes

import (
	authController "secure/challenge-2/infrastructure/restapi/controllers/auth"

	"github.com/gin-gonic/gin"
)

// AuthRoutes is a function that contains all routes of the auth
func AuthRoutes(router *gin.RouterGroup, controller *authController.Controller) {

	routerAuth := router.Group("/auth")
	{
		routerAuth.POST("/login", controller.Login)
		routerAuth.POST("/access-token", controller.GetAccessTokenByRefreshToken)
	}

}
