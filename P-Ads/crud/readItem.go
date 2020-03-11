package crud

import (
	"Publisher/P-Ads/constant"
	"Publisher/P-Ads/model"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)


func ReadItem() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	// snippet-end:[dynamodb.go.read_item.session]

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(constant.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String("400"),
			},

		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	adsItem := model.Ads{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &adsItem)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	adsItemJson, err := json.Marshal(adsItem)


	fmt.Println("Found item:")
	fmt.Println(string(adsItemJson))
}
// snippet-end:[dynamodb.go.read_item]
