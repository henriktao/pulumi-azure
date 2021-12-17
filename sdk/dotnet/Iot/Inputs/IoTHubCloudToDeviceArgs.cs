// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Iot.Inputs
{

    public sealed class IoTHubCloudToDeviceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default time to live for cloud-to-device messages, specified as an [ISO 8601 timespan duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This value must be between 1 minute and 48 hours, and evaluates to `PT1H` by default.
        /// </summary>
        [Input("defaultTtl")]
        public Input<string>? DefaultTtl { get; set; }

        [Input("feedbacks")]
        private InputList<Inputs.IoTHubCloudToDeviceFeedbackArgs>? _feedbacks;

        /// <summary>
        /// A `feedback` block as defined below.
        /// </summary>
        public InputList<Inputs.IoTHubCloudToDeviceFeedbackArgs> Feedbacks
        {
            get => _feedbacks ?? (_feedbacks = new InputList<Inputs.IoTHubCloudToDeviceFeedbackArgs>());
            set => _feedbacks = value;
        }

        /// <summary>
        /// The maximum delivery count for cloud-to-device per-device queues. This value must be between `1` and `100`, and evaluates to `10` by default.
        /// </summary>
        [Input("maxDeliveryCount")]
        public Input<int>? MaxDeliveryCount { get; set; }

        public IoTHubCloudToDeviceArgs()
        {
        }
    }
}
