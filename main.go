package main

import (
	"fmt"
	"github.com/cwxstat/examples-go-aws/dynamo"
)

func main() {
	d := dynamo.NewDynamo("mmcPKSK")

	d.Put("pk0", "skey", "good", dynamo.Doc("location", "aws"))
	result, err := d.Get("My Data", "Something")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result.Doc)
	fmt.Println(result.PK)
	fmt.Println(result.SK)
	fmt.Println(result.Status)

}
