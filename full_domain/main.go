package main

import "full_domain/di"

func main() {
	server := di.Init()
	server.StartServer()
}
