// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataFactory.Inputs
{

    public sealed class TriggerTumblingWindowRetryGetArgs : Pulumi.ResourceArgs
    {
        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        [Input("interval")]
        public Input<int>? Interval { get; set; }

        public TriggerTumblingWindowRetryGetArgs()
        {
        }
    }
}
