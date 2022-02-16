// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.SiteRecovery
{
    public static class GetReplicationPolicy
    {
        /// <summary>
        /// Use this data source to access information about an existing Azure Site Recovery replication policy.
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
        ///         var policy = Output.Create(Azure.SiteRecovery.GetReplicationPolicy.InvokeAsync(new Azure.SiteRecovery.GetReplicationPolicyArgs
        ///         {
        ///             Name = "replication-policy",
        ///             RecoveryVaultName = "tfex-recovery_vault",
        ///             ResourceGroupName = "tfex-resource_group",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetReplicationPolicyResult> InvokeAsync(GetReplicationPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetReplicationPolicyResult>("azure:siterecovery/getReplicationPolicy:getReplicationPolicy", args ?? new GetReplicationPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Azure Site Recovery replication policy.
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
        ///         var policy = Output.Create(Azure.SiteRecovery.GetReplicationPolicy.InvokeAsync(new Azure.SiteRecovery.GetReplicationPolicyArgs
        ///         {
        ///             Name = "replication-policy",
        ///             RecoveryVaultName = "tfex-recovery_vault",
        ///             ResourceGroupName = "tfex-resource_group",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetReplicationPolicyResult> Invoke(GetReplicationPolicyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetReplicationPolicyResult>("azure:siterecovery/getReplicationPolicy:getReplicationPolicy", args ?? new GetReplicationPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReplicationPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Azure Site Recovery replication policy.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Recovery Services Vault that the Azure Site Recovery replication policy is associated witth.
        /// </summary>
        [Input("recoveryVaultName", required: true)]
        public string RecoveryVaultName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the associated Azure Site Recovery replication policy resides.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetReplicationPolicyArgs()
        {
        }
    }

    public sealed class GetReplicationPolicyInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Azure Site Recovery replication policy.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Recovery Services Vault that the Azure Site Recovery replication policy is associated witth.
        /// </summary>
        [Input("recoveryVaultName", required: true)]
        public Input<string> RecoveryVaultName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the associated Azure Site Recovery replication policy resides.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetReplicationPolicyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetReplicationPolicyResult
    {
        /// <summary>
        /// Specifies the frequency (in minutes) at which to create application consistent recovery.
        /// </summary>
        public readonly int ApplicationConsistentSnapshotFrequencyInMinutes;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// The duration in minutes for which the recovery points need to be stored.
        /// </summary>
        public readonly int RecoveryPointRetentionInMinutes;
        public readonly string RecoveryVaultName;
        public readonly string ResourceGroupName;

        [OutputConstructor]
        private GetReplicationPolicyResult(
            int applicationConsistentSnapshotFrequencyInMinutes,

            string id,

            string name,

            int recoveryPointRetentionInMinutes,

            string recoveryVaultName,

            string resourceGroupName)
        {
            ApplicationConsistentSnapshotFrequencyInMinutes = applicationConsistentSnapshotFrequencyInMinutes;
            Id = id;
            Name = name;
            RecoveryPointRetentionInMinutes = recoveryPointRetentionInMinutes;
            RecoveryVaultName = recoveryVaultName;
            ResourceGroupName = resourceGroupName;
        }
    }
}
