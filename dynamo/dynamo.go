package dynamo

import (
	"github.com/mchirico/go-aws/client"
	db "github.com/mchirico/go-aws/dynamodb"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Dynamo struct {
	cfg   aws.Config
	table string
}

func NewDynamo(table string) *Dynamo {
	d := &Dynamo{
		cfg:   client.Config(),
		table: table,
	}
	return d
}

func Doc(location, aws string) *db.Doc {
	d := &db.Doc{}
	d.Location = location
	d.AWS = aws
	return d
}

func (d *Dynamo) Get(pkey, skey string) (*db.PKSK, error) {

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
		TableName:       &d.table,
		AttributesToGet: []string{"PK", "Doc", "SK", "Status"},
	}
	result, err := db.Get(d.cfg, input)
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

func (d *Dynamo) Put(pkey, skey, status string, doc *db.Doc) error {

	p := &db.PKSK{}
	p.PK = pkey
	p.SK = skey
	p.Status = status
	p.Doc = *doc

	av, err := attributevalue.MarshalMap(p)
	if err != nil {
		return err
	}
	_,err = db.Put(d.cfg, "mmcPKSK", av)
	if err != nil {
		return err
	}
	return nil
}
