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

type ArticleContent struct {
	Content string `json:"content"`
}
