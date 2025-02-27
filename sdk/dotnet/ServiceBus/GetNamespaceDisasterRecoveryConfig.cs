// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ServiceBus
{
    public static class GetNamespaceDisasterRecoveryConfig
    {
        public static Task<GetNamespaceDisasterRecoveryConfigResult> InvokeAsync(GetNamespaceDisasterRecoveryConfigArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNamespaceDisasterRecoveryConfigResult>("azure:servicebus/getNamespaceDisasterRecoveryConfig:getNamespaceDisasterRecoveryConfig", args ?? new GetNamespaceDisasterRecoveryConfigArgs(), options.WithDefaults());

        public static Output<GetNamespaceDisasterRecoveryConfigResult> Invoke(GetNamespaceDisasterRecoveryConfigInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNamespaceDisasterRecoveryConfigResult>("azure:servicebus/getNamespaceDisasterRecoveryConfig:getNamespaceDisasterRecoveryConfig", args ?? new GetNamespaceDisasterRecoveryConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNamespaceDisasterRecoveryConfigArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNamespaceDisasterRecoveryConfigArgs()
        {
        }
    }

    public sealed class GetNamespaceDisasterRecoveryConfigInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetNamespaceDisasterRecoveryConfigInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNamespaceDisasterRecoveryConfigResult
    {
        public readonly string DefaultPrimaryKey;
        public readonly string DefaultSecondaryKey;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string NamespaceName;
        public readonly string PartnerNamespaceId;
        public readonly string PrimaryConnectionStringAlias;
        public readonly string ResourceGroupName;
        public readonly string SecondaryConnectionStringAlias;

        [OutputConstructor]
        private GetNamespaceDisasterRecoveryConfigResult(
            string defaultPrimaryKey,

            string defaultSecondaryKey,

            string id,

            string name,

            string namespaceName,

            string partnerNamespaceId,

            string primaryConnectionStringAlias,

            string resourceGroupName,

            string secondaryConnectionStringAlias)
        {
            DefaultPrimaryKey = defaultPrimaryKey;
            DefaultSecondaryKey = defaultSecondaryKey;
            Id = id;
            Name = name;
            NamespaceName = namespaceName;
            PartnerNamespaceId = partnerNamespaceId;
            PrimaryConnectionStringAlias = primaryConnectionStringAlias;
            ResourceGroupName = resourceGroupName;
            SecondaryConnectionStringAlias = secondaryConnectionStringAlias;
        }
    }
}
