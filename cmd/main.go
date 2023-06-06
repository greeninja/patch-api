package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/greeninja/patch-api/pkg/common/db"
	"github.com/greeninja/patch-api/pkg/patches"
)

func main() {
	router := gin.Default()
	dbHandler := db.Init()

	listenHost := flag.String("host", "localhost", "Host address to listen on")
	listenPort := flag.String("port", "8080", "Port to listen on.")

	flag.Parse()

	patches.RegisterRoutes(router, dbHandler)

	router.Run(*listenHost + ":" + *listenPort)

}
