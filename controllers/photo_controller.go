package controllers

import (
	"final-task-pbi-rakamin/database"
	"final-task-pbi-rakamin/helpers"
	"final-task-pbi-rakamin/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GetPhoto(c *gin.Context) {
	var (
		photos []models.Photo
	)

	db := database.GetDB()
	err := db.Preload("User").Find(&photos).Error // Preload: untuk relasi dengan models user

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &photos,
	})
}

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint64(userData["id"].(float64))

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&Photo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid JSON",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&Photo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid form data",
				"message": err.Error(),
			})
			return
		}
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint64(userData["id"].(float64))

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&Photo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid JSON",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&Photo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid form data",
				"message": err.Error(),
			})
			return
		}
	}

	Photo.UserID = userID
	Photo.ID = uint64(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoURL: Photo.PhotoURL}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))

	err := db.Delete(&Photo, photoId).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
