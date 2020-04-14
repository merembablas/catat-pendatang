package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merembablas/catat-pendatang/middleware"
	"github.com/merembablas/catat-pendatang/models"
	"github.com/merembablas/catat-pendatang/people"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PeopleHandler  represent the http handler for people
type PeopleHandler struct {
}

// NewPeopleHandler will initialize the people resources endpoint
func NewPeopleHandler(r *gin.Engine, pu people.Usecase, middl *middleware.Middlewares) {
	api := r.Group("/api")
	{
		api.POST("/people", func(c *gin.Context) {
			var people models.People
			c.BindJSON(&people)
			peopleID, err := pu.CreatePeople(people)

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"data": peopleID})
			}
		})

		api.GET("/people", middl.APIAuth, func(c *gin.Context) {
			name := c.Query("name")
			fromProvinceID := c.Query("from_province_id")
			fromRegencyID := c.Query("from_regency_id")
			fromDistrictID := c.Query("from_district_id")
			fromVillageID := c.Query("from_village_id")
			provinceID := c.Query("province_id")
			regencyID := c.Query("regency_id")
			districtID := c.Query("district_id")
			villageID := c.Query("village_id")
			havenName := c.Query("haven_name")
			surveyorID := c.Query("surveyor_id")
			arrivalDateFrom := c.Query("arrival_date_from")
			arrivalDateTo := c.Query("arrival_date_to")

			filter := bson.M{}
			if name != "" {
				filter["name"] = bson.M{"$regex": primitive.Regex{Pattern: name, Options: "i"}}
			}
			if fromProvinceID != "" {
				filter["origin.province_id"] = fromProvinceID
			}
			if fromRegencyID != "" {
				filter["origin.regency_id"] = fromRegencyID
			}
			if fromDistrictID != "" {
				filter["origin.district_id"] = fromDistrictID
			}
			if fromVillageID != "" {
				filter["origin.village_id"] = fromVillageID
			}
			if provinceID != "" {
				filter["haven.address.province_id"] = provinceID
			}
			if regencyID != "" {
				filter["haven.address.regency_id"] = regencyID
			}
			if districtID != "" {
				filter["haven.address.district_id"] = districtID
			}
			if villageID != "" {
				filter["haven.address.village_id"] = villageID
			}
			if havenName != "" {
				filter["haven.name"] = bson.M{"$regex": primitive.Regex{Pattern: havenName, Options: "i"}}
			}
			if surveyorID != "" {
				filter["surveyor_id"], _ = primitive.ObjectIDFromHex(surveyorID)
			}
			if arrivalDateFrom != "" && arrivalDateTo != "" {
				filter["arrival_date"] = bson.M{"$gte": arrivalDateFrom, "$lte": arrivalDateTo}
			}

			people, err := pu.Peoples(filter)

			if err != nil {
				people = []*models.People{}
			}

			c.JSON(http.StatusOK, gin.H{"data": people})
		})

		api.GET("/people/:id", middl.APIAuth, func(c *gin.Context) {
			people, err := pu.People(c.Param("id"))

			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"data": nil})
			} else {
				c.JSON(http.StatusOK, gin.H{"data": people})
			}
		})
	}
}
