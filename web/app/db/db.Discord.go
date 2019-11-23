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

func DiscordScan() []model.Player {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)
	input := &dynamodb.ScanInput{
		TableName: aws.String("discordpy"),
	}

	result, err := svc.Scan(input)
	if err != nil {
		panic(err.Error())
	}

	var Mainlist []model.Player
	stuff := result.Items
	for _, i := range stuff {

		things := model.Player{}

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

func PostPlayers(names model.Player) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(names.Name),
			},
			"userid": {
				S: aws.String(names.Userid),
			},
			"ip":{
				S:aws.String(names.Ip),
			},
		},
		TableName: aws.String("discordpy"),
	}

	_, err := svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

}