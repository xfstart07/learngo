// Author: xufei
// Date: 2018/12/19

package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type searchController struct{}

func (c *searchController) SearchHandle(w http.ResponseWriter, r *http.Request) {
	log.Printf("search Handler")

	q := r.URL.Query().Get("q")
	pageQuery := r.URL.Query().Get("page")
	page, _ := strconv.Atoi(pageQuery)

	hits, err := service.SearchService.Profile(q, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	json.NewEncoder(w).Encode(hits)
}
