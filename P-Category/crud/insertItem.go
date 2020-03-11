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

	categoryList :=[] model.Category{}

	categoryNames := model.GetCategoriesName()
	categoryIds := model.GetCategoriesIds()
	for i, categoryName := range categoryNames {

		categoryViewConfigs := []model.ViewConfig{}
		viewConfiguration := model.ViewConfig{ActionTitle: "Action Title", ViewIndex: 1, ViewTitle: "View Title", ViewType: "H_LIST"}
		categoryViewConfigs = append(categoryViewConfigs, viewConfiguration)

		category := model.Category{Name:categoryName, Id:categoryIds[i], ViewConfiguration: categoryViewConfigs}

		av, err := dynamodbattribute.MarshalMap(category)
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
		categoryList = append(categoryList, category)
	}


	insertedCategories, _ := json.Marshal(categoryList)

	fmt.Println("Successfully added '" + string(insertedCategories))
}
