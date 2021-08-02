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
    public sealed class TriggerTumblingWindowPipeline
    {
        public readonly string Name;
        public readonly ImmutableDictionary<string, string>? Parameters;

        [OutputConstructor]
        private TriggerTumblingWindowPipeline(
            string name,

            ImmutableDictionary<string, string>? parameters)
        {
            Name = name;
            Parameters = parameters;
        }
    }
}
