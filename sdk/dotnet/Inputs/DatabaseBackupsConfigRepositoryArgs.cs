// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pgedge.Pgedge.Inputs
{

    public sealed class DatabaseBackupsConfigRepositoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Azure account.
        /// </summary>
        [Input("azureAccount")]
        public Input<string>? AzureAccount { get; set; }

        /// <summary>
        /// Azure container.
        /// </summary>
        [Input("azureContainer")]
        public Input<string>? AzureContainer { get; set; }

        /// <summary>
        /// Azure endpoint.
        /// </summary>
        [Input("azureEndpoint")]
        public Input<string>? AzureEndpoint { get; set; }

        /// <summary>
        /// ID of the backup store.
        /// </summary>
        [Input("backupStoreId")]
        public Input<string>? BackupStoreId { get; set; }

        /// <summary>
        /// Base path for the repository.
        /// </summary>
        [Input("basePath")]
        public Input<string>? BasePath { get; set; }

        /// <summary>
        /// GCS bucket name.
        /// </summary>
        [Input("gcsBucket")]
        public Input<string>? GcsBucket { get; set; }

        /// <summary>
        /// GCS endpoint.
        /// </summary>
        [Input("gcsEndpoint")]
        public Input<string>? GcsEndpoint { get; set; }

        /// <summary>
        /// ID of the database
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Retention period for full backups.
        /// </summary>
        [Input("retentionFull")]
        public Input<int>? RetentionFull { get; set; }

        /// <summary>
        /// Type of retention for full backups.
        /// </summary>
        [Input("retentionFullType")]
        public Input<string>? RetentionFullType { get; set; }

        /// <summary>
        /// S3 bucket name.
        /// </summary>
        [Input("s3Bucket")]
        public Input<string>? S3Bucket { get; set; }

        /// <summary>
        /// S3 endpoint.
        /// </summary>
        [Input("s3Endpoint")]
        public Input<string>? S3Endpoint { get; set; }

        /// <summary>
        /// S3 region.
        /// </summary>
        [Input("s3Region")]
        public Input<string>? S3Region { get; set; }

        /// <summary>
        /// Type of the repository.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DatabaseBackupsConfigRepositoryArgs()
        {
        }
        public static new DatabaseBackupsConfigRepositoryArgs Empty => new DatabaseBackupsConfigRepositoryArgs();
    }
}