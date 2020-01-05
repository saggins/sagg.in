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

var svc *dynamodb.DynamoDB
var tableName string
//BLOG

func init() {
	//Creates aws session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
	}))
	svc = dynamodb.New(sess)

	tableName = "saggweb"
}

//AUTH

func StoreSession() {
	
}

func GetSession()  {
	
}

// General Blog stuff

func GetBlobByID(id int) model.Item {
	newid := string(id)
	return GetRaws(newid)
}

func GetRaws(id string) model.Item {

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

	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
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


//DISCORD

func DiscordScan() []model.Player {

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

//WHITELIST

func PostPlayers(names model.Player) {


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
func WhitelistScan() []model.Whitelist {

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

// MC SHOP

func ShopScan() []model.MCShop {

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



