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

    public sealed class DatabaseNodeConnectionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Database of the node
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// Host of the node
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Password of the node
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Port of the node
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Username of the node
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public DatabaseNodeConnectionGetArgs()
        {
        }
        public static new DatabaseNodeConnectionGetArgs Empty => new DatabaseNodeConnectionGetArgs();
    }
}
