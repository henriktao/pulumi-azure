// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MySql
{
    public static class GetServer
    {
        /// <summary>
        /// Use this data source to access information about an existing MySQL Server.
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
        ///         var example = Output.Create(Azure.MySql.GetServer.InvokeAsync(new Azure.MySql.GetServerArgs
        ///         {
        ///             Name = "existingMySqlServer",
        ///             ResourceGroupName = "existingResGroup",
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
        public static Task<GetServerResult> InvokeAsync(GetServerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerResult>("azure:mysql/getServer:getServer", args ?? new GetServerArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing MySQL Server.
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
        ///         var example = Output.Create(Azure.MySql.GetServer.InvokeAsync(new Azure.MySql.GetServerArgs
        ///         {
        ///             Name = "existingMySqlServer",
        ///             ResourceGroupName = "existingResGroup",
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
        public static Output<GetServerResult> Invoke(GetServerInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetServerResult>("azure:mysql/getServer:getServer", args ?? new GetServerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the MySQL Server.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group for the MySQL Server.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetServerArgs()
        {
        }
    }

    public sealed class GetServerInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the MySQL Server.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group for the MySQL Server.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetServerInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServerResult
    {
        /// <summary>
        /// The Administrator Login for the MySQL Server.
        /// </summary>
        public readonly string AdministratorLogin;
        /// <summary>
        /// The auto grow setting for this MySQL Server.
        /// </summary>
        public readonly bool AutoGrowEnabled;
        /// <summary>
        /// The backup retention days for this MySQL server.
        /// </summary>
        public readonly int BackupRetentionDays;
        /// <summary>
        /// The FQDN of the MySQL Server.
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// The geo redundant backup setting for this MySQL Server.
        /// </summary>
        public readonly bool GeoRedundantBackupEnabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// An `identity` block for this MySQL server as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServerIdentityResult> Identities;
        /// <summary>
        /// Whether or not infrastructure is encrypted for this MySQL Server.
        /// </summary>
        public readonly bool InfrastructureEncryptionEnabled;
        /// <summary>
        /// The Azure location where the resource exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        /// <summary>
        /// Whether or not public network access is allowed for this MySQL Server.
        /// </summary>
        public readonly bool PublicNetworkAccessEnabled;
        public readonly string ResourceGroupName;
        public readonly string RestorePointInTime;
        /// <summary>
        /// The SKU Name for this MySQL Server.
        /// </summary>
        public readonly string SkuName;
        /// <summary>
        /// Specifies if SSL should be enforced on connections for this MySQL Server.
        /// </summary>
        public readonly bool SslEnforcementEnabled;
        /// <summary>
        /// The minimum TLS version to support for this MySQL Server.
        /// </summary>
        public readonly string SslMinimalTlsVersionEnforced;
        /// <summary>
        /// Max storage allowed for this MySQL Server.
        /// </summary>
        public readonly int StorageMb;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// ---
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// Threat detection policy configuration, known in the API as Server Security Alerts Policy. The `threat_detection_policy` block exports fields documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServerThreatDetectionPolicyResult> ThreatDetectionPolicies;
        /// <summary>
        /// The version of this MySQL Server.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetServerResult(
            string administratorLogin,

            bool autoGrowEnabled,

            int backupRetentionDays,

            string fqdn,

            bool geoRedundantBackupEnabled,

            string id,

            ImmutableArray<Outputs.GetServerIdentityResult> identities,

            bool infrastructureEncryptionEnabled,

            string location,

            string name,

            bool publicNetworkAccessEnabled,

            string resourceGroupName,

            string restorePointInTime,

            string skuName,

            bool sslEnforcementEnabled,

            string sslMinimalTlsVersionEnforced,

            int storageMb,

            ImmutableDictionary<string, string> tags,

            ImmutableArray<Outputs.GetServerThreatDetectionPolicyResult> threatDetectionPolicies,

            string version)
        {
            AdministratorLogin = administratorLogin;
            AutoGrowEnabled = autoGrowEnabled;
            BackupRetentionDays = backupRetentionDays;
            Fqdn = fqdn;
            GeoRedundantBackupEnabled = geoRedundantBackupEnabled;
            Id = id;
            Identities = identities;
            InfrastructureEncryptionEnabled = infrastructureEncryptionEnabled;
            Location = location;
            Name = name;
            PublicNetworkAccessEnabled = publicNetworkAccessEnabled;
            ResourceGroupName = resourceGroupName;
            RestorePointInTime = restorePointInTime;
            SkuName = skuName;
            SslEnforcementEnabled = sslEnforcementEnabled;
            SslMinimalTlsVersionEnforced = sslMinimalTlsVersionEnforced;
            StorageMb = storageMb;
            Tags = tags;
            ThreatDetectionPolicies = threatDetectionPolicies;
            Version = version;
        }
    }
}
