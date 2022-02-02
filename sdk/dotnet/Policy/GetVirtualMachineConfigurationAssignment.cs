// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Policy
{
    public static class GetVirtualMachineConfigurationAssignment
    {
        /// <summary>
        /// Use this data source to access information about an existing Guest Configuration Policy.
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
        ///         var example = Output.Create(Azure.Policy.GetVirtualMachineConfigurationAssignment.InvokeAsync(new Azure.Policy.GetVirtualMachineConfigurationAssignmentArgs
        ///         {
        ///             Name = "AzureWindowsBaseline",
        ///             ResourceGroupName = "example-RG",
        ///             VirtualMachineName = "example-vm",
        ///         }));
        ///         this.ComplianceStatus = example.Apply(example =&gt; example.ComplianceStatus);
        ///     }
        /// 
        ///     [Output("complianceStatus")]
        ///     public Output&lt;string&gt; ComplianceStatus { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVirtualMachineConfigurationAssignmentResult> InvokeAsync(GetVirtualMachineConfigurationAssignmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualMachineConfigurationAssignmentResult>("azure:policy/getVirtualMachineConfigurationAssignment:getVirtualMachineConfigurationAssignment", args ?? new GetVirtualMachineConfigurationAssignmentArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Guest Configuration Policy.
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
        ///         var example = Output.Create(Azure.Policy.GetVirtualMachineConfigurationAssignment.InvokeAsync(new Azure.Policy.GetVirtualMachineConfigurationAssignmentArgs
        ///         {
        ///             Name = "AzureWindowsBaseline",
        ///             ResourceGroupName = "example-RG",
        ///             VirtualMachineName = "example-vm",
        ///         }));
        ///         this.ComplianceStatus = example.Apply(example =&gt; example.ComplianceStatus);
        ///     }
        /// 
        ///     [Output("complianceStatus")]
        ///     public Output&lt;string&gt; ComplianceStatus { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetVirtualMachineConfigurationAssignmentResult> Invoke(GetVirtualMachineConfigurationAssignmentInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVirtualMachineConfigurationAssignmentResult>("azure:policy/getVirtualMachineConfigurationAssignment:getVirtualMachineConfigurationAssignment", args ?? new GetVirtualMachineConfigurationAssignmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualMachineConfigurationAssignmentArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Guest Configuration Assignment.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the Name of the Resource Group where the Guest Configuration Assignment exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Only retrieve Policy Set Definitions from this Management Group.
        /// </summary>
        [Input("virtualMachineName", required: true)]
        public string VirtualMachineName { get; set; } = null!;

        public GetVirtualMachineConfigurationAssignmentArgs()
        {
        }
    }

    public sealed class GetVirtualMachineConfigurationAssignmentInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Guest Configuration Assignment.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the Name of the Resource Group where the Guest Configuration Assignment exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Only retrieve Policy Set Definitions from this Management Group.
        /// </summary>
        [Input("virtualMachineName", required: true)]
        public Input<string> VirtualMachineName { get; set; } = null!;

        public GetVirtualMachineConfigurationAssignmentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVirtualMachineConfigurationAssignmentResult
    {
        /// <summary>
        /// Combined hash of the configuration package and parameters.
        /// </summary>
        public readonly string AssignmentHash;
        /// <summary>
        /// A value indicating compliance status of the machine for the assigned guest configuration. Possible return values are `Compliant`, `NonCompliant` and `Pending`.
        /// </summary>
        public readonly string ComplianceStatus;
        /// <summary>
        /// The content hash for the Guest Configuration package.
        /// </summary>
        public readonly string ContentHash;
        /// <summary>
        /// The content URI where the Guest Configuration package is stored.
        /// </summary>
        public readonly string ContentUri;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Date and time, in RFC3339 format, when the machines compliance status was last checked.
        /// </summary>
        public readonly string LastComplianceStatusChecked;
        /// <summary>
        /// The ID of the latest report for the guest configuration assignment.
        /// </summary>
        public readonly string LatestReportId;
        public readonly string Name;
        public readonly string ResourceGroupName;
        public readonly string VirtualMachineName;

        [OutputConstructor]
        private GetVirtualMachineConfigurationAssignmentResult(
            string assignmentHash,

            string complianceStatus,

            string contentHash,

            string contentUri,

            string id,

            string lastComplianceStatusChecked,

            string latestReportId,

            string name,

            string resourceGroupName,

            string virtualMachineName)
        {
            AssignmentHash = assignmentHash;
            ComplianceStatus = complianceStatus;
            ContentHash = contentHash;
            ContentUri = contentUri;
            Id = id;
            LastComplianceStatusChecked = lastComplianceStatusChecked;
            LatestReportId = latestReportId;
            Name = name;
            ResourceGroupName = resourceGroupName;
            VirtualMachineName = virtualMachineName;
        }
    }
}
