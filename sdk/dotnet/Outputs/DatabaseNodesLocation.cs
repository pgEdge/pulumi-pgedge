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
    public sealed class DatabaseNodesLocation
    {
        public readonly string? City;
        public readonly string? Code;
        public readonly string? Country;
        public readonly double? Latitude;
        public readonly double? Longitude;
        public readonly string? MetroCode;
        public readonly string? Name;
        public readonly string? PostalCode;
        public readonly string? Region;
        public readonly string? RegionCode;
        public readonly string? Timezone;

        [OutputConstructor]
        private DatabaseNodesLocation(
            string? city,

            string? code,

            string? country,

            double? latitude,

            double? longitude,

            string? metroCode,

            string? name,

            string? postalCode,

            string? region,

            string? regionCode,

            string? timezone)
        {
            City = city;
            Code = code;
            Country = country;
            Latitude = latitude;
            Longitude = longitude;
            MetroCode = metroCode;
            Name = name;
            PostalCode = postalCode;
            Region = region;
            RegionCode = regionCode;
            Timezone = timezone;
        }
    }
}