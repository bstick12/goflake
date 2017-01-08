package main

import (
	"fmt"
	. "github.com/bstick12/goflake"
)

func main() {
	generator := GoFlakeInstanceUsingUnique("D01Z01")
	uuid := generator.GetBase64UUID()
	fmt.Println(uuid)
}
