package user

import "github.com/merembablas/catat-pendatang/models"

// Repository represent the user's repository contract
type Repository interface {
	GetUsers() ([]*models.User, error)
	GetUser(ID string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	CreateUser(user models.Register) (string, error)
}
