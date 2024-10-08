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
        /// Backup store IDs of the cluster
        /// </summary>
        public readonly ImmutableArray<string> BackupStoreIds;
        /// <summary>
        /// Capacity of the cluster
        /// </summary>
        public readonly int Capacity;
        /// <summary>
        /// Cloud account ID of the cluster
        /// </summary>
        public readonly string CloudAccountId;
        /// <summary>
        /// Created at of the cluster
        /// </summary>
        public readonly string CreatedAt;
        public readonly ImmutableArray<Outputs.GetClustersClusterFirewallRuleResult> FirewallRules;
        /// <summary>
        /// ID of the cluster
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the cluster
        /// </summary>
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetClustersClusterNetworkResult> Networks;
        /// <summary>
        /// Node location of the cluster. Must be either 'public' or 'private'.
        /// </summary>
        public readonly string NodeLocation;
        public readonly ImmutableArray<Outputs.GetClustersClusterNodeResult> Nodes;
        public readonly ImmutableArray<string> Regions;
        /// <summary>
        /// Resource tags of the cluster
        /// </summary>
        public readonly ImmutableDictionary<string, string> ResourceTags;
        /// <summary>
        /// SSH key ID of the cluster
        /// </summary>
        public readonly string SshKeyId;
        /// <summary>
        /// Status of the cluster
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetClustersClusterResult(
            ImmutableArray<string> backupStoreIds,

            int capacity,

            string cloudAccountId,

            string createdAt,

            ImmutableArray<Outputs.GetClustersClusterFirewallRuleResult> firewallRules,

            string id,

            string name,

            ImmutableArray<Outputs.GetClustersClusterNetworkResult> networks,

            string nodeLocation,

            ImmutableArray<Outputs.GetClustersClusterNodeResult> nodes,

            ImmutableArray<string> regions,

            ImmutableDictionary<string, string> resourceTags,

            string sshKeyId,

            string status)
        {
            BackupStoreIds = backupStoreIds;
            Capacity = capacity;
            CloudAccountId = cloudAccountId;
            CreatedAt = createdAt;
            FirewallRules = firewallRules;
            Id = id;
            Name = name;
            Networks = networks;
            NodeLocation = nodeLocation;
            Nodes = nodes;
            Regions = regions;
            ResourceTags = resourceTags;
            SshKeyId = sshKeyId;
            Status = status;
        }
    }
}
