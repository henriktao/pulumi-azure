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
    public sealed class BudgetManagementGroupFilter
    {
        /// <summary>
        /// One or more `dimension` blocks as defined below to filter the budget on.
        /// </summary>
        public readonly ImmutableArray<Outputs.BudgetManagementGroupFilterDimension> Dimensions;
        /// <summary>
        /// A `not` block as defined below to filter the budget on.
        /// </summary>
        public readonly Outputs.BudgetManagementGroupFilterNot? Not;
        /// <summary>
        /// One or more `tag` blocks as defined below to filter the budget on.
        /// </summary>
        public readonly ImmutableArray<Outputs.BudgetManagementGroupFilterTag> Tags;

        [OutputConstructor]
        private BudgetManagementGroupFilter(
            ImmutableArray<Outputs.BudgetManagementGroupFilterDimension> dimensions,

            Outputs.BudgetManagementGroupFilterNot? not,

            ImmutableArray<Outputs.BudgetManagementGroupFilterTag> tags)
        {
            Dimensions = dimensions;
            Not = not;
            Tags = tags;
        }
    }
}
