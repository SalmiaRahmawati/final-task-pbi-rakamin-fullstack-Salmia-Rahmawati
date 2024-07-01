package router

import (
	"final-task-pbi-rakamin/controllers"
	"final-task-pbi-rakamin/database"
	"final-task-pbi-rakamin/middlewares"
	"final-task-pbi-rakamin/models"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userModel := models.NewUserModel(database.DB)
	userController := controllers.NewUserController(userModel)

	userRouter := r.Group("/users")
	{
		userRouter.POST("/signUp", userController.SignUp)
		userRouter.POST("/login", userController.Login)
		userRouter.GET("", userController.GetUsers)
		userRouter.GET("/:id", userController.GetUsersById)
		userRouter.PUT("/:id", userController.UpdateUser)
		userRouter.DELETE("/:id", userController.DeleteUser)
	}

	photoModel := models.NewPhotoModel(database.DB)
	photoController := controllers.NewPhotoController(photoModel)

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", photoController.UploadPhoto)
	}
	return r
}

// type UserRouter interface {
// 	Mount()
// }

// type userRouterImpl struct {
// 	v           *gin.RouterGroup
// 	controllers controllers.UserController
// }

// func NewUserRouter(v *gin.RouterGroup, controllers controllers.UserController) UserRouter {
// 	return &userRouterImpl{v: v, controllers: controllers}
// }

// func (u *userRouterImpl) Mount() {
// 	u.v.POST("/register", u.controllers.SignUp)
// }

// func StartApp() *gin.Engine {
// 	r := gin.Default()

// 	// userModel := models.NewUserModel(database.DB)
// 	//userController := controllers.NewUserController(userModel)

// 	userRouter := r.Group("/users")
// 	{
// 		userRouter.POST("/signup", u.controllers.SignUp)
// 	}
// 	return r
// }

// func StartApp() *gin.Engine {
// 	r := gin.Default()

// 	userRouter := r.Group("/users")
// 	{
// 		userRouter.POST("/signUp", controllers.SignUp)
// 	}
// 	return r
// }
