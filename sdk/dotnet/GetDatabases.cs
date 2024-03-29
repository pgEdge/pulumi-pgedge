// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pgedge.Pgedge
{
    public static class GetDatabases
    {
        /// <summary>
        /// Interface with the pgEdge service API.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Pgedge = Pulumi.Pgedge;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Pgedge.GetDatabases.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDatabasesResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabasesResult>("pgedge:index/getDatabases:getDatabases", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Interface with the pgEdge service API.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Pgedge = Pulumi.Pgedge;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Pgedge.GetDatabases.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDatabasesResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabasesResult>("pgedge:index/getDatabases:getDatabases", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetDatabasesResult
    {
        public readonly ImmutableArray<Outputs.GetDatabasesDatabaseResult> Databases;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetDatabasesResult(
            ImmutableArray<Outputs.GetDatabasesDatabaseResult> databases,

            string id)
        {
            Databases = databases;
            Id = id;
        }
    }
}
