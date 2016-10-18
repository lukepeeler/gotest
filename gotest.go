
package main

import (
	"fmt"

	"github.com/lukepeeler/goconfig"
)


var testInt = -1

func main() {

	goconfig.IntVar(&testInt, "testInt", testInt, "THIS IS A TEST")

	goconfig.Parse()

	fmt.Println("testInt is", testInt)
	// goconfig.Test()
}
