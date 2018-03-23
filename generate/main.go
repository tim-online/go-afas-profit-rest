package main

import (
	"log"
)

var (
	pkg = "main"
)

func main() {
	g := Generator{}
	err := g.All()
	if err != nil {
		log.Fatal(err)
	}
}
