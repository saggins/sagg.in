package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

//Item yes
type Item struct {
	Blobs      []string `json:"blobs"`
	Title      string   `json:"title"`
	ID         string   `json:"id"`
	BlobsTitle []string `json:"blobstitle"`
}

func getRaws(id string) Item {
	//Creates aws session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)

	//Get Item

	//Parameters
	tableName := "saggweb"

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(tableName),
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

func getAllPages() []Item {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)
	input := &dynamodb.ScanInput{
		TableName: aws.String("saggweb"),
	}

	result, err := svc.Scan(input)
	if err != nil {
		panic(err.Error())
	}

	var TextBlob []Item

	for _, i := range result.Items {

		things := Item{}
		err = dynamodbattribute.UnmarshalMap(i, &things)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		TextBlob = append(TextBlob, things)

	}

	//limit := len(result.Items)

	//for i:=0;i<=limit;i++{

	//}

	return TextBlob

}
