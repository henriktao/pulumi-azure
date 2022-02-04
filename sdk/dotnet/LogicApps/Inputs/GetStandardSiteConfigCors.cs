// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.LogicApps.Inputs
{

    public sealed class GetStandardSiteConfigCorsArgs : Pulumi.InvokeArgs
    {
        [Input("allowedOrigins", required: true)]
        private List<string>? _allowedOrigins;
        public List<string> AllowedOrigins
        {
            get => _allowedOrigins ?? (_allowedOrigins = new List<string>());
            set => _allowedOrigins = value;
        }

        [Input("supportCredentials")]
        public bool? SupportCredentials { get; set; }

        public GetStandardSiteConfigCorsArgs()
        {
        }
    }
}
