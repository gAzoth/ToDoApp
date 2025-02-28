package main

import (
	"github.com/gAzoth/todoapp/router"
)

func main() {
	r := router.SetupRouter()

	r.run(":8080")
}
