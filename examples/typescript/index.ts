import * as pgedge from "@pgEdge/pulumi-pgedge";

var cluster = new pgedge.Cluster("clusterCreate",{
  name: "n1",
  cloudAccountId: "", // cloud account id
  regions:["us-east-2"],
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