package usecase

import (
	"github.com/merembablas/catat-pendatang/models"
	"github.com/merembablas/catat-pendatang/people"
	"go.mongodb.org/mongo-driver/bson"
)

type peopleUsecase struct {
	peopleRepo people.Repository
}

// NewPeopleUsecase will create new an people usecase
func NewPeopleUsecase(peopleRepo people.Repository) people.Usecase {
	return &peopleUsecase{peopleRepo}
}

func (pu *peopleUsecase) Peoples(filter bson.M) ([]*models.People, error) {
	people, err := pu.peopleRepo.GetPeoples(filter)
	return people, err
}

func (pu *peopleUsecase) People(ID string) (*models.People, error) {
	return pu.peopleRepo.GetPeople(ID)
}

func (pu *peopleUsecase) CreatePeople(people models.People) (string, error) {
	peopleID, err := pu.peopleRepo.CreatePeople(people)
	return peopleID, err
}
