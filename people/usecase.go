package people

import (
	"github.com/merembablas/catat-pendatang/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Usecase represent people
type Usecase interface {
	People(ID string) (*models.People, error)
	Peoples(filter bson.M) ([]*models.People, error)
	CreatePeople(people models.People) (string, error)
}
