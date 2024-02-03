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
    public sealed class DatabaseNodeConnection
    {
        /// <summary>
        /// Database of the node
        /// </summary>
        public readonly string? Database;
        /// <summary>
        /// Host of the node
        /// </summary>
        public readonly string? Host;
        /// <summary>
        /// Password of the node
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Port of the node
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// Username of the node
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private DatabaseNodeConnection(
            string? database,

            string? host,

            string? password,

            int? port,

            string? username)
        {
            Database = database;
            Host = host;
            Password = password;
            Port = port;
            Username = username;
        }
    }
}
