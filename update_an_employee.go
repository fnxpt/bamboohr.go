package bamboohr

import (
	"encoding/json"
	"fmt"

	"github.com/AkihikoITOH/bamboohr.go/types"
)

// UpdateAnEmployee updates an employee with provided changes.
// For more details, refer to https://www.bamboohr.com/api/documentation/employees.php
func (api *API) UpdateAnEmployee(employee *types.Employee) error {
	path := fmt.Sprintf("/employees/%d", employee.ID)

	data, err := json.Marshal(employee)
	if err != nil {
		return err
	}

	_, err = api.post(path, data)

	return err
}
