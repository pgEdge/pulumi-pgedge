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
    public sealed class GetDatabasesDatabaseNodeLocationResult
    {
        /// <summary>
        /// City of the location
        /// </summary>
        public readonly string City;
        /// <summary>
        /// Code of the location
        /// </summary>
        public readonly string Code;
        /// <summary>
        /// Country of the location
        /// </summary>
        public readonly string Country;
        /// <summary>
        /// Latitude of the location
        /// </summary>
        public readonly double Latitude;
        /// <summary>
        /// Longitude of the location
        /// </summary>
        public readonly double Longitude;
        /// <summary>
        /// Metro code of the location
        /// </summary>
        public readonly string MetroCode;
        /// <summary>
        /// Name of the component
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Postal code of the location
        /// </summary>
        public readonly string PostalCode;
        /// <summary>
        /// Region of the location
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Region code of the location
        /// </summary>
        public readonly string RegionCode;
        /// <summary>
        /// Timezone of the location
        /// </summary>
        public readonly string Timezone;

        [OutputConstructor]
        private GetDatabasesDatabaseNodeLocationResult(
            string city,

            string code,

            string country,

            double latitude,

            double longitude,

            string metroCode,

            string name,

            string postalCode,

            string region,

            string regionCode,

            string timezone)
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
