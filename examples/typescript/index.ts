import * as pulumi from "@pulumi/pulumi";
import * as pgedge from "@pgEdge/pulumi-pgedge";

const config = new pulumi.Config();
const baseUrl = config.get("baseUrl") || "https://api.pgedge.com";

const sshKey = new pgedge.SSHKey("exampleSSHKey", {
  name: "examplessh72",
  publicKey: "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAICXfT63i04t5fvvlGeoUoVG71123mvbbwFDkyxvyXbYQNhKP/rSeLY your@your-MacBook-Pro.local",
});

const cloudAccount = new pgedge.CloudAccount("exampleCloudAccount", {
  name: "my-aws-account",
  type: "aws",
  description: "My AWS Cloud Account",
  credentials: {
    role_arn: "arn:aws:iam::1123122:role/pgedge-123wvv",
  },
}, { dependsOn: sshKey });

const backupStore = new pgedge.BackupStore("testBackupStore", {
  name: "testst117ore",
  cloudAccountId: cloudAccount.id,
  region: "ap-northeast-1",
}, { dependsOn: cloudAccount });

const cluster = new pgedge.Cluster("exampleCluster", {
  name: "finalcluster4",
  cloudAccountId: cloudAccount.id,
  regions: ["ap-northeast-1", "ap-northeast-3", "ap-northeast-2"],
  nodeLocation: "public",
  sshKeyId: sshKey.id,
  nodes: [
    {
      name: "n1",
      region: "ap-northeast-1",
      instanceType: "t4g.medium",
      volumeSize: 20,
      volumeType: "gp2",
    },
    {
      name: "n2",
      region: "ap-northeast-3",
      instanceType: "t4g.medium",
      volumeSize: 20,
      volumeType: "gp2",
    },
    {
      name: "n3",
      region: "ap-northeast-2",
      instanceType: "t4g.medium",
      volumeSize: 20,
      volumeType: "gp2",
    },
  ],
  networks: [
    {
      region: "ap-northeast-1",
      cidr: "10.1.0.0/16",
      publicSubnets: ["10.1.0.0/24"],
    },
    {
      region: "ap-northeast-3",
      cidr: "10.2.0.0/16",
      publicSubnets: ["10.2.0.0/24"],
    },
    {
      region: "ap-northeast-2",
      cidr: "10.3.0.0/16",
      publicSubnets: ["10.3.0.0/24"],
    },
  ],
  // backupStoreIds: [backupStore.id],
  firewallRules: [
    {
      name: "postgres",
      port: 5432,
      sources: ["0.0.0.0/0"],
    },
  ],
}, { dependsOn: cloudAccount });

const database = new pgedge.Database("exampleDatabase", {
  name: "exampledb12",
  clusterId: cluster.id,
  options: [
    "install:northwind",
    "rest:enabled",
    "autoddl:enabled",
  ],
  extensions: {
    autoManage: true,
    requesteds: [
      "postgis",
    ],
  },
  nodes: [
    { name: "n1" },
    { name: "n2" },
    { name: "n3" },
  ],
  backups: {
    provider: "pgbackrest",
  },
}, { dependsOn: cluster });

// Outputs
export const sshKeyId = sshKey.id;
export const backupStoreId = backupStore.id;
export const cloudAccountId = cloudAccount.id;
export const clusterId = cluster.id;
export const clusterStatus = cluster.status;
export const clusterCreatedAt = cluster.createdAt;
export const databaseStatus = database.status;

// Helper function to resolve Output<string> to string
function resolveOutput(output: pulumi.Output<string>): pulumi.Output<string> {
  return output.apply(v => v);
}

// Log outputs
sshKeyId.apply(id => console.log(`SSH Key ID: ${id}`));
backupStoreId.apply(id => console.log(`Backup Store ID: ${id}`));
cloudAccountId.apply(id => console.log(`Cloud Account ID: ${id}`));
clusterId.apply(id => console.log(`Cluster ID: ${id}`));
clusterStatus.apply(status => console.log(`Cluster Status: ${status}`));
clusterCreatedAt.apply(createdAt => console.log(`Cluster Created At: ${createdAt}`));
databaseStatus.apply(status => console.log(`Database Status: ${status}`));