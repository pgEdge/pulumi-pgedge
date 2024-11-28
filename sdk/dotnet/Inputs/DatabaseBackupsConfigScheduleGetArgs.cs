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

    public sealed class DatabaseBackupsConfigScheduleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cron expression for the schedule.
        /// </summary>
        [Input("cronExpression", required: true)]
        public Input<string> CronExpression { get; set; } = null!;

        /// <summary>
        /// Unique identifier for the backup config.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// Repository type (e.g., s3, gcs, azure).
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DatabaseBackupsConfigScheduleGetArgs()
        {
        }
        public static new DatabaseBackupsConfigScheduleGetArgs Empty => new DatabaseBackupsConfigScheduleGetArgs();
    }
}
