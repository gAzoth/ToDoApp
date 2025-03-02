package main

import (
	"github.com/gAzoth/todoapp/router"
)

func main() {
	r := router.SetupRouter()

	r.Run(":8080")
}
