package main

import (
	"context"
	"github.com/jayatwork/helper_utils/distributedGO-vanSickel/app/log"
	"github.com/jayatwork/helper_utils/distributedGO-vanSickel/app/service"
	stlog "log"
	"fmt"
)

func main() {
	log.Run("./app.log")

	host, port := "localhost", "4055"

	ctx, err := service.Start(
		context.Background(), 
		"Log Service",
		host,
		port, 
		log.RegisterHandlers)

	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down log service")
}