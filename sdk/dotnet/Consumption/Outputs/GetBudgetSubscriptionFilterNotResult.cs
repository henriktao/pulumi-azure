// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Consumption.Outputs
{

    [OutputType]
    public sealed class GetBudgetSubscriptionFilterNotResult
    {
        /// <summary>
        /// A `dimension` block as defined above.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBudgetSubscriptionFilterNotDimensionResult> Dimensions;
        /// <summary>
        /// A `tag` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBudgetSubscriptionFilterNotTagResult> Tags;

        [OutputConstructor]
        private GetBudgetSubscriptionFilterNotResult(
            ImmutableArray<Outputs.GetBudgetSubscriptionFilterNotDimensionResult> dimensions,

            ImmutableArray<Outputs.GetBudgetSubscriptionFilterNotTagResult> tags)
        {
            Dimensions = dimensions;
            Tags = tags;
        }
    }
}
