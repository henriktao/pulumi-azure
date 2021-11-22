// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Azure.PostgreSql
{
    public static class GetFlexibleServer
    {
        /// <summary>
        /// Use this data source to access information about an existing PostgreSQL Flexible Server.
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
        ///         var example = Output.Create(Azure.PostgreSql.GetFlexibleServer.InvokeAsync(new Azure.PostgreSql.GetFlexibleServerArgs
        ///         {
        ///             Name = "existing-postgresql-fs",
        ///             ResourceGroupName = "existing-postgresql-resgroup",
        ///         }));
        ///         this.Id = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFlexibleServerResult> InvokeAsync(GetFlexibleServerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFlexibleServerResult>("azure:postgresql/getFlexibleServer:getFlexibleServer", args ?? new GetFlexibleServerArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to access information about an existing PostgreSQL Flexible Server.
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
        ///         var example = Output.Create(Azure.PostgreSql.GetFlexibleServer.InvokeAsync(new Azure.PostgreSql.GetFlexibleServerArgs
        ///         {
        ///             Name = "existing-postgresql-fs",
        ///             ResourceGroupName = "existing-postgresql-resgroup",
        ///         }));
        ///         this.Id = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFlexibleServerResult> Invoke(GetFlexibleServerInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFlexibleServerResult>("azure:postgresql/getFlexibleServer:getFlexibleServer", args ?? new GetFlexibleServerInvokeArgs(), options.WithVersion());
    }


    public sealed class GetFlexibleServerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this PostgreSQL Flexible Server.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the PostgreSQL Flexible Server exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetFlexibleServerArgs()
        {
        }
    }

    public sealed class GetFlexibleServerInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this PostgreSQL Flexible Server.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the PostgreSQL Flexible Server exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetFlexibleServerInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFlexibleServerResult
    {
        /// <summary>
        /// The Administrator Login for the PostgreSQL Flexible Server.
        /// </summary>
        public readonly string AdministratorLogin;
        /// <summary>
        /// The backup retention days for the PostgreSQL Flexible Server.
        /// </summary>
        public readonly int BackupRetentionDays;
        /// <summary>
        /// The status showing whether the data encryption is enabled with a customer-managed key.
        /// </summary>
        public readonly string CmkEnabled;
        /// <summary>
        /// The ID of the virtual network subnet to create the PostgreSQL Flexible Server.
        /// </summary>
        public readonly string DelegatedSubnetId;
        /// <summary>
        /// The FQDN of the PostgreSQL Flexible Server.
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure Region where the PostgreSQL Flexible Server exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        /// <summary>
        /// Is public network access enabled?
        /// </summary>
        public readonly bool PublicNetworkAccessEnabled;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The SKU Name for the PostgreSQL Flexible Server. The name of the SKU, follows the `tier` + `name` pattern (e.g. `B_Standard_B1ms`, `GP_Standard_D2s_v3`, `MO_Standard_E4s_v3`).
        /// </summary>
        public readonly string SkuName;
        /// <summary>
        /// The max storage allowed for the PostgreSQL Flexible Server.
        /// </summary>
        public readonly int StorageMb;
        /// <summary>
        /// A mapping of tags assigned to the PostgreSQL Flexible Server.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The version of PostgreSQL Flexible Server to use.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetFlexibleServerResult(
            string administratorLogin,

            int backupRetentionDays,

            string cmkEnabled,

            string delegatedSubnetId,

            string fqdn,

            string id,

            string location,

            string name,

            bool publicNetworkAccessEnabled,

            string resourceGroupName,

            string skuName,

            int storageMb,

            ImmutableDictionary<string, string> tags,

            string version)
        {
            AdministratorLogin = administratorLogin;
            BackupRetentionDays = backupRetentionDays;
            CmkEnabled = cmkEnabled;
            DelegatedSubnetId = delegatedSubnetId;
            Fqdn = fqdn;
            Id = id;
            Location = location;
            Name = name;
            PublicNetworkAccessEnabled = publicNetworkAccessEnabled;
            ResourceGroupName = resourceGroupName;
            SkuName = skuName;
            StorageMb = storageMb;
            Tags = tags;
            Version = version;
        }
    }
}
