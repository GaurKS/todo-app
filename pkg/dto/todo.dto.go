package dtos

// todo - DTOs
type Todo struct {
	ID     				 string        `json:"id"`
	Title          string        `json:"title,omitempty"`
	TodoStatus     string        `json:"todoStatus,omitempty"`
	Description    string        `json:"description,omitempty"`
	CreatedBy      string        `json:"createdBy,omitempty"`
}