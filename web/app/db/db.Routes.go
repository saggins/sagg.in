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

//Item yes

func GetBlobByID(id int) model.Item {
	newid := string(id)
	return GetRaws(newid)
}

func GetRaws(id string) model.Item {

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

	item := model.Item{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	return item
}

func GetAllPages() []model.Item {

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

	var TextBlob []model.Item

	for _, i := range result.Items {

		things := model.Item{}
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
