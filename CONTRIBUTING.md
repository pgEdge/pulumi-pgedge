# Contributing to pulumi-pgedge

Thanks for your interest in contributing! This document covers how to set up a
development environment, iterate on changes locally for each language SDK, and
get your contributions merged.

## Repository layout

```
provider/        Bridge code (Terraform-provider-pgedge -> Pulumi). Go.
provider/shim/   Thin wrapper that pins the upstream Terraform provider.
sdk/<lang>/      Generated SDK source. Committed; regenerated via `make regen`.
examples/<lang>/ End-to-end examples. Reference setup for local testing.
docs/            Pulumi Registry docs.
```

You almost never edit files under `sdk/<lang>/` by hand -- they're regenerated
from the upstream Terraform provider schema and the bridge config in
`provider/resources.go`.

## Prerequisites

Pin to these versions to match CI. Older versions of `golangci-lint` in
particular will fail with cryptic Go-toolchain mismatch errors.

| Tool             | Version           | Notes                                       |
| ---------------- | ----------------- | ------------------------------------------- |
| Go               | 1.25.9            | Matches the toolchain CI installs           |
| `golangci-lint`  | v2.11.4           | Must be built with Go 1.25+                 |
| Pulumi CLI       | 3.234.0+          | `brew install pulumi` or get.pulumi.com     |
| `pulumictl`      | latest            | https://github.com/pulumi/pulumictl         |
| Node.js          | Active LTS (>=18) | Use `nvm` if juggling versions              |
| Yarn             | 1.22.x (classic)  | `npm i -g yarn`                             |
| TypeScript       | latest            | Installed via the Node SDK's deps           |
| Python           | 3.9+              | For the Python SDK and example              |
| .NET SDK         | 6.0+              | For the .NET SDK and example                |

Quick install on macOS (Homebrew):

```sh
brew install go golangci-lint pulumi node yarn python@3.11 dotnet
go install github.com/pulumi/pulumictl/cmd/pulumictl@latest
```

## Initial setup

```sh
git clone https://github.com/pgEdge/pulumi-pgedge.git
cd pulumi-pgedge
git remote add upstream https://github.com/pgEdge/pulumi-pgedge.git
git checkout -b feature/your-feature-name
```

## Pre-flight checks

Run these locally before pushing -- they mirror what CI runs:

```sh
make lint                    # golangci-lint on provider/
cd provider && go test ./... # provider unit tests
cd sdk && go build ./...     # Go SDK still compiles
```

> Note: provider unit tests and the Go SDK smoke build are also exposed as
> `make test_provider` and `make smoke_go_sdk` once that work merges from the
> security-updates branch.

## Inner dev loops by language

Pick the SDK you're testing against. Each loop ends with `pulumi up` against an
example project, so you'll need pgEdge Cloud credentials:

```sh
export PGEDGE_CLIENT_ID="your-client-id"
export PGEDGE_CLIENT_SECRET="your-client-secret"
# Optional, for non-prod APIs:
export PGEDGE_BASE_URL="https://your-api-host"
```

The provider binary always lives at `./bin/pulumi-resource-pgedge` after a
build. Pulumi finds it via `PATH`, so every loop ends with the same export:

```sh
export PATH=$PATH:$(pwd)/bin
```

### TypeScript

```sh
make dev_typescript                        # builds provider + Node SDK + registers yarn link
export PATH=$PATH:$(pwd)/bin

cd examples/typescript
yarn install
yarn link "@pgEdge/pulumi-pgedge"
pulumi up
```

If you're testing in a Pulumi project you generated yourself with `pulumi new
typescript`, the template's `Pulumi.yaml` includes a `plugins:` block:

```yaml
plugins:
  providers:
    - name: pgedge
      path: ./node_modules/@pgEdge/pulumi-pgedge
```

That path is only valid for the published npm package (which bundles the
provider binary). For a yarn-linked dev SDK, **delete this block** -- otherwise
you'll get `loading PulumiPlugin.yaml: no such file or directory`. The
`examples/typescript/Pulumi.yaml` already omits it; use that as a reference.

To iterate after a code change:

```sh
make dev_typescript        # rebuild
cd examples/typescript && pulumi up
```

### Go

```sh
make dev_go                                # builds provider + Go SDK
export PATH=$PATH:$(pwd)/bin

cd examples/go
pulumi up
```

The example's `go.mod` already contains:

```
replace github.com/pgEdge/pulumi-pgedge/sdk => ../../sdk
```

If you're testing in your own Go Pulumi project, add the same `replace`
directive pointing at this repo's `sdk/` directory.

### Python

```sh
make dev_python                            # builds provider + Python SDK
export PATH=$PATH:$(pwd)/bin

