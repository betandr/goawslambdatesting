package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type mockedPutItem struct {
	dynamodbiface.DynamoDBAPI
	Response dynamodb.PutItemOutput
}

func (d mockedPutItem) PutItem(in *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return &d.Response, nil
}

func TestLambdaHandler(t *testing.T) {
	t.Run("Successful Request", func(t *testing.T) {
		mock := mockedPutItem{
			Response: dynamodb.PutItemOutput{},
		}

		d := dependencies{
			ddb:   mock,
			table: "test_table",
		}

		_, err := d.handler((events.APIGatewayProxyRequest{}))
		if err != nil {
			t.Fatal(err)
		}
	})
}
