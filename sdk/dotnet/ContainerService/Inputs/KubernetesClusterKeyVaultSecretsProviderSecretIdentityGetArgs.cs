// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService.Inputs
{

    public sealed class KubernetesClusterKeyVaultSecretsProviderSecretIdentityGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Client ID of the user-defined Managed Identity to be assigned to the Kubelets. If not specified a Managed Identity is created automatically.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The Object ID of the user-defined Managed Identity assigned to the Kubelets.If not specified a Managed Identity is created automatically.
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        /// <summary>
        /// The ID of a user assigned identity.
        /// </summary>
        [Input("userAssignedIdentityId")]
        public Input<string>? UserAssignedIdentityId { get; set; }

        public KubernetesClusterKeyVaultSecretsProviderSecretIdentityGetArgs()
        {
        }
    }
}
