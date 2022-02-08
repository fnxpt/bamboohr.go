package bamboohr

import (
	"encoding/json"
	"fmt"

	"github.com/AkihikoITOH/bamboohr.go/types"
)

// UpdateAnEmployee updates an employee with the given employee number.
// For more details, refer to https://www.bamboohr.com/api/documentation/employees.php
func (api *API) UpdateAnEmployee(number string, employee *types.EmployeeUpdate) error {
	path := fmt.Sprintf("/employees/%s", number)
	body, _ := json.Marshal(employee)

	_, err := api.post(path, body)

	return err
}
