package patches

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greeninja/patch-api/pkg/common/models"
)

func (h handler) GetPatchesByTime(c *gin.Context) {
	var patches []models.Patches

	pstart := c.Param("pstart")

	if result := h.DB.Where("patch_start LIKE ?", pstart+"%").Find(&patches); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.IndentedJSON(http.StatusOK, &patches)
}
