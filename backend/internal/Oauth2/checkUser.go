package oauth2

import (
	db "Astral/internal/jwt/database"
	md "Astral/internal/jwt/models"
	"fmt"

	"github.com/google/uuid"
)

type GGUser struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
}

func CheckUser(user GGUser) error {
	var userDB md.User
	fmt.Println(userDB)
	result := db.GlobalDB.Where("email = ?", user.Email).Find(&userDB)
	if result.RowsAffected == 0 {
		userDB = md.User{
			UserId: uuid.New().String(),
			Email:   user.Email,
			Name:    user.Name,
		}
		fmt.Println(userDB)
		res := db.GlobalDB.Create(&userDB)
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}
