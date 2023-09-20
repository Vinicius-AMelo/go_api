package studentModel

type Student struct {
	ID      int      `json:"id" validate:"required"`
	Name    string   `json:"name" validate:"required"`
	Age     int      `json:"age" validate:"required"`
	Course  string   `json:"course" validate:"required"`
	Classes []string `json:"classes" validate:"required"`
}
