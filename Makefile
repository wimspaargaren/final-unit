.PHONY: all test

GOBIN = $(GOPATH)/bin
GOTEST = $(GOBIN)/gotest

# Runs lint
lint:
	@echo Linting...
	@golangci-lint  -v --concurrency=3 --config=.golangci.yml --issues-exit-code=0 run \
	--out-format=colored-line-number ${GOLANGCI_ADDITONAL_ARGS}

# Runs unit tests
test: | $(GOTEST)
	@mkdir -p reports
	LOGFORMAT=ASCII gotest -covermode=count -p=4 -v -coverprofile reports/codecoverage_all.cov --tags=${GO_TEST_BUILD_TAGS} `go list ./...`
	@echo "Done running tests"
	@go tool cover -func=reports/codecoverage_all.cov > reports/functioncoverage.out
	@go tool cover -html=reports/codecoverage_all.cov -o reports/coverage.html
	@echo "View report at $(PWD)/reports/coverage.html"
	@tail -n 1 reports/functioncoverage.out

# runs e2e tests
test.e2e: | $(GOTEST)
	@go test --tags=e2e -v ./e2e/... 
	@echo "Done running tests"

#  Creates coverage report
coverage-report:
	@open reports/coverage.html

# Installs current version of generator to cmdline
install:
	@go install ${GOPATH}/src/github.com/wimspaargaren/final-unit/cmd/finalunit  
