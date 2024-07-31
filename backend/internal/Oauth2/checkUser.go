package oauth2

import (
	db "Astral/internal/jwt/database"
	md "Astral/internal/jwt/models"
	"fmt"
)

type GGUser struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
}

func CheckUser(user GGUser) error {
	var userDB md.User
	fmt.Println(userDB)
	result := db.GlobalDB.Where("email = ?", user.Email).Find(&userDB)
	if result.RowsAffected == 0 {
		userDB = md.User{
			Email: user.Email,
			Name:  user.Name,
		}
		fmt.Println(userDB)
		res := db.GlobalDB.Create(&userDB)
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}
