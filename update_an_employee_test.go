package bamboohr_test

import (
	"fmt"
	"net/http"
	"testing"

	"gotest.tools/assert"

	bamboohr "github.com/AkihikoITOH/bamboohr.go"
	"github.com/AkihikoITOH/bamboohr.go/types"
)

func Example_updateAnEmployee() {
	conf := bamboohr.NewConfig("api-domain", "api-key")
	api := bamboohr.NewAPI(conf)
	employee := types.Employee{
		ID:        99999,
		FirstName: "John",
		LastName:  "Doe",
	}
	api.UpdateAnEmployee(&employee)
}

func TestUpdateAnEmployee(t *testing.T) {
	client := NewMockClient(employeeData)
	api := &bamboohr.API{Config: mockAPIConfig, HTTPClient: client}

	employee := types.Employee{
		ID:        99999,
		FirstName: "John",
		LastName:  "Doe",
	}

	err := api.UpdateAnEmployee(&employee)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodPost)
	assert.Equal(t, client.Request.URL.String(), "https://api.bamboohr.com/api/gateway.php/test/v1/employees/99999")

	SharedRequestTests(t, client.Request)
}
