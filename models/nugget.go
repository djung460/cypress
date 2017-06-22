package models

import "encoding/json"

type Nugget struct {
	ID         int    `json:"id"`
	AuthorID   int    `json:"author_id"`
	CategoryID string `json:"category_id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	CreatedAt  string `json:"created_at"`
	EditedAt   string `json:"edited_at"`
}

type Nuggets []Nugget

func (ngt Nugget) MarshalJSON() ([]byte, error) {
	return json.Marshal(ngt)
	// return json.Marshal(map[string]string{
	// 	"id":          string(ngt.ID),
	// 	"author_id":   string(ngt.AuthorID),
	// 	"category_id": ngt.CategoryID,
	// 	"title":       ngt.Title,
	// 	"body":        ngt.Body,
	// 	"created_at":  ngt.CreatedAt,
	// 	"edited_at":   ngt.EditedAt,
	// })
}
