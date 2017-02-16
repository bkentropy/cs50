package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	s := flag.Int64("seed", 5, "seed number")
	numbPtr := flag.Int("iterations", 10, "an int64")
	flag.Parse()
	rand.Seed(*s)
	for i := 0; i < *numbPtr; i++ {
		fmt.Println(rand.Intn(10))
	}

}
