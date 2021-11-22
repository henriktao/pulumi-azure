// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Azure.Billing
{
    public static class GetEnrollmentAccountScope
    {
        /// <summary>
        /// Use this data source to access information about an existing Enrollment Account Billing Scope.
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
        ///         var example = Output.Create(Azure.Billing.GetEnrollmentAccountScope.InvokeAsync(new Azure.Billing.GetEnrollmentAccountScopeArgs
        ///         {
        ///             BillingAccountName = "existing",
        ///             EnrollmentAccountName = "existing",
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
        public static Task<GetEnrollmentAccountScopeResult> InvokeAsync(GetEnrollmentAccountScopeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEnrollmentAccountScopeResult>("azure:billing/getEnrollmentAccountScope:getEnrollmentAccountScope", args ?? new GetEnrollmentAccountScopeArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to access information about an existing Enrollment Account Billing Scope.
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
        ///         var example = Output.Create(Azure.Billing.GetEnrollmentAccountScope.InvokeAsync(new Azure.Billing.GetEnrollmentAccountScopeArgs
        ///         {
        ///             BillingAccountName = "existing",
        ///             EnrollmentAccountName = "existing",
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
        public static Output<GetEnrollmentAccountScopeResult> Invoke(GetEnrollmentAccountScopeInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEnrollmentAccountScopeResult>("azure:billing/getEnrollmentAccountScope:getEnrollmentAccountScope", args ?? new GetEnrollmentAccountScopeInvokeArgs(), options.WithVersion());
    }


    public sealed class GetEnrollmentAccountScopeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Billing Account Name of the Enterprise Account.
        /// </summary>
        [Input("billingAccountName", required: true)]
        public string BillingAccountName { get; set; } = null!;

        /// <summary>
        /// The Enrollment Account Name in the above Enterprise Account.
        /// </summary>
        [Input("enrollmentAccountName", required: true)]
        public string EnrollmentAccountName { get; set; } = null!;

        public GetEnrollmentAccountScopeArgs()
        {
        }
    }

    public sealed class GetEnrollmentAccountScopeInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Billing Account Name of the Enterprise Account.
        /// </summary>
        [Input("billingAccountName", required: true)]
        public Input<string> BillingAccountName { get; set; } = null!;

        /// <summary>
        /// The Enrollment Account Name in the above Enterprise Account.
        /// </summary>
        [Input("enrollmentAccountName", required: true)]
        public Input<string> EnrollmentAccountName { get; set; } = null!;

        public GetEnrollmentAccountScopeInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEnrollmentAccountScopeResult
    {
        public readonly string BillingAccountName;
        public readonly string EnrollmentAccountName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetEnrollmentAccountScopeResult(
            string billingAccountName,

            string enrollmentAccountName,

            string id)
        {
            BillingAccountName = billingAccountName;
            EnrollmentAccountName = enrollmentAccountName;
            Id = id;
        }
    }
}
