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

    public sealed class DatabaseNodesArgs : global::Pulumi.ResourceArgs
    {
        [Input("connection")]
        public Input<Inputs.DatabaseNodesConnectionArgs>? Connection { get; set; }

        [Input("extensions")]
        public Input<Inputs.DatabaseNodesExtensionsArgs>? Extensions { get; set; }

        [Input("location")]
        public Input<Inputs.DatabaseNodesLocationArgs>? Location { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("region")]
        public Input<Inputs.DatabaseNodesRegionArgs>? Region { get; set; }

        public DatabaseNodesArgs()
        {
        }
        public static new DatabaseNodesArgs Empty => new DatabaseNodesArgs();
    }
}
