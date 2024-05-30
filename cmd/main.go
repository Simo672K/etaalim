package main

import (
	"ETaalim/internals"
	"log"
	"os"
)

func main() {
	if err := internals.RunServer(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
