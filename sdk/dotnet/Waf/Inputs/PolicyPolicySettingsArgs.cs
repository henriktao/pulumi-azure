// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Waf.Inputs
{

    public sealed class PolicyPolicySettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes if the policy is in enabled state or disabled state. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The File Upload Limit in MB. Accepted values are in the range `1` to `4000`. Defaults to `100`.
        /// </summary>
        [Input("fileUploadLimitInMb")]
        public Input<int>? FileUploadLimitInMb { get; set; }

        /// <summary>
        /// The Maximum Request Body Size in KB.  Accepted values are in the range `8` to `2000`. Defaults to `128`.
        /// </summary>
        [Input("maxRequestBodySizeInKb")]
        public Input<int>? MaxRequestBodySizeInKb { get; set; }

        /// <summary>
        /// Describes if it is in detection mode or prevention mode at the policy level. Valid values are `Detection` and `Prevention`. Defaults to `Prevention`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Is Request Body Inspection enabled? Defaults to `true`.
        /// </summary>
        [Input("requestBodyCheck")]
        public Input<bool>? RequestBodyCheck { get; set; }

        public PolicyPolicySettingsArgs()
        {
        }
    }
}
