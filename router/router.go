package router

import (
	"final-task-pbi-rakamin/controllers"
	"final-task-pbi-rakamin/database"
	"final-task-pbi-rakamin/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	// inisialisasi koneksi database
	database.StartDB()

	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.GET("/", controllers.GetUser)
		userRouter.PUT("/:userId", controllers.UserUpdate)
		userRouter.DELETE("/:userId", controllers.UserDelete)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetPhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", controllers.DeletePhoto)
	}
	return r
}
