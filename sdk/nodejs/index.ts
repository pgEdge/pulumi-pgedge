// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { BackupStoreArgs, BackupStoreState } from "./backupStore";
export type BackupStore = import("./backupStore").BackupStore;
export const BackupStore: typeof import("./backupStore").BackupStore = null as any;
utilities.lazyLoad(exports, ["BackupStore"], () => require("./backupStore"));

export { CloudAccountArgs, CloudAccountState } from "./cloudAccount";
export type CloudAccount = import("./cloudAccount").CloudAccount;
export const CloudAccount: typeof import("./cloudAccount").CloudAccount = null as any;
utilities.lazyLoad(exports, ["CloudAccount"], () => require("./cloudAccount"));

export { ClusterArgs, ClusterState } from "./cluster";
export type Cluster = import("./cluster").Cluster;
export const Cluster: typeof import("./cluster").Cluster = null as any;
utilities.lazyLoad(exports, ["Cluster"], () => require("./cluster"));

export { DatabaseArgs, DatabaseState } from "./database";
export type Database = import("./database").Database;
export const Database: typeof import("./database").Database = null as any;
utilities.lazyLoad(exports, ["Database"], () => require("./database"));

export { GetBackupStoresResult } from "./getBackupStores";
export const getBackupStores: typeof import("./getBackupStores").getBackupStores = null as any;
export const getBackupStoresOutput: typeof import("./getBackupStores").getBackupStoresOutput = null as any;
utilities.lazyLoad(exports, ["getBackupStores","getBackupStoresOutput"], () => require("./getBackupStores"));

export { GetCloudAccountsResult } from "./getCloudAccounts";
export const getCloudAccounts: typeof import("./getCloudAccounts").getCloudAccounts = null as any;
export const getCloudAccountsOutput: typeof import("./getCloudAccounts").getCloudAccountsOutput = null as any;
utilities.lazyLoad(exports, ["getCloudAccounts","getCloudAccountsOutput"], () => require("./getCloudAccounts"));

export { GetClustersResult } from "./getClusters";
export const getClusters: typeof import("./getClusters").getClusters = null as any;
export const getClustersOutput: typeof import("./getClusters").getClustersOutput = null as any;
utilities.lazyLoad(exports, ["getClusters","getClustersOutput"], () => require("./getClusters"));

export { GetDatabasesResult } from "./getDatabases";
export const getDatabases: typeof import("./getDatabases").getDatabases = null as any;
export const getDatabasesOutput: typeof import("./getDatabases").getDatabasesOutput = null as any;
utilities.lazyLoad(exports, ["getDatabases","getDatabasesOutput"], () => require("./getDatabases"));

export { GetSSHKeysResult } from "./getSSHKeys";
export const getSSHKeys: typeof import("./getSSHKeys").getSSHKeys = null as any;
export const getSSHKeysOutput: typeof import("./getSSHKeys").getSSHKeysOutput = null as any;
utilities.lazyLoad(exports, ["getSSHKeys","getSSHKeysOutput"], () => require("./getSSHKeys"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { SSHKeyArgs, SSHKeyState } from "./sshkey";
export type SSHKey = import("./sshkey").SSHKey;
export const SSHKey: typeof import("./sshkey").SSHKey = null as any;
utilities.lazyLoad(exports, ["SSHKey"], () => require("./sshkey"));


// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "pgedge:index/backupStore:BackupStore":
                return new BackupStore(name, <any>undefined, { urn })
            case "pgedge:index/cloudAccount:CloudAccount":
                return new CloudAccount(name, <any>undefined, { urn })
            case "pgedge:index/cluster:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "pgedge:index/database:Database":
                return new Database(name, <any>undefined, { urn })
            case "pgedge:index/sSHKey:SSHKey":
                return new SSHKey(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("pgedge", "index/backupStore", _module)
pulumi.runtime.registerResourceModule("pgedge", "index/cloudAccount", _module)
pulumi.runtime.registerResourceModule("pgedge", "index/cluster", _module)
pulumi.runtime.registerResourceModule("pgedge", "index/database", _module)
pulumi.runtime.registerResourceModule("pgedge", "index/sSHKey", _module)
pulumi.runtime.registerResourcePackage("pgedge", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:pgedge") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
