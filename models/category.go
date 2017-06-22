package models

import "encoding/json"

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (cat Category) MarshalJSON() ([]byte, error) {
	return json.Marshal(cat)
}
