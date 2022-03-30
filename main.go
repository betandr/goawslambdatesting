package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type dependencies struct {
	ddb   dynamodbiface.DynamoDBAPI
	table string
}

func (d *dependencies) handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if d.ddb == nil {
		sess := session.Must(session.NewSession())

		d.ddb = dynamodb.New(sess)
		d.table = "SomeTable"
	}

	return events.APIGatewayProxyResponse{
		Body:       "Hello, World!",
		StatusCode: 200,
	}, nil
}

func main() {
	deps := dependencies{}
	lambda.Start(deps.handler)
}
