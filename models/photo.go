package models

import (
	"final-task-pbi-rakamin/app"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type PhotoModel struct {
	db *gorm.DB
}

func NewPhotoModel(db *gorm.DB) *PhotoModel {
	return &PhotoModel{
		db: db,
	}
}

func (m *PhotoModel) UploadPhoto(photoCreate app.PhotoInput, file multipart.File, fileHeader *multipart.FileHeader) (*app.Photo, error) {
	photoCreate.PhotoURL = "placeholder"
	_, err := govalidator.ValidateStruct(photoCreate)
	if err != nil {
		return nil, err
	}

	//untuk menentukan direktori penyimpanan foto
	photoDir := "uploads/photos"
	if _, err := os.Stat(photoDir); os.IsNotExist(err) {
		os.MkdirAll(photoDir, os.ModePerm)
	}

	//Memberi nama pada foto
	photoUrl := filepath.Join(photoDir, fileHeader.Filename)

	// Menyimpan foto ke foder
	out, err := os.Create(photoUrl)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	_, err = file.Seek(0, 0) // Reset file pointer
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(out, file); err != nil {
		return nil, err
	}

	photo := app.Photo{
		Title:    photoCreate.Title,
		Caption:  photoCreate.Caption,
		UserID:   photoCreate.UserID,
		PhotoURL: photoUrl,
	}

	if err := m.db.Create(&photo).Error; err != nil {
		return nil, err
	}

	return &photo, nil
}

func (m *PhotoModel) Update(photoCreate app.PhotoInput, file multipart.File, fileHeader *multipart.FileHeader) (*app.Photo, error) {
	(id uint64, userUpdate *app.User) (*app.User, error) {
		var user app.User
		if err := m.db.Where("id = ?", id).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("user not found")
			}
			return nil, err
		}
	
		user.Username = userUpdate.Username
		user.Email = userUpdate.Email
		if userUpdate.Password != "" {
			user.Password = userUpdate.Password
		}
	
		if err := m.db.Save(&user).Error; err != nil {
			return nil, err
		}
	
		return &user, nil
	}
}