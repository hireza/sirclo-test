run:
	@docker-compose up -d --build

log:
	@docker-compose logs -f

unit-test:
	@go test ./...

integration-test: run
	@cd domain/integration_tests && go test -v ./...

test-coverage:
	@go test ./... -coverprofile=cover.out.tmp
	@cat cover.out.tmp | grep -v "_mock.go" > cover.out
	@go tool cover -html=cover.out -o cover.html
	@open cover.html

generate-mock:
	@echo "GENERATING ..."
	@echo "- weight"
	@mockgen -destination=domain/mocks/weight_mock.go -package=mocks -source=domain/weight.go
	@echo "DONE ..."