package patches

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db, *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("patches")
	routes.Post("/", h.AddPatch)
	routes.GET("/", h.GetPatches)
	routes.GET("/:id", h.GetPatch)
	routes.PUT("/:id", h.UpdatePatch)
	routes.DELETE("/:id", h.DeletePatch)
}