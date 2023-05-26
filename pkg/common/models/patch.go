package models

import "gorm.io/gorm"

type Patches struct {
	gorm.Model
	Server            string `json:"Server"`
	PatchStart        string `json:"PatchStart"`
	PreCheckScheduled string `json:"PreCheckScheduled"`
	PreCheckStatus    string `json:"PreCheckStatus"`
	PatchScheduled    string `json:"PatchScheduled"`
	Status            string `json:"Status"`
}
