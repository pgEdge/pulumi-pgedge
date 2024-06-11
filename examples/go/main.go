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
		// Options:   pulumi.StringArray{pulumi.String("replication = '1'")},
	})

	if err != nil {
		return err
	}

	return nil
})

}
