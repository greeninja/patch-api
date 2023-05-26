package patches

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greeninja/patch-api/pkg/common/models"
)

type UpdatePatchReqBody struct {
	PreCheckScheduled string `json: "PreCheckScheduled"`
	PreCheckStatus    string `json: "PreCheckStatus"`
	Status            string `json: "Status"`
}

func (h handler) UpdatePatch(c *gin.Context) {
	id := c.Param("id")
	body := UpdatePatchReqBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var patch models.Patches

	if result := h.DB.First(&patch, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	patch.PreCheckScheduled = body.PreCheckScheduled
	patch.PreCheckStatus = body.PreCheckStatus
	patch.Status = body.Status

	h.DB.Save(&patch)

	c.IndentedJSON(http.StatusOK, &patch)
}
