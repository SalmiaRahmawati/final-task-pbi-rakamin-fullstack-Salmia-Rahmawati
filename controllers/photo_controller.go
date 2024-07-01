package controllers

import (
	"final-task-pbi-rakamin/app"
	"final-task-pbi-rakamin/helpers"
	"final-task-pbi-rakamin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type PhotoController struct {
	photoModel *models.PhotoModel
}

func NewPhotoController(photoModel *models.PhotoModel) *PhotoController {
	return &PhotoController{
		photoModel: photoModel,
	}
}

func (p *PhotoController) UploadPhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	photoInput := app.PhotoInput{}
	userId := uint64(userData["userID"].(float64))
	photoInput.UserID = userId

	if contentType == appJSON {
		c.ShouldBindJSON(&photoInput)
	} else {
		c.ShouldBind(&photoInput)
	}

	// Mengecek jika foto sudah ter-upload
	file, fileHeader, err := c.Request.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}
	defer file.Close()

	photo, err := p.photoModel.UploadPhoto(photoInput, file, fileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createdPhoto := app.PhotoOutput{
		GormApp:  photo.GormApp,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID:   photo.UserID,
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo uploaded successfully", "photo": createdPhoto})

}
