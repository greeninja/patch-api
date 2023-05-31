package models

import "gorm.io/gorm"

type Patches struct {
	gorm.Model
	Server            string `json:"Server"`
	PatchStart        string `json:"PatchStart"`
	PreCheckScheduled string `json:"PreCheckScheduled"`
	PreCheckStatus    string `json:"PreCheckStatus"`
	PreCheckJobID     string `json:"PreCheckJobID"`
	PatchScheduled    string `json:"PatchScheduled"`
	PatchScheduleID   string `json:"PatchScheduleID"`
	PatchJobID        string `json:"PatchJobID"`
	Status            string `json:"Status"`
}
