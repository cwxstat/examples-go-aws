package main

import (
	"fmt"
	"github.com/mchirico/go-aws/client"
	db "github.com/mchirico/go-aws/dynamodb"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func Get(table string, pkey, skey string) (*db.PKSK, error) {

	type KEY struct {
		PK string `json:"PK"`
		SK string `json:"SK"`
	}

	key, _ := attributevalue.MarshalMap(&KEY{
		PK: pkey,
		SK: skey,
	})

	input := &dynamodb.GetItemInput{
		Key:             key,
		TableName:       &table,
		AttributesToGet: []string{"PK", "Doc", "SK", "Status"},
	}
	result, err := db.Get(client.Config(), input)
	if err != nil {
		return nil, err
	}
	p := &db.PKSK{}
	err = attributevalue.UnmarshalMap(result.Item, p)
	if err != nil {
		return nil, err
	}
	return p, nil

}

func main() {
	result, err := Get("mmcPKSK", "My Data", "Something")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result.Doc)
	fmt.Println(result.PK)
	fmt.Println(result.SK)
	fmt.Println(result.Status)

}
