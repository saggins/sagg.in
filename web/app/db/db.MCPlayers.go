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

func WhitelistScan() []model.Whitelist {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)
	input := &dynamodb.ScanInput{
		TableName: aws.String("whitelist"),
	}

	result, err := svc.Scan(input)
	if err != nil {
		panic(err.Error())
	}

	var Mainlist []model.Whitelist
	stuff := result.Items
	for _, i := range stuff {

		things := model.Whitelist{}

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

func PostNames(names model.Whitelist) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(names.Mcuser),
			},
			"uuid": {
				S: aws.String(names.Mcuuid),
			},
			"rname":{
				S:aws.String(names.Name),
			},
			"ip":{
				S:aws.String(names.Ip),
			},
		},
		TableName: aws.String("whitelist"),
	}

	_, err := svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
