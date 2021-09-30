// Author: xufei
// Date: 2020/6/2

package main

import (
	"context"
	"log"
)

func main() {

	c := BuildClient([]string{"http://weixinote.dev:2379"})

	// set "/foo" key with "bar" value
	log.Print("Setting '/foo' key with 'bar' value")
	resp, err := c.Put(context.Background(), "/foo", "bar", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Set is done. Metadata is %q\n", resp)
	}

	// get "/foo" key's value
	log.Print("Getting '/foo' key value")
	resp2, err := c.Get(context.Background(), "/foo", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Get is done. Metadata is %q\n", resp)
		// print value
		log.Printf("%+v\n", resp2.Kvs)
	}
}
