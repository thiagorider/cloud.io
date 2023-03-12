all: help

.PHONY : help
help : Makefile
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<command>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Command

.PHONY: config-docs
config-docs: ## Configure environment for documentation build
	pip3 install --upgrade pip pipenv && \
	pipenv --rm || true && \
	pipenv lock --python=$(shell which python3) && \
	pipenv sync

.PHONY: build-docs
build-docs: ## Build documentation
	pipenv run mkdocs build --clean

.PHONY: serve-docs
serve-docs: ## Preview documentation (local)
	pipenv run mkdocs serve