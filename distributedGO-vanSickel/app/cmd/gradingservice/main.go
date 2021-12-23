package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/jayatwork/helper_utils/distributedGO-vanSickel/app/grades"
	"github.com/jayatwork/helper_utils/distributedGO-vanSickel/app/registry"
	"github.com/jayatwork/helper_utils/distributedGO-vanSickel/app/service"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.GradingService
	r.ServiceURL = serviceAddress

	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		grades.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
