package usecase

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/merembablas/catat-pendatang/models"
	"github.com/merembablas/catat-pendatang/user"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo user.Repository
}

// NewUserUsecase will create new an user usecase
func NewUserUsecase(u user.Repository) user.Usecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (usr *userUsecase) Login(login models.Login) (string, error) {
	user, err := usr.userRepo.GetUserByUsername(login.Username)
	if err != nil {
		return "", errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return "", errors.New("Wrong password")
	}

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"Id":  user.ID,
		"exp": time.Now().Add(time.Hour * 2880).Unix(),
	}

	tokenString, err := token.SignedString([]byte(viper.GetString("SECRET")))

	return tokenString, err
}

func (usr *userUsecase) Register(reg models.Register) (string, error) {
	_, err := usr.userRepo.GetUserByUsername(reg.Username)
	if err == nil {
		return "", errors.New("Username taken")
	}

	userID, err := usr.userRepo.CreateUser(reg)
	if err != nil {
		return userID, err
	}

	return userID, err
}

func (usr *userUsecase) Users() ([]*models.User, error) {
	users, err := usr.userRepo.GetUsers()
	return users, err
}

func (usr *userUsecase) User(ID string) (*models.User, error) {
	return usr.userRepo.GetUser(ID)
}
