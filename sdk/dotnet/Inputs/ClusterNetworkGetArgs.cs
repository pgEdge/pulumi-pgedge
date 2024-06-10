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

    public sealed class ClusterNetworkGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CIDR range for the network
        /// </summary>
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        /// <summary>
        /// Is the network externally defined
        /// </summary>
        [Input("external")]
        public Input<bool>? External { get; set; }

        /// <summary>
        /// ID of the network, if externally defined
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// Name of the network
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("privateSubnets")]
        private InputList<string>? _privateSubnets;
        public InputList<string> PrivateSubnets
        {
            get => _privateSubnets ?? (_privateSubnets = new InputList<string>());
            set => _privateSubnets = value;
        }

        [Input("publicSubnets")]
        private InputList<string>? _publicSubnets;
        public InputList<string> PublicSubnets
        {
            get => _publicSubnets ?? (_publicSubnets = new InputList<string>());
            set => _publicSubnets = value;
        }

        /// <summary>
        /// Region of the network
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public ClusterNetworkGetArgs()
        {
        }
        public static new ClusterNetworkGetArgs Empty => new ClusterNetworkGetArgs();
    }
}
