package helpers

import (
	"fmt"
	"testing"
)

func TestGenUid(t *testing.T) {
	for i := 0; i < 15; i++ {
		fmt.Println(RandomNumber(100000, 999999))
	}
	fmt.Println("==========")
	for i := 0; i < 15; i++ {
		fmt.Println(RandomAlphaStr(12))
	}
	fmt.Println("==========")
	for i := 0; i < 15; i++ {
		fmt.Println(RandomAlphaNumStr(12))
	}
	fmt.Println("==========")
	for i := 0; i < 15; i++ {
		fmt.Println(RandomAlphaDashStr(10))
	}
}
