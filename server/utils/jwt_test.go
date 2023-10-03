package utils

import (
	"fmt"
	"testing"
)

func TestJWT(t *testing.T) {
	jwtStr, err := ReleaseToken(1)
	if err != nil {
		panic(err)
	}
	_, claims, err := ParseToken(jwtStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(claims.Id == "1")
}
