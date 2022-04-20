package bamboohr

import (
	"github.com/AkihikoITOH/bamboohr.go/types"
)

// GetFields fetches all fields from the BambooHR API.
// For more details, refer to https://documentation.bamboohr.com/reference/metadata-get-a-list-of-fields
func (api *API) GetFields() (*[]types.Field, error) {
	path := "/meta/fields/"

	data, err := api.get(path)

	if err != nil {
		return nil, err
	}

	return types.FieldsFromJSON(data)
}
