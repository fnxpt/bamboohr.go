package bamboohr

import (
	"encoding/json"

	"github.com/AkihikoITOH/bamboohr.go/types"
)

func (api *API) GetReportEmployees(fieldList []string) (*types.EmployeeDirectory, error) {
	path := "/reports/custom?format=JSON&onlyCurrent=true"

	input := &types.Report{
		Name:   "Employees Report",
		Fields: fieldList,
	}
	body, _ := json.Marshal(input)

	data, err := api.post(path, body)
	if err != nil {
		return nil, err
	}

	return types.NewEmployeeDirectoryFromJSON(data)
}
