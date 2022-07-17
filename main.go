package main

import (
	"fmt"
	"github.com/cwxstat/examples-go-aws/dynamo"
)

func main() {
	result, err := dynamo.Get("mmcPKSK", "My Data", "Something")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result.Doc)
	fmt.Println(result.PK)
	fmt.Println(result.SK)
	fmt.Println(result.Status)

}
