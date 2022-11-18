

#################### DEVELOPMENT ####################
.PHONY: run/dev
run/dev:
	@go run ./cmd/api -cors-trusted-origins="http://localhost:3000"
