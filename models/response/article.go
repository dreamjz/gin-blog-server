package response

import (
	"time"
)

type ArticleListResult struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Author     string    `json:"author"`
	Title      string    `json:"title" `
	Importance int       `json:"importance"`
	Status     int       `json:"status"`
}

type ArticleDetail struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Author     string    `json:"author"`
	Title      string    `json:"title" `
	Summary    string    `json:"summary"`
	Importance int       `json:"importance"`
	Status     int       `json:"status"`
	Content    string    `json:"content"`
}
