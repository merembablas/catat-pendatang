package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Address represent ddress model
type Address struct {
	ProvinceID string `bson:"province_id" json:"province_id"`
	RegencyID  string `bson:"regency_id" json:"regency_id"`
	DistrictID string `bson:"district_id" json:"district_id"`
	VillageID  string `bson:"village_id" json:"village_id"`
	RT         string `json:"RT"`
	RW         string `json:"RW"`
	Street     string `json:"street"`
}

// Haven represent haven model
type Haven struct {
	Name         string  `json:"name"`
	Phone        string  `json:"phone"`
	Address      Address `json:"address"`
	Relationship string  `json:"relationship"`
	RtName       string  `json:"rt_name"`
}

// PeopleProfile represent person model
type PeopleProfile struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	CaseHistory string `bson:"case_history" json:"case_history"`
}

// People represent people model
type People struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	IDstr       string             `json:"-"`
	Name        string             `json:"name"`
	IDNumber    string             `bson:"id_number" json:"id_number"`
	Address     string             `json:"address"`
	Phone       string             `json:"phone"`
	Origin      Address            `json:"origin"`
	Purpose     string             `json:"purpose"`
	PeopleCount int                `bson:"people_count" json:"people_count"`
	People      []PeopleProfile    `bson:"people" json:"people"`
	ArrivalDate time.Time          `bson:"arrival_date" json:"arrival_date"`
	Haven       Haven              `json:"haven"`
	SurveyorID  primitive.ObjectID `bson:"surveyor_id" json:"surveyor_id"`
	SubmittedAt time.Time          `bson:"submitted_at" json:"submitted_at"`
}

func (p *People) submitDate() string {
	return p.SubmittedAt.Format("02-01-2006")
}
