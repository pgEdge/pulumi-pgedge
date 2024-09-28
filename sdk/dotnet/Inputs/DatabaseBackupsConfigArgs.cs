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

    public sealed class DatabaseBackupsConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the database
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the node.
        /// </summary>
        [Input("nodeName")]
        public Input<string>? NodeName { get; set; }

        [Input("repositories")]
        private InputList<Inputs.DatabaseBackupsConfigRepositoryArgs>? _repositories;

        /// <summary>
        /// List of backup repositories.
        /// </summary>
        public InputList<Inputs.DatabaseBackupsConfigRepositoryArgs> Repositories
        {
            get => _repositories ?? (_repositories = new InputList<Inputs.DatabaseBackupsConfigRepositoryArgs>());
            set => _repositories = value;
        }

        [Input("schedules")]
        private InputList<Inputs.DatabaseBackupsConfigScheduleArgs>? _schedules;

        /// <summary>
        /// List of backup schedules.
        /// </summary>
        public InputList<Inputs.DatabaseBackupsConfigScheduleArgs> Schedules
        {
            get => _schedules ?? (_schedules = new InputList<Inputs.DatabaseBackupsConfigScheduleArgs>());
            set => _schedules = value;
        }

        public DatabaseBackupsConfigArgs()
        {
        }
        public static new DatabaseBackupsConfigArgs Empty => new DatabaseBackupsConfigArgs();
    }
}