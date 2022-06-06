package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
)

var n int

func init() {
	flag.IntVar(&n, "n", 1, "id numbers")
}

func main() {
	flag.Parse()
	for i := 0; i < n; i++ {
		fmt.Println(UUID())
	}
}

func UUID() string {
	return uuid.New().String()
}
