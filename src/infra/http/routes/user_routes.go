package routes

import (
	core "github.com/eduardo-paes/cashflow/core/entities"
	"github.com/gin-gonic/gin"
)

// ConfigureUserRoutes configures routes related to users
func ConfigureUserRoutes(baseRoute *gin.RouterGroup, userController core.UserService) {
	// Use the Gin router group for /user
	userGroup := baseRoute.Group("/user")

	// POST /user
	userGroup.POST("", func(c *gin.Context) {
		userController.Create(c.Writer, c.Request)
	})

	// DELETE /user/:id
	userGroup.DELETE("/:id", func(c *gin.Context) {
		userController.Delete(c)
	})

	// GET /users
	userGroup.GET("/:id", func(c *gin.Context) {
		userController.GetOne(c)
	})

	// PUT /user/:id
	userGroup.PUT("/:id", func(c *gin.Context) {
		userController.Update(c)
	})
}

// ConfigureAuthRoutes configures routes related to authentication
func ConfigureAuthRoutes(baseRoute *gin.RouterGroup, userController core.UserService) {
	// Use the Gin router group for /auth
	authGroup := baseRoute.Group("/auth")

	// POST /auth
	authGroup.POST("/login", func(c *gin.Context) {
		userController.Login(c.Writer, c.Request)
	})
}
