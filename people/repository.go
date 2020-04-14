package people

import (
	"github.com/merembablas/catat-pendatang/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Repository represent people
type Repository interface {
	GetPeoples(filter bson.M) ([]*models.People, error)
	GetPeople(ID string) (*models.People, error)
	CreatePeople(people models.People) (string, error)
}
