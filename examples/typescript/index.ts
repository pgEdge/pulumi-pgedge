import * as pulumi from "@pulumi/pulumi";
import * as pgedge from "@pgEdge/pulumi-pgedge";

const config = new pulumi.Config();
const baseUrl = config.get("baseUrl") || "https://api.pgedge.com";

const sshKey = new pgedge.SSHKey("exampleSSHKey", {
  name: "example",
  publicKey: "ssh-ed25519 AAAAC3NzaC1lZsdw877237ICXfT63i04t5fvvlGesddwed21VG7DkyxvyXbYQNhKP/rSeLY user@example.com",
});

const cloudAccount = new pgedge.CloudAccount("exampleCloudAccount", {
  name: "my-aws-account",
  type: "aws",
  description: "My AWS Cloud Account",
  credentials: {
    role_arn: "arn:aws:iam::21112529deae39:role/pgedge-135232c",
  },
}, { dependsOn: sshKey });

const backupStore = new pgedge.BackupStore("exampleBackupStore", {
  name: "example",
  cloudAccountId: cloudAccount.id,
  region: "us-west-2",
}, { dependsOn: cloudAccount });

const cluster = new pgedge.Cluster("exampleCluster", {
  name: "example",
  cloudAccountId: cloudAccount.id,
  regions: ["us-west-2", "us-east-1", "eu-central-1"],
  nodeLocation: "public",
  sshKeyId: sshKey.id,
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
  // backupStoreIds: [backupStore.id],
  firewallRules: [
    {
      name: "postgres",
      port: 5432,
      sources: ["123.456.789.0/32"],
    },
  ],
}, { dependsOn: backupStore });

const database = new pgedge.Database("exampleDatabase", {
  name: "example",
  clusterId: cluster.id,
  // configVersion: "12.8.3",
  options: [
    "install:northwind",
    "rest:enabled",
    "autoddl:enabled",
  ],
  extensions: {
    autoManage: true,
    requesteds: [
      "postgis",
      "vector"
    ],
  },
  nodes: [
    { name: "n1" },
    { name: "n2" },
    { name: "n3" },
  ],
  backups: {
    provider: "pgbackrest",
    configs: [
      {
        id: "default",
        nodeName: "n1",
        schedules: [
          {
            id: "daily",
            cronExpression: "15 * * * ",
            type: "daily",
          },
        ]
      }
    ]
  },
}, { dependsOn: cluster });

// Outputs
export const sshKeyId = sshKey.id;
export const backupStoreId = backupStore.id;
export const cloudAccountId = cloudAccount.id;
export const clusterId = cluster.id;
export const clusterStatus = cluster.status;
export const clusterCreatedAt = cluster.createdAt;
export const databaseId = database.id;

// Log outputs
sshKeyId.apply(id => console.log(`SSH Key ID: ${id}`));
backupStoreId.apply(id => console.log(`Backup Store ID: ${id}`));
cloudAccountId.apply(id => console.log(`Cloud Account ID: ${id}`));
clusterId.apply(id => console.log(`Cluster ID: ${id}`));
clusterStatus.apply(status => console.log(`Cluster Status: ${status}`));
clusterCreatedAt.apply(createdAt => console.log(`Cluster Created At: ${createdAt}`));
databaseId.apply(id => console.log(`Database id: ${id}`));