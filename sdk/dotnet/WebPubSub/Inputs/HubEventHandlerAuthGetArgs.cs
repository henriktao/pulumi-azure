// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.WebPubSub.Inputs
{

    public sealed class HubEventHandlerAuthGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify the identity ID of the target resource.
        /// </summary>
        [Input("managedIdentityId", required: true)]
        public Input<string> ManagedIdentityId { get; set; } = null!;

        public HubEventHandlerAuthGetArgs()
        {
        }
    }
}
