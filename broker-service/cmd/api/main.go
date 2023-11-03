package main

import (
	"broker/config"
	"broker/server"
)

func main() {
	c := config.Get()

	server.RunWith(c)
}
