// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ClusterFirewallRule {
    name: string;
    port: number;
    sources: string[];
}

export interface ClusterNetwork {
    /**
     * CIDR of the network
     */
    cidr: string;
    /**
     * Whether the network is external
     */
    external: boolean;
    /**
     * External ID of the network
     */
    externalId: string;
    /**
     * Name of the network
     */
    name: string;
    /**
     * List of private subnets
     */
    privateSubnets: string[];
    /**
     * List of public subnets
     */
    publicSubnets: string[];
    /**
     * Region of the network
     */
    region: string;
}

export interface ClusterNode {
    availabilityZone: string;
    instanceType: string;
    name: string;
    region: string;
    volumeIops: number;
    volumeSize: number;
    volumeType: string;
}

export interface DatabaseBackups {
    /**
     * List of backup configurations.
     */
    configs: outputs.DatabaseBackupsConfig[];
    /**
     * The backup provider.
     */
    provider: string;
}

export interface DatabaseBackupsConfig {
    /**
     * Unique identifier for the backup config.
     */
    id: string;
    /**
     * Name of the node.
     */
    nodeName: string;
    /**
     * List of backup repositories.
     */
    repositories: outputs.DatabaseBackupsConfigRepository[];
    /**
     * List of backup schedules.
     */
    schedules?: outputs.DatabaseBackupsConfigSchedule[];
}

export interface DatabaseBackupsConfigRepository {
    /**
     * Azure account for azure-type repositories.
     */
    azureAccount: string;
    /**
     * Azure container for azure-type repositories.
     */
    azureContainer: string;
    /**
     * Azure endpoint for azure-type repositories.
     */
    azureEndpoint: string;
    /**
     * ID of the backup store to use. If specified, other fields will be automatically populated.
     */
    backupStoreId?: string;
    /**
     * Base path for the repository.
     */
    basePath: string;
    /**
     * GCS bucket name for gcs-type repositories.
     */
    gcsBucket: string;
    /**
     * GCS endpoint for gcs-type repositories.
     */
    gcsEndpoint: string;
    /**
     * Unique identifier for the backup config.
     */
    id: string;
    /**
     * Retention period for full backups.
     */
    retentionFull: number;
    /**
     * Type of retention for full backups.
     */
    retentionFullType: string;
    /**
     * S3 bucket name for s3-type repositories.
     */
    s3Bucket: string;
    /**
     * S3 endpoint for s3-type repositories.
     */
    s3Endpoint: string;
    /**
     * S3 region for s3-type repositories.
     */
    s3Region: string;
    /**
     * Repository type (e.g., s3, gcs, azure).
     */
    type: string;
}

export interface DatabaseBackupsConfigSchedule {
    /**
     * Cron expression for the schedule.
     */
    cronExpression: string;
    /**
     * Unique identifier for the backup config.
     */
    id: string;
    /**
     * Repository type (e.g., s3, gcs, azure).
     */
    type: string;
}

export interface DatabaseComponent {
    id: string;
    name: string;
    releaseDate: string;
    status: string;
    version: string;
}

export interface DatabaseExtensions {
    autoManage?: boolean;
    availables: string[];
    requesteds: string[];
}

export interface DatabaseNodes {
    connection: outputs.DatabaseNodesConnection;
    extensions: outputs.DatabaseNodesExtensions;
    location: outputs.DatabaseNodesLocation;
    name: string;
    region: outputs.DatabaseNodesRegion;
}

export interface DatabaseNodesConnection {
    database: string;
    externalIpAddress: string;
    host: string;
    internalHost: string;
    internalIpAddress: string;
    password: string;
    port: number;
    username: string;
}

export interface DatabaseNodesExtensions {
    errors: {[key: string]: string};
    installeds: string[];
}

export interface DatabaseNodesLocation {
    city: string;
    code: string;
    country: string;
    latitude: number;
    longitude: number;
    metroCode: string;
    name: string;
    postalCode: string;
    region: string;
    regionCode: string;
    timezone: string;
}

export interface DatabaseNodesRegion {
    active: boolean;
    availabilityZones: string[];
    cloud: string;
    code: string;
    name: string;
    parent: string;
}

