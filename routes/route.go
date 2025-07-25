package routes

import (
	"backend/controllers"
	middlewares "backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	router.POST("/api/register", controllers.Register)
	router.POST("/api/login", controllers.Login)
	router.GET("/api/users", middlewares.AuthMiddleware(), controllers.FindUsers)

	router.Static("/uploads", "./uploads")
	router.GET("/api/products", controllers.GetProduct)
	router.POST("/api/products", controllers.CreateProduct)
	router.PUT("/api/products/:id", controllers.UpdateProduct)
	router.DELETE("/api/products/:id", controllers.DeleteProduct)

	return router
}

