package patches

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("patch")
	routes.POST("/", h.AddPatch)
	routes.GET("/", h.GetPatches)
	routes.GET("/:id", h.GetPatch)
	routes.PUT("/:id", h.UpdatePatch)
	routes.DELETE("/:id", h.DeletePatch)
	routes.GET("/date/:pstart", h.GetPatchesByTime)
	routes.PUT("/precheck/:id", h.UpdatePatchPreCheck)
	routes.PUT("/patched/:id", h.UpdatePatchPatched)
}
