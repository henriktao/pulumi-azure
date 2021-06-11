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
    public sealed class BudgetSubscriptionFilterNotDimension
    {
        /// <summary>
        /// The name of the column to use for the filter. The allowed values are `ChargeType`, `Frequency`, `InvoiceId`, `Meter`, `MeterCategory`, `MeterSubCategory`, `PartNumber`, `PricingModel`, `Product`, `ProductOrderId`, `ProductOrderName`, `PublisherType`, `ReservationId`, `ReservationName`, `ResourceGroupName`, `ResourceGuid`, `ResourceId`, `ResourceLocation`, `ResourceType`, `ServiceFamily`, `ServiceName`, `UnitOfMeasure`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The operator to use for comparison. The allowed values are `In`.
        /// </summary>
        public readonly string? Operator;
        /// <summary>
        /// Specifies a list of values for the column.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private BudgetSubscriptionFilterNotDimension(
            string name,

            string? @operator,

            ImmutableArray<string> values)
        {
            Name = name;
            Operator = @operator;
            Values = values;
        }
    }
}
