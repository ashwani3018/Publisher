package crud

import (
	"Publisher/P-Article/constant"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
	"time"
)

type Article struct {
	ArticleID string `json:ArticleID`
	PublishTime int64 `json:PublishTime`
	OriginTime int64 `json:OriginTime`
	Title string `json:Title`
	Author           []string `json:Author`
	ArticleLink      string   `json:ArticleLink`
	ShortDescription string   `json:ShortDescription`
	Description      string   `json:Description`
	AudioLink        string   `json:AudioLink`
	VideoLink        string   `json:VideoLink`
	YoutubeId        string   `json:YoutubeId`
	ThumbLink        string   `json:ThumbLink`
	Location         string   `json:Location`
	Type             string   `json:Type`
	categoryId       string   `json:CategoryID`
	MediaID          string   `json:MediaID`
	IsFree           bool     `json:IsFree`
}

func InsertItem() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	//////////////////////////////////////////////////////////////////

	article := Article{}
	article.ArticleID = "AID_1"
	article.MediaID = "MID_1"
	article.PublishTime = time.Now().Unix() // Current time in Epoch Format
	article.OriginTime = time.Now().Unix()  // Current time in Epoch Format
	article.Title = "Article Title 1"
	article.ShortDescription = "Article Short Description 1"
	article.Description = "Article Description 1"
	article.ArticleLink = "Article Link"
	article.VideoLink = "Video Link"
	article.YoutubeId = "Youtube ID"
	article.categoryId = "Article Category ID"

	// Marshalling CategoryItem for Dynamo DB
	dynamoCategoryItem, err := dynamodbattribute.MarshalMap(article)
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
