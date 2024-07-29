package controllers

import (
	"final-task-pbi-rakamin/database"
	"final-task-pbi-rakamin/helpers"
	"final-task-pbi-rakamin/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid JSON",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid form data",
				"message": err.Error(),
			})
			return
		}
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to create user",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         User.ID,
		"username":   User.Username,
		"email":      User.Email,
		"created_at": User.CreatedAt,
		"updated_at": User.UpdatedAt,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	loginUser := models.User{}
	password := ""

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&loginUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid JSON",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&loginUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid form data",
				"message": err.Error(),
			})
			return
		}
	}

	password = loginUser.Password

	userFromDB := models.User{}
	err := db.Debug().Where("email = ?", loginUser.Email).Take(&userFromDB).Error

	if err != nil {
		log.Println("Error finding user with email:", loginUser.Email, "Error:", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(userFromDB.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(userFromDB.ID, userFromDB.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func GetUser(c *gin.Context) {
	var (
		users []models.User
	)

	db := database.GetDB()
	err := db.Find(&users).Error

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func UserUpdate(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	User := models.User{}

	id := c.Param("userId")

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid JSON",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid form data",
				"message": err.Error(),
			})
			return
		}
	}

	// Cari user berdasarkan id
	userID := models.User{}
	err := db.First(&userID, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{ // apabila user tidak ditemukan
			"error":   "User not found",
			"message": err.Error(),
		})
		return
	}

	// Update data user
	err = db.Model(&userID).Updates(User).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ // untuk kegagalan update
			"error":   "Failed to update user",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         User.ID,
		"email":      User.Email,
		"username":   User.Username,
		"created_at": User.CreatedAt,
		"updated_at": User.UpdatedAt,
	})
}

func UserDelete(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	id := c.Param("userId")

	err := db.First(&User, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = db.Delete(&User).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
