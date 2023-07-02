package main

import (
	"fmt"
	"log"
	"net/http"
	"ref/config"
	"ref/controller"
	"ref/docs"
	"ref/entity"
	"ref/infrastructure/driver"
	"ref/infrastructure/middleware"
	"ref/infrastructure/repogitory"
	"ref/usecase"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title		ref
// @description	ref
// @version		1.0
func main() {
	// Load config
	conf := config.Load()
	db := driver.NewDB(conf)
	tokenDriver := driver.NewTokenDriver(conf)

	ingredientRepogitory := repogitory.NewIngredientRepogitory(tokenDriver)
	componentRepogitory := repogitory.NewComponentRepogitory(tokenDriver)

	ingredientUseCase := usecase.NewIngredientUseCase(ingredientRepogitory)
	componentUseCase := usecase.NewComponentUseCase(componentRepogitory)

	ingredientController := controller.NewIngredientController(ingredientUseCase)
	componentController := controller.NewComponentController(componentUseCase)

	// Setup webserver
	app := gin.Default()
	app.Use(middleware.Transaction(db))

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "It works")
	})

	api := app.Group("/api/v1")

	ingredientRouter := api.Group("/ingredients")
	ingredientRouter.GET("/", handleResponse(ingredientController.GetIngredients))

	componentRouter := api.Group("/components")
	componentRouter.GET("/", handleResponse(componentController.GetComponents))
	componentRouter.GET("/:componentId", handleResponse(componentController.GetComponentByID))

	runApp(app, conf)
}

func runApp(app *gin.Engine, conf *config.Config) {
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", conf.Port)
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.Run(fmt.Sprintf("%s:%d", conf.Hostname, conf.Port))

	log.Println(fmt.Sprintf("http://localhost:%d", conf.Port))
	log.Println(fmt.Sprintf("http://localhost:%d/swagger/index.html", conf.Port))
}

func handleResponse(f func(ctx *gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := f(c)
		if err != nil {
			e, ok := err.(*entity.Error)
			if ok {
				c.JSON(e.Code, entity.ErrorResponse{Message: err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, entity.ErrorResponse{Message: err.Error()})
			}
			c.Abort()
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}
