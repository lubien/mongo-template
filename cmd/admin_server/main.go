package main

import (
	"github.com/fly-apps/mongo-flex/internal/api"
)

func main() {
	if err := api.StartHttpServer(); err != nil {
		panic(err)
	}
}
