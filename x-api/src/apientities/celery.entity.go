package apientities

import "time"

type CeleryStatus struct {
	Upload  string `json:"upload" example:"SUCCESS" format:"string"`
	Ml      string `json:"ml" example:"SUCCESS" format:"string"`
	General string `json:"general" example:"SUCCESS" format:"string"`
}

type CeleryTimes struct {
	StartUpload time.Time `json:"start_upload" example:"2021-12-29 10:11:22" format:"string"`
	EndUpload   time.Time `json:"end_upload" example:"2021-12-29 10:11:22" format:"string"`
	StartMl     time.Time `json:"start_ml" example:"2021-12-29 10:11:22" format:"string"`
	EndMl       time.Time `json:"end_ml" example:"2021-12-29 10:11:22" format:"string"`
}

type CeleryStatusResult struct {
	TaskId string       `json:"task_id" example:"550e8400-e29b-41d4-a716-446655440000" format:"string"`
	Status CeleryStatus `json:"status"`
	Times  CeleryTimes  `json:"times"`
}
