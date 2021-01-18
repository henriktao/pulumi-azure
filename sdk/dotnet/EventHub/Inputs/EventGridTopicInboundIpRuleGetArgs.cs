// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventHub.Inputs
{

    public sealed class EventGridTopicInboundIpRuleGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action to take when the rule is matched. Possible values are `Allow`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// The ip mask (CIDR) to match on.
        /// </summary>
        [Input("ipMask", required: true)]
        public Input<string> IpMask { get; set; } = null!;

        public EventGridTopicInboundIpRuleGetArgs()
        {
        }
    }
}
