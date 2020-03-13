package crud

import (
	"Publisher/P-Media/constant"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
)

type GalleryStruct struct {
	ImgUrl    string `json:ImgUrl`
	VideoUrl  string `json:VideoUrl`
	YoutubeId string `json:YoutubeId`
	Caption   string `json:Caption`
}

type Media struct {
	MediaID string `json:MediaID`
	ArticleID string `json:ArticleID`
	PublishTime int64 `json:PublishTime`
	Gallery []GalleryStruct `json:Gallery`
}

func InsertItem() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	//////////////////////////////////////////////////////////////////

	media := Media{}
	media.MediaID = "MID_106"
	media.ArticleID = "AID_105"
	//media.PublishTime = time.Now().Unix() // Current time in Epoch Format

	// Creating Gallery
	galleryData := []GalleryStruct{}
	galleryStruct := GalleryStruct{}
	galleryStruct.Caption = "Media Caption"
	galleryStruct.ImgUrl = "Media Img Url"
	galleryStruct.VideoUrl = "Media Video Url"
	galleryStruct.YoutubeId = "YOu tube Video Id"
	galleryData = append(galleryData, galleryStruct)

	media.Gallery = galleryData

	// Marshalling CategoryItem for Dynamo DB
	dynamoCategoryItem, err := dynamodbattribute.MarshalMap(media)
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
