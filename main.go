package main

import (
	"context"
	"errors"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func HandleRequest(ctx context.Context, event events.APIGatewayCustomAuthorizerRequestTypeRequest) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {

	// Initilize the Firebase App
	opt := option.WithCredentialsFile("firebase-credential.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logrus.Fatalln("error initializing app: %v\n", err)
	}

	authResponse := events.APIGatewayV2CustomAuthorizerSimpleResponse{
		IsAuthorized: false,
	}

	authClient, err := app.Auth(ctx)

	if err != nil {
		logrus.Fatalln("error initializing app: %v\n", err)
	}

	token := event.Headers["authorization"]

	splittedToken := strings.Split(token, " ")

	if len(splittedToken) != 2 {
		logrus.Info("Auth Key not correctly structed.")
		return authResponse, errors.New("Unauthorized")
	}

	_, err = authClient.VerifyIDTokenAndCheckRevoked(ctx, splittedToken[1])
	if err != nil {
		logrus.Info("Vlidation Failed")
		return authResponse, errors.New("Unauthorized")
	}

	authResponse.IsAuthorized = true
	return authResponse, nil
}

func main() {
	lambda.Start(HandleRequest)
}
