package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jayatwork/helper_utils/distributedGO-vanSickel/app/registry"
)

func main() {

	http.Handle("/services", &registry.RegistryService{})

	ctx, cancel := context.WithCancel(context.Background())
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now())
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

	<-ctx.Done()
	fmt.Println("Shutting down the registry service")
}
