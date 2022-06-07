package main

import (
	"fmt"
	"mymod/internal/method"
	"mymod/internal/server"
)

func main() {
	err := method.FileIntoMap()
	if err != nil {
		fmt.Println(err)
	}
	server.Server()
}
