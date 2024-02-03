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
    public sealed class GetClustersClusterNodeGroupsAwResult
    {
        /// <summary>
        /// Availability zones of the AWS node group
        /// </summary>
        public readonly ImmutableArray<string> AvailabilityZones;
        /// <summary>
        /// CIDR of the AWS node group
        /// </summary>
        public readonly string Cidr;
        /// <summary>
        /// Instance type of the AWS node group
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// Node location of the AWS node group
        /// </summary>
        public readonly string NodeLocation;
        public readonly ImmutableArray<Outputs.GetClustersClusterNodeGroupsAwNodeResult> Nodes;
        public readonly ImmutableArray<string> PrivateSubnets;
        public readonly ImmutableArray<string> PublicSubnets;
        /// <summary>
        /// Region of the AWS node group
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Volume IOPS of the AWS node group
        /// </summary>
        public readonly int VolumeIops;
        /// <summary>
        /// Volume size of the AWS node group
        /// </summary>
        public readonly int VolumeSize;
        /// <summary>
        /// Volume type of the AWS node group
        /// </summary>
        public readonly string VolumeType;

        [OutputConstructor]
        private GetClustersClusterNodeGroupsAwResult(
            ImmutableArray<string> availabilityZones,

            string cidr,

            string instanceType,

            string nodeLocation,

            ImmutableArray<Outputs.GetClustersClusterNodeGroupsAwNodeResult> nodes,

            ImmutableArray<string> privateSubnets,

            ImmutableArray<string> publicSubnets,

            string region,

            int volumeIops,

            int volumeSize,

            string volumeType)
        {
            AvailabilityZones = availabilityZones;
            Cidr = cidr;
            InstanceType = instanceType;
            NodeLocation = nodeLocation;
            Nodes = nodes;
            PrivateSubnets = privateSubnets;
            PublicSubnets = publicSubnets;
            Region = region;
            VolumeIops = volumeIops;
            VolumeSize = volumeSize;
            VolumeType = volumeType;
        }
    }
}
