package patches

import (
	"net/http"

	"github.com/greeninja/patch-api/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) DeletePatch(c *gin.Context) {
	id := c.Param("id")

	var patch models.Patches

	if result := h.DB.First(&patch, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&patch)

	c.Status(http.StatusOK)
}
