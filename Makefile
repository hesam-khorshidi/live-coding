wire:
	@export PATH="$$PATH:$$(go env GOPATH)/bin" && wire ./cmd/app && echo "✅ done!"

swagger:
	swag init --generalInfo=infra/http_server.go --output=./docs