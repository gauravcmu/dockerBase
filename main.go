package main

import (
	"github.com/dockerBase/ping"
	"github.com/dockerBase/router"
)

func main() {
	rtr := router.NewRouter()
	rtrgroup := rtr.Group("/v1.0")

	ping.Routes(rtrgroup)
}
