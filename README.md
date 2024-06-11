# pgEdge Pulumi Provider

Welcome to the pgEdge pulumi provider repository. This repository contains the Pulumi resource provider for pgEdge.

## Getting Started

To get started with the pgEdge Pulumi provider, you'll need to install the Pulumi CLI and the pgEdge provider SDK for your language of choice. Below are the installation instructions and example usage in Go and TypeScript.

### Installation

#### Pulumi CLI

Ensure you have the Pulumi CLI installed. You can install it using the following command:
```bash
curl -fsSL https://get.pulumi.com | sh -s -- --version dev
```

#### Go SDK

To install the pgEdge SDK for Go, use the following command in your application:
```bash
go get -u github.com/pgEdge/pulumi-pgedge/sdk/go/pgedge
```

#### TypeScript SDK

To install the pgEdge SDK for TypeScript, use the following command in your application:
```bash
npm install @pgEdge/pulumi-pgedge
```

## Usage

### Go Example

Below is an example of how to use the pgEdge Pulumi provider in a Go program:

```go
package main

import (
    "github.com/pgEdge/pulumi-pgedge/sdk/go/pgedge"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {

        cluster, err := pgedge.NewCluster(ctx, "cluster", &pgedge.ClusterArgs{
            Name:           pulumi.String("n1"),
            CloudAccountId: pulumi.String(""), // cloud account id
            Regions: pulumi.StringArray{
                pulumi.String("us-east-2"),
            },
            FirewallRules: pgedge.ClusterFirewallRuleArray{
                &pgedge.ClusterFirewallRuleArgs{
                    Port: pulumi.Int(5432),
                    Sources: pulumi.ToStringArray([]string{
                        "0.0.0.0/0",
                    }),
                },
            },
            Nodes: pgedge.ClusterNodeArray{
                &pgedge.ClusterNodeArgs{
                    InstanceType:     pulumi.String("t4g.small"),
                    Name:             pulumi.String("n1"),
                    VolumeSize:       pulumi.Int(30),
                    Region:           pulumi.String("us-east-2"),
                    AvailabilityZone: pulumi.String("us-east-2a"),
                    VolumeType:       pulumi.String("gp2"),
                },
            },
            Networks: pgedge.ClusterNetworkArray{
                &pgedge.ClusterNetworkArgs{
                    Region: pulumi.String("us-east-2"),
                    Cidr:   pulumi.String("10.2.0.0/16"),
                    PublicSubnets: pulumi.ToStringArray([]string{
                        "10.2.1.0/24",
                    }),
                },
            },
        })

        if err != nil {
            return err
        }

        _, err = pgedge.NewDatabase(ctx, "database", &pgedge.DatabaseArgs{
            ClusterId: cluster.ID(),
            Name:      pulumi.String("defaultdb"),
        })

        if err != nil {
            return err
        }

        return nil
    })
}
```

### TypeScript Example

Below is an example of how to use the pgEdge Pulumi provider in a TypeScript program:

```ts
import * as pgedge from "@pgEdge/pulumi-pgedge";

var cluster = new pgedge.Cluster("cluster", {
  name: "n1",
  cloudAccountId: "", // cloud account id
  regions: ["us-east-2"],
  firewallRules: [
    {
      port: 5432,
      sources: ["0.0.0.0/0"],
    }
  ],
  nodes: [
    {
      name: "n1",
      instanceType: "t4g.medium",
      volumeSize: 30,
      region: "us-east-2",
      availabilityZone: "us-east-2a",
      volumeType: "gp2",
    }
  ],
  networks: [
    {
      region: "us-east-2",
      cidr: "10.2.0.0/16",
      publicSubnets: ["10.2.1.0/24"],
    }
  ],
});

new pgedge.Database("databaseCreate", {
  name: "defaultdb",
  clusterId: cluster.id,
  options: [""],
});
```

### Additional Examples
For more detailed examples, please refer to the examples directory in this repository:
- [examples](examples)

## License
This project is licensed under the terms of the Apache License. See the [LICENSE](LICENSE) file for details.

Thank you for using the Pulumi pgEdge provider!