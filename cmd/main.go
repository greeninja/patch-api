package main

import (
	"/home/ncampion/git/github/greeninja/patch-api/pkg/common/db"
	"/home/ncampion/git/github/greeninja/patch-api/pkg/patches"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	dbHandler := db.Init()

	patches.RegisterRoutes(router, dbHandler)

	router.Run("localhost:8080")

}
