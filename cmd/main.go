package main

import (
	"github.com/gin-gonic/gin"
	"github.com/greeninja/patch-api/pkg/common/db"
	"github.com/greeninja/patch-api/pkg/patches"
)

func main() {
	router := gin.Default()
	dbHandler := db.Init()

	patches.RegisterRoutes(router, dbHandler)

	router.Run("localhost:8080")

}
