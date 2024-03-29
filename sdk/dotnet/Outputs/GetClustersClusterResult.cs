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
    public sealed class GetClustersClusterResult
    {
        /// <summary>
        /// Cloud account ID of the cluster
        /// </summary>
        public readonly string CloudAccountId;
        /// <summary>
        /// Created at of the cluster
        /// </summary>
        public readonly string CreatedAt;
        public readonly ImmutableArray<Outputs.GetClustersClusterFirewallResult> Firewalls;
        /// <summary>
        /// ID of the cluster
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the cluster
        /// </summary>
        public readonly string Name;
        public readonly Outputs.GetClustersClusterNodeGroupsResult NodeGroups;
        /// <summary>
        /// Status of the cluster
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetClustersClusterResult(
            string cloudAccountId,

            string createdAt,

            ImmutableArray<Outputs.GetClustersClusterFirewallResult> firewalls,

            string id,

            string name,

            Outputs.GetClustersClusterNodeGroupsResult nodeGroups,

            string status)
        {
            CloudAccountId = cloudAccountId;
            CreatedAt = createdAt;
            Firewalls = firewalls;
            Id = id;
            Name = name;
            NodeGroups = nodeGroups;
            Status = status;
        }
    }
}
