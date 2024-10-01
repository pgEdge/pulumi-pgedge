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

    public sealed class DatabaseExtensionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("autoManage")]
        public Input<bool>? AutoManage { get; set; }

        [Input("availables")]
        private InputList<string>? _availables;
        public InputList<string> Availables
        {
            get => _availables ?? (_availables = new InputList<string>());
            set => _availables = value;
        }

        [Input("requesteds")]
        private InputList<string>? _requesteds;
        public InputList<string> Requesteds
        {
            get => _requesteds ?? (_requesteds = new InputList<string>());
            set => _requesteds = value;
        }

        public DatabaseExtensionsArgs()
        {
        }
        public static new DatabaseExtensionsArgs Empty => new DatabaseExtensionsArgs();
    }
}
