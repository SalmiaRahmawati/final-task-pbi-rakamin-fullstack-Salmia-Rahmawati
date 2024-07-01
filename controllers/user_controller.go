package controllers

import (
	"final-task-pbi-rakamin/app"
	"final-task-pbi-rakamin/helpers"
	"final-task-pbi-rakamin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userModel *models.UserModel
}

func NewUserController(userModel *models.UserModel) *UserController {
	return &UserController{
		userModel: userModel,
	}
}

var (
	appJSON = "application/json"
)

func (u *UserController) SignUp(c *gin.Context) {
	// db := database.GetDB()
	contentType := helpers.GetContentType(c)
	// _, _ = db, contentType
	userSignUp := app.UserSignUp{}

	if contentType == appJSON {
		c.ShouldBindJSON(&userSignUp)
	} else {
		c.ShouldBind(&userSignUp)
	}

	user, err := u.userModel.CreateUser(userSignUp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createdUser := app.UserSignUpOutput{
		GormApp:  user.GormApp,
		Username: userSignUp.Username,
		Email:    userSignUp.Email,
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": createdUser})
}

func (u *UserController) Login(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userLogin := app.UserLogin{}

	if contentType == appJSON {
		c.ShouldBindJSON(&userLogin)
	} else {
		c.ShouldBind(&userLogin)
	}

	user, err := u.userModel.GetUsersByEmail(userLogin.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid email or password"})
		return
	}

	pass := helpers.CheckPasswordHash(userLogin.Password, user.Password)
	if !pass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token, err := helpers.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (u *UserController) GetUsers(c *gin.Context) {
	users, err := u.userModel.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (u *UserController) GetUsersById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if id == 0 || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid required param",
		})
		return
	}

	user := app.User{}
	users, err := u.userModel.GetUsersByID(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (u *UserController) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid required param",
		})
		return
	}

	contentType := helpers.GetContentType(c)
	user := app.User{}
	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	updatedUser, err := u.userModel.UpdateUser(id, &user)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    updatedUser,
	})
}

func (u *UserController) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid required param",
		})
		return
	}

	contentType := helpers.GetContentType(c)
	user := app.User{}
	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	deleteUser, err := u.userModel.DeleteUser(id, &user)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User successfully deleted",
		"user":    deleteUser,
	})
}

// func SignUp(c *gin.Context) {
// 	var userSignUp app.UserSignUp
// 	if err := c.ShouldBindJSON(&userSignUp); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	hashedPassword, err := helpers.HashPass(userSignUp.Password)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	user := &app.User{
// 		Username: userSignUp.Username,
// 		Email:    userSignUp.Email,
// 		Password: hashedPassword,
// 	}

// 	userModel := models.NewUserModel(database.DB)
// 	createdUser, err := userModel.CreateUser(user)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": createdUser})
// }

// func Login(c *gin.Context) {
// 	var user app.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	dbUser, err := models.GetUsersByEmail(user.Email)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
// 		return
// 	}

// 	if !helpers.CheckPasswordHash(user.Password, dbUser.Password) {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
// 		return
// 	}

// 	token, err := helpers.GenerateJWT(dbUser)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"token": token})
// }

// type UserController struct {
// 	userModel models.UserModel
// }

// func NewUserController(userModel models.UserModel) *UserController {
// 	return &UserController{
// 		userModel: userModel,
// 	}
// }

// userModel := models.NewUserModel(db * gorm.DB)
// createdUser, err := userModel.CreateUser(user)
// if err != nil {
// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	return
// }
// ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": createdUser})

// hashedPassword, err := helpers.HashPass(userSignUp.Password)
// if err != nil {
// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	return
// }

// func (uc *UserController) GetUsersByID(ctx *gin.Context) {
// 	id, err := strconv.Atoi(ctx.Param("id"))
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"status":  "error",
// 			"message": "Invalid user ID",
// 		})
// 		return
// 	}

// 	user, err := uc.userModel.GetUsersByID(uint64(id))
// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, gin.H{
// 			"status":  "error",
// 			"message": "User not found",
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"status": "OK",
// 		"data":   user,
// 	})
// }
