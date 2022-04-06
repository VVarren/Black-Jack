//This is the test file where i test stuff

package main

import (
	"fmt"
	"strings"
)

func main() {
	var first string
	fmt.Scanln(&first)
	first = strings.ToUpper(first)
	fmt.Println(first)
}
