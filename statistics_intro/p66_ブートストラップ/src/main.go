package main

import (
	"fmt"
	"math/rand"
	"time"
)

type height struct {
	man   int
	woman int
}

func main() {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	fmt.Println(random.Intn(100))
}
