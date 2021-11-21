// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring.Outputs
{

    [OutputType]
    public sealed class LogzMonitorPlan
    {
        /// <summary>
        /// Different billing cycles. Possible values are `MONTHLY` or `WEEKLY`. Changing this forces a new logz Monitor to be created.
        /// </summary>
        public readonly string BillingCycle;
        /// <summary>
        /// Date when plan was applied. Changing this forces a new logz Monitor to be created.
        /// </summary>
        public readonly string EffectiveDate;
        /// <summary>
        /// Plan id as published by Logz. Possible values are `100gb14days`. Changing this forces a new logz Monitor to be created.
        /// </summary>
        public readonly string PlanId;
        /// <summary>
        /// Different usage types. Possible values are `PAYG` or `COMMITTED`. Changing this forces a new logz Monitor to be created.
        /// </summary>
        public readonly string UsageType;

        [OutputConstructor]
        private LogzMonitorPlan(
            string billingCycle,

            string effectiveDate,

            string planId,

            string usageType)
        {
            BillingCycle = billingCycle;
            EffectiveDate = effectiveDate;
            PlanId = planId;
            UsageType = usageType;
        }
    }
}
