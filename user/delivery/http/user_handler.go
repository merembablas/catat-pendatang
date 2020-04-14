package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merembablas/catat-pendatang/middleware"
	"github.com/merembablas/catat-pendatang/models"
	"github.com/merembablas/catat-pendatang/user"
)

// UserHandler  represent the http handler for user
type UserHandler struct {
}

// NewUserHandler will initialize the user resources endpoint
func NewUserHandler(r *gin.Engine, us user.Usecase, middl *middleware.Middlewares) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"status": "Halo",
		})
	})

	api := r.Group("/api")
	{
		api.POST("/login", func(c *gin.Context) {
			var login models.Login
			c.BindJSON(&login)
			token, err := us.Login(login)

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"data": token})
			}
		})

		api.POST("/user", func(c *gin.Context) {
			var user models.Register
			c.BindJSON(&user)
			UserID, err := us.Register(user)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"data": UserID})
			}
		})

		api.GET("/users", middl.APIAuth, func(c *gin.Context) {
			users, err := us.Users()

			if err != nil {
				users = []*models.User{}
			}

			c.JSON(http.StatusOK, gin.H{"data": users})
		})

		api.GET("/user/:id", middl.APIAuth, func(c *gin.Context) {
			user, err := us.User(c.Param("id"))

			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"data": nil})
			} else {
				c.JSON(http.StatusOK, gin.H{"data": user})
			}
		})

		api.GET("/me", middl.APIAuth, func(c *gin.Context) {
			user, err := us.User(c.MustGet("UserID").(string))

			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"data": nil})
			} else {
				c.JSON(http.StatusOK, gin.H{"data": user})
			}
		})
	}
}
