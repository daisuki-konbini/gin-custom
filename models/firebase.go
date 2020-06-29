package models

import (
	"gin/config"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/api/option"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

var app *firebase.App
var authClient *auth.Client

func init() {
	var err error
	opt := option.WithCredentialsFile(config.GetConfig().GetString("firebase.credentials"))
	app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}
}

//GetAuthClient ...
func GetAuthClient(ctx *gin.Context) (*auth.Client, error) {
	var err error
	authClient, err = app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	return authClient, nil
}
