package crud

import (
	"Publisher/P-SimilarNews/constant"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
	"time"
)

type SimilarNewsStruct struct {
	ArticleID string `json:ArticleID`
	PublishTime int64 `json:PublishTime`
	SimilarNews []string `json:SimilarNews`
}

func InsertItem() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	//////////////////////////////////////////////////////////////////

	similarNewsStruct := SimilarNewsStruct{}
	similarNewsStruct.ArticleID = "AID_106"
	similarNewsStruct.PublishTime = time.Now().Unix() // Current time in Epoch Format
	similarNewsStruct.SimilarNews = append(similarNewsStruct.SimilarNews, "AID_163")
	similarNewsStruct.SimilarNews = append(similarNewsStruct.SimilarNews, "AID_162")
	similarNewsStruct.SimilarNews = append(similarNewsStruct.SimilarNews, "AID_169")
	similarNewsStruct.SimilarNews = append(similarNewsStruct.SimilarNews, "AID_164")

	// Marshalling CategoryItem for Dynamo DB
	dynamoCategoryItem, err := dynamodbattribute.MarshalMap(similarNewsStruct)
	if err != nil {
		fmt.Println("Got error marshalling new movie item:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Creating final structure of CategoryItem to insert into Dynamo DB
	input := &dynamodb.PutItemInput{
		Item:      dynamoCategoryItem,
		TableName: aws.String(constant.TableName),
	}

	// Finally Inserting CategoryItem in Dynamo DB
	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully added '")
}
