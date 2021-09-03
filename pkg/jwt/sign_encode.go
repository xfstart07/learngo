// Author: Xu Fei
// Date: 2018/7/30
package main

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type User struct {
	Nonce string `json:"nonce"`
	Orig  string `json:"orig"`
	jwt.StandardClaims
}

func main() {
	payload := &User{
		Nonce:          uuid.New().String(),
		Orig:           "body",
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), payload)

	tokenStr, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", tokenStr)
}
