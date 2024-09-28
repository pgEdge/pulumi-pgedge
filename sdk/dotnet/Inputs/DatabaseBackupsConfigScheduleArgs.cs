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

    public sealed class DatabaseBackupsConfigScheduleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cron expression for the schedule.
        /// </summary>
        [Input("cronExpression")]
        public Input<string>? CronExpression { get; set; }

        /// <summary>
        /// ID of the database
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Type of the schedule.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DatabaseBackupsConfigScheduleArgs()
        {
        }
        public static new DatabaseBackupsConfigScheduleArgs Empty => new DatabaseBackupsConfigScheduleArgs();
    }
}