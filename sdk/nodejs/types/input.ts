// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ClusterFirewallRule {
    name: pulumi.Input<string>;
    port: pulumi.Input<number>;
    sources: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ClusterNetwork {
    /**
     * CIDR of the network
     */
    cidr: pulumi.Input<string>;
    /**
     * Whether the network is external
     */
    external?: pulumi.Input<boolean>;
    /**
     * External ID of the network
     */
    externalId?: pulumi.Input<string>;
    /**
     * Name of the network
     */
    name?: pulumi.Input<string>;
    /**
     * List of private subnets
     */
    privateSubnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of public subnets
     */
    publicSubnets: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Region of the network
     */
    region: pulumi.Input<string>;
}

export interface ClusterNode {
    availabilityZone?: pulumi.Input<string>;
    instanceType: pulumi.Input<string>;
    name: pulumi.Input<string>;
    region: pulumi.Input<string>;
    volumeIops?: pulumi.Input<number>;
    volumeSize?: pulumi.Input<number>;
    volumeType?: pulumi.Input<string>;
}

export interface DatabaseBackups {
    /**
     * List of backup configurations.
     */
    configs?: pulumi.Input<pulumi.Input<inputs.DatabaseBackupsConfig>[]>;
    /**
     * The backup provider.
     */
    provider?: pulumi.Input<string>;
}

export interface DatabaseBackupsConfig {
    /**
     * Unique identifier for the backup config.
     */
    id?: pulumi.Input<string>;
    /**
     * Name of the node.
     */
    nodeName?: pulumi.Input<string>;
    /**
     * List of backup repositories.
     */
    repositories?: pulumi.Input<pulumi.Input<inputs.DatabaseBackupsConfigRepository>[]>;
    /**
     * List of backup schedules.
     */
    schedules?: pulumi.Input<pulumi.Input<inputs.DatabaseBackupsConfigSchedule>[]>;
}

export interface DatabaseBackupsConfigRepository {
    /**
     * Azure account for azure-type repositories.
     */
    azureAccount?: pulumi.Input<string>;
    /**
     * Azure container for azure-type repositories.
     */
    azureContainer?: pulumi.Input<string>;
    /**
     * Azure endpoint for azure-type repositories.
     */
    azureEndpoint?: pulumi.Input<string>;
    /**
     * ID of the backup store to use. If specified, other fields will be automatically populated.
     */
    backupStoreId?: pulumi.Input<string>;
    /**
     * Base path for the repository.
     */
    basePath?: pulumi.Input<string>;
    /**
     * GCS bucket name for gcs-type repositories.
     */
    gcsBucket?: pulumi.Input<string>;
    /**
     * GCS endpoint for gcs-type repositories.
     */
    gcsEndpoint?: pulumi.Input<string>;
    /**
     * Unique identifier for the backup config.
     */
    id?: pulumi.Input<string>;
    /**
     * Retention period for full backups.
     */
    retentionFull?: pulumi.Input<number>;
    /**
     * Type of retention for full backups.
     */
    retentionFullType?: pulumi.Input<string>;
    /**
     * S3 bucket name for s3-type repositories.
     */
    s3Bucket?: pulumi.Input<string>;
    /**
     * S3 endpoint for s3-type repositories.
     */
    s3Endpoint?: pulumi.Input<string>;
    /**
     * S3 region for s3-type repositories.
     */
    s3Region?: pulumi.Input<string>;
    /**
     * Repository type (e.g., s3, gcs, azure).
     */
    type?: pulumi.Input<string>;
}

export interface DatabaseBackupsConfigSchedule {
    /**
     * Cron expression for the schedule.
     */
    cronExpression: pulumi.Input<string>;
    /**
     * Unique identifier for the backup config.
     */
    id: pulumi.Input<string>;
    /**
     * Repository type (e.g., s3, gcs, azure).
     */
    type: pulumi.Input<string>;
}

export interface DatabaseComponent {
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    releaseDate?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    version?: pulumi.Input<string>;
}

export interface DatabaseExtensions {
    autoManage?: pulumi.Input<boolean>;
    availables?: pulumi.Input<pulumi.Input<string>[]>;
    requesteds?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface DatabaseNodes {
    connection?: pulumi.Input<inputs.DatabaseNodesConnection>;
    extensions?: pulumi.Input<inputs.DatabaseNodesExtensions>;
    location?: pulumi.Input<inputs.DatabaseNodesLocation>;
    name: pulumi.Input<string>;
    region?: pulumi.Input<inputs.DatabaseNodesRegion>;
}

export interface DatabaseNodesConnection {
    database?: pulumi.Input<string>;
    externalIpAddress?: pulumi.Input<string>;
    host?: pulumi.Input<string>;
    internalHost?: pulumi.Input<string>;
    internalIpAddress?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    username?: pulumi.Input<string>;
}

export interface DatabaseNodesExtensions {
    errors?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    installeds?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface DatabaseNodesLocation {
    city?: pulumi.Input<string>;
    code?: pulumi.Input<string>;
    country?: pulumi.Input<string>;
    latitude?: pulumi.Input<number>;
    longitude?: pulumi.Input<number>;
    metroCode?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    postalCode?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    regionCode?: pulumi.Input<string>;
    timezone?: pulumi.Input<string>;
}

export interface DatabaseNodesRegion {
    active?: pulumi.Input<boolean>;
    availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    cloud?: pulumi.Input<string>;
    code?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    parent?: pulumi.Input<string>;
}

export interface DatabaseRole {
    bypassRls?: pulumi.Input<boolean>;
    connectionLimit?: pulumi.Input<number>;
    createDb?: pulumi.Input<boolean>;
    createRole?: pulumi.Input<boolean>;
    inherit?: pulumi.Input<boolean>;
    login?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    replication?: pulumi.Input<boolean>;
    superuser?: pulumi.Input<boolean>;
}

