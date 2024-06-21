package main

import (
	"fmt"
	"github.com/samber/lo"
)

func main() {
	fmt.Println(lo.PascalCase("\n\taccount-service AND AUTH SERVICE Int8Value HTTPCode200 Timeout:1000ms__\n\t"))
	fmt.Println(lo.CamelCase("\n\taccount-service AND AUTH SERVICE Int8Value HTTPCode200 Timeout:1000ms__\n\t"))
	fmt.Println(lo.SnakeCase("\n\taccount-service AND AUTH SERVICE Int8Value HTTPCode200 Timeout:1000ms__\n\t"))
	fmt.Println(lo.KebabCase("\n\taccount-service AND AUTH SERVICE Int8Value HTTPCode200 Timeout:1000ms__\n\t"))
}
