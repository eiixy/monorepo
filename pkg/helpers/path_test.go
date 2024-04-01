package helpers

import (
	"fmt"
	"testing"
)

func TestPath(t *testing.T) {
	fmt.Println(Path())
	fmt.Println(Path(""))
	fmt.Println(Path("/"))
	fmt.Println(Path("test.go"))
	fmt.Println(Path("./test.go"))
	fmt.Println(Path("/test.go"))
	fmt.Println(Path("cmd", "enums.go"))
	fmt.Println(Path("./cmd", "./enums.go"))
	fmt.Println(Path("/cmd", "enums.go"))
}
