package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Register represent register model
type Register struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name       string             `json:"name"`
	Username   string             `json:"username"`
	Password   string             `json:"password"`
	Role       string             `json:"role"`
	Phone      string             `json:"phone"`
	Level      string             `json:"level"`
	ProvinceID string             `bson:"province_id" json:"province_id"`
	RegencyID  string             `bson:"regency_id" json:"regency_id"`
	DistrictID string             `bson:"district_id" json:"district_id"`
	VillageID  string             `bson:"village_id" json:"village_id"`
}
