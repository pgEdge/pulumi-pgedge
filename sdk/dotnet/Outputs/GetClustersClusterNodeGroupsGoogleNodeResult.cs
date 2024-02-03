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
    public sealed class GetClustersClusterNodeGroupsGoogleNodeResult
    {
        /// <summary>
        /// Display name of the node
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// IP address of the node
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// Is the node active
        /// </summary>
        public readonly bool IsActive;

        [OutputConstructor]
        private GetClustersClusterNodeGroupsGoogleNodeResult(
            string displayName,

            string ipAddress,

            bool isActive)
        {
            DisplayName = displayName;
            IpAddress = ipAddress;
            IsActive = isActive;
        }
    }
}
