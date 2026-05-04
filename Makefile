PROJECT_NAME := pgedge Package

SHELL            := /bin/bash
PACK             := pgedge
ORG              := pgEdge
PROJECT          := github.com/${ORG}/pulumi-${PACK}
NODE_MODULE_NAME := @pgEdge/pulumi-${PACK}
TF_NAME          := ${PACK}
PROVIDER_PATH    := provider
VERSION_PATH     := ${PROVIDER_PATH}/pkg/version.Version

TFGEN           := pulumi-tfgen-${PACK}
PROVIDER        := pulumi-resource-${PACK}
VERSION         := $(shell pulumictl get version)

TESTPARALLELISM := 4

WORKING_DIR     := $(shell pwd)

OS := $(shell uname)

.PHONY: development regen provider build_sdks build_nodejs build_go build_dotnet build_python \
	dev_typescript dev_go dev_python dev_dotnet lint cleanup

# regen rebuilds the provider, lints it, and regenerates every language SDK.
# Use this after editing the upstream Terraform provider or the bridge config in
# provider/resources.go to refresh the committed SDK source under sdk/<lang>/.
regen:: install_plugins provider lint_provider build_sdks install_sdks # Regenerate provider + all SDK source

# Back-compat alias. Note: unlike previous versions, this no longer runs `cleanup`
# at the end -- you'll be left with a working provider binary and yarn link
# registered. Run `make cleanup` separately if you want to wipe dev state.
development:: regen

# Required for the codegen action that runs in pulumi/pulumi and pulumi/pulumi-terraform-bridge
build:: install_plugins provider build_sdks install_sdks
only_build:: build

tfgen:: install_plugins
	(cd provider && go build -o $(WORKING_DIR)/bin/${TFGEN} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${TFGEN})
	$(WORKING_DIR)/bin/${TFGEN} schema --out provider/cmd/${PROVIDER}
	(cd provider && VERSION=$(VERSION) go generate cmd/${PROVIDER}/main.go)

provider:: tfgen install_plugins # build the provider binary
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${PROVIDER})

build_sdks:: install_plugins provider build_nodejs build_go build_dotnet build_python

build_nodejs:: VERSION := $(shell pulumictl get version --language javascript)
build_nodejs:: install_plugins tfgen # build the node sdk
	$(WORKING_DIR)/bin/$(TFGEN) nodejs --overlays provider/overlays/nodejs --out sdk/nodejs/
	cd sdk/nodejs/ && \
        yarn install && \
        yarn run tsc && \
        cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
		sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" ./bin/package.json

build_python:: PYPI_VERSION := $(shell pulumictl get version --language python)
build_python:: install_plugins tfgen # build the python sdk
	$(WORKING_DIR)/bin/$(TFGEN) python --overlays provider/overlays/python --out sdk/python/
	cd sdk/python/ && \
        cp ../../README.md . && \
        python3 setup.py clean --all 2>/dev/null && \
        rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
        sed -i.bak -e 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' ./bin/setup.py && \
        rm ./bin/setup.py.bak && \
        cd ./bin && python3 setup.py build sdist

build_dotnet:: DOTNET_VERSION := $(shell pulumictl get version --language dotnet)
build_dotnet:: install_plugins tfgen # build the dotnet sdk
	pulumictl get version --language dotnet
	$(WORKING_DIR)/bin/$(TFGEN) dotnet --overlays provider/overlays/dotnet --out sdk/dotnet/
	cd sdk/dotnet/ && \
		echo "${DOTNET_VERSION}" >version.txt && \
        dotnet build /p:Version=${DOTNET_VERSION}

build_go:: install_plugins tfgen # build the go sdk
	$(WORKING_DIR)/bin/$(TFGEN) go --overlays provider/overlays/go --out sdk/go/

lint_provider:: provider # lint the provider code
	cd provider && golangci-lint run -c ../.golangci.yml

lint:: lint_provider # alias for lint_provider

cleanup:: # cleans up the temporary directory
	rm -rf ~/.pulumi/plugins
	rm -rf ~/.config/yarn/link/@pgEdge/pulumi-pgedge/
	rm -r $(WORKING_DIR)/bin
	rm -f provider/cmd/${PROVIDER}/schema.go

help::
	@grep '^[^.#]\+:\s\+.*#' Makefile | \
 	sed "s/\(.\+\):\s*\(.*\) #\s*\(.*\)/`printf "\033[93m"`\1`printf "\033[0m"`	\3 [\2]/" | \
 	expand -t20

clean::
	rm -rf sdk/{nodejs,go,dotnet,python}

install_plugins::
	[ -x $(shell which pulumi) ] || curl -fsSL https://get.pulumi.com | sh
	pulumi plugin install resource random 4.3.1

install_dotnet_sdk::
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

install_python_sdk::

install_go_sdk::

install_nodejs_sdk::
	cd $(WORKING_DIR)/sdk/nodejs/bin && yarn link

install_sdks:: install_nodejs_sdk install_dotnet_sdk install_python_sdk

# === Per-language dev loops ===
# Each target builds just enough to test the named language locally and prints
# the next manual steps (linking, env vars). They intentionally do not auto-clean.
dev_typescript:: provider build_nodejs install_nodejs_sdk # build provider + Node SDK and register yarn link
	@echo ""
	@echo "Next steps:"
	@echo "  export PATH=\$$PATH:$(WORKING_DIR)/bin"
	@echo "  cd <your-project> && yarn link \"@pgEdge/pulumi-pgedge\""
	@echo "  # In your project's Pulumi.yaml, remove any 'plugins:' block that pins"
	@echo "  # path: ./node_modules/@pgEdge/pulumi-pgedge -- it's only valid for the"
	@echo "  # published npm package, not a yarn-linked dev SDK."

dev_go:: provider build_go # build provider + Go SDK
	@echo ""
	@echo "Next steps:"
	@echo "  export PATH=\$$PATH:$(WORKING_DIR)/bin"
	@echo "  # In your example go.mod, add:"
	@echo "  #   replace github.com/pgEdge/pulumi-pgedge/sdk => $(WORKING_DIR)/sdk"
	@echo "  # (see examples/go/go.mod for the canonical setup)."

dev_python:: provider build_python # build provider + Python SDK
	@echo ""
	@echo "Next steps:"
	@echo "  export PATH=\$$PATH:$(WORKING_DIR)/bin"
	@echo "  # From your project's virtualenv:"
	@echo "  pip install -e $(WORKING_DIR)/sdk/python/bin"

dev_dotnet:: provider build_dotnet install_dotnet_sdk # build provider + .NET SDK and stage nupkgs
	@echo ""
	@echo "Next steps:"
	@echo "  export PATH=\$$PATH:$(WORKING_DIR)/bin"
	@echo "  # Add this repo's nuget/ as a local source for your project:"
	@echo "  dotnet nuget add source $(WORKING_DIR)/nuget --name pgedge-local"
	@echo "  # Then reference the locally built version in your csproj."

test::
	cd examples && go test -v -tags=all -parallel ${TESTPARALLELISM} -timeout 2h

release::
	@if [ -n "$(CUSTOM_VERSION)" ]; then \
		VERSION="$(CUSTOM_VERSION)"; \
	elif [ -z "$(VERSION)" ]; then \
		echo "VERSION is not set. Either set VERSION or use CUSTOM_VERSION=x.x.x"; \
		exit 1; \
	fi; \
	echo "Creating release for version v$$VERSION"; \
	git tag -a v$$VERSION -m "Release v$$VERSION"; \
	git push origin v$$VERSION; \
	echo "Release v$$VERSION has been created and pushed"
