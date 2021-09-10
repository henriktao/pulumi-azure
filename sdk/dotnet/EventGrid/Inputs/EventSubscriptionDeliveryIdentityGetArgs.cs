// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventGrid.Inputs
{

    public sealed class EventSubscriptionDeliveryIdentityGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the type of Managed Service Identity that is used for event delivery. Allowed value is `SystemAssigned`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public EventSubscriptionDeliveryIdentityGetArgs()
        {
        }
    }
}
