package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"log"
	"os"
	"strings"
)

type Response struct {
	Message string `json:"message"`
}

type EsaRequestBody struct {
	Kind string `json:"kind"`
	User struct {
		Name       string `json:"name"`
		ScreenName string `json:"screen_name"`
		Icon       struct {
			URL    string `json:"url"`
			ThumbS struct {
				URL string `json:"url"`
			} `json:"thumb_s"`
			ThumbMs struct {
				URL string `json:"url"`
			} `json:"thumb_ms"`
			ThumbM struct {
				URL string `json:"url"`
			} `json:"thumb_m"`
			Emoji struct {
				URL string `json:"url"`
			} `json:"emoji"`
			ThumbL struct {
				URL string `json:"url"`
			} `json:"thumb_l"`
			ThumbEsa struct {
				URL string `json:"url"`
			} `json:"thumb_esa"`
		} `json:"icon"`
	} `json:"user"`
	Team struct {
		Name string `json:"name"`
	} `json:"team"`
	Post struct {
		BodyMd   string `json:"body_md"`
		BodyHTML string `json:"body_html"`
		Message  string `json:"message"`
		Wip      bool   `json:"wip"`
		Number   int    `json:"number"`
		Name     string `json:"name"`
		URL      string `json:"url"`
		DiffURL  string `json:"diff_url"`
	} `json:"post"`
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	lc, ok := lambdacontext.FromContext(ctx)
	if !ok {
		return Response{Message: ""}, errors.New("Cannot find LambdaContext")
	}
	fmt.Println(request.Body)
	var esaRequest EsaRequestBody
	json.Unmarshal([]byte(request.Body), &esaRequest)
	fmt.Println(esaRequest.Post.Name)
	fmt.Println(strings.HasPrefix(esaRequest.Post.Name, os.Getenv("ESA_TARGET_PATH")))

	fmt.Println(lc.ClientContext.Custom)
	log.Printf("log:InvokedFunctionArn = %s\n", lc.InvokedFunctionArn)
	fmt.Printf("fmt:InvokedFunctionArn = %s\n", lc.InvokedFunctionArn)
	log.Printf("\"%s\" executes on \"%s\".\n", os.Getenv("LAMBDA_TASK_ROOT"), os.Getenv("AWS_EXECUTION_ENV"))

	return Response{
		Message: "function finished",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
