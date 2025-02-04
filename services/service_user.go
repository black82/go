package services

import (
	"awesomeProject/config"
	"awesomeProject/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	return users, result.Error
}

func CreateUser(user *models.User) error {
	result := config.DB.Create(user)
	return result.Error
}

func GetUserByID(id string) (*models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	return &user, result.Error
}

func UpdateUser(user *models.User) error {
	result := config.DB.Save(user)
	return result.Error
}

func DeleteUser(id string) error {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return err
	}
	return config.DB.Delete(&user).Error
}
