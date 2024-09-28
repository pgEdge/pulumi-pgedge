// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pgedge.Pgedge.Outputs
{

    [OutputType]
    public sealed class DatabaseBackups
    {
        /// <summary>
        /// List of backup configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.DatabaseBackupsConfig> Configs;
        /// <summary>
        /// The backup provider.
        /// </summary>
        public readonly string? Provider;

        [OutputConstructor]
        private DatabaseBackups(
            ImmutableArray<Outputs.DatabaseBackupsConfig> configs,

            string? provider)
        {
            Configs = configs;
            Provider = provider;
        }
    }
}