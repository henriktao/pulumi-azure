// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.CosmosDB.Inputs
{

    public sealed class AccountAnalyticalStorageGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The schema type of the Analytical Storage for this Cosmos DB account. Possible values are `FullFidelity` and `WellDefined`.
        /// </summary>
        [Input("schemaType", required: true)]
        public Input<string> SchemaType { get; set; } = null!;

        public AccountAnalyticalStorageGetArgs()
        {
        }
    }
}
