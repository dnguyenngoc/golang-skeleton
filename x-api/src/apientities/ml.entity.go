package apientities

type MlFace struct {
	TaskId   string `json:"task_id" example:"4234-3243-23-24234-23442"`
	Status   string `json:"status" example:"SUCCESS"`
	FileUrl  string `json:"file_url" example:"http://localhost:8080/storage/images/1.jpg"`
	FilePath string `json:"file_path" example:"storage/images/1.jpg"`
}
