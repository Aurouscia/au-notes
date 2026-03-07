package main

import (
	"fmt"

	"example.com/mylib/utils"
)

func main() {
	var greetResult = utils.Greet("Workspace")
	fmt.Println(greetResult)
}
