package middleware

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Middlewares represent data struct for middlewares
type Middlewares struct {
}

// CORS will handle CORS
func (m *Middlewares) CORS(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

// APIAuth will handle authentification for API
func (m *Middlewares) APIAuth(c *gin.Context) {
	token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString(`SECRET`)), nil
	})

	if token != nil && err == nil {
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			c.Set("UserID", claims["Id"].(string))
			c.Next()
		}
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}

// WebAuth will handle authentification for web
func (m *Middlewares) WebAuth(c *gin.Context) {
	session := sessions.Default(c)
	var tokenString string
	if session.Get("token") != nil {
		tokenString = session.Get("token").(string)
	}

	if tokenString != "" {
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(viper.GetString(`SECRET`)), nil
		})

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			c.Set("UserID", claims["Id"].(string))
			c.Next()
		}
	} else {
		session.Set("error", "Silahkan login kembali")
		session.Save()
		c.Redirect(http.StatusFound, "/")
	}
}

// InitMiddleware initialize middleware
func InitMiddleware() *Middlewares {
	return &Middlewares{}
}
