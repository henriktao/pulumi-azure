// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Iot.Outputs
{

    [OutputType]
    public sealed class IoTHubFallbackRoute
    {
        /// <summary>
        /// The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to true by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
        /// </summary>
        public readonly string? Condition;
        /// <summary>
        /// Used to specify whether the fallback route is enabled.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The endpoints to which messages that satisfy the condition are routed. Currently only 1 endpoint is allowed.
        /// </summary>
        public readonly ImmutableArray<string> EndpointNames;
        /// <summary>
        /// The source that the routing rule is to be applied to, such as `DeviceMessages`. Possible values include: `Invalid`, `DeviceMessages`, `TwinChangeEvents`, `DeviceLifecycleEvents`, `DeviceConnectionStateEvents`, `DeviceJobLifecycleEvents`.
        /// </summary>
        public readonly string? Source;

        [OutputConstructor]
        private IoTHubFallbackRoute(
            string? condition,

            bool? enabled,

            ImmutableArray<string> endpointNames,

            string? source)
        {
            Condition = condition;
            Enabled = enabled;
            EndpointNames = endpointNames;
            Source = source;
        }
    }
}
