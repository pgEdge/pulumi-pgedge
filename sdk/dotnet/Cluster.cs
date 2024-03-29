// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pgedge.Pgedge
{
    /// <summary>
    /// Interface with the pgEdge service API for clusters.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Pgedge = Pgedge.Pgedge;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Pgedge.Cluster("example", new()
    ///     {
    ///         CloudAccountId = "",
    ///         Firewalls = new[]
    ///         {
    ///             new Pgedge.Inputs.ClusterFirewallArgs
    ///             {
    ///                 Port = 5432,
    ///                 Sources = new[]
    ///                 {
    ///                     "0.0.0.0/0",
    ///                 },
    ///                 Type = "postgres",
    ///             },
    ///         },
    ///         NodeGroups = new Pgedge.Inputs.ClusterNodeGroupsArgs
    ///         {
    ///             Aws = new[]
    ///             {
    ///                 new Pgedge.Inputs.ClusterNodeGroupsAwArgs
    ///                 {
    ///                     InstanceType = "t4g.small",
    ///                     Region = "us-west-2",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [PgedgeResourceType("pgedge:index/cluster:Cluster")]
    public partial class Cluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cloud account ID of the cluster
        /// </summary>
        [Output("cloudAccountId")]
        public Output<string> CloudAccountId { get; private set; } = null!;

        /// <summary>
        /// Created at of the cluster
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("firewalls")]
        public Output<ImmutableArray<Outputs.ClusterFirewall>> Firewalls { get; private set; } = null!;

        /// <summary>
        /// Name of the cluster
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("nodeGroups")]
        public Output<Outputs.ClusterNodeGroups> NodeGroups { get; private set; } = null!;

        /// <summary>
        /// Status of the cluster
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("pgedge:index/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("pgedge:index/cluster:Cluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, state, options);
        }
    }

    public sealed class ClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud account ID of the cluster
        /// </summary>
        [Input("cloudAccountId", required: true)]
        public Input<string> CloudAccountId { get; set; } = null!;

        [Input("firewalls")]
        private InputList<Inputs.ClusterFirewallArgs>? _firewalls;
        public InputList<Inputs.ClusterFirewallArgs> Firewalls
        {
            get => _firewalls ?? (_firewalls = new InputList<Inputs.ClusterFirewallArgs>());
            set => _firewalls = value;
        }

        /// <summary>
        /// Name of the cluster
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nodeGroups")]
        public Input<Inputs.ClusterNodeGroupsArgs>? NodeGroups { get; set; }

        public ClusterArgs()
        {
        }
        public static new ClusterArgs Empty => new ClusterArgs();
    }

    public sealed class ClusterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud account ID of the cluster
        /// </summary>
        [Input("cloudAccountId")]
        public Input<string>? CloudAccountId { get; set; }

        /// <summary>
        /// Created at of the cluster
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("firewalls")]
        private InputList<Inputs.ClusterFirewallGetArgs>? _firewalls;
        public InputList<Inputs.ClusterFirewallGetArgs> Firewalls
        {
            get => _firewalls ?? (_firewalls = new InputList<Inputs.ClusterFirewallGetArgs>());
            set => _firewalls = value;
        }

        /// <summary>
        /// Name of the cluster
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nodeGroups")]
        public Input<Inputs.ClusterNodeGroupsGetArgs>? NodeGroups { get; set; }

        /// <summary>
        /// Status of the cluster
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ClusterState()
        {
        }
        public static new ClusterState Empty => new ClusterState();
    }
}
