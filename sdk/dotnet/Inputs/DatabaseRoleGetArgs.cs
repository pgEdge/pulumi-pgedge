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

    public sealed class DatabaseRoleGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("bypassRls")]
        public Input<bool>? BypassRls { get; set; }

        [Input("connectionLimit")]
        public Input<int>? ConnectionLimit { get; set; }

        [Input("createDb")]
        public Input<bool>? CreateDb { get; set; }

        [Input("createRole")]
        public Input<bool>? CreateRole { get; set; }

        [Input("inherit")]
        public Input<bool>? Inherit { get; set; }

        [Input("login")]
        public Input<bool>? Login { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("replication")]
        public Input<bool>? Replication { get; set; }

        [Input("superuser")]
        public Input<bool>? Superuser { get; set; }

        public DatabaseRoleGetArgs()
        {
        }
        public static new DatabaseRoleGetArgs Empty => new DatabaseRoleGetArgs();
    }
}
