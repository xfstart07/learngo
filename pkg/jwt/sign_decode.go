// Author: Xu Fei
// Date: 2018/7/30
package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	str := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJub25jZSI6ImRjMzc4ODYyMTI5MjRlODE5MGFjZDQ5OWJjOGMwMDBjIiwib3JpZyI6ImJvZHkifQ.X_5ozm9b6bEQkOL9F5TCZcWYm_kS2wjeoaWbHBLGo8c"
	//str := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzb21lIjoicGF5bG9hZCJ9.Joh1R2dYzkRvDkqv3sygm5YyK8Gi4ShZqbhK2gxcs2U"

	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	fmt.Println(err)

	fmt.Printf("%+v", token)
}
