package patches

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greeninja/patch-api/pkg/common/models"
)

type AddPatchReqBody struct {
	Server            string `json: "server"`
	PatchStart        string `json: "patch_start"`
	PreCheckScheduled int    `json: "pre_check_scheduled"`
	PreCheckStatus    string `json: "pre_check_status"`
	PatchScheduled    int    `json: "patch_scheduled"`
	Status            string `json: status`
}

func (h handler) AddPatch(c *gin.Context) {
	body := AddPatchReqBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var patch models.Patches

	patch.Server = body.Server
	patch.PatchStart = body.PatchStart
	patch.PreCheckScheduled = body.PreCheckScheduled
	patch.PreCheckStatus = body.PreCheckStatus
	patch.PatchScheduled = body.PatchScheduled
	patch.Status = body.Status

	if result := h.DB.Create(&patch); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &patch)
}
