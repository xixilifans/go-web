package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "go tour", "help info")
	flag.StringVar(&name, "n", "go tour", "info")
	flag.Parse()

	log.Printf("name: %s", name)
}
