// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Automation.Inputs
{

    public sealed class AccountIdentityArgs : Pulumi.ResourceArgs
    {
        [Input("identityIds")]
        private InputList<string>? _identityIds;

        /// <summary>
        /// The ID of the User Assigned Identity which should be assigned to this Automation Account.
        /// </summary>
        public InputList<string> IdentityIds
        {
            get => _identityIds ?? (_identityIds = new InputList<string>());
            set => _identityIds = value;
        }

        [Input("principalId")]
        public Input<string>? PrincipalId { get; set; }

        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The type of identity used for the automation account. Possible values are `SystemAssigned`, `UserAssigned` and `SystemAssigned, UserAssigned`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public AccountIdentityArgs()
        {
        }
    }
}
