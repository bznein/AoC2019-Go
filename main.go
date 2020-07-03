package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	for {
		fmt.Println("Random: %+v", rand.Intn(3))
	}
	for {
		fmt.Println("Random: %+v", rand.Intn(3))
	}

}
