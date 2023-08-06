package handler

type CreateOpeningRequest struct {
	Role     string `json:"role" validate:"required"`
	Company  string `json:"company" validate:"required"`
	Location string `json:"location" validate:"required"`
	Remote   *bool  `json:"remote" validate:"required"`
	Link     string `json:"link" validate:"required"`
	Salary   int64  `json:"salary" validate:"required,min=1"`
}

type UpdateOpeningRequest struct {
	Role     string `json:"role" validate:"omitempty"`
	Company  string `json:"company" validate:"omitempty"`
	Location string `json:"location" validate:"omitempty"`
	Remote   *bool  `json:"remote" validate:"omitempty"`
	Link     string `json:"link" validate:"omitempty"`
	Salary   int64  `json:"salary" validate:"omitempty,min=1"`
}
