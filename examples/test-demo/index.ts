// import * as pulumi from "@pulumi/pulumi";
import * as pgedge from "@pgEdge/pulumi-pgedge";

const databases = pgedge.getDatabases()

const clusters = pgedge.getClusters()

// new pgedge.Database("databaseCreate", {
//     name: "",
//     clusterId: "",
//     options: [""],
// })

new pgedge.Cluster("clusterCreate", {
  name: "test1069",
  cloudAccountId: "",
  firewalls: [
    {
      type: "https",
      port: 5432,
      sources: ["0.0.0.0/0"],
    }
  ],
  nodeGroups: {
    aws: [
      {
        region: "us-west-2",
        instanceType: "t4g.small",
        availabilityZones: [
          "us-west-2a",
        ],
        nodes: [
          {
            displayName: "n1",
            isActive: true
          }
        ]
      },
    ]
  }
})

export const clusterNames = clusters.then(clusters => clusters.clusters.map(cluster => cluster.name))
// export const databaseNames = databases.then(dbs => dbs.databases.map(db => db.name))