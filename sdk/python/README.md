# pgEdge Pulumi Provider

<img alt="pgEdge" src="https://pgedge-public-assets.s3.amazonaws.com/product/images/pgedge_mark.svg" width="100px">

The official Pulumi provider for [pgEdge Cloud](https://www.pgedge.com/cloud), designed to simplify the management of pgEdge Cloud resources using infrastructure as code.

- **Documentation:** [pgEdge Pulumi Docs](https://www.pulumi.com/registry/packages/pgedge/)
- **Website:** [pgEdge](https://www.pgedge.com/)
- **Discuss:** [GitHub Issues](https://github.com/pgEdge/pulumi-pgedge/issues)

## Prerequisites

Before you begin, ensure you have the following tools installed:

- [Pulumi CLI](https://www.pulumi.com/docs/get-started/install/)
- [Go](https://golang.org/doc/install) (version 1.18 or later)
- [pulumictl](https://github.com/pulumi/pulumictl)
- [golangci-lint](https://golangci-lint.run/usage/install/)
- [Node.js](https://nodejs.org/) (Active LTS or maintenance version, we recommend using [nvm](https://github.com/nvm-sh/nvm) to manage Node.js installations)
- [Yarn](https://yarnpkg.com/getting-started/install)
- [TypeScript](https://www.typescriptlang.org/download)

## Installation

To use this provider, you need to have Pulumi installed. If you haven't already, you can [install Pulumi here](https://www.pulumi.com/docs/get-started/install/).

### Go

```bash
go get github.com/pgEdge/pulumi-pgedge/sdk/go/pgedge
```

### Node.js

```bash
npm install @pgEdge/pulumi-pgedge
```

## Configuration

Before using the provider, you need to configure your pgEdge credentials. Set the following environment variables:

```sh
export PGEDGE_CLIENT_ID="your-client-id"
export PGEDGE_CLIENT_SECRET="your-client-secret"
```

These credentials authenticate the Pulumi provider with your pgEdge Cloud account.

## Getting Started

### Creating a New Pulumi Project

1. Create a new directory for your project:

```bash
mkdir pgedge-pulumi-project && cd pgedge-pulumi-project
```

2. Initialize a new Pulumi project:

```bash
pulumi new typescript
```

Follow the prompts to set up your project.

3. Install the pgEdge provider:

```bash
npm install @pgedge/pulumi-pgedge
```

4. Update your `Pulumi.yaml` file to include the pgEdge provider:

```yaml
name: pgedge-pulumi-project
runtime: nodejs
description: A new Pulumi project using pgEdge
plugins:
  providers:
    - name: pgedge
      path: ./node_modules/@pgEdge/pulumi-pgedge
```

### Writing Your Pulumi Program

Replace the contents of `index.ts` with the following:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as pgedge from "@pgEdge/pulumi-pgedge";

// Create an SSH Key
const sshKey = new pgedge.SSHKey("exampleSSHKey", {
  name: "example",
  publicKey: "ssh-ed25519 AAAAC3NzaC1lZsdw877237ICXfT63i04t5fvvlGesddwed21VG7DkyxvyXbYQNhKP/rSeLY user@example.com",
});

// Create a Cloud Account
const cloudAccount = new pgedge.CloudAccount("exampleCloudAccount", {
  name: "my-aws-account",
  type: "aws",
  description: "My AWS Cloud Account",
  credentials: {
    role_arn: "arn:aws:iam::21112529deae39:role/pgedge-135232c",
  },
}, { 
  dependsOn: [sshKey]
});

// Create a Backup Store
const backupStore = new pgedge.BackupStore("exampleBackupStore", {
  name: "example",
  cloudAccountId: cloudAccount.id,
  region: "us-west-2",
}, { 
  dependsOn: [cloudAccount]
});

// Create a Cluster
const cluster = new pgedge.Cluster("exampleCluster", {
  name: "example",
  cloudAccountId: cloudAccount.id,
  regions: ["us-west-2", "us-east-1", "eu-central-1"],
  nodeLocation: "public",
  sshKeyId: sshKey.id,
  backupStoreIds: [backupStore.id],
  nodes: [
    {
      name: "n1",
      region: "us-west-2",
      instanceType: "r6g.medium",
      volumeSize: 100,
      volumeType: "gp2",
    },
    {
      name: "n2",
      region: "us-east-1",
      instanceType: "r6g.medium",
      volumeSize: 100,
      volumeType: "gp2",
    },
    {
      name: "n3",
      region: "eu-central-1",
      instanceType: "r6g.medium",
      volumeSize: 100,
      volumeType: "gp2",
    },
  ],
  networks: [
    {
      region: "us-west-2",
      cidr: "10.1.0.0/16",
      publicSubnets: ["10.1.0.0/24"],
      // privateSubnets: ["10.1.0.0/24"],
    },
    {
      region: "us-east-1",
      cidr: "10.2.0.0/16",
      publicSubnets: ["10.2.0.0/24"],
      // privateSubnets: ["10.2.0.0/24"],
    },
    {
      region: "eu-central-1",
      cidr: "10.3.0.0/16",
      publicSubnets: ["10.3.0.0/24"],
      // privateSubnets: ["10.3.0.0/24"],
    },
  ],
  firewallRules: [
    {
      name: "postgres",
      port: 5432,
      sources: ["123.456.789.0/32"],
    },
  ],
}, { 
  dependsOn: [sshKey, cloudAccount, backupStore]
});

// Create a Database
const database = new pgedge.Database("exampleDatabase", {
  name: "example",
  clusterId: cluster.id,
  options: [
    "autoddl:enabled",
    "install:northwind",
    "rest:enabled",
  ],
  extensions: {
    autoManage: true,
    requesteds: [
      "postgis",
    ],
  },
  nodes: {
    n1: {
      name: "n1",
    },
    n2: {
      name: "n2",
    },
    n3: {
      name: "n3",
    },
  },
  backups: {
    provider: "pgbackrest",
    configs: [
      {
        id: "default",
        schedules: [
          {
            id: "daily-full-backup",
            cronExpression: "0 1 * * *",
            type: "full",
          },
          {
            id: "hourly-incr-backup",
            cronExpression: "0 * * * ?",
            type: "incr",
          },
        ]
      },
    ]
  },
}, { 
  dependsOn: [cluster]
});

// Export resource IDs
export const sshKeyId = sshKey.id;
export const cloudAccountId = cloudAccount.id;
export const backupStoreId = backupStore.id;
export const clusterId = cluster.id;
export const clusterStatus = cluster.status;
export const clusterCreatedAt = cluster.createdAt;
export const databaseId = database.id;
export const databaseStatus = database.status;

// Optional: Log outputs
pulumi.all([
  sshKeyId,
  cloudAccountId,
  backupStoreId,
  clusterId,
  clusterStatus,
  clusterCreatedAt,
  databaseId,
  databaseStatus
]).apply(([
  sshId,
  accountId,
  backupId,
  cId,
  cStatus,
  cCreatedAt,
  dbId,
  dbStatus
]) => {
  console.log({
    sshKeyId: sshId,
    cloudAccountId: accountId,
    backupStoreId: backupId,
    clusterId: cId,
    clusterStatus: cStatus,
    clusterCreatedAt: cCreatedAt,
    databaseId: dbId,
    databaseStatus: dbStatus
  });
});
```

### Deploying Your Infrastructure

To deploy your infrastructure:

1. Set up your pgEdge credentials as environment variables.
2. Run the following command:

```bash
pulumi up
```

Review the changes and confirm the deployment.

## Updating Resources

### Updating a Database

To update a database, you can modify properties such as `options`, `extensions`, or `nodes`. Here's an example of adding a new extension and removing a node. (Make sure to update one property at a time):

```typescript
const database = new pgedge.Database("exampleDatabase", {
    // ... other properties ...
    options: [
        "install:northwind",
        "rest:enabled",
        "autoddl:enabled",
        "cloudwatch_metrics:enabled", // New option
    ],
    extensions: {
        autoManage: true,
        requesteds: [
            "postgis",
            "vector", // New extension
        ],
    },
    nodes: {
        n1: {
            name: "n1",
        },
        n3: {
            name: "n3",
        },
    },
    // ... other properties ...
});
```

### Updating a Cluster

To update an existing cluster, such as adding or removing nodes, you can modify the `nodes`, `regions`, and `networks` arrays in your Pulumi program. Here's an example of removing a node:

```typescript
const cluster = new pgedge.Cluster("exampleCluster", {
    // ... other properties ...
    nodes: [
        {
            name: "n1",
            region: "us-west-2",
            instanceType: "r6g.medium",
            volumeSize: 100,
            volumeType: "gp2",
        },
        {
            name: "n3",
            region: "eu-central-1",
            instanceType: "r6g.medium",
            volumeSize: 100,
            volumeType: "gp2",
        },
    ],
    regions: ["us-west-2", "eu-central-1"],
    networks: [
        {
            region: "us-west-2",
            cidr: "10.1.0.0/16",
            publicSubnets: ["10.1.0.0/24"],
            // privateSubnets: ["10.1.0.0/24"],
        },
        {
            region: "eu-central-1",
            cidr: "10.3.0.0/16",
            publicSubnets: ["10.3.0.0/24"],
            // privateSubnets: ["10.3.0.0/24"],
        },
    ],
    // ... other properties ...
});
```

After making these changes, run `pulumi up` to apply the updates to your infrastructure.

You can find more examples in the [examples](examples/) directory.

## Contributing

We welcome contributions from the community. Please review our [contribution guidelines](CONTRIBUTING.md) for more information on how to get started.

## License

This project is licensed under the Apache License. See the [LICENSE](LICENSE) file for details.