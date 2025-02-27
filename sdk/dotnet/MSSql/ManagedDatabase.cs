// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MSSql
{
    /// <summary>
    /// Manages an Azure SQL Azure Managed Database for a SQL Managed Instance.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Azure.MSSql.ManagedDatabase("example", new Azure.MSSql.ManagedDatabaseArgs
    ///         {
    ///             ManagedInstanceId = azurerm_mssql_managed_instance.Example.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// SQL Managed Databases can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:mssql/managedDatabase:ManagedDatabase example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/managedInstances/myserver/databases/mydatabase
    /// ```
    /// </summary>
    [AzureResourceType("azure:mssql/managedDatabase:ManagedDatabase")]
    public partial class ManagedDatabase : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Azure SQL Managed Instance on which to create this Managed Database. Changing this forces a new resource to be created.
        /// </summary>
        [Output("managedInstanceId")]
        public Output<string> ManagedInstanceId { get; private set; } = null!;

        /// <summary>
        /// The name of the Managed Database to create. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedDatabase resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedDatabase(string name, ManagedDatabaseArgs args, CustomResourceOptions? options = null)
            : base("azure:mssql/managedDatabase:ManagedDatabase", name, args ?? new ManagedDatabaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedDatabase(string name, Input<string> id, ManagedDatabaseState? state = null, CustomResourceOptions? options = null)
            : base("azure:mssql/managedDatabase:ManagedDatabase", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ManagedDatabase resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedDatabase Get(string name, Input<string> id, ManagedDatabaseState? state = null, CustomResourceOptions? options = null)
        {
            return new ManagedDatabase(name, id, state, options);
        }
    }

    public sealed class ManagedDatabaseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Azure SQL Managed Instance on which to create this Managed Database. Changing this forces a new resource to be created.
        /// </summary>
        [Input("managedInstanceId", required: true)]
        public Input<string> ManagedInstanceId { get; set; } = null!;

        /// <summary>
        /// The name of the Managed Database to create. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ManagedDatabaseArgs()
        {
        }
    }

    public sealed class ManagedDatabaseState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Azure SQL Managed Instance on which to create this Managed Database. Changing this forces a new resource to be created.
        /// </summary>
        [Input("managedInstanceId")]
        public Input<string>? ManagedInstanceId { get; set; }

        /// <summary>
        /// The name of the Managed Database to create. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ManagedDatabaseState()
        {
        }
    }
}
