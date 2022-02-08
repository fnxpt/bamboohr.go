package types

import (
	"encoding/json"
)

// EmployeeDirectory contains a list of fields and a list of employees provided by the API.
// (https://www.bamboohr.com/api/documentation/employees.php)
type EmployeeDirectory struct {
	Employees []*Employee `json:"employees"`
	Fields    []*Field    `json:"fields"`
}

// NewEmployeeDirectoryFromJSON parses the given JSON (as byte sequence) and returns a new EmployeeDirectory.
func NewEmployeeDirectoryFromJSON(data []byte) (*EmployeeDirectory, error) {
	var employees EmployeeDirectory
	err := json.Unmarshal(data, &employees)
	return &employees, err
}
