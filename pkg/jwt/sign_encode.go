// Author: Xu Fei
// Date: 2018/7/30
package main

import (
	"encoding/hex"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Nonce string `json:"nonce"`
	Orig  string `json:"orig"`
	jwt.StandardClaims
}

func main() {
	payload := &User{
		Nonce: GetUUIDV4(),
		Orig:  "body",
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), payload)

	tokenStr, err := token.SignedString([]byte("secret"))

	fmt.Println(err)

	fmt.Printf("%+v", tokenStr)

}

func GetUUIDV4() (uuidHex string) {
	uuidV4 := uuid.NewV4()
	uuidHex = hex.EncodeToString(uuidV4.Bytes())
	return
}
