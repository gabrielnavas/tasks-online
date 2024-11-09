package dtos

type CreateTaskDto struct {
	Description string `json:"description"`
}

type UpdateTaskDto struct {
	Done        bool   `json:"done"`
	Description string `json:"description"`
}
