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
    public sealed class LogzTagRuleTagFilter
    {
        /// <summary>
        /// The action for a filtering tag. Possible values are "Include" and "Exclude" is allowed. Note that the `Exclude` takes priority over the `Include`.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// The name of this `tag_filter`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The value of this `tag_filter`.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private LogzTagRuleTagFilter(
            string action,

            string name,

            string? value)
        {
            Action = action;
            Name = name;
            Value = value;
        }
    }
}
