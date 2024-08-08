package oauth2

import (
	db "Astral/internal/App/database"
	md "Astral/internal/jwt/models"
	ws "Astral/internal/workspace/models"

	"github.com/google/uuid"
)

type GGUser struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
}

func CheckUser(user GGUser) error {
	var userDB md.User
	id := uuid.New().String()

	result := db.GlobalDB.Where("email = ?", user.Email).Find(&userDB)
	if result.RowsAffected == 0 {
		userDB = md.User{
			UserId:      id,
			Email:       user.Email,
			Name:        user.Name,
			DisplayName: user.DisplayName,
			WorkspaceID: 1,
			WorkSpace: ws.Workspace{
				Key:         id,
				Title:       "First workspace",
				Description: "Test",
				LogoColor:   "#ffffff",
			},
		}
		res := db.GlobalDB.Create(&userDB)
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}
