// Author: xufei
// Date: 2019/2/26

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"gopkg.in/olivere/elastic.v6"
)

func main() {
	url := "http://weixinote.dev:9201"
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)),
	)
	if err != nil {
		panic(err)
	}

	ping, pingInt, _ := client.Ping(url).Do(context.Background())
	fmt.Println(ping, pingInt)
}
