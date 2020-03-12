package crud

import (
	"Publisher/P-Category/constant"
	"Publisher/P-Category/model"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
)

func InsertItem() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	//////////////////////////////////////////////////////////////////

	// Category Items Slice
	categoryItemList :=[] model.CategoryItem{}

	// Category Mock Name Slice
	categoryItemNames := model.GetCategoriesName()

	// Category Mock Id Slice
	categoryItemIds := model.GetCategoriesIds()

	// For Loop to add Category data with fields
	for i, categoryName := range categoryItemNames {

		// View Configuration for SubCategory of CategoryItem
		inCategoryItemViewConfig := model.ViewConfig{}
		inCategoryItemViewConfig.ViewTitle = "View Title"
		inCategoryItemViewConfig.ActionTitle = "Action Title"
		inCategoryItemViewConfig.ViewIndex = 0
		inCategoryItemViewConfig.ViewType = constant.VT_H_PAGER
		inCategoryItemViewConfig.Redirection = "On SubSection Tab-Slider"

		// SubCategory for SubCategory Slice
		subCategoryItem := model.DefaultSubCategoryItem();
		subCategoryItem.InCategoryItemViewConfig = inCategoryItemViewConfig

		// Creating SubCategory Slice for CategoryItem
		subCategoryList := []model.SubCategoryItem{}
		subCategoryList = append(subCategoryList, subCategoryItem)

		// Creating Category Item
		categoryItem := model.DefaultCategoryItem()
		categoryItem.Name = categoryName
		categoryItem.Id = categoryItemIds[i]
		categoryItem.SubCategory = subCategoryList

		nonGroupedItemViewConfigSlice := []model.ViewConfig{}
		nonGroupedItemViewConfig := model.DefaultViewConfig()
		nonGroupedItemViewConfig.ViewTitle = "Item View Title"
		nonGroupedItemViewConfig.ActionTitle = "Item Action Title"
		nonGroupedItemViewConfig.ViewIndex = 0
		nonGroupedItemViewConfig.ViewType = constant.VT_ITEM
		nonGroupedItemViewConfig.Redirection = "On Detail Page"
		nonGroupedItemViewConfigSlice = append(nonGroupedItemViewConfigSlice, nonGroupedItemViewConfig)

		categoryItem.NonGroupedItemViewConfig = nonGroupedItemViewConfigSlice


		// Marshalling CategoryItem for Dynamo DB
		dynamoCategoryItem, err := dynamodbattribute.MarshalMap(categoryItem)
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
		categoryItemList = append(categoryItemList, categoryItem)
	}


	insertedCategories, _ := json.Marshal(categoryItemList)

	fmt.Println("Successfully added '" + string(insertedCategories))
}
