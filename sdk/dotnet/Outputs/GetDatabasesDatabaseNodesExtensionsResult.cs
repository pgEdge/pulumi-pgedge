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
    public sealed class GetDatabasesDatabaseNodesExtensionsResult
    {
        public readonly ImmutableDictionary<string, string> Errors;
        public readonly ImmutableArray<string> Installeds;

        [OutputConstructor]
        private GetDatabasesDatabaseNodesExtensionsResult(
            ImmutableDictionary<string, string> errors,

            ImmutableArray<string> installeds)
        {
            Errors = errors;
            Installeds = installeds;
        }
    }
}
