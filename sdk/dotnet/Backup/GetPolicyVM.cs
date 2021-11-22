// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Azure.Backup
{
    public static class GetPolicyVM
    {
        /// <summary>
        /// Use this data source to access information about an existing VM Backup Policy.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var policy = Output.Create(Azure.Backup.GetPolicyVM.InvokeAsync(new Azure.Backup.GetPolicyVMArgs
        ///         {
        ///             Name = "policy",
        ///             RecoveryVaultName = "recovery_vault",
        ///             ResourceGroupName = "resource_group",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPolicyVMResult> InvokeAsync(GetPolicyVMArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyVMResult>("azure:backup/getPolicyVM:getPolicyVM", args ?? new GetPolicyVMArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to access information about an existing VM Backup Policy.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var policy = Output.Create(Azure.Backup.GetPolicyVM.InvokeAsync(new Azure.Backup.GetPolicyVMArgs
        ///         {
        ///             Name = "policy",
        ///             RecoveryVaultName = "recovery_vault",
        ///             ResourceGroupName = "resource_group",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetPolicyVMResult> Invoke(GetPolicyVMInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPolicyVMResult>("azure:backup/getPolicyVM:getPolicyVM", args ?? new GetPolicyVMInvokeArgs(), options.WithVersion());
    }


    public sealed class GetPolicyVMArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the VM Backup Policy.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Recovery Services Vault.
        /// </summary>
        [Input("recoveryVaultName", required: true)]
        public string RecoveryVaultName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the VM Backup Policy resides.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPolicyVMArgs()
        {
        }
    }

    public sealed class GetPolicyVMInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the VM Backup Policy.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Recovery Services Vault.
        /// </summary>
        [Input("recoveryVaultName", required: true)]
        public Input<string> RecoveryVaultName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the VM Backup Policy resides.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetPolicyVMInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPolicyVMResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string RecoveryVaultName;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetPolicyVMResult(
            string id,

            string name,

            string recoveryVaultName,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags)
        {
            Id = id;
            Name = name;
            RecoveryVaultName = recoveryVaultName;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
        }
    }
}
