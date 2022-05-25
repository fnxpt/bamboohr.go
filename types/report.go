package types

type Report struct {
	Name    string                       `json:"name"`
	Filters map[string]map[string]string `json:"filters,omitempty"`
	Fields  []string                     `json:"fields,omitempty"`
}
