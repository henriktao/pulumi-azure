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
    public sealed class IoTHubFileUpload
    {
        /// <summary>
        /// The connection string for the Azure Storage account to which files are uploaded.
        /// </summary>
        public readonly string ConnectionString;
        /// <summary>
        /// The name of the root container where you upload files. The container need not exist but should be creatable using the connection_string specified.
        /// </summary>
        public readonly string ContainerName;
        /// <summary>
        /// The period of time for which a file upload notification message is available to consume before it is expired by the IoT hub, specified as an [ISO 8601 timespan duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This value must be between 1 minute and 48 hours, and evaluates to `PT1H` by default.
        /// </summary>
        public readonly string? DefaultTtl;
        /// <summary>
        /// The lock duration for the file upload notifications queue, specified as an [ISO 8601 timespan duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This value must be between 5 and 300 seconds, and evaluates to `PT1M` by default.
        /// </summary>
        public readonly string? LockDuration;
        /// <summary>
        /// The number of times the IoT hub attempts to deliver a file upload notification message. It evaluates to `10` by default.
        /// </summary>
        public readonly int? MaxDeliveryCount;
        /// <summary>
        /// Used to specify whether file notifications are sent to IoT Hub on upload. It evaluates to false by default.
        /// </summary>
        public readonly bool? Notifications;
        /// <summary>
        /// The period of time for which the SAS URI generated by IoT Hub for file upload is valid, specified as an [ISO 8601 timespan duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This value must be between 1 minute and 24 hours, and evaluates to `PT1H` by default.
        /// </summary>
        public readonly string? SasTtl;

        [OutputConstructor]
        private IoTHubFileUpload(
            string connectionString,

            string containerName,

            string? defaultTtl,

            string? lockDuration,

            int? maxDeliveryCount,

            bool? notifications,

            string? sasTtl)
        {
            ConnectionString = connectionString;
            ContainerName = containerName;
            DefaultTtl = defaultTtl;
            LockDuration = lockDuration;
            MaxDeliveryCount = maxDeliveryCount;
            Notifications = notifications;
            SasTtl = sasTtl;
        }
    }
}
