package main

import (
	"mymod/internal/method"
	"mymod/internal/server"
)

func main() {
	method.FileIntoMap()
	server.Server()
}
