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

    public sealed class ClusterFirewallRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the network
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Port whose traffic is allowed
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        [Input("sources", required: true)]
        private InputList<string>? _sources;

        /// <summary>
        /// CIDRs and/or IP addresses allowed
        /// </summary>
        public InputList<string> Sources
        {
            get => _sources ?? (_sources = new InputList<string>());
            set => _sources = value;
        }

        public ClusterFirewallRuleGetArgs()
        {
        }
        public static new ClusterFirewallRuleGetArgs Empty => new ClusterFirewallRuleGetArgs();
    }
}
