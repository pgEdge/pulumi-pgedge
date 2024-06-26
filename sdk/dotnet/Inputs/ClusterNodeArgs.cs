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

    public sealed class ClusterNodeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud provider availability zone name
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Instance type used for the node
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// Node name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("options")]
        private InputList<string>? _options;
        public InputList<string> Options
        {
            get => _options ?? (_options = new InputList<string>());
            set => _options = value;
        }

        /// <summary>
        /// Cloud provider region
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// Volume IOPS of the node data volume
        /// </summary>
        [Input("volumeIops")]
        public Input<int>? VolumeIops { get; set; }

        /// <summary>
        /// Volume size of the node data volume
        /// </summary>
        [Input("volumeSize")]
        public Input<int>? VolumeSize { get; set; }

        /// <summary>
        /// Volume type of the node data volume
        /// </summary>
        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public ClusterNodeArgs()
        {
        }
        public static new ClusterNodeArgs Empty => new ClusterNodeArgs();
    }
}
