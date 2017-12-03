package main

import "./server"

func main() {
	myServer := server.NewServer()
	myServer.Run()
}