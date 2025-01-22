.PHONY: generate
generate:
	go generate ./ent

.PHONY: api
api:
	oapi-codegen -generate types -o restapi/types.gen.go -package restapi docs/openapi.yml
	oapi-codegen -generate chi-server -o restapi/api.gen.go -package restapi docs/openapi.yml
