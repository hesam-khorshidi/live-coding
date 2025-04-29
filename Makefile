wire:
	@export PATH="$$PATH:$$(go env GOPATH)/bin" && wire ./cmd/app && echo "✅ done!"