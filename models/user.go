package models

import (
	"errors"
	"final-task-pbi-rakamin/app"

	"gorm.io/gorm"
)

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

func (m *UserModel) CreateUser(userSignUp app.UserSignUp) (app.User, error) {
	user := app.User{
		Username: userSignUp.Username,
		Email:    userSignUp.Email,
		Password: userSignUp.Password,
	}
	err := m.db.Create(&user).Error
	return user, err
}

func (m *UserModel) GetUsersByEmail(email string) (app.User, error) {
	var user app.User
	err := m.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (m *UserModel) GetUsers() ([]app.User, error) {
	var users []app.User

	if err := m.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (m *UserModel) GetUsersByID(id uint64) (*app.User, error) {
	var user app.User
	if err := m.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (m *UserModel) UpdateUser(id uint64, userUpdate *app.User) (*app.User, error) {
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

func (m *UserModel) DeleteUser(id uint64, userDelete *app.User) (*app.User, error) {
	var user app.User
	if err := m.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if err := m.db.Delete(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
