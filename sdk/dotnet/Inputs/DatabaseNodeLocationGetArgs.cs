// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pgedge.Pgedge.Inputs
{

    public sealed class DatabaseNodeLocationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Code of the location
        /// </summary>
        [Input("code")]
        public Input<string>? Code { get; set; }

        /// <summary>
        /// Country of the location
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// Latitude of the location
        /// </summary>
        [Input("latitude")]
        public Input<double>? Latitude { get; set; }

        /// <summary>
        /// Longitude of the location
        /// </summary>
        [Input("longitude")]
        public Input<double>? Longitude { get; set; }

        /// <summary>
        /// Name of the location
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Region of the location
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public DatabaseNodeLocationGetArgs()
        {
        }
        public static new DatabaseNodeLocationGetArgs Empty => new DatabaseNodeLocationGetArgs();
    }
}
