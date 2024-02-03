import * as pulumi from "@pulumi/pulumi";
import * as pgedge from "@pgEdge/pulumi-pgedge";

const databases = pgedge.getDatabases()

new pgedge.Database("databaseCreate", {
    name: "",
    clusterId: "",
    options: [""],
})

export const databaseNames = databases.then(dbs => dbs.databases.map(db => db.name))