// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Azure.Storage
{
    public static class GetPolicy
    {
        /// <summary>
        /// Use this data source to access information about an existing Storage Management Policy.
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
        ///         var exampleAccount = Output.Create(Azure.Storage.GetAccount.InvokeAsync(new Azure.Storage.GetAccountArgs
        ///         {
        ///             Name = "storageaccountname",
        ///             ResourceGroupName = "resourcegroupname",
        ///         }));
        ///         var examplePolicy = Output.Create(Azure.Storage.GetPolicy.InvokeAsync(new Azure.Storage.GetPolicyArgs
        ///         {
        ///             StorageAccountId = azurerm_storage_account.Example.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPolicyResult> InvokeAsync(GetPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyResult>("azure:storage/getPolicy:getPolicy", args ?? new GetPolicyArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to access information about an existing Storage Management Policy.
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
        ///         var exampleAccount = Output.Create(Azure.Storage.GetAccount.InvokeAsync(new Azure.Storage.GetAccountArgs
        ///         {
        ///             Name = "storageaccountname",
        ///             ResourceGroupName = "resourcegroupname",
        ///         }));
        ///         var examplePolicy = Output.Create(Azure.Storage.GetPolicy.InvokeAsync(new Azure.Storage.GetPolicyArgs
        ///         {
        ///             StorageAccountId = azurerm_storage_account.Example.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetPolicyResult> Invoke(GetPolicyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPolicyResult>("azure:storage/getPolicy:getPolicy", args ?? new GetPolicyInvokeArgs(), options.WithVersion());
    }


    public sealed class GetPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the id of the storage account to retrieve the management policy for.
        /// </summary>
        [Input("storageAccountId", required: true)]
        public string StorageAccountId { get; set; } = null!;

        public GetPolicyArgs()
        {
        }
    }

    public sealed class GetPolicyInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the id of the storage account to retrieve the management policy for.
        /// </summary>
        [Input("storageAccountId", required: true)]
        public Input<string> StorageAccountId { get; set; } = null!;

        public GetPolicyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPolicyResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A `rule` block as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPolicyRuleResult> Rules;
        public readonly string StorageAccountId;

        [OutputConstructor]
        private GetPolicyResult(
            string id,

            ImmutableArray<Outputs.GetPolicyRuleResult> rules,

            string storageAccountId)
        {
            Id = id;
            Rules = rules;
            StorageAccountId = storageAccountId;
        }
    }
}
