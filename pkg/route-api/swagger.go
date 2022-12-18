package route_api

import (
	_ "embed"
)

//go:embed route_api.swagger.json
var SwaggerJSON []byte
