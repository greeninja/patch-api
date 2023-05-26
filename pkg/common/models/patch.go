package models

import "gorm.io/gorm"

type Patches struct {
	gorm.Model
	Server            string `json: "server"`
	PatchStart        string `json: "patch_start"`
	PreCheckScheduled string `json: "pre_check_scheduled"`
	PreCheckStatus    string `json: "pre_check_status"`
	PatchScheduled    string `json: "patch_scheduled"`
	Status            string `json: status`
}
