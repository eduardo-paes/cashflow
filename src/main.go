package main

import (
	"log"
	"net/http"

	"github.com/eduardo-paes/cashflow/infra/data"
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

// @title CashFlow API
// @version 1.0
// @description API for managing expenses in CashFlow application.
// @host localhost:3008
// @basePath /api/v1
// @schemes http
// @BasePath /api/v1
func main() {
	// Openning database connection
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

	// Home
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "CashFlow Restful API, v1")
	})

	// Injecting core services
	expenseController := ioc.ConfigExpenseDI(conn)
	endpoints := routes.ConfigureExpenseRoutes(router, expenseController)

	// Serving API
	port := viper.GetString("server.port")
	log.Printf("API is running on port: %v", port)
	log.Printf("You can access the Swagger via: http://localhost:%v/swagger/index.html", port)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: endpoints,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}