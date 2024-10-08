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
    public sealed class GetDatabasesDatabaseNodesResult
    {
        /// <summary>
        /// Node connection details
        /// </summary>
        public readonly Outputs.GetDatabasesDatabaseNodesConnectionResult Connection;
        /// <summary>
        /// Extensions configuration for the database
        /// </summary>
        public readonly Outputs.GetDatabasesDatabaseNodesExtensionsResult Extensions;
        /// <summary>
        /// Node location
        /// </summary>
        public readonly Outputs.GetDatabasesDatabaseNodesLocationResult Location;
        /// <summary>
        /// Component name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Node region
        /// </summary>
        public readonly Outputs.GetDatabasesDatabaseNodesRegionResult Region;

        [OutputConstructor]
        private GetDatabasesDatabaseNodesResult(
            Outputs.GetDatabasesDatabaseNodesConnectionResult connection,

            Outputs.GetDatabasesDatabaseNodesExtensionsResult extensions,

            Outputs.GetDatabasesDatabaseNodesLocationResult location,

            string name,

            Outputs.GetDatabasesDatabaseNodesRegionResult region)
        {
            Connection = connection;
            Extensions = extensions;
            Location = location;
            Name = name;
            Region = region;
        }
    }
}
