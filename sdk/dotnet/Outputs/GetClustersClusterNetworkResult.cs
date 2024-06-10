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
    public sealed class GetClustersClusterNetworkResult
    {
        /// <summary>
        /// CIDR of the AWS node group
        /// </summary>
        public readonly string Cidr;
        /// <summary>
        /// Is the network external
        /// </summary>
        public readonly bool External;
        /// <summary>
        /// External ID of the network
        /// </summary>
        public readonly string ExternalId;
        /// <summary>
        /// IP address of the node
        /// </summary>
        public readonly string Name;
        public readonly ImmutableArray<string> PrivateSubnets;
        public readonly ImmutableArray<string> PublicSubnets;
        /// <summary>
        /// Region of the network
        /// </summary>
        public readonly string Region;

        [OutputConstructor]
        private GetClustersClusterNetworkResult(
            string cidr,

            bool external,

            string externalId,

            string name,

            ImmutableArray<string> privateSubnets,

            ImmutableArray<string> publicSubnets,

            string region)
        {
            Cidr = cidr;
            External = external;
            ExternalId = externalId;
            Name = name;
            PrivateSubnets = privateSubnets;
            PublicSubnets = publicSubnets;
            Region = region;
        }
    }
}
