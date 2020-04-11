package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merembablas/catat-pendatang/address"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// AddressHandler  represent the http handler for address
type AddressHandler struct {
	addressUcase address.Usecase
}

// NewAddressHandler will initialize the user resources endpoint
func NewAddressHandler(r *gin.Engine, au address.Usecase) {
	api := r.Group("/api")
	{
		api.GET("/provinces", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": au.Provinces()})
		})

		api.GET("/province/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": au.Province(c.Param("id"))})
		})
	}
}
