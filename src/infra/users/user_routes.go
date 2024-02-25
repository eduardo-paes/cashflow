package users

import (
	core "github.com/eduardo-paes/cashflow/core/users"
	"github.com/gin-gonic/gin"
)

// ConfigureUserRoutes configures routes related to users
func ConfigureUserRoutes(baseRoute *gin.RouterGroup, userController core.UserService) {
	// Use the Gin router group for /user
	userGroup := baseRoute.Group("/user")

	// POST /user
	userGroup.POST("", func(ctx *gin.Context) {
		userController.Create(ctx)
	})

	// DELETE /user/:id
	userGroup.DELETE("/:id", func(ctx *gin.Context) {
		userController.Delete(ctx)
	})

	// GET /users
	userGroup.GET("/:id", func(ctx *gin.Context) {
		userController.GetOne(ctx)
	})

	// PUT /user/:id
	userGroup.PUT("/:id", func(ctx *gin.Context) {
		userController.Update(ctx)
	})
}

// ConfigureAuthRoutes configures routes related to authentication
func ConfigureAuthRoutes(baseRoute *gin.RouterGroup, userController core.UserService) {
	// Use the Gin router group for /auth
	authGroup := baseRoute.Group("/auth")

	// POST /auth
	authGroup.POST("/login", func(ctx *gin.Context) {
		userController.Login(ctx)
	})
}
