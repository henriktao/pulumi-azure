// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.CosmosDB.Outputs
{

    [OutputType]
    public sealed class AccountRestore
    {
        /// <summary>
        /// A `database` block as defined below. Changing this forces a new resource to be created.
        /// </summary>
        public readonly ImmutableArray<Outputs.AccountRestoreDatabase> Databases;
        /// <summary>
        /// The creation time of the database or the collection (Datetime Format `RFC 3339`). Changing this forces a new resource to be created.
        /// </summary>
        public readonly string RestoreTimestampInUtc;
        /// <summary>
        /// The resource ID of the restorable database account from which the restore has to be initiated. The example is `/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{restorableDatabaseAccountName}`. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string SourceCosmosdbAccountId;

        [OutputConstructor]
        private AccountRestore(
            ImmutableArray<Outputs.AccountRestoreDatabase> databases,

            string restoreTimestampInUtc,

            string sourceCosmosdbAccountId)
        {
            Databases = databases;
            RestoreTimestampInUtc = restoreTimestampInUtc;
            SourceCosmosdbAccountId = sourceCosmosdbAccountId;
        }
    }
}
