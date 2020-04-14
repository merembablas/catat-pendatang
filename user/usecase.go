package user

import (
	"github.com/merembablas/catat-pendatang/models"
)

// Usecase represent the user
type Usecase interface {
	Login(login models.Login) (string, error)
	Register(user models.Register) (string, error)
	Users() ([]*models.User, error)
	User(ID string) (*models.User, error)
}
