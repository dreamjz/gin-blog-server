package models

type User struct {
	Model
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedBy string `json:"createdBy"`
	UpdatedBy string `json:"updatedBy"`
}
