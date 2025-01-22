.PHONY: ent
ent:
	go generate ./ent

.PHONY: api
api:
	oapi-codegen -generate types -o restapi/types.gen.go -package restapi docs/openapi.yml
	oapi-codegen -generate chi-server -o restapi/api.gen.go -package restapi docs/openapi.yml

.PHONY: docker
docker:
	docker build -t test-server .
	docker run --rm -p 80:80 test-server

.PHONY: test
test:
# test with coverage
	go test -v -cover ./...
