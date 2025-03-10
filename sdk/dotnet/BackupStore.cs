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
    [PgedgeResourceType("pgedge:index/backupStore:BackupStore")]
    public partial class BackupStore : global::Pulumi.CustomResource
    {
        [Output("cloudAccountId")]
        public Output<string> CloudAccountId { get; private set; } = null!;

        [Output("cloudAccountType")]
        public Output<string> CloudAccountType { get; private set; } = null!;

        [Output("clusterIds")]
        public Output<ImmutableArray<string>> ClusterIds { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("properties")]
        public Output<ImmutableDictionary<string, string>> Properties { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a BackupStore resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackupStore(string name, BackupStoreArgs args, CustomResourceOptions? options = null)
            : base("pgedge:index/backupStore:BackupStore", name, args ?? new BackupStoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackupStore(string name, Input<string> id, BackupStoreState? state = null, CustomResourceOptions? options = null)
            : base("pgedge:index/backupStore:BackupStore", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pgEdge/pulumi-pgedge",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BackupStore resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackupStore Get(string name, Input<string> id, BackupStoreState? state = null, CustomResourceOptions? options = null)
        {
            return new BackupStore(name, id, state, options);
        }
    }

    public sealed class BackupStoreArgs : global::Pulumi.ResourceArgs
    {
        [Input("cloudAccountId", required: true)]
        public Input<string> CloudAccountId { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public BackupStoreArgs()
        {
        }
        public static new BackupStoreArgs Empty => new BackupStoreArgs();
    }

    public sealed class BackupStoreState : global::Pulumi.ResourceArgs
    {
        [Input("cloudAccountId")]
        public Input<string>? CloudAccountId { get; set; }

        [Input("cloudAccountType")]
        public Input<string>? CloudAccountType { get; set; }

        [Input("clusterIds")]
        private InputList<string>? _clusterIds;
        public InputList<string> ClusterIds
        {
            get => _clusterIds ?? (_clusterIds = new InputList<string>());
            set => _clusterIds = value;
        }

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("properties")]
        private InputMap<string>? _properties;
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public BackupStoreState()
        {
        }
        public static new BackupStoreState Empty => new BackupStoreState();
    }
}
