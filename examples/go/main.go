package main

import (
	"os"

	"github.com/pgEdge/pulumi-pgedge/sdk/go/pgedge"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

	cluster, err := pgedge.NewCluster(ctx, "cluster", &pgedge.ClusterArgs{
		Name:           pulumi.String("testcluster"),
		CloudAccountId: pulumi.String(os.Getenv("PGEDGE_CLOUD_ACCOUNT_ID")),
		Regions: pulumi.StringArray{
			pulumi.String("us-east-1"),
			pulumi.String("us-east-2"),
			pulumi.String("eu-central-1"),
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
				Region:           pulumi.String("us-east-1"),
				AvailabilityZone: pulumi.String("us-east-1a"),
				VolumeType:       pulumi.String("gp2"),
			},
			&pgedge.ClusterNodeArgs{
				InstanceType:     pulumi.String("t4g.small"),
				Name:             pulumi.String("n2"),
				VolumeSize:       pulumi.Int(30),
				Region:           pulumi.String("us-east-2"),
				AvailabilityZone: pulumi.String("us-east-2a"),
				VolumeType:       pulumi.String("gp2"),
			},
			&pgedge.ClusterNodeArgs{
				InstanceType:     pulumi.String("t4g.small"),
				Name:             pulumi.String("n2"),
				VolumeSize:       pulumi.Int(30),
				Region:           pulumi.String("eu-central-1"),
				AvailabilityZone: pulumi.String("eu-central-1a"),
				VolumeType:       pulumi.String("gp2"),
			},
		},

		Networks: pgedge.ClusterNetworkArray{
			&pgedge.ClusterNetworkArgs{
				Region: pulumi.String("us-east-1"),
				Cidr:   pulumi.String("10.1.0.0/16"),
				PublicSubnets: pulumi.ToStringArray([]string{
					"10.1.1.0/24",
				}),
			},
			&pgedge.ClusterNetworkArgs{
				Region: pulumi.String("us-east-2"),
				Cidr:   pulumi.String("10.2.0.0/16"),
				PublicSubnets: pulumi.ToStringArray([]string{
					"10.2.1.0/24",
				}),
			},
			&pgedge.ClusterNetworkArgs{
				Region: pulumi.String("eu-central-1"),
				Cidr:   pulumi.String("10.3.0.0/16"),
				PublicSubnets: pulumi.ToStringArray([]string{
					"10.3.1.0/24",
				}),
			},
		},
	})

	if err != nil {
		return err
	}

	_, err = pgedge.NewDatabase(ctx, "database", &pgedge.DatabaseArgs{
		ClusterId: cluster.ID(),
		Name:      pulumi.String("testdb"),
		Options:   pulumi.StringArray{pulumi.String("install:northwind"), pulumi.String("autoddl:enabled")},
	})

	if err != nil {
		return err
	}

	return nil
})

}
