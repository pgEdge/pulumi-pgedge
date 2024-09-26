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
    public sealed class GetDatabasesDatabaseNodeRegionResult
    {
        /// <summary>
        /// Active status of the region
        /// </summary>
        public readonly bool Active;
        /// <summary>
        /// Availability zones of the region
        /// </summary>
        public readonly ImmutableArray<string> AvailabilityZones;
        /// <summary>
        /// Cloud provider of the region
        /// </summary>
        public readonly string Cloud;
        /// <summary>
        /// Code of the location
        /// </summary>
        public readonly string Code;
        /// <summary>
        /// Name of the component
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Parent region
        /// </summary>
        public readonly string Parent;

        [OutputConstructor]
        private GetDatabasesDatabaseNodeRegionResult(
            bool active,

            ImmutableArray<string> availabilityZones,

            string cloud,

            string code,

            string name,

            string parent)
        {
            Active = active;
            AvailabilityZones = availabilityZones;
            Cloud = cloud;
            Code = code;
            Name = name;
            Parent = parent;
        }
    }
}