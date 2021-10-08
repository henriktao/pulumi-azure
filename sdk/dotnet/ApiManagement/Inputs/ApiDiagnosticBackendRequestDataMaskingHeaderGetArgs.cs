// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement.Inputs
{

    public sealed class ApiDiagnosticBackendRequestDataMaskingHeaderGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The data masking mode. Possible values are `Mask` and `Hide` for `query_params`. The only possible value is `Mask` for `headers`.
        /// </summary>
        [Input("mode", required: true)]
        public Input<string> Mode { get; set; } = null!;

        /// <summary>
        /// The name of the header or the query parameter to mask.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ApiDiagnosticBackendRequestDataMaskingHeaderGetArgs()
        {
        }
    }
}
