package main

import (
	"context"
	"log"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	addressHttpHandler "github.com/merembablas/catat-pendatang/address/delivery/http"
	addressRepository "github.com/merembablas/catat-pendatang/address/repository"
	addressUsecase "github.com/merembablas/catat-pendatang/address/usecase"
	"github.com/merembablas/catat-pendatang/middleware"
	userHttpHandler "github.com/merembablas/catat-pendatang/user/delivery/http"
	userRepository "github.com/merembablas/catat-pendatang/user/repository"
	userUsecase "github.com/merembablas/catat-pendatang/user/usecase"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
}

func main() {
	mConn := getMongoClient()
	ur := userRepository.NewMongoUserRepository(mConn)
	ar := addressRepository.NewMongoAddressRepository(mConn)

	uu := userUsecase.NewUserUsecase(ur)
	au := addressUsecase.NewAddressUsecase(ar)

	r := gin.Default()
	middl := middleware.InitMiddleware()
	store := cookie.NewStore([]byte(viper.GetString("SECRET")))

	r.Use(middl.CORS)
	r.Static("/assets", viper.GetString("STATIC_PATH"))
	r.HTMLRender = loadTemplates(viper.GetString("TEMPLATE_PATH"))
	r.Use(sessions.Sessions("data", store))

	userHttpHandler.NewUserHandler(r, uu, middl)
	addressHttpHandler.NewAddressHandler(r, au)

	r.Run(":8989")

}

func getMongoClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(viper.GetString("MONGO_URI"))
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	publicLayouts, err := filepath.Glob(templatesDir + "/layouts/public-base.html")
	if err != nil {
		panic(err.Error())
	}

	publics, err := filepath.Glob(templatesDir + "/public/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our publicLayouts/ and public/ directories
	for _, public := range publics {
		layoutCopy := make([]string, len(publicLayouts))
		copy(layoutCopy, publicLayouts)
		files := append(layoutCopy, public)
		r.AddFromFiles(filepath.Base(public), files...)
	}

	dashboardLayouts, err := filepath.Glob(templatesDir + "/layouts/dashboard-base.html")
	if err != nil {
		panic(err.Error())
	}

	admins, err := filepath.Glob(templatesDir + "/dashboard/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our dashboardLayouts/ and dashboard/ directories
	for _, admin := range admins {
		layoutCopy := make([]string, len(dashboardLayouts))
		copy(layoutCopy, dashboardLayouts)
		files := append(layoutCopy, admin)
		r.AddFromFiles(filepath.Base(admin), files...)
	}
	return r
}
