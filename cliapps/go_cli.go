package main

import (
	"flag"
	"fmt"
)

var myFlag string

func init() {
	flag.StringVar(&myFlag, "name", "greg", "Please provide your name.")
	flag.Parse()
}

func main() {
	fmt.Println("Your provided name is: ", myFlag)
}
