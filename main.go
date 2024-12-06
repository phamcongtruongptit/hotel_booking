package main

import (
	"context"
	"fmt"
	"github.com/samber/do"
	"hotel_booking/api"
	"hotel_booking/conf"
	"hotel_booking/connection"
	"hotel_booking/log"
	"hotel_booking/repository"
	"hotel_booking/service"
)

func main() {
	injector := do.New()
	defer func() {
		_ = injector.Shutdown()
	}()

	conf.Inject(injector)
	//utils.Inject(injector)
	connection.Inject(injector)
	repository.Inject(injector)
	service.Inject(injector)

	r, err := api.InitRouter(injector)
	if err != nil {
		panic(err)
	}

	cf := do.MustInvoke[*conf.Config](injector)
	addr := fmt.Sprintf(":%v", cf.ApiService.Port)
	log.Infow(context.Background(), fmt.Sprintf("start api server at %v", addr))
	_ = r.Run(addr)
}
