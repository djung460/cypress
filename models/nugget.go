package models

import "encoding/json"

type Nugget struct {
	Username string `json:"username"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Body     string `json:"body"`
}

type Nuggets []Nugget

func (ngt Nugget) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"title":    ngt.Title,
		"category": ngt.Category,
		"username": ngt.Username,
		"body":     ngt.Body,
	})
}
