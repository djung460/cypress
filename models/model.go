package models

import "encoding/json"

// Model mapping between json and go values
type Model interface {
	// A model just needs to implement this
	json.Marshaler
}
