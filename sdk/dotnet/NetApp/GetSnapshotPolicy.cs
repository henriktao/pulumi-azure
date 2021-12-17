// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Azure.NetApp
{
    public static class GetSnapshotPolicy
    {
        /// <summary>
        /// Uses this data source to access information about an existing NetApp Snapshot Policy.
        /// 
        /// ## NetApp Snapshot Policy Usage
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.NetApp.GetSnapshotPolicy.InvokeAsync(new Azure.NetApp.GetSnapshotPolicyArgs
        ///         {
        ///             ResourceGroupName = "acctestRG",
        ///             AccountName = "acctestnetappaccount",
        ///             Name = "example-snapshot-policy",
        ///         }));
        ///         this.Id = example.Apply(example =&gt; example.Id);
        ///         this.Name = example.Apply(example =&gt; example.Name);
        ///         this.Enabled = example.Apply(example =&gt; example.Enabled);
        ///         this.HourlySchedule = example.Apply(example =&gt; example.HourlySchedules);
        ///         this.DailySchedule = example.Apply(example =&gt; example.DailySchedules);
        ///         this.WeeklySchedule = example.Apply(example =&gt; example.WeeklySchedules);
        ///         this.MonthlySchedule = example.Apply(example =&gt; example.MonthlySchedules);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        ///     [Output("name")]
        ///     public Output&lt;string&gt; Name { get; set; }
        ///     [Output("enabled")]
        ///     public Output&lt;string&gt; Enabled { get; set; }
        ///     [Output("hourlySchedule")]
        ///     public Output&lt;string&gt; HourlySchedule { get; set; }
        ///     [Output("dailySchedule")]
        ///     public Output&lt;string&gt; DailySchedule { get; set; }
        ///     [Output("weeklySchedule")]
        ///     public Output&lt;string&gt; WeeklySchedule { get; set; }
        ///     [Output("monthlySchedule")]
        ///     public Output&lt;string&gt; MonthlySchedule { get; set; }
        /// }
        /// ```
        /// </summary>
        public static Task<GetSnapshotPolicyResult> InvokeAsync(GetSnapshotPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotPolicyResult>("azure:netapp/getSnapshotPolicy:getSnapshotPolicy", args ?? new GetSnapshotPolicyArgs(), options.WithVersion());

        /// <summary>
        /// Uses this data source to access information about an existing NetApp Snapshot Policy.
        /// 
        /// ## NetApp Snapshot Policy Usage
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.NetApp.GetSnapshotPolicy.InvokeAsync(new Azure.NetApp.GetSnapshotPolicyArgs
        ///         {
        ///             ResourceGroupName = "acctestRG",
        ///             AccountName = "acctestnetappaccount",
        ///             Name = "example-snapshot-policy",
        ///         }));
        ///         this.Id = example.Apply(example =&gt; example.Id);
        ///         this.Name = example.Apply(example =&gt; example.Name);
        ///         this.Enabled = example.Apply(example =&gt; example.Enabled);
        ///         this.HourlySchedule = example.Apply(example =&gt; example.HourlySchedules);
        ///         this.DailySchedule = example.Apply(example =&gt; example.DailySchedules);
        ///         this.WeeklySchedule = example.Apply(example =&gt; example.WeeklySchedules);
        ///         this.MonthlySchedule = example.Apply(example =&gt; example.MonthlySchedules);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        ///     [Output("name")]
        ///     public Output&lt;string&gt; Name { get; set; }
        ///     [Output("enabled")]
        ///     public Output&lt;string&gt; Enabled { get; set; }
        ///     [Output("hourlySchedule")]
        ///     public Output&lt;string&gt; HourlySchedule { get; set; }
        ///     [Output("dailySchedule")]
        ///     public Output&lt;string&gt; DailySchedule { get; set; }
        ///     [Output("weeklySchedule")]
        ///     public Output&lt;string&gt; WeeklySchedule { get; set; }
        ///     [Output("monthlySchedule")]
        ///     public Output&lt;string&gt; MonthlySchedule { get; set; }
        /// }
        /// ```
        /// </summary>
        public static Output<GetSnapshotPolicyResult> Invoke(GetSnapshotPolicyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSnapshotPolicyResult>("azure:netapp/getSnapshotPolicy:getSnapshotPolicy", args ?? new GetSnapshotPolicyInvokeArgs(), options.WithVersion());
    }


    public sealed class GetSnapshotPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the NetApp account where the NetApp Snapshot Policy exists.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the NetApp Snapshot Policy.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the NetApp Snapshot Policy exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetSnapshotPolicyArgs()
        {
        }
    }

    public sealed class GetSnapshotPolicyInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the NetApp account where the NetApp Snapshot Policy exists.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the NetApp Snapshot Policy.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the NetApp Snapshot Policy exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetSnapshotPolicyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSnapshotPolicyResult
    {
        /// <summary>
        /// The name of the NetApp Account in which the NetApp Snapshot Policy was created.
        /// </summary>
        public readonly string AccountName;
        /// <summary>
        /// Daily snapshot schedule.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSnapshotPolicyDailyScheduleResult> DailySchedules;
        /// <summary>
        /// Defines that the NetApp Snapshot Policy is enabled or not.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// Hourly snapshot schedule.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSnapshotPolicyHourlyScheduleResult> HourlySchedules;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specifies the supported Azure location where the resource exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// List of the days of the month when the snapshots will be created.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSnapshotPolicyMonthlyScheduleResult> MonthlySchedules;
        /// <summary>
        /// The name of the NetApp Snapshot Policy.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name of the resource group where the NetApp Snapshot Policy should be created.
        /// </summary>
        public readonly string ResourceGroupName;
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// Weekly snapshot schedule.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSnapshotPolicyWeeklyScheduleResult> WeeklySchedules;

        [OutputConstructor]
        private GetSnapshotPolicyResult(
            string accountName,

            ImmutableArray<Outputs.GetSnapshotPolicyDailyScheduleResult> dailySchedules,

            bool enabled,

            ImmutableArray<Outputs.GetSnapshotPolicyHourlyScheduleResult> hourlySchedules,

            string id,

            string location,

            ImmutableArray<Outputs.GetSnapshotPolicyMonthlyScheduleResult> monthlySchedules,

            string name,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags,

            ImmutableArray<Outputs.GetSnapshotPolicyWeeklyScheduleResult> weeklySchedules)
        {
            AccountName = accountName;
            DailySchedules = dailySchedules;
            Enabled = enabled;
            HourlySchedules = hourlySchedules;
            Id = id;
            Location = location;
            MonthlySchedules = monthlySchedules;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
            WeeklySchedules = weeklySchedules;
        }
    }
}
