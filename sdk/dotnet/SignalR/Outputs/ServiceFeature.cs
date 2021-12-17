// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.SignalR.Outputs
{

    [OutputType]
    public sealed class ServiceFeature
    {
        /// <summary>
        /// The kind of Feature. Possible values are `EnableConnectivityLogs`, `EnableMessagingLogs`, `EnableLiveTrace` and `ServiceMode`.
        /// </summary>
        public readonly string Flag;
        /// <summary>
        /// A value of a feature flag. Possible values are `Classic`, `Default` and `Serverless`.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ServiceFeature(
            string flag,

            string value)
        {
            Flag = flag;
            Value = value;
        }
    }
}
