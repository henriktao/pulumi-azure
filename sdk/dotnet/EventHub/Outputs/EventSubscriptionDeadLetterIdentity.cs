// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventHub.Outputs
{

    [OutputType]
    public sealed class EventSubscriptionDeadLetterIdentity
    {
        /// <summary>
        /// Specifies the type of Managed Service Identity that is used for dead lettering. Allowed value is `SystemAssigned`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private EventSubscriptionDeadLetterIdentity(string type)
        {
            Type = type;
        }
    }
}
