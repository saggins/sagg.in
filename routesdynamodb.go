package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)
//Item yes
type Item struct {
	title string
	blobs []string
	id    int
}

func getRaws(id string) Item {
	//Creates aws session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)

	//Get Item

	//Parameters
	tableName := "websites"

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(id),
			},
		},
	})

	if err != nil {

		panic(err.Error())
	}

	item := Item{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	return item
}

func getAllPages() []Item{

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)

	var TextBlob []Item

	err := svc.ScanPages(&dynamodb.ScanInput{
		TableName: aws.String("websites"),
	}, func(page *dynamodb.ScanOutput, last bool) bool {
		blob := []Item{}
	
		err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &blob)
		if err != nil {
			 panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
		}
	
		TextBlob = append(TextBlob, blob...)
	
		return true // keep paging
	})
	if err == nil{ //Keep Close Eye
		panic(fmt.Sprintf("error"))
	}

	return TextBlob

}
