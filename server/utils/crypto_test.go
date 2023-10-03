package utils

import (
	"fmt"
	"testing"
)

func TestAES(t *testing.T) {
	const plain = `nZxcj0nPW6PH31qBWp2vaw==`
	fmt.Println(DecodeAESWithKey("your 44 char key", plain))
}
