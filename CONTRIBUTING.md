# Contributing to pulumi-pgedge

We appreciate your interest in contributing to the pulumi-pgedge project! This document provides guidelines and instructions for contributing to this repository.

## Prerequisites

Before you begin, ensure you have the following tools installed:

- [Pulumi CLI](https://www.pulumi.com/docs/get-started/install/)
- [Go](https://golang.org/doc/install) (version 1.18 or later)
- [pulumictl](https://github.com/pulumi/pulumictl)
- [golangci-lint](https://golangci-lint.run/usage/install/)
- [Node.js](https://nodejs.org/) (Active LTS or maintenance version, we recommend using [nvm](https://github.com/nvm-sh/nvm) to manage Node.js installations)
- [Yarn](https://yarnpkg.com/getting-started/install)
- [TypeScript](https://www.typescriptlang.org/download)
- [Python](https://www.python.org/downloads/) (Python 3)
- [.NET SDK](https://dotnet.microsoft.com/download)

## Setting Up Your Development Environment

1. Fork the repository and clone your fork:
   ```
   git clone https://github.com/pgEdge/pulumi-pgedge.git
   cd pulumi-pgedge
   ```

2. Add the upstream repository as a remote:
   ```
   git remote add upstream https://github.com/pgEdge/pulumi-pgedge.git
   ```

3. Create a new branch for your changes:
   ```
   git checkout -b feature/your-feature-name
   ```

## Updating the Terraform Provider

To migrate to a newer Terraform version:

1. Update the pgedge Terraform provider in the `provider/shim` directory:
   ```
   cd provider/shim
   go get -u github.com/pgEdge/terraform-provider-pgedge
   go mod tidy
   go build
   ```

2. Change to the `provider` directory and update dependencies:
   ```
   cd ../
   go mod tidy
   go build
   ```

3. Generate schemas, binaries, and SDKs:
   ```
   make tfgen
   make build
   ```

## Testing Your Changes

1. Set your `PATH` to include the `/bin` directory of the codebase for the current terminal session:
   ```
   export PATH=$PATH:/path/to/pulumi-pgedge/bin
   ```

2. Test the examples:

   For TypeScript:
   ```
   cd examples/typescript
   yarn install
   yarn link "@pgEdge/pulumi-pgedge"
   # Make changes to index.ts
   pulumi up
   ```

   For Go:
   ```
   cd examples/go
   # Add this line to go.mod:
   # replace github.com/pgEdge/pulumi-pgedge/sdk => ../../sdk
   pulumi up
   ```

   Note: If you encounter errors, try copying the binary to your GOPATH:
   ```
   cp bin/pulumi-resource-pgedge $GOPATH/bin
   ```

## Submitting Your Changes

1. Commit your changes with a clear and descriptive commit message:
   ```
   git commit -am "Add feature: your feature description"
   ```

2. Push your changes to your fork:
   ```
   git push origin feature/your-feature-name
   ```

3. Create a pull request from your fork to the main repository.

4. Wait for the maintainers to review your pull request. They may ask for changes or clarifications.

## Code Style and Guidelines

- Follow the existing code style in the project.
- Write clear, concise, and well-documented code.
- Include unit tests for new features or bug fixes.
- Update documentation as necessary.

## Reporting Issues

If you find a bug or have a suggestion for improvement:

1. Check if the issue already exists in the [GitHub Issues](https://github.com/pgEdge/pulumi-pgedge/issues).
2. If not, create a new issue, providing as much detail as possible.

## Getting Help

If you need help or have questions, feel free to:

- Open an issue for discussion
- Reach out to the maintainers

Thank you for contributing to pulumi-pgedge!