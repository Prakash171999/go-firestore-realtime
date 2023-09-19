package main

import (
	routes "github.com/prakash171999/go-realtime-firestore/api/routes"
)

func main() {
	r := routes.SetupRouter()

	r.Run()
}
