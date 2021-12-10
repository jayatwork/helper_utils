package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/jayatwork/helper_utils/distributedGO-vanSickel/app/log"
	"github.com/jayatwork/helper_utils/distributedGO-vanSickel/app/registry"
	"github.com/jayatwork/helper_utils/distributedGO-vanSickel/app/service"
)

func main() {
	log.Run("./app.log")

	host, port := "localhost", "4055"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.LogService
	r.ServiceURL = serviceAddress

	ctx, err := service.Start(
		context.Background(),
		r,
		host,
		port,
		log.RegisterHandlers,
	)

	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down log service")
}
