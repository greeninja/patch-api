package patches

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"/home/ncampion/git/github/greeninja/patch-api/pkg/common/models"
)

func (h handler) GetPatch(c *gin.Context) {
	id := c.Param("id")

	var patch models.Patches

	if result := h.DB.First(&patch, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &patch)
}
