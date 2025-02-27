// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventGrid.Inputs
{

    public sealed class TopicIdentityArgs : Pulumi.ResourceArgs
    {
        [Input("identityIds")]
        private InputList<string>? _identityIds;

        /// <summary>
        /// Specifies a list of user managed identity ids to be assigned. Required if `type` is `UserAssigned`.
        /// </summary>
        public InputList<string> IdentityIds
        {
            get => _identityIds ?? (_identityIds = new InputList<string>());
            set => _identityIds = value;
        }

        /// <summary>
        /// Specifies the Principal ID of the System Assigned Managed Service Identity that is configured on this Event Grid Topic.
        /// </summary>
        [Input("principalId")]
        public Input<string>? PrincipalId { get; set; }

        /// <summary>
        /// Specifies the Tenant ID of the System Assigned Managed Service Identity that is configured on this Event Grid Topic.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Specifies the identity type of Event Grid Topic. Possible values are `SystemAssigned` (where Azure will generate a Principal for you) or `UserAssigned` where you can specify the User Assigned Managed Identity IDs in the `identity_ids` field.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public TopicIdentityArgs()
        {
        }
    }
}
