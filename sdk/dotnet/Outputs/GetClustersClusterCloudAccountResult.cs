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
    public sealed class GetClustersClusterCloudAccountResult
    {
        /// <summary>
        /// Display name of the node
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IP address of the node
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Type of the node
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetClustersClusterCloudAccountResult(
            string id,

            string name,

            string type)
        {
            Id = id;
            Name = name;
            Type = type;
        }
    }
}
