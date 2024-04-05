package cmd

import (
	"log"
	"videoChat/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
