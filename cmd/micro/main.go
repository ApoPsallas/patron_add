package main

import (
	"fmt"
	rhttp "hellomicroservice/http"
	"os"

	"github.com/thebeatapp/patron"
	"github.com/thebeatapp/patron/log"
	phttp "github.com/thebeatapp/patron/sync/http"
)

var (
	version = "dev"
)

func main() {
	name := "micro"

	err := patron.Setup(name, version)
	if err != nil {
		fmt.Printf("failed to set up logging: %v", err)
		os.Exit(1)
	}
	// Set up HTTP routes
	routes := []phttp.Route{
		phttp.NewGetRoute("/marco", rhttp.HandleMarco, true),
		phttp.NewGetRoute("/add", rhttp.HandleAddition, true),
	}
	srv, err := patron.New(name, version, patron.Routes(routes))

	if err != nil {
		log.Fatalf("failed to create service %v", err)
	}

	err = srv.Run()
	if err != nil {
		log.Fatalf("failed to run service %v", err)
	}
}
