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
    public sealed class DatabaseNodesRegion
    {
        public readonly bool? Active;
        public readonly ImmutableArray<string> AvailabilityZones;
        public readonly string? Cloud;
        public readonly string? Code;
        public readonly string? Name;
        public readonly string? Parent;

        [OutputConstructor]
        private DatabaseNodesRegion(
            bool? active,

            ImmutableArray<string> availabilityZones,

            string? cloud,

            string? code,

            string? name,

            string? parent)
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
