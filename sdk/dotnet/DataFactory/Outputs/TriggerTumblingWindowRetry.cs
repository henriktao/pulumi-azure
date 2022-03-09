// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataFactory.Outputs
{

    [OutputType]
    public sealed class TriggerTumblingWindowRetry
    {
        /// <summary>
        /// The maximum retry attempts if the pipeline run failed.
        /// </summary>
        public readonly int Count;
        /// <summary>
        /// The Interval in seconds between each retry if the pipeline run failed.
        /// </summary>
        public readonly int? Interval;

        [OutputConstructor]
        private TriggerTumblingWindowRetry(
            int count,

            int? interval)
        {
            Count = count;
            Interval = interval;
        }
    }
}
