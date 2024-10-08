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

    public sealed class DatabaseBackupsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("configs")]
        private InputList<Inputs.DatabaseBackupsConfigGetArgs>? _configs;

        /// <summary>
        /// List of backup configurations.
        /// </summary>
        public InputList<Inputs.DatabaseBackupsConfigGetArgs> Configs
        {
            get => _configs ?? (_configs = new InputList<Inputs.DatabaseBackupsConfigGetArgs>());
            set => _configs = value;
        }

        /// <summary>
        /// The backup provider.
        /// </summary>
        [Input("provider")]
        public Input<string>? Provider { get; set; }

        public DatabaseBackupsGetArgs()
        {
        }
        public static new DatabaseBackupsGetArgs Empty => new DatabaseBackupsGetArgs();
    }
}
