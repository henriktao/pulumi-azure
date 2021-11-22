// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Azure.Advisor
{
    public static class GetRecommendations
    {
        /// <summary>
        /// Use this data source to access information about an existing Advisor Recommendations.
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
        ///         var example = Output.Create(Azure.Advisor.GetRecommendations.InvokeAsync(new Azure.Advisor.GetRecommendationsArgs
        ///         {
        ///             FilterByCategories = 
        ///             {
        ///                 "security",
        ///                 "cost",
        ///             },
        ///             FilterByResourceGroups = 
        ///             {
        ///                 "example-resgroups",
        ///             },
        ///         }));
        ///         this.Recommendations = example.Apply(example =&gt; example.Recommendations);
        ///     }
        /// 
        ///     [Output("recommendations")]
        ///     public Output&lt;string&gt; Recommendations { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRecommendationsResult> InvokeAsync(GetRecommendationsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRecommendationsResult>("azure:advisor/getRecommendations:getRecommendations", args ?? new GetRecommendationsArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to access information about an existing Advisor Recommendations.
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
        ///         var example = Output.Create(Azure.Advisor.GetRecommendations.InvokeAsync(new Azure.Advisor.GetRecommendationsArgs
        ///         {
        ///             FilterByCategories = 
        ///             {
        ///                 "security",
        ///                 "cost",
        ///             },
        ///             FilterByResourceGroups = 
        ///             {
        ///                 "example-resgroups",
        ///             },
        ///         }));
        ///         this.Recommendations = example.Apply(example =&gt; example.Recommendations);
        ///     }
        /// 
        ///     [Output("recommendations")]
        ///     public Output&lt;string&gt; Recommendations { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRecommendationsResult> Invoke(GetRecommendationsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRecommendationsResult>("azure:advisor/getRecommendations:getRecommendations", args ?? new GetRecommendationsInvokeArgs(), options.WithVersion());
    }


    public sealed class GetRecommendationsArgs : Pulumi.InvokeArgs
    {
        [Input("filterByCategories")]
        private List<string>? _filterByCategories;

        /// <summary>
        /// Specifies a list of categories in which the Advisor Recommendations will be listed. Possible values are `HighAvailability`, `Security`, `Performance`, `Cost` and `OperationalExcellence`.
        /// </summary>
        public List<string> FilterByCategories
        {
            get => _filterByCategories ?? (_filterByCategories = new List<string>());
            set => _filterByCategories = value;
        }

        [Input("filterByResourceGroups")]
        private List<string>? _filterByResourceGroups;

        /// <summary>
        /// Specifies a list of resource groups about which the Advisor Recommendations will be listed.
        /// </summary>
        public List<string> FilterByResourceGroups
        {
            get => _filterByResourceGroups ?? (_filterByResourceGroups = new List<string>());
            set => _filterByResourceGroups = value;
        }

        public GetRecommendationsArgs()
        {
        }
    }

    public sealed class GetRecommendationsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("filterByCategories")]
        private InputList<string>? _filterByCategories;

        /// <summary>
        /// Specifies a list of categories in which the Advisor Recommendations will be listed. Possible values are `HighAvailability`, `Security`, `Performance`, `Cost` and `OperationalExcellence`.
        /// </summary>
        public InputList<string> FilterByCategories
        {
            get => _filterByCategories ?? (_filterByCategories = new InputList<string>());
            set => _filterByCategories = value;
        }

        [Input("filterByResourceGroups")]
        private InputList<string>? _filterByResourceGroups;

        /// <summary>
        /// Specifies a list of resource groups about which the Advisor Recommendations will be listed.
        /// </summary>
        public InputList<string> FilterByResourceGroups
        {
            get => _filterByResourceGroups ?? (_filterByResourceGroups = new InputList<string>());
            set => _filterByResourceGroups = value;
        }

        public GetRecommendationsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRecommendationsResult
    {
        public readonly ImmutableArray<string> FilterByCategories;
        public readonly ImmutableArray<string> FilterByResourceGroups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// One or more `recommendations` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRecommendationsRecommendationResult> Recommendations;

        [OutputConstructor]
        private GetRecommendationsResult(
            ImmutableArray<string> filterByCategories,

            ImmutableArray<string> filterByResourceGroups,

            string id,

            ImmutableArray<Outputs.GetRecommendationsRecommendationResult> recommendations)
        {
            FilterByCategories = filterByCategories;
            FilterByResourceGroups = filterByResourceGroups;
            Id = id;
            Recommendations = recommendations;
        }
    }
}
