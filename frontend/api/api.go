// Author: xufei
// Date: 2018/12/19

package api

import (
	"log"
	"net/http"
)

func NewApi(host string) error {
	c := &searchController{}
	http.HandleFunc("/search", c.SearchHandle)

	log.Println("listening...")
	return http.ListenAndServe(host, nil)
}