cd examples/python
python3 -m venv venv && source venv/bin/activate
pip install -r requirements.txt
pip install -e ../../sdk/python/bin        # install the locally built SDK over the pinned one
pulumi up
```

The editable install (`pip install -e`) means re-running `make dev_python`
picks up new generated source without a reinstall.

### .NET

```sh
make dev_dotnet                            # builds provider + .NET SDK + stages nupkgs in ./nuget
export PATH=$PATH:$(pwd)/bin

# Once per machine: register the local nuget source
dotnet nuget add source $(pwd)/nuget --name pgedge-local

cd examples/dotnet
# Update dotnet.csproj's <PackageReference Version="..."/> to match the
# locally built nupkg (see ./nuget/Pgedge.Pgedge.<version>.nupkg).
dotnet restore
pulumi up
```

## Regenerating SDKs after a schema change

After bumping the upstream Terraform provider, editing `provider/resources.go`,
or changing `provider/shim/`:

```sh
make regen
```

This rebuilds the provider binary, regenerates every language SDK, and lints
the bridge code. Commit the regenerated source under `sdk/<lang>/`.

> `make development` is a back-compat alias for `make regen`. Earlier versions
> ran `make cleanup` at the end, which deleted the very binaries and yarn links
> needed for testing. That step has been removed; run `make cleanup` manually
> if you want to wipe dev state.

## Updating the upstream Terraform provider

To migrate to a newer version of `terraform-provider-pgedge`:

1. Update the shim:
   ```sh
   cd provider/shim
   go get -u github.com/pgEdge/terraform-provider-pgedge
   go mod tidy
   go build
   ```
2. Update the provider:
   ```sh
   cd ../        # provider/
   go mod tidy
   go build
   ```
3. Regenerate everything:
   ```sh
   cd ../        # repo root
   make regen
   ```
4. Commit changes under `provider/shim/`, `provider/`, and `sdk/`.

### Testing against an unreleased Terraform provider

**Local copy** -- in `provider/shim/go.mod`, add:

```
replace github.com/pgEdge/terraform-provider-pgedge => /path/to/local/terraform-provider-pgedge
```

**Branch on GitHub**:

```sh
cd provider/shim
go get github.com/pgEdge/terraform-provider-pgedge@<branch>
go mod tidy
```

Either way, run `make regen` afterwards. Revert the `go.mod` changes before
committing unless they're meant to ship.

## Make target reference

| Target                | Purpose                                                          |
| --------------------- | ---------------------------------------------------------------- |
| `make dev_typescript` | Build provider + Node SDK + register yarn link. Ready to test.   |
| `make dev_go`         | Build provider + Go SDK.                                         |
| `make dev_python`     | Build provider + Python SDK.                                     |
| `make dev_dotnet`     | Build provider + .NET SDK + stage nupkgs.                        |
| `make regen`          | Rebuild provider, lint, regenerate every SDK source. Pre-commit. |
| `make lint`           | Run `golangci-lint` against `provider/`. Alias for `lint_provider`. |
| `make provider`       | Build just the provider binary into `./bin/`.                    |
| `make build_nodejs`   | Build just the Node SDK.                                         |
| `make build_go`       | Build just the Go SDK source.                                    |
| `make build_python`   | Build just the Python SDK + sdist.                               |
| `make build_dotnet`   | Build just the .NET SDK + nupkg.                                 |
| `make tfgen`          | Regenerate the schema only.                                      |
| `make test`           | Run integration tests under `examples/` (long; needs creds).     |
| `make cleanup`        | Wipe `./bin`, the yarn link registration, and `~/.pulumi/plugins`. |
| `make clean`          | Wipe regenerated SDK source under `sdk/`.                        |

## Submitting changes

1. Commit with a clear, descriptive message:
   ```sh
   git commit -am "Add feature: your feature description"
   ```
2. Push to your fork:
   ```sh
   git push origin feature/your-feature-name
   ```
3. Open a PR against `main`. Make sure:
   - `make lint` passes locally
   - Provider unit tests pass (`cd provider && go test ./...`)
   - The Go SDK still compiles (`cd sdk && go build ./...`)
   - You've manually exercised at least one example end-to-end if your change
     affects runtime behavior

## Code style

- Follow the existing style in the package you're editing.
- Don't hand-edit generated SDK source under `sdk/<lang>/`; change the bridge
  config or the upstream provider and regenerate.
- Add unit tests for non-trivial provider logic.

## Reporting issues

Check the [GitHub Issues](https://github.com/pgEdge/pulumi-pgedge/issues) tab
first. If your issue isn't there, open a new one with reproduction steps and
the version you're running.

## Release process

Releases are managed by the pgEdge team.

1. Make sure everything builds:
   ```sh
   make build
   ```
2. Tag and push:
   ```sh
   make release
   ```
   This creates and pushes a `v<version>` tag, which triggers GitHub Actions
   to build and publish the SDKs to npm, PyPI, and NuGet.
