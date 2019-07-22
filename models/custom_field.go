package models

// AllowedValue - Allow values
type AllowedValue struct {
	Self string `json:"self"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CustomField - Custom Field
type CustomField struct {
	Required bool `json:"required"`
	Schema   struct {
		Type     string `json:"type"`
		Custom   string `json:"custom"`
		CustomID int    `json:"customId"`
		Items    string `json:"items"`
		System   string `json:"system"`
	} `json:"schema"`
	Name          string         `json:"name"`
	Key           string         `json:"key"`
	Operations    []string       `json:"operations"`
	AllowedValues []AllowedValue `json:"allowedValues"`
}

// FieldResponse - Field Response
type FieldResponse struct {
	Fields map[string]CustomField `json:"fields"`
}
