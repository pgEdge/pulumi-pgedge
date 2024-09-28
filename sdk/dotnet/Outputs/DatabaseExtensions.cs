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
    public sealed class DatabaseExtensions
    {
        public readonly bool? AutoManage;
        public readonly ImmutableArray<string> Availables;
        public readonly ImmutableArray<string> Requesteds;

        [OutputConstructor]
        private DatabaseExtensions(
            bool? autoManage,

            ImmutableArray<string> availables,

            ImmutableArray<string> requesteds)
        {
            AutoManage = autoManage;
            Availables = availables;
            Requesteds = requesteds;
        }
    }
}