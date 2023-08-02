package main

import (
	"log"

	"github.com/ljsea6/go-clean-architecture/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
