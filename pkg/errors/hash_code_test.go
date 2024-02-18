package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestErr2HashCode(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := errors.New(fmt.Sprintf("Access denied for user'db_user'@'127.0.0.%d'", i))
		code, msg := Err2HashCode(err)
		fmt.Println(code, msg)
	}
}
