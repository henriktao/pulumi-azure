// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.CosmosDB
{
    public static class GetRestorableDatabaseAccounts
    {
        /// <summary>
        /// Use this data source to access information about Cosmos DB Restorable Database Accounts.
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
        ///         var example = Output.Create(Azure.CosmosDB.GetRestorableDatabaseAccounts.InvokeAsync(new Azure.CosmosDB.GetRestorableDatabaseAccountsArgs
        ///         {
        ///             Name = "example-ca",
        ///             Location = "West Europe",
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
        public static Task<GetRestorableDatabaseAccountsResult> InvokeAsync(GetRestorableDatabaseAccountsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRestorableDatabaseAccountsResult>("azure:cosmosdb/getRestorableDatabaseAccounts:getRestorableDatabaseAccounts", args ?? new GetRestorableDatabaseAccountsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about Cosmos DB Restorable Database Accounts.
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
        ///         var example = Output.Create(Azure.CosmosDB.GetRestorableDatabaseAccounts.InvokeAsync(new Azure.CosmosDB.GetRestorableDatabaseAccountsArgs
        ///         {
        ///             Name = "example-ca",
        ///             Location = "West Europe",
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
        public static Output<GetRestorableDatabaseAccountsResult> Invoke(GetRestorableDatabaseAccountsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRestorableDatabaseAccountsResult>("azure:cosmosdb/getRestorableDatabaseAccounts:getRestorableDatabaseAccounts", args ?? new GetRestorableDatabaseAccountsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRestorableDatabaseAccountsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The location where the Cosmos DB Database Account.
        /// </summary>
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        /// <summary>
        /// The name of this Cosmos DB Database Account.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetRestorableDatabaseAccountsArgs()
        {
        }
    }

    public sealed class GetRestorableDatabaseAccountsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The location where the Cosmos DB Database Account.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of this Cosmos DB Database Account.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetRestorableDatabaseAccountsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRestorableDatabaseAccountsResult
    {
        /// <summary>
        /// One or more `accounts` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRestorableDatabaseAccountsAccountResult> Accounts;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The location of the regional Cosmos DB Restorable Database Account.
        /// </summary>
        public readonly string Location;
        public readonly string Name;

        [OutputConstructor]
        private GetRestorableDatabaseAccountsResult(
            ImmutableArray<Outputs.GetRestorableDatabaseAccountsAccountResult> accounts,

            string id,

            string location,

            string name)
        {
            Accounts = accounts;
            Id = id;
            Location = location;
            Name = name;
        }
    }
}
