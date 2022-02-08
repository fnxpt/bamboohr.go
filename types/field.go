package types

import (
	"encoding/json"
)

// Field represents a BambooHR API field.
type Field struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

func FieldsFromJSON(data []byte) (*[]Field, error) {
	var fields []Field
	err := json.Unmarshal(data, &fields)
	return &fields, err
}
