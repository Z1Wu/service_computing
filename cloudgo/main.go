package main

import (
	"github.com/Z1Wu/testweb/service"
)

func main() {
	n := service.NewServer()
	n.Run(":8080")
}
