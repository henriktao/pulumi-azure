// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement.Outputs
{

    [OutputType]
    public sealed class ApiDiagnosticFrontendRequestDataMaskingHeader
    {
        /// <summary>
        /// The data masking mode. Possible values are `Mask` and `Hide` for `query_params`. The only possible value is `Mask` for `headers`.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// The name of the header or the query parameter to mask.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ApiDiagnosticFrontendRequestDataMaskingHeader(
            string mode,

            string value)
        {
            Mode = mode;
            Value = value;
        }
    }
}
