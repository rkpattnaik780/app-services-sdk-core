.PHONY: lint
lint: # Use spectral to lint an openapi spec file, set the TARGET env variable as a path to the spec file
	@rhoasapi lint ${TARGET}

.PHONY: generate
generate: # Generate all SDKs from the current openapi specs in .openapi
	bash scripts/generate-go.sh
	bash scripts/generate-java.sh
	bash scripts/generate-js.sh
	bash scripts/generate-py.sh

.PHONY: gnerate-go
generate-go: # Generate the Go SDK from the current openapi specs in .openapi
	bash scripts/generate-go.sh

.PHONY: gnerate-java
generate-java: # Generate the Java SDK from the current openapi specs in .openapi
	bash scripts/generate-java.sh	

.PHONY: gnerate-js
generate-js: # Generate the Typescript/JavaScript SDK from the current openapi specs in .openapi
	bash scripts/generate-js.sh	

.PHONY: gnerate-py
generate-py: # Generate the Python SDK from the current openapi specs in .openapi
	bash scripts/generate-pt.sh	

.PHONY: help
help: # Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done