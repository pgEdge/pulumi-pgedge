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
    public sealed class DatabaseNodesConnection
    {
        public readonly string? Database;
        public readonly string? ExternalIpAddress;
        public readonly string? Host;
        public readonly string? InternalHost;
        public readonly string? InternalIpAddress;
        public readonly string? Password;
        public readonly int? Port;
        public readonly string? Username;

        [OutputConstructor]
        private DatabaseNodesConnection(
            string? database,

            string? externalIpAddress,

            string? host,

            string? internalHost,

            string? internalIpAddress,

            string? password,

            int? port,

            string? username)
        {
            Database = database;
            ExternalIpAddress = externalIpAddress;
            Host = host;
            InternalHost = internalHost;
            InternalIpAddress = internalIpAddress;
            Password = password;
            Port = port;
            Username = username;
        }
    }
}
