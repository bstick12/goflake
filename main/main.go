package main

import (
	. "github.com/bstick12/goflake"
	"fmt"
)

func main() {
	generator := GoFlakeInstanceUsingUnique("D01Z01")
	uuid := generator.GetBase64UUID()
	fmt.Println(uuid)
}
