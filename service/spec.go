package service

import (
	_ "embed"
	"encoding/json"
)

//go:embed openapi.json
var openapiSpec []byte
var spec map[string]any

func init() {
	json.Unmarshal(openapiSpec, &spec)
}

func GetOpenAPISpec() map[string]any {
	return spec
}
