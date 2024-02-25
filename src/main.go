package main

import (
	"log"
	"net/http"

	"github.com/eduardo-paes/cashflow/infra/data"
	"github.com/eduardo-paes/cashflow/infra/http/middleware"
	"github.com/eduardo-paes/cashflow/infra/http/routes"
	"github.com/eduardo-paes/cashflow/infra/ioc"
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
	conn, dbErr := data.GetConnection()
	if dbErr != nil {
		log.Printf("Error opening DB connection: %v", dbErr)
		return
	}

	// Running migrations
	mgErr := data.RunMigrations(conn)
	if mgErr != nil {
		log.Printf("Error running migrations: %v", mgErr)
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
	expenseController := ioc.ConfigExpenseDI(conn)
	userController := ioc.ConfigUserDI(conn, viper.GetString("jwt.secret"))

	// Home
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "CashFlow Restful API, v1")
	})

	// Authenticated routes
	mainGroup := router.Group("/api/v1")

	// Unauthenticated routes
	routes.ConfigureUserRoutes(mainGroup, userController)
	routes.ConfigureAuthRoutes(mainGroup, userController)

	// Authenticated routes
	mainGroup.Use(middleware.AuthMiddleware())
	{
		// Injecting core services
		routes.ConfigureExpenseRoutes(mainGroup, expenseController)
	}

	// Serving API
	port := viper.GetString("server.port")
	log.Printf("API is running on port: %v", port)
	log.Printf("You can access the Swagger via: http://localhost:%v/swagger/index.html", port)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
