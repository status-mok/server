package server_api

import (
	_ "embed"
)

//go:embed server_api.swagger.json
var SwaggerJSON []byte
