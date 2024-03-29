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
    public sealed class GetClustersClusterFirewallResult
    {
        /// <summary>
        /// Port for the firewall rule
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Sources for the firewall rule
        /// </summary>
        public readonly ImmutableArray<string> Sources;
        /// <summary>
        /// Type of the firewall rule
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetClustersClusterFirewallResult(
            int port,

            ImmutableArray<string> sources,

            string type)
        {
            Port = port;
            Sources = sources;
            Type = type;
        }
    }
}
