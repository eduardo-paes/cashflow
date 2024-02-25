package main

import (
	"github.com/eduardo-paes/cashflow/infra/configs"
	"github.com/eduardo-paes/cashflow/infra/expenses"
	"github.com/eduardo-paes/cashflow/infra/users"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

//	@title			CashFlow API
//	@version		1.0
//	@description	API for managing expenses in CashFlow application.

//	@contact.name	Eduardo Paes
//	@contact.url	https://twitter.com/edpaes
//	@contact.email	eduardo-paes@outlook.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host						localhost:3008
// @basePath					/api/v1
// @schemes					http
// @security					BasicAuth
// @security					BearerToken
// @securityDefinitions.apiKey	JWT
// @in							header
// @name						token
func main() {
	// Open database connection
	conn, err := configs.GetConnection()
	if err != nil {
		log.Printf("Error opening DB connection: %v", err)
		return
	}

	// Running migrations
	err = configs.RunMigrations(conn)
	if err != nil {
		log.Printf("Error running migrations: %v", err)
		return
	}

	// Creating routes
	router := gin.Default()

	// Serve static files from the docs folder
	router.Static("/docs", "./docs")

	// Add Swagger
	swaggerURL := ginSwagger.URL("/docs/swagger.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))

	// Create controllers
	expenseController := expenses.ConfigExpenseDI(conn)
	userController := users.ConfigUserDI(
		conn,
		viper.GetString("security.jwtSecret"),
		viper.GetString("security.passwordSalt"))

	// Home
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "CashFlow Restful API, v1")
	})

	// Authenticated routes
	mainGroup := router.Group("/api/v1")

	// Unauthenticated routes
	users.ConfigureAuthRoutes(mainGroup, userController)

	// Authenticated routes
	mainGroup.Use(configs.AuthMiddleware())
	{
		// Injecting core services
		expenses.ConfigureExpenseRoutes(mainGroup, expenseController)
		users.ConfigureUserRoutes(mainGroup, userController)
	}

	// Serving API
	port := viper.GetString("server.port")
	log.Printf("API is running on port: %v", port)
	log.Printf("You can access the Swagger via: http://localhost:%v/swagger/index.html", port)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
