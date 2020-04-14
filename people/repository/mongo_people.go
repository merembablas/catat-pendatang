package repository

import (
	"context"
	"log"

	"github.com/merembablas/catat-pendatang/models"
	"github.com/merembablas/catat-pendatang/people"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type peopleMongoRepository struct {
	client *mongo.Client
}

// NewMongoPeopleRepository will create object that represent people repository
func NewMongoPeopleRepository(client *mongo.Client) people.Repository {
	return &peopleMongoRepository{client}
}

func (pr *peopleMongoRepository) GetPeoples(filter bson.M) ([]*models.People, error) {
	var peoples []*models.People
	collection := pr.client.Database(viper.GetString("MONGO_DB")).Collection("people")
	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var people models.People
		err = cur.Decode(&people)
		if err != nil {
			continue
		}
		peoples = append(peoples, &people)
	}

	return peoples, err
}

func (pr *peopleMongoRepository) GetPeople(ID string) (*models.People, error) {
	objID, _ := primitive.ObjectIDFromHex(ID)

	var people models.People
	collection := pr.client.Database(viper.GetString("MONGO_DB")).Collection("people")
	err := collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&people)

	return &people, err
}

func (pr *peopleMongoRepository) CreatePeople(people models.People) (string, error) {
	collection := pr.client.Database(viper.GetString("MONGO_DB")).Collection("people")
	result, err := collection.InsertOne(context.TODO(), people)

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), err
	}

	return "", err
}
