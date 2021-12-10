// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Sql.Inputs
{

    public sealed class ManagedInstanceFailoverGroupReadWriteEndpointFailoverPolicyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Applies only if `mode` is `Automatic`. The grace period in minutes before failover with data loss is attempted.
        /// </summary>
        [Input("graceMinutes")]
        public Input<int>? GraceMinutes { get; set; }

        /// <summary>
        /// The failover mode. Possible values are `Manual`, `Automatic`
        /// </summary>
        [Input("mode", required: true)]
        public Input<string> Mode { get; set; } = null!;

        public ManagedInstanceFailoverGroupReadWriteEndpointFailoverPolicyGetArgs()
        {
        }
    }
}
