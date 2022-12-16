package endpoint_api

import (
	_ "embed"
)

//go:embed endpoint_api.swagger.json
var SwaggerJSON []byte
