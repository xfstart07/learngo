// Author: xufei
// Date: 2020/4/24

package main

import (
	"fmt"

	"github.com/jrallison/go-workers"
)

func myJob(message *workers.Msg) {
	fmt.Println(message)
	// do something with your message
	// message.Jid()
	// message.Args() is a wrapper around go-simplejson (http://godoc.org/github.com/bitly/go-simplejson)
}

func main() {
	workers.Configure(map[string]string{
		// location of redis instance
		"server": "192.168.4.71:6379",
		// instance of the database
		"database": "0",
		// number of connections to keep open with redis
		"pool": "3",
		// unique process id for this instance of workers (for proper recovery of inprogress jobs on crash)
		"process": "1",
	})

	// pull messages from "myqueue" with concurrency of 10
	workers.Process("myqueue", myJob, 10)

	workers.Start()

}