export interface DatabaseRole {
    bypassRls: boolean;
    connectionLimit: number;
    createDb: boolean;
    createRole: boolean;
    inherit: boolean;
    login: boolean;
    name: string;
    replication: boolean;
    superuser: boolean;
}

export interface GetBackupStoresBackupStore {
    cloudAccountId: string;
    cloudAccountType: string;
    clusterIds: string[];
    createdAt: string;
    id: string;
    name: string;
    properties: {[key: string]: string};
    status: string;
}

export interface GetCloudAccountsCloudAccount {
    /**
     * Creation time of the cloud account
     */
    createdAt: string;
    /**
     * Description of the cloud account
     */
    description: string;
    /**
     * ID of the cloud account
     */
    id: string;
    /**
     * Name of the cloud account
     */
    name: string;
    /**
     * Additional properties of the cloud account
     */
    properties: {[key: string]: string};
    /**
     * Type of the cloud account (e.g., AWS, Azure, GCP)
     */
    type: string;
}

export interface GetClustersCluster {
    /**
     * Backup store IDs of the cluster
     */
    backupStoreIds: string[];
    /**
     * Capacity of the cluster
     */
    capacity: number;
    /**
     * Cloud account ID of the cluster
     */
    cloudAccountId: string;
    /**
     * Created at of the cluster
     */
    createdAt: string;
    firewallRules: outputs.GetClustersClusterFirewallRule[];
    /**
     * ID of the cluster
     */
    id: string;
    /**
     * Name of the cluster
     */
    name: string;
    networks: outputs.GetClustersClusterNetwork[];
    /**
     * Node location of the cluster. Must be either 'public' or 'private'.
     */
    nodeLocation: string;
    nodes: outputs.GetClustersClusterNode[];
    regions: string[];
    /**
     * Resource tags of the cluster
     */
    resourceTags: {[key: string]: string};
    /**
     * SSH key ID of the cluster
     */
    sshKeyId: string;
    /**
     * Status of the cluster
     */
    status: string;
}

export interface GetClustersClusterFirewallRule {
    /**
     * Name of the firewall rule
     */
    name: string;
    /**
     * Port for the firewall rule
     */
    port: number;
    /**
     * Sources for the firewall rule
     */
    sources: string[];
}

export interface GetClustersClusterNetwork {
    /**
     * CIDR of the network
     */
    cidr: string;
    /**
     * Is the network external
     */
    external: boolean;
    /**
     * External ID of the network
     */
    externalId: string;
    /**
     * Name of the firewall rule
     */
    name: string;
    privateSubnets: string[];
    publicSubnets: string[];
    /**
     * Region of the network
     */
    region: string;
}

export interface GetClustersClusterNode {
    /**
     * Cloud provider availability zone name
     */
    availabilityZone: string;
    /**
     * Instance type used for the node
     */
    instanceType: string;
    /**
     * Name of the firewall rule
     */
    name: string;
    options: string[];
    /**
     * Region of the network
     */
    region: string;
    /**
     * Volume IOPS of the node data volume
     */
    volumeIops: number;
    /**
     * Volume size of the node data volume
     */
    volumeSize: number;
    /**
     * Volume type of the node data volume
     */
    volumeType: string;
}

export interface GetDatabasesDatabase {
    /**
     * Backup configuration for the database
     */
    backups: outputs.GetDatabasesDatabaseBackups;
    /**
     * ID of the cluster this database belongs to
     */
    clusterId: string;
    /**
     * Components of the database
     */
    components: outputs.GetDatabasesDatabaseComponent[];
    /**
     * Configuration version of the database
     */
    configVersion: string;
    /**
     * Creation timestamp of the database
     */
    createdAt: string;
    /**
     * Display name for the database. Maximum length is 25 characters.
     */
    displayName: string;
    /**
     * Domain of the database
     */
    domain: string;
    /**
     * Extensions configuration for the database
     */
    extensions: outputs.GetDatabasesDatabaseExtensions;
    /**
     * ID of the database
     */
    id: string;
    /**
     * Name of the database
     */
    name: string;
    /**
     * Map of nodes in the database
     */
    nodes: {[key: string]: outputs.GetDatabasesDatabaseNodes};
    /**
     * Options for the database
     */
    options: string[];
    /**
     * PostgreSQL version of the database
     */
    pgVersion: string;
    /**
     * Roles in the database
     */
    roles: outputs.GetDatabasesDatabaseRole[];
    /**
     * Status of the database
     */
    status: string;
}

