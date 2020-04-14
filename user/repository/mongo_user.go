package repository

import (
	"context"
	"errors"

	"github.com/merembablas/catat-pendatang/models"
	"github.com/merembablas/catat-pendatang/user"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type userMongoRepository struct {
	client *mongo.Client
}

// NewMongoUserRepository will create an object that represent the article.Repository interface
func NewMongoUserRepository(client *mongo.Client) user.Repository {
	return &userMongoRepository{client}
}

func (usr *userMongoRepository) GetUsers() ([]*models.User, error) {
	var users []*models.User
	collection := usr.client.Database(viper.GetString("MONGO_DB")).Collection("users")
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var user models.User
		err = cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, err
}

func (usr *userMongoRepository) GetUser(ID string) (*models.User, error) {
	objID, _ := primitive.ObjectIDFromHex(ID)

	var user models.User
	collection := usr.client.Database(viper.GetString("MONGO_DB")).Collection("users")
	err := collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)

	return &user, err
}

func (usr *userMongoRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	collection := usr.client.Database(viper.GetString("MONGO_DB")).Collection("users")
	err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)

	return &user, err
}

func (usr *userMongoRepository) CreateUser(user models.Register) (string, error) {
	var err error
	user.ID = primitive.NewObjectID()
	user.Password, err = hashAndSalt([]byte(user.Password))

	if err != nil {
		return "", errors.New("Cannot create encrypted password")
	}

	collection := usr.client.Database(viper.GetString("MONGO_DB")).Collection("users")
	result, err := collection.InsertOne(context.TODO(), user)

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), err
	}

	return "", err
}

func hashAndSalt(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), err
}
