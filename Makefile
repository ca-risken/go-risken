.PHONY: all
all: help

.PHONY: help
help:
	@echo "You can use sub-command, Usage: make <sub-command>"
	@echo "\n---------------- sub-command list ----------------"
	@cat Makefile | grep -e '^\.PHONY: .*$$' | grep -v -e "all" -e "help" | sed -e 's/^\.PHONY: //g' | sed -e 's/^/- /g' | sort

.PHONY: lint
lint:
	yamllint -c .yamllint generator.yaml

.PHONY: generate
generate:
	go generate ./...
	go run tools/code-generator/main.go

.PHONY: finding-example
finding-example:
	go run examples/finding/main.go

.PHONY: iam-example
iam-example:
	go run examples/iam/main.go
