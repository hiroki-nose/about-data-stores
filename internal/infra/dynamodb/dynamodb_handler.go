package dynamodb

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamodbHandler struct {
	DbService *dynamodb.DynamoDB
}

func NewDynamodbHandler() DynamodbHandler {
	dbSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return DynamodbHandler{
		DbService: dynamodb.New(dbSession),
	}
}
func (dbc DynamodbHandler) Get(tableName string, keyName string, hashkey string) (map[string]*dynamodb.AttributeValue, error) {
	fmt.Println("Hashkey: ", hashkey)
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			keyName: {
				S: aws.String(hashkey),
			},
		},
	}

	response, err := dbc.DbService.GetItem(input)
	if err != nil {
		fmt.Println("Got error calling GetItem:")
		fmt.Println(err.Error())
	}
	return response.Item, err
}