export interface GetDatabasesDatabaseBackups {
    /**
     * Backup configurations
     */
    configs: outputs.GetDatabasesDatabaseBackupsConfig[];
    /**
     * Backup provider
     */
    provider: string;
}

export interface GetDatabasesDatabaseBackupsConfig {
    /**
     * Backup configuration ID
     */
    id: string;
    /**
     * Node name
     */
    nodeName: string;
    /**
     * Backup repositories
     */
    repositories: outputs.GetDatabasesDatabaseBackupsConfigRepository[];
    /**
     * Backup schedules
     */
    schedules: outputs.GetDatabasesDatabaseBackupsConfigSchedule[];
}

export interface GetDatabasesDatabaseBackupsConfigRepository {
    azureAccount: string;
    azureContainer: string;
    azureEndpoint: string;
    backupStoreId: string;
    basePath: string;
    gcsBucket: string;
    gcsEndpoint: string;
    /**
     * Backup configuration ID
     */
    id: string;
    retentionFull: number;
    retentionFullType: string;
    s3Bucket: string;
    s3Endpoint: string;
    s3Region: string;
    type: string;
}

export interface GetDatabasesDatabaseBackupsConfigSchedule {
    cronExpression: string;
    /**
     * Backup configuration ID
     */
    id: string;
    type: string;
}

export interface GetDatabasesDatabaseComponent {
    /**
     * Backup configuration ID
     */
    id: string;
    /**
     * Component name
     */
    name: string;
    /**
     * Component release date
     */
    releaseDate: string;
    /**
     * Component status
     */
    status: string;
    /**
     * Component version
     */
    version: string;
}

export interface GetDatabasesDatabaseExtensions {
    /**
     * Auto-manage extensions
     */
    autoManage: boolean;
    /**
     * Available extensions
     */
    availables: string[];
    /**
     * Requested extensions
     */
    requesteds: string[];
}

export interface GetDatabasesDatabaseNodes {
    /**
     * Node connection details
     */
    connection: outputs.GetDatabasesDatabaseNodesConnection;
    /**
     * Extensions configuration for the database
     */
    extensions: outputs.GetDatabasesDatabaseNodesExtensions;
    /**
     * Node location
     */
    location: outputs.GetDatabasesDatabaseNodesLocation;
    /**
     * Component name
     */
    name: string;
    /**
     * Node region
     */
    region: outputs.GetDatabasesDatabaseNodesRegion;
}

export interface GetDatabasesDatabaseNodesConnection {
    database: string;
    externalIpAddress: string;
    host: string;
    internalHost: string;
    internalIpAddress: string;
    password: string;
    port: number;
    username: string;
}

export interface GetDatabasesDatabaseNodesExtensions {
    errors: {[key: string]: string};
    installeds: string[];
}

export interface GetDatabasesDatabaseNodesLocation {
    city: string;
    code: string;
    country: string;
    latitude: number;
    longitude: number;
    metroCode: string;
    /**
     * Component name
     */
    name: string;
    postalCode: string;
    region: string;
    regionCode: string;
    timezone: string;
}

export interface GetDatabasesDatabaseNodesRegion {
    active: boolean;
    availabilityZones: string[];
    cloud: string;
    code: string;
    /**
     * Component name
     */
    name: string;
    parent: string;
}

export interface GetDatabasesDatabaseRole {
    bypassRls: boolean;
    connectionLimit: number;
    createDb: boolean;
    createRole: boolean;
    inherit: boolean;
    login: boolean;
    /**
     * Component name
     */
    name: string;
    replication: boolean;
    superuser: boolean;
}

export interface GetSSHKeysSshKey {
    /**
     * Creation time of the SSH key
     */
    createdAt: string;
    /**
     * ID of the SSH key
     */
    id: string;
    /**
     * Name of the SSH key
     */
    name: string;
    /**
     * Public key
     */
    publicKey: string;
}

