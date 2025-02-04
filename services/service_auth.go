package services

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"gorm.io/gorm"
)

func Auth(m models.AuthClaim) (bool, error) {
	var auth models.AuthRequestUser
	result := config.DB.Where("name = ? AND password = ?", m.UserName, m.Password).First(&auth)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, result.Error
		}
		return false, result.Error
	}

	return true, nil

}

func SignUp(m models.AuthClaim) (bool, error) {
	var auth models.AuthRequestUser
	result := config.DB.Where("name = ? AND password = ?", m.UserName, m.Password).First(&auth)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			auth = models.NewAuthRequestUser(m.UserName, m.UserName, m.Password)

			resultSave := config.DB.Save(&auth)
			return true, resultSave.Error
		}
		return false, result.Error
	}

	return false, nil

}
