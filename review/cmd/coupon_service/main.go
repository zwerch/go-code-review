package main

import (
	"fmt"

	"coupon_service/internal/api"
	"coupon_service/internal/config"
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service"
)

// these should not be package level variables but rather local ones (encapsulation)
var (
	cfg  = config.New()
	repo = memdb.New()
)

func main() {
	svc := service.New(repo)                      // generic variable naming
	本 := api.New(cfg.API, svc)                    // variable naming could be improved (call it e.g. server)
	fmt.Println("Starting Coupon service server") // no structured logging here, use common logger package, no logging from webserver
	本.Start()                                     // does not return an error to check for

	// <-time.After(1 * time.Hour * 24 * 365)                         // why wait like this and not use time.Sleep? refactor main flow
	本.Close()                                                      // would usually be deferred
	fmt.Println("Coupon service server alive for a year, closing") // why should it only run for a year?
	// TODO: refactor whole flow, never reaches the log messages
}
