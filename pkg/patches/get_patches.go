package patches

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greeninja/patch-api/pkg/common/models"
)

func (h handler) GetPatches(c *gin.Context) {
	var patches []models.Patches

	if result := h.DB.Find(&patches); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &patches)
}
