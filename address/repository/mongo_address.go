package repository

import (
	"context"

	"github.com/merembablas/catat-pendatang/address"
	"github.com/merembablas/catat-pendatang/models"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type addressMongoRepository struct {
	client *mongo.Client
}

// NewMongoAddressRepository will create address repository object
func NewMongoAddressRepository(client *mongo.Client) address.Repository {
	return &addressMongoRepository{client}
}

func (addr *addressMongoRepository) GetProvinces() ([]*models.Province, error) {
	var provinces []*models.Province
	collection := addr.client.Database(viper.GetString("MONGO_DB")).Collection("provinces")
	cur, err := collection.Find(context.TODO(), bson.M{}, options.Find().SetProjection(bson.M{"regencies": 0}))

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var province models.Province
		err = cur.Decode(&province)
		if err != nil {
			return nil, err
		}
		provinces = append(provinces, &province)
	}

	return provinces, err
}

func (addr *addressMongoRepository) GetProvince(ID string) *models.Province {
	var province models.Province
	collection := addr.client.Database(viper.GetString("MONGO_DB")).Collection("provinces")
	collection.FindOne(context.TODO(), bson.M{"id": ID}).Decode(&province)

	return &province
}
