package main

import (
	"log"

	"mymod/internal/helpers"
	"mymod/internal/server"
)

func main() {
	if err := helpers.ReadFileIntoMap(); err != nil {
		log.Println(err)
	}
	if err := server.Server(); err != nil {
		log.Println(err)
	}
}
