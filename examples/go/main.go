package main

import (
	"fmt"

	"github.com/pgEdge/pulumi-pgedge/sdk/go/pgedge"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Configuration
		cfg := config.New(ctx, "")
		baseUrl := cfg.Get("baseUrl")
		if baseUrl == "" {
			baseUrl = "https://api.pgedge.com"
		}

		// SSH Key
		sshKey, err := pgedge.NewSSHKey(ctx, "exampleSSHKey", &pgedge.SSHKeyArgs{
			Name:      pulumi.String("example"),
			PublicKey: pulumi.String("ssh-ed25519 AAAAC3NzaC1lZsdw877237ICXfT63i04t5fvvlGesddwed21VG7DkyxvyXbYQNhKP/rSeLY user@example.com"),
		})
		if err != nil {
			return err
		}

		// Cloud Account
		cloudAccount, err := pgedge.NewCloudAccount(ctx, "exampleCloudAccount", &pgedge.CloudAccountArgs{
			Name:        pulumi.String("my-aws-account"),
			Type:        pulumi.String("aws"),
			Description: pulumi.String("My AWS Cloud Account"),
			Credentials: pulumi.StringMap{
				"role_arn": pulumi.String("arn:aws:iam::1234567890:role/pgedge-13we22c"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{sshKey}))
		if err != nil {
			return err
		}

		// Backup Store
		backupStore, err := pgedge.NewBackupStore(ctx, "exampleBackupStore", &pgedge.BackupStoreArgs{
			Name:           pulumi.String("example"),
			CloudAccountId: cloudAccount.ID(),
			Region:         pulumi.String("us-west-2"),
		}, pulumi.DependsOn([]pulumi.Resource{cloudAccount}))
		if err != nil {
			return err
		}

		// Cluster
		cluster, err := pgedge.NewCluster(ctx, "exampleCluster", &pgedge.ClusterArgs{
			Name:           pulumi.String("example"),
			CloudAccountId: cloudAccount.ID(),
			Regions:        pulumi.ToStringArray([]string{"us-west-2", "us-east-1", "eu-central-1"}),
			NodeLocation:   pulumi.String("public"),
			SshKeyId:       sshKey.ID(),
			Nodes: pgedge.ClusterNodeArray{
				&pgedge.ClusterNodeArgs{
					Name:         pulumi.String("n1"),
					Region:       pulumi.String("us-west-2"),
					InstanceType: pulumi.String("r6g.medium"),
					VolumeSize:   pulumi.Int(100),
					VolumeType:   pulumi.String("gp2"),
				},
				&pgedge.ClusterNodeArgs{
					Name:         pulumi.String("n2"),
					Region:       pulumi.String("us-east-1"),
					InstanceType: pulumi.String("r6g.medium"),
					VolumeSize:   pulumi.Int(100),
					VolumeType:   pulumi.String("gp2"),
				},
				&pgedge.ClusterNodeArgs{
					Name:         pulumi.String("n3"),
					Region:       pulumi.String("eu-central-1"),
					InstanceType: pulumi.String("r6g.medium"),
					VolumeSize:   pulumi.Int(100),
					VolumeType:   pulumi.String("gp2"),
				},
			},
			Networks: pgedge.ClusterNetworkArray{
				&pgedge.ClusterNetworkArgs{
					Region:        pulumi.String("us-west-2"),
					Cidr:          pulumi.String("10.1.0.0/16"),
					PublicSubnets: pulumi.ToStringArray([]string{"10.1.0.0/24"}),
					// PrivateSubnets: pulumi.ToStringArray([]string{"10.1.1.0/24"}),
				},
				&pgedge.ClusterNetworkArgs{
					Region:        pulumi.String("us-east-1"),
					Cidr:          pulumi.String("10.2.0.0/16"),
					PublicSubnets: pulumi.ToStringArray([]string{"10.2.0.0/24"}),
					// PrivateSubnets: pulumi.ToStringArray([]string{"10.2.1.0/24"}),
				},
				&pgedge.ClusterNetworkArgs{
					Region:        pulumi.String("eu-central-1"),
					Cidr:          pulumi.String("10.3.0.0/16"),
					PublicSubnets: pulumi.ToStringArray([]string{"10.3.0.0/24"}),
					// PrivateSubnets: pulumi.ToStringArray([]string{"10.3.1.0/24"}),
				},
			},
			FirewallRules: pgedge.ClusterFirewallRuleArray{
				&pgedge.ClusterFirewallRuleArgs{
					Name:    pulumi.String("postgres"),
					Port:    pulumi.Int(5432),
					Sources: pulumi.ToStringArray([]string{"151.302.23.96/32"}),
				},
			},
			// backupStoreIds: pulumi.ToStringArray([]string{backupStore.ID()}),
		}, pulumi.DependsOn([]pulumi.Resource{backupStore}))
		if err != nil {
			return err
		}

		// Database
		database, err := pgedge.NewDatabase(ctx, "exampleDatabase", &pgedge.DatabaseArgs{
			Name:          pulumi.String("example"),
			ClusterId:     cluster.ID(),
			// ConfigVersion: pulumi.String("12.8.3"),
			Options:       pulumi.ToStringArray([]string{"install:northwind", "rest:enabled", "autoddl:enabled"}),
			Extensions: &pgedge.DatabaseExtensionsArgs{
				AutoManage: pulumi.Bool(true),
				Requesteds: pulumi.ToStringArray([]string{"postgis"}),
			},
			Nodes: pgedge.DatabaseNodesMap{
				"n1": &pgedge.DatabaseNodesArgs{
					Name:       pulumi.String("n1"),
				},
				"n2": &pgedge.DatabaseNodesArgs{
					Name:       pulumi.String("n1"),
				},
				"n3": &pgedge.DatabaseNodesArgs{
					Name:       pulumi.String("n3"),
				},
			},
			Backups: &pgedge.DatabaseBackupsArgs{
				Provider: pulumi.String("pgbackrest"),
				Configs:  pgedge.DatabaseBackupsConfigArray{&pgedge.DatabaseBackupsConfigArgs{
					Id: pulumi.String("default"),
					NodeName: pulumi.String("n1"),
					Schedules: pgedge.DatabaseBackupsConfigScheduleArray{&pgedge.DatabaseBackupsConfigScheduleArgs{
						Id: pulumi.String("daily-full-backup"),
						CronExpression: pulumi.String("15 * * * *"),
						Type: pulumi.String("full"),
					}},
				}},
			},
		}, pulumi.DependsOn([]pulumi.Resource{cluster}))
		if err != nil {
			return err
		}

		// Outputs
		ctx.Export("sshKeyId", sshKey.ID())
		ctx.Export("backupStoreId", backupStore.ID())
		ctx.Export("cloudAccountId", cloudAccount.ID())
		ctx.Export("clusterId", cluster.ID())
		ctx.Export("clusterStatus", cluster.Status)
		ctx.Export("clusterCreatedAt", cluster.CreatedAt)
		ctx.Export("databaseStatus", database.Status)

		// Log outputs
		sshKey.ID().ApplyT(func(id string) error {
			fmt.Printf("SSH Key ID: %s\n", id)
			return nil
		})
		backupStore.ID().ApplyT(func(id string) error {
			fmt.Printf("Backup Store ID: %s\n", id)
			return nil
		})
		cloudAccount.ID().ApplyT(func(id string) error {
			fmt.Printf("Cloud Account ID: %s\n", id)
			return nil
		})
		cluster.ID().ApplyT(func(id string) error {
			fmt.Printf("Cluster ID: %s\n", id)
			return nil
		})
		cluster.Status.ApplyT(func(status string) error {
			fmt.Printf("Cluster Status: %s\n", status)
			return nil
		})
		cluster.CreatedAt.ApplyT(func(createdAt string) error {
			fmt.Printf("Cluster Created At: %s\n", createdAt)
			return nil
		})
		database.Status.ApplyT(func(status string) error {
			fmt.Printf("Database Status: %s\n", status)
			return nil
		})

		return nil
	})
}