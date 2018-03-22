package swagger

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// LoadModule is used for getting swagger module
func LoadModule(c echo.Context) error {
	if c.Param("module") == "api.json" {
		return c.Blob(http.StatusOK, "application/json", []byte(resourceListingJson))
	}
	fmt.Println(c.Request().URL.String())
	apiKey := c.Request().URL.String()
	fmt.Println(apiKey)
	content := make([]byte, 0)

	if json, ok := apiDescriptionsJson[apiKey]; ok {
		content = []byte(json)
	}

	return c.Blob(http.StatusOK, "application/json", content)
}
