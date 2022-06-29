SHELL:=/bin/bash

changelog_args=-o CHANGELOG.md -tag-filter-pattern '^v'
test_command=richgo test ./... $(TEST_ARGS) -v --cover

lint:
	golangci-lint run --concurrency 4 --print-issued-lines=false --exclude-use-default=false --enable=revive --enable=goimports  --enable=unconvert --fix

changelog:
ifdef version
	$(eval changelog_args=--next-tag $(version) $(changelog_args))
endif
	git-chglog $(changelog_args)

check-gotest:
ifeq (, $(shell which richgo))
	$(warning "richgo is not installed, falling back to plain go test")
	$(eval TEST_BIN=go test)
else
	$(eval TEST_BIN=richgo test)
endif
ifdef test_run
	$(eval TEST_ARGS := -run $(test_run))
endif
	$(eval test_command=$(TEST_BIN) ./... $(TEST_ARGS) -v --cover)

test-only: check-gotest
	$(test_command)

test: lint test-only

.PHONY: lint changelog check-gotest test-only test