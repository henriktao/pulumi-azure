// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Iot.Inputs
{

    public sealed class IoTHubIdentityGetArgs : Pulumi.ResourceArgs
    {
        [Input("identityIds")]
        private InputList<string>? _identityIds;

        /// <summary>
        /// A list of User Managed Identity ID's which should be assigned to the Iot Hub.
        /// </summary>
        public InputList<string> IdentityIds
        {
            get => _identityIds ?? (_identityIds = new InputList<string>());
            set => _identityIds = value;
        }

        /// <summary>
        /// The ID of the System Managed Service Principal.
        /// </summary>
        [Input("principalId")]
        public Input<string>? PrincipalId { get; set; }

        /// <summary>
        /// The ID of the Tenant the System Managed Service Principal is assigned in.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The type of Managed Identity which should be assigned to the Iot Hub. Possible values are `SystemAssigned`, `UserAssigned` and `SystemAssigned, UserAssigned`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public IoTHubIdentityGetArgs()
        {
        }
    }
}
