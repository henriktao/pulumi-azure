// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Automation
{
    public static class GetDateTimeVariable
    {
        /// <summary>
        /// Use this data source to access information about an existing Automation Datetime Variable.
        /// 
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
        ///         var example = Output.Create(Azure.Automation.GetDateTimeVariable.InvokeAsync(new Azure.Automation.GetDateTimeVariableArgs
        ///         {
        ///             Name = "tfex-example-var",
        ///             ResourceGroupName = "tfex-example-rg",
        ///             AutomationAccountName = "tfex-example-account",
        ///         }));
        ///         this.VariableId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("variableId")]
        ///     public Output&lt;string&gt; VariableId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDateTimeVariableResult> InvokeAsync(GetDateTimeVariableArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDateTimeVariableResult>("azure:automation/getDateTimeVariable:getDateTimeVariable", args ?? new GetDateTimeVariableArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Automation Datetime Variable.
        /// 
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
        ///         var example = Output.Create(Azure.Automation.GetDateTimeVariable.InvokeAsync(new Azure.Automation.GetDateTimeVariableArgs
        ///         {
        ///             Name = "tfex-example-var",
        ///             ResourceGroupName = "tfex-example-rg",
        ///             AutomationAccountName = "tfex-example-account",
        ///         }));
        ///         this.VariableId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("variableId")]
        ///     public Output&lt;string&gt; VariableId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDateTimeVariableResult> Invoke(GetDateTimeVariableInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDateTimeVariableResult>("azure:automation/getDateTimeVariable:getDateTimeVariable", args ?? new GetDateTimeVariableInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDateTimeVariableArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the automation account in which the Automation Variable exists.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public string AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Automation Variable.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the automation account exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDateTimeVariableArgs()
        {
        }
    }

    public sealed class GetDateTimeVariableInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the automation account in which the Automation Variable exists.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public Input<string> AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Automation Variable.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the automation account exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetDateTimeVariableInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDateTimeVariableResult
    {
        public readonly string AutomationAccountName;
        /// <summary>
        /// The description of the Automation Variable.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Specifies if the Automation Variable is encrypted. Defaults to `false`.
        /// </summary>
        public readonly bool Encrypted;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The value of the Automation Variable in the [RFC3339 Section 5.6 Internet Date/Time Format](https://tools.ietf.org/html/rfc3339#section-5.6).
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetDateTimeVariableResult(
            string automationAccountName,

            string description,

            bool encrypted,

            string id,

            string name,

            string resourceGroupName,

            string value)
        {
            AutomationAccountName = automationAccountName;
            Description = description;
            Encrypted = encrypted;
            Id = id;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Value = value;
        }
    }
}
