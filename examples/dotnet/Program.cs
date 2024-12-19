using System.Collections.Generic;
using Pulumi;
using Pgedge.Pgedge;


return await Deployment.RunAsync(() =>
{
    var config = new Pulumi.Config();
    var baseUrl = config.Get("baseUrl") ?? "https://api.pgedge.com";

    // Create SSH Key
    var sshKey = new SSHKey("exampleSSHKey", new SSHKeyArgs
    {
        Name = "example",
        PublicKey = "ssh-ed25519 AAAAC3NzaC1lZsdw877237ICXfT63i04t5fvvlGesddwed21VG7DkyxvyXbYQNhKP/rSeLY user@example.com"
    });

    // Create Cloud Account
    var cloudAccount = new CloudAccount("exampleCloudAccount", new CloudAccountArgs
    {
        Name = "my-aws-account",
        Type = "aws",
        Description = "My AWS Cloud Account",
        Credentials = new Dictionary<string, string>
        {
            { "role_arn", "arn:aws:iam::21112529deae39:role/pgedge-135232c" }
        }
    }, new CustomResourceOptions { DependsOn = { sshKey } });

    // Create Backup Store
    var backupStore = new BackupStore("testBackupStore", new BackupStoreArgs
    {
        Name = "example",
        CloudAccountId = cloudAccount.Id,
        Region = "us-west-2"
    }, new CustomResourceOptions { DependsOn = { cloudAccount } });

    // Create Cluster
    var cluster = new Cluster("exampleCluster", new ClusterArgs
    {
        Name = "example",
        CloudAccountId = cloudAccount.Id,
        Regions = new[] { "us-west-2", "us-east-1", "eu-central-1" },
        NodeLocation = "public",
        SshKeyId = sshKey.Id,
        Nodes = new InputList<Pgedge.Pgedge.Inputs.ClusterNodeArgs>
    {
        new Pgedge.Pgedge.Inputs.ClusterNodeArgs
        {
            Name = "n1",
            Region = "us-west-2",
            InstanceType = "r6g.medium",
            VolumeSize = 100,
            VolumeType = "gp2"
        },
        new Pgedge.Pgedge.Inputs.ClusterNodeArgs
        {
            Name = "n2",
            Region = "us-east-1",
            InstanceType = "r6g.medium",
            VolumeSize = 100,
            VolumeType = "gp2"
        },
        new Pgedge.Pgedge.Inputs.ClusterNodeArgs
        {
            Name = "n3",
            Region = "eu-central-1",
            InstanceType = "r6g.medium",
            VolumeSize = 100,
            VolumeType = "gp2"
        }
    },
        Networks = new InputList<Pgedge.Pgedge.Inputs.ClusterNetworkArgs>
    {
        new Pgedge.Pgedge.Inputs.ClusterNetworkArgs
        {
            Region = "us-west-2",
            Cidr = "10.1.0.0/16",
            PublicSubnets = new[] { "10.1.0.0/24" }
            // PrivateSubnets = new[] { "10.1.1.0/24" }
        },
        new Pgedge.Pgedge.Inputs.ClusterNetworkArgs
        {
            Region = "us-east-1",
            Cidr = "10.2.0.0/16",
            PublicSubnets = new[] { "10.2.0.0/24" }
            // PrivateSubnets = new[] { "10.2.1.0/24" }
        },
        new Pgedge.Pgedge.Inputs.ClusterNetworkArgs
        {
            Region = "eu-central-1",
            Cidr = "10.3.0.0/16",
            PublicSubnets = new[] { "10.3.0.0/24" }
            // PrivateSubnets = new[] { "10.3.1.0/24" }
        }
    },
        FirewallRules = new InputList<Pgedge.Pgedge.Inputs.ClusterFirewallRuleArgs>
    {
        new Pgedge.Pgedge.Inputs.ClusterFirewallRuleArgs
        {
            Name = "postgres",
            Port = 5432,
            Sources = new[] { "0.0.0.0/0" }
        }
    }
    }, new CustomResourceOptions { DependsOn = { backupStore } });


    // Create Database
    var database = new Database("exampleDatabase", new DatabaseArgs
    {
        Name = "example",
        ClusterId = cluster.Id,
        Options = new[] { "install:northwind", "rest:enabled" },
        Nodes = new InputMap<Pgedge.Pgedge.Inputs.DatabaseNodesArgs>
    {
        { "n1", new Pgedge.Pgedge.Inputs.DatabaseNodesArgs { Name = "n1" } },
        { "n2", new Pgedge.Pgedge.Inputs.DatabaseNodesArgs { Name = "n2" } },
        { "n3", new Pgedge.Pgedge.Inputs.DatabaseNodesArgs { Name = "n3" } }
    },
        Backups = new Pgedge.Pgedge.Inputs.DatabaseBackupsArgs
        {
            Provider = "pgbackrest",
            Configs = new[]
        {
            new Pgedge.Pgedge.Inputs.DatabaseBackupsConfigArgs
            {
                Id = "default",
                NodeName = "n1",
                Schedules = new[]
                {
                    new Pgedge.Pgedge.Inputs.DatabaseBackupsConfigScheduleArgs
                    {
                        Id = "daily-full-backup",
                        CronExpression = "0 6 * * ?",
                        Type = "full"
                    }
                },
                Repositories = new[]
                {
                    new Pgedge.Pgedge.Inputs.DatabaseBackupsConfigRepositoryArgs
                    {
                        BackupStoreId = backupStore.Id,
                        RetentionFull = 7
                    }
                }
            }
        }
        }
    }, new CustomResourceOptions { DependsOn = { cluster } });


    // Export values
    return new Dictionary<string, object?>
    {
        ["sshKeyId"] = sshKey.Id,
        ["backupStoreId"] = backupStore.Id,
        ["cloudAccountId"] = cloudAccount.Id,
        ["clusterId"] = cluster.Id,
        ["clusterStatus"] = cluster.Status,
        ["clusterCreatedAt"] = cluster.CreatedAt,
        ["databaseId"] = database.Id,
        ["databaseNodes"] = database.Nodes,
        ["configVersion"] = database.ConfigVersion,
        ["databaseBackups"] = database.Backups,
        ["databaseBackupsConfigs"] = database.Backups.Apply(backups => Output.Create(backups.Configs)),
        ["databaseExtensions"] = database.Extensions,
        ["databaseOptions"] = database,
        ["databaseDomain"] = database.Domain
    };
});