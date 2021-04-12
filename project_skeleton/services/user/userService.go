package user

import (
	"errors"
	"usercurd/models"

	"gorm.io/gorm"
)

type response struct {
	data string `json:"data"`
}

func CreateUser(u models.User) models.User {
	models.DB.Create(&u)
	return u
}

func GetUser(id int) (models.User, error) {
	var user models.User
	err := models.DB.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("err")
		}
		return user, err
	}
	return user, nil
}

func GetUsers() []models.User {
	var users []models.User
	models.DB.Find(&users)
	return users
}

func UpdateUser(user models.User, id int) (models.User, error) {
	user.ID = uint(id)
	var u models.User
	err := models.DB.First(&u, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return u, errors.New("err")
	} else {
		models.DB.Model(&u).Updates(user)
		return user, err
	}
}

func DeleteUser(id int) (string, error) {
	var user models.User
	err := models.DB.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "user not exists", errors.New("err")
		}
		return "internal server error", err
	}
	models.DB.Delete(&user, id)
	return "user deleted successfully ", err
}
