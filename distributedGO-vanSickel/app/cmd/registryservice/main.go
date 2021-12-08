package main

import (
	"github.com/jayatwork/helper_utils/distributedGO-vanSickel/app/registry"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.Handle("/services", &registry.RegistryService{})

	ctx, cancel := context.Cancel(context.Background())
	defer cancel()
	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("Registry service started. Press any key to stop.")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
	}()

	<-ctx.Done
	fmt.Println("Shutting down the registry service")
}