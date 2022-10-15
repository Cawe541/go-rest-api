package main

import "github.com/Cawe541/go-rest-api.git/server"

func main() {
	Server := server.NewServer()

	Server.Run()
}
