package main

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event events.APIGatewayCustomAuthorizerRequestTypeRequest) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {

	authResponse := events.APIGatewayV2CustomAuthorizerSimpleResponse{
		IsAuthorized: false,
	}

	fmt.Println(event.Headers["br authorization"])

	token := event.Headers["authorization"]

	fmt.Println(token)

	splittedToken := strings.Split(token, " ")

	fmt.Println("")

	switch strings.ToLower(splittedToken[1]) {
	case "allow":
		authResponse.IsAuthorized = true
		authResponse.Context = map[string]interface{}{
			"userId": "Test",
		}
		return authResponse, nil
	case "deny":
		return authResponse, errors.New("Unauthorized")
	case "unauthorized":
		return authResponse, errors.New("Unauthorized")
	default:
		return authResponse, errors.New("Unauthorized")
	}
}

func main() {
	lambda.Start(HandleRequest)
}
