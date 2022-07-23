package pkg

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type HashedTokenByUID struct {
	UID         int    `json:"uid"`
	HashedToken string `json:"hashed_token"`
}

var HashedTokens []HashedTokenByUID

func hashToken(token string) []byte {
	hashString := sha256.New()
	hashString.Write([]byte(token))
	bs := hashString.Sum(nil)
	return bs
}

func AuthenticateUser(reqToken string) int {
	byteHashed := hashToken(reqToken)
	stringHashed := fmt.Sprintf("%x", byteHashed)
	fmt.Println(stringHashed)
	for _, hToken := range HashedTokens {
		if strings.Compare(hToken.HashedToken, stringHashed) == 0 {
			return hToken.UID
		}
	}
	return -1
}
