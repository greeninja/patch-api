package patches

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"/home/ncampion/git/github/greeninja/patch-api/pkg/common/models"
)

type UpdatePatchReqBody struct {
	PreCheckScheduled int    `json: "pre_check_scheduled"`
	PreCheckStatus    string `json: "pre_check_status"`
	Status            string `json: status`
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

	c.JSON(http.StatusOK, &patch)
}
