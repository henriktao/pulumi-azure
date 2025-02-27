// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.SignalR.Inputs
{

    public sealed class ServiceFeatureGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The kind of Feature. Possible values are `EnableConnectivityLogs`, `EnableMessagingLogs`, `EnableLiveTrace` and `ServiceMode`.
        /// </summary>
        [Input("flag", required: true)]
        public Input<string> Flag { get; set; } = null!;

        /// <summary>
        /// A value of a feature flag. Possible values are `Classic`, `Default` and `Serverless`.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ServiceFeatureGetArgs()
        {
        }
    }
}
