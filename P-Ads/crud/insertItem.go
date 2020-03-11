package crud

import (
	"Publisher/P-Ads/constant"
	"Publisher/P-Ads/model"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
	"strconv"
)

func InsertItem() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	//////////////////////////////////////////////////////////////////

	// Each Category Logic
	categories := []model.Category{}
	for j := 0; j<3; j++ {
		// Category Ads Array
		catAds := []model.AdsItem{}
		// Category Size
		categorySize := 5
		// Each Category All Ads Initializing
		for i := 0; i < categorySize; i++ {
			categoryAds := model.AdsItem{Position: i, AdsID: "AdsID "+strconv.Itoa(j)+strconv.Itoa(i),
				Type: constant.ADS_TYPE_300x250}
			catAds = append(catAds, categoryAds)
		}
		// Each Sub-Category Init
		category := model.Category{ID : "categoryId "+strconv.Itoa(j), CategoryAds:catAds}
		categories = append(categories, category)
	}

	// Each Sub-Category Ads Logic
	subCategories := []model.SubCategory{}
	for j := 0; j<3; j++ {
		// Category Ads Array
		subcatAds := []model.AdsItem{}
		// Category Size
		categorySize := 5
		// Each Category All Ads Initializing
		for i := 0; i < categorySize; i++ {
			subCategoryAds := model.AdsItem{Position: i, AdsID: "AdsID "+strconv.Itoa(j)+strconv.Itoa(i), Type: constant.ADS_TYPE_300x250}
			subcatAds = append(subcatAds, subCategoryAds)
		}

		// Each Sub-Category Init
		subCategory := model.SubCategory{ID : "subCategoryId "+strconv.Itoa(j), CategoryAds:subcatAds}
		subCategories = append(subCategories, subCategory)
	}

	bottomBanner := model.BottomBanner{AdsID: "Bottom Banner Ads Id", TotalRequest:20}
	fullScreen := model.FullScreen{AdsID: "FullScreen Ads Id", TotalRequest:20, ShowInDetail : true, ShowInListing:true}

	// Final Ads Model
	ads := model.Ads{ID: 400}

	ads.Category = categories
	ads.SubCategory = subCategories
	ads.BottomBanner = bottomBanner
	ads.FullScreen = fullScreen

	av, err := dynamodbattribute.MarshalMap(ads)
	if err != nil {
		fmt.Println("Got error marshalling new movie item:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(constant.TableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}


	fmt.Println("Successfully added '" + constant.TableName)
}
