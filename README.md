# pgEdge Cloud Pulumi Provider

<img alt="pgEdge" src="https://pgedge-public-assets.s3.amazonaws.com/product/images/pgedge_mark.svg" width="100px">

The official Pulumi provider for [pgEdge Cloud](https://www.pgedge.com/cloud),
designed to simplify the management of pgEdge Cloud resources using
infrastructure as code.

- **Documentation:** [pgEdge Pulumi Docs](https://www.pulumi.com/registry/packages/pgedge/)
- **Website:** [pgEdge](https://www.pgedge.com/)
- **Issues:** [GitHub Issues](https://github.com/pgEdge/pulumi-pgedge/issues)
- **Contributing / local development:** [CONTRIBUTING.md](CONTRIBUTING.md)

## Installation

You'll need the [Pulumi CLI](https://www.pulumi.com/docs/get-started/install/)
and a runtime for your language of choice.

### TypeScript / JavaScript

```bash
npm install @pgEdge/pulumi-pgedge
```

### Go

```bash
go get github.com/pgEdge/pulumi-pgedge/sdk/go/pgedge
```

### Python

```bash
pip install pgEdge_pulumi_pgedge
```

### .NET

```bash
dotnet add package Pgedge.Pgedge
```

## Configuration

Authenticate by exporting credentials from a pgEdge Cloud API client:

```sh
export PGEDGE_CLIENT_ID="your-client-id"
export PGEDGE_CLIENT_SECRET="your-client-secret"
# Optional, only when targeting a non-default API host:
export PGEDGE_BASE_URL="https://your-api-host"
```

Equivalently, you can set them via Pulumi config:

```sh
pulumi config set pgedge:clientId      "your-client-id" --secret
pulumi config set pgedge:clientSecret  "your-client-secret" --secret
pulumi config set pgedge:baseUrl       "https://your-api-host"
```

## Quick start

Create a new Pulumi project and add the provider:

```bash
mkdir my-pgedge-project && cd my-pgedge-project
pulumi new typescript
npm install @pgEdge/pulumi-pgedge
```

A minimal program -- one SSH key and one cluster on an existing cloud account:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as pgedge from "@pgEdge/pulumi-pgedge";

const sshKey = new pgedge.SSHKey("example", {
  name: "example",
  publicKey: "ssh-ed25519 AAAA... user@example.com",
});

const cluster = new pgedge.Cluster("example", {
  name: "example",
  cloudAccountId: "<your-cloud-account-id>",
  regions: ["us-west-2"],
  nodeLocation: "public",
  sshKeyId: sshKey.id,
  nodes: [{
    name: "n1",
    region: "us-west-2",
    instanceType: "r6g.medium",
    volumeSize: 100,
    volumeType: "gp2",
  }],
  networks: [{
    region: "us-west-2",
    cidr: "10.1.0.0/16",
    publicSubnets: ["10.1.0.0/24"],
  }],
  firewallRules: [{
    name: "postgres",
    port: 5432,
    sources: ["0.0.0.0/0"],
  }],
});

export const clusterId = cluster.id;
```

Deploy:

```bash
pulumi up
```

## Full examples

See the [`examples/`](examples/) directory for end-to-end programs in each
supported language, including cloud account creation, backup stores, multi-
region clusters, databases, extensions, and scheduled backups:

- TypeScript: [`examples/typescript/index.ts`](examples/typescript/index.ts)
- Go: [`examples/go/main.go`](examples/go/main.go)
- Python: [`examples/python/__main__.py`](examples/python/__main__.py)
- .NET: [`examples/dotnet/Program.cs`](examples/dotnet/Program.cs)

## Updating resources

You can modify properties of an existing resource and run `pulumi up` to apply
the changes. A few resource-specific notes:

- **Database** -- update `options`, `extensions`, or `nodes` one property at a
  time. The provider applies changes idempotently.
- **Cluster** -- to add or remove nodes, edit the `nodes`, `regions`, and
  `networks` arrays together so they stay consistent.

See the [`examples/typescript/index.ts`](examples/typescript/index.ts) for a
worked example showing both.

## Contributing

Contributions are welcome -- see [CONTRIBUTING.md](CONTRIBUTING.md) for the
full local development guide, including per-language inner-loop instructions,
required tool versions, and the pre-flight checks CI runs.

## License

Licensed under the Apache License 2.0. See [LICENSE](LICENSE).
