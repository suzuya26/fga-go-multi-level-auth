package router

import (
	"sesi_12/controllers"
	"sesi_12/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.GET("/:productId", middlewares.ProductAuth(), controllers.GetProduct)
		productRouter.PUT("/:productId", middlewares.LevelCheck(), controllers.UpdateProduct)
		productRouter.DELETE("/:productId", middlewares.LevelCheck(), controllers.DeleteProduct)
		productRouter.GET("/", middlewares.LevelCheck(), controllers.GetAllProduct)
	}

	return r
}
