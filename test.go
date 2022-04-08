//This is the test file where i test stuff

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for true {
		caller()

	}
}
func caller() {
	time.Sleep(1 * time.Second)
	fmt.Println(rand.Intn(4))
}
