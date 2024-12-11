import pulumi
from pgEdge_pulumi_pgedge import (
    SSHKey, 
    CloudAccount, 
    BackupStore, 
    Cluster, 
    Database,
    CloudAccountArgs,
    BackupStoreArgs,
    ClusterArgs,
    DatabaseArgs,
    SSHKeyArgs
)

# Get config
config = pulumi.Config()
base_url = config.get("baseUrl") or "https://api.pgedge.com"

# Create SSH Key
ssh_key = SSHKey("exampleSSHKey",
    SSHKeyArgs(
        name="example",
        public_key="ssh-ed25519 AAAAC3NzaC1lZsdw877237ICXfT63i04t5fvvlGesddwed21VG7DkyxvyXbYQNhKP/rSeLY user@example.com"
    )
)

# Create Cloud Account
cloud_account = CloudAccount("exampleCloudAccount",
    CloudAccountArgs(
        name="my-aws-account",
        type="aws",
        description="My AWS Cloud Account",
        credentials={
            "role_arn": "arn:aws:iam::21112529deae39:role/pgedge-135232c"
        }
    ),
    opts=pulumi.ResourceOptions(depends_on=[ssh_key])
)

# Create Backup Store
backup_store = BackupStore("testBackupStore",
    BackupStoreArgs(
        name="example",
        cloud_account_id=cloud_account.id,
        region="us-west-2"
    ),
    opts=pulumi.ResourceOptions(depends_on=[cloud_account])
)

# Create Cluster
cluster = Cluster("exampleCluster",
    ClusterArgs(
        name="example",
        cloud_account_id=cloud_account.id,
        regions=["us-west-2", "us-east-1", "us-east-1"],
        node_location="public",
        ssh_key_id=ssh_key.id,
        nodes=[
            {
                "name": "n1",
                "region": "us-west-2",
                "instance_type": "r6g.medium",
                "volume_size": 100,
                "volume_type": "gp2"
            },
            {
                "name": "n2",
                "region": "us-east-1",
                "instance_type": "r6g.medium",
                "volume_size": 100,
                "volume_type": "gp2"
            },
            {
                "name": "n3",
                "region": "us-east-1",
                "instance_type": "r6g.medium",
                "volume_size": 100,
                "volume_type": "gp2"
            }
        ],
        networks=[
            {
                "region": "us-west-2",
                "cidr": "10.1.0.0/16",
                "public_subnets": ["10.1.0.0/24"]
            },
            {
                "region": "us-east-1",
                "cidr": "10.2.0.0/16",
                "public_subnets": ["10.2.0.0/24"]
            },
            {
                "region": "us-east-1",
                "cidr": "10.3.0.0/16",
                "public_subnets": ["10.3.0.0/24"]
            }
        ],
        backup_store_ids=[backup_store.id],
        firewall_rules=[
            {
                "name": "postgres",
                "port": 5432,
                "sources": ["0.0.0.0/0"]
            }
        ]
    ),
    opts=pulumi.ResourceOptions(depends_on=[backup_store])
)

# Create Database
database = Database("exampleDatabase",
    DatabaseArgs(
        name="example",
        cluster_id=cluster.id,
        options=["install:northwind", "rest:enabled"],
        nodes={
            "n1": {
                "name": "n1"
            },
            "n2": {
                "name": "n2"
            },
            "n3": {
                "name": "n3"
            }
        },
        backups={
            "provider": "pgbackrest",
            "configs": [{
                "id": "default",
                "node_name": "n1",
                "schedules": [{
                    "id": "daily-full-backup",
                    "cron_expression": "0 6 * * ?",
                    "type": "full"
                }],
                "repositories": [{
                    "backup_store_id": backup_store.id,
                    "retention_full": 7
                }]
            }]
        }
    ),
    opts=pulumi.ResourceOptions(depends_on=[cluster])
)

# Export values
pulumi.export("sshKeyId", ssh_key.id)
pulumi.export("backupStoreId", backup_store.id)
pulumi.export("cloudAccountId", cloud_account.id)
pulumi.export("clusterId", cluster.id)
pulumi.export("clusterStatus", cluster.status)
pulumi.export("clusterCreatedAt", cluster.created_at)
pulumi.export("databaseId", database.id)
pulumi.export("databaseNodes", database.nodes)
pulumi.export("configVersion", database.config_version)
pulumi.export("databaseBackups", database.backups)
pulumi.export("databaseBackupsConfigs", database.backups["configs"])
pulumi.export("databaseBackupsSchedules", database.backups["configs"][0]["schedules"])
pulumi.export("databaseExtensions", database.extensions)
pulumi.export("databaseOptions", database)
pulumi.export("databaseDomain", database.domain)