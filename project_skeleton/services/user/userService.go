package user

import (
	"fmt"
	"usercurd/models"
)

func CreateUser(u models.User) models.User {
	models.DB.Create(&u)
	return u
}

func GetUser(id int) models.User {

	var user models.User
	models.DB.First(&user, id)
	return user
}

func GetUsers() []models.User {
	var users []models.User
	models.DB.Find(&users)
	return users
}

func UpdateUser(u models.User, id int) models.User {

	fmt.Print("44444444444444 in updated  u", u)
	fmt.Print("44444444444444 u", id)
	u.ID = uint(id)
	

	models.DB.Save(&u)
	return u
}

func DeleteUser(id int) string {
	var user models.User
	models.DB.Delete(&user, id)
	return "User deleted successfully "
}
