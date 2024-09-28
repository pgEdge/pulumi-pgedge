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
    public sealed class GetCloudAccountsCloudAccountResult
    {
        /// <summary>
        /// Creation time of the cloud account
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Description of the cloud account
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// ID of the cloud account
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the cloud account
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Additional properties of the cloud account
        /// </summary>
        public readonly ImmutableDictionary<string, string> Properties;
        /// <summary>
        /// Type of the cloud account (e.g., AWS, Azure, GCP)
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Last update time of the cloud account
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetCloudAccountsCloudAccountResult(
            string createdAt,

            string description,

            string id,

            string name,

            ImmutableDictionary<string, string> properties,

            string type,

            string updatedAt)
        {
            CreatedAt = createdAt;
            Description = description;
            Id = id;
            Name = name;
            Properties = properties;
            Type = type;
            UpdatedAt = updatedAt;
        }
    }
}