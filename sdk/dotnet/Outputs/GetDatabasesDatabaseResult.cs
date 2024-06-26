// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pgedge.Pgedge.Outputs
{

    [OutputType]
    public sealed class GetDatabasesDatabaseResult
    {
        /// <summary>
        /// Updated at of the database
        /// </summary>
        public readonly string ClusterId;
        public readonly ImmutableArray<Outputs.GetDatabasesDatabaseComponentResult> Components;
        /// <summary>
        /// Config version of the database
        /// </summary>
        public readonly string? ConfigVersion;
        /// <summary>
        /// Created at of the database
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Domain of the database
        /// </summary>
        public readonly string Domain;
        public readonly Outputs.GetDatabasesDatabaseExtensionsResult Extensions;
        /// <summary>
        /// ID of the database
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the database
        /// </summary>
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetDatabasesDatabaseNodeResult> Nodes;
        /// <summary>
        /// Options for creating the database
        /// </summary>
        public readonly ImmutableArray<string> Options;
        /// <summary>
        /// Postgres version of the database
        /// </summary>
        public readonly string PgVersion;
        public readonly ImmutableArray<Outputs.GetDatabasesDatabaseRoleResult> Roles;
        /// <summary>
        /// Status of the database
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Storage used of the database
        /// </summary>
        public readonly int StorageUsed;
        public readonly ImmutableArray<Outputs.GetDatabasesDatabaseTableResult> Tables;
        /// <summary>
        /// Updated at of the database
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetDatabasesDatabaseResult(
            string clusterId,

            ImmutableArray<Outputs.GetDatabasesDatabaseComponentResult> components,

            string? configVersion,

            string createdAt,

            string domain,

            Outputs.GetDatabasesDatabaseExtensionsResult extensions,

            string id,

            string name,

            ImmutableArray<Outputs.GetDatabasesDatabaseNodeResult> nodes,

            ImmutableArray<string> options,

            string pgVersion,

            ImmutableArray<Outputs.GetDatabasesDatabaseRoleResult> roles,

            string status,

            int storageUsed,

            ImmutableArray<Outputs.GetDatabasesDatabaseTableResult> tables,

            string updatedAt)
        {
            ClusterId = clusterId;
            Components = components;
            ConfigVersion = configVersion;
            CreatedAt = createdAt;
            Domain = domain;
            Extensions = extensions;
            Id = id;
            Name = name;
            Nodes = nodes;
            Options = options;
            PgVersion = pgVersion;
            Roles = roles;
            Status = status;
            StorageUsed = storageUsed;
            Tables = tables;
            UpdatedAt = updatedAt;
        }
    }
}
