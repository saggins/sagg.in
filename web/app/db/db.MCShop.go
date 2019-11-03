package db

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	model "github.com/win32prog/sagg.in/web/app/models"
)

func ShopScan() []model.MCShop {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)
	input := &dynamodb.ScanInput{
		TableName: aws.String("shops"),
	}

	result, err := svc.Scan(input)
	if err != nil {
		panic(err.Error())
	}

	var Mainlist []model.MCShop
	stuff := result.Items
	for _, i := range stuff {

		things := model.MCShop{}

		err = dynamodbattribute.UnmarshalMap(i, &things)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		Mainlist = append(Mainlist, things)

	}

	return Mainlist

}

func PostShops(Shops model.MCShop) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(Shops.Name),
			},
			"item": {
				S: aws.String(Shops.Item),
			},
			"price":{
				S: aws.String(Shops.Price),
			},
		},
		TableName: aws.String("shops"),
	}

	_, err := svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
func DelShops(Shops model.MCShop) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	input := &dynamodb.DeleteItemInput {
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(Shops.Name),
			},
			"item": {
				S: aws.String(Shops.Item),
			},
		},
		TableName: aws.String("shops"),
	}

	_, err := svc.DeleteItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
