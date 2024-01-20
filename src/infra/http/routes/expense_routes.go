package routes

import (
	"net/http"
	"strconv"

	core "github.com/eduardo-paes/cashflow/core/entities"
	"github.com/gin-gonic/gin"
)

// ConfigureExpenseRoutes configures routes related to expenses
func ConfigureExpenseRoutes(router *gin.Engine, expenseController core.ExpenseService) *gin.Engine {
	// Use the Gin router group for /expense
	baseRoute := router.Group("/api/v1")
	expenseGroup := baseRoute.Group("/expense")

	// POST /expense
	expenseGroup.POST("", func(c *gin.Context) {
		expenseController.Create(c.Writer, c.Request)
	})

	// DELETE /expense/:id
	expenseGroup.DELETE("/:id", func(c *gin.Context) {
		_, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		expenseController.Delete(c.Writer, c.Request)
	})

	// GET /expenses
	expenseGroup.GET("", func(c *gin.Context) {
		expenseController.GetOneOrMany(c.Writer, c.Request)
	})

	// PUT /expense/:id
	expenseGroup.PUT("/:id", func(c *gin.Context) {
		_, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		expenseController.Update(c.Writer, c.Request)
	})

	return router
}
