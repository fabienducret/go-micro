package main

import (
	"broker/config"
	"broker/server"
)

func main() {
	c := config.Get()

	if err := server.RunWith(c); err != nil {
		panic(err)
	}
}
