package models

type Article struct {
	Model
	Author     string `json:"author"    binding:"required"`
	Title      string `json:"title"     binding:"required"`
	Summary    string `json:"summary"`
	Content    string `json:"content"   binding:"required"`
	Importance int    `json:"importance"`
	Status     *int   `json:"status"    binding:"required"`
	CreatedBy  string `json:"createdBy"`
	UpdatedBy  string `json:"UpdatedBy"`
}
