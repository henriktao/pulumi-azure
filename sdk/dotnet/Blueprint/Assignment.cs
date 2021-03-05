// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Blueprint
{
    /// <summary>
    /// Manages a Blueprint Assignment resource
    /// 
    /// &gt; **NOTE:** Azure Blueprints are in Preview and potentially subject to breaking change without notice.
    /// 
    /// &gt; **NOTE:** Azure Blueprint Assignments can only be applied to Subscriptions.  Assignments to Management Groups is not currently supported by the service or by this provider.
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
    ///         var current = Output.Create(Azure.Core.GetClientConfig.InvokeAsync());
    ///         var exampleSubscription = Output.Create(Azure.Core.GetSubscription.InvokeAsync());
    ///         var exampleDefinition = exampleSubscription.Apply(exampleSubscription =&gt; Output.Create(Azure.Blueprint.GetDefinition.InvokeAsync(new Azure.Blueprint.GetDefinitionArgs
    ///         {
    ///             Name = "exampleBlueprint",
    ///             ScopeId = exampleSubscription.Id,
    ///         })));
    ///         var examplePublishedVersion = Output.Tuple(exampleDefinition, exampleDefinition).Apply(values =&gt;
    ///         {
    ///             var exampleDefinition = values.Item1;
    ///             var exampleDefinition1 = values.Item2;
    ///             return Output.Create(Azure.Blueprint.GetPublishedVersion.InvokeAsync(new Azure.Blueprint.GetPublishedVersionArgs
    ///             {
    ///                 ScopeId = exampleDefinition.ScopeId,
    ///                 BlueprintName = exampleDefinition1.Name,
    ///                 Version = "v1.0.0",
    ///             }));
    ///         });
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///             Tags = 
    ///             {
    ///                 { "Environment", "example" },
    ///             },
    ///         });
    ///         var exampleUserAssignedIdentity = new Azure.Authorization.UserAssignedIdentity("exampleUserAssignedIdentity", new Azure.Authorization.UserAssignedIdentityArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///         });
    ///         var @operator = new Azure.Authorization.Assignment("operator", new Azure.Authorization.AssignmentArgs
    ///         {
    ///             Scope = exampleSubscription.Apply(exampleSubscription =&gt; exampleSubscription.Id),
    ///             RoleDefinitionName = "Blueprint Operator",
    ///             PrincipalId = exampleUserAssignedIdentity.PrincipalId,
    ///         });
    ///         var owner = new Azure.Authorization.Assignment("owner", new Azure.Authorization.AssignmentArgs
    ///         {
    ///             Scope = exampleSubscription.Apply(exampleSubscription =&gt; exampleSubscription.Id),
    ///             RoleDefinitionName = "Owner",
    ///             PrincipalId = exampleUserAssignedIdentity.PrincipalId,
    ///         });
    ///         var exampleAssignment = new Azure.Blueprint.Assignment("exampleAssignment", new Azure.Blueprint.AssignmentArgs
    ///         {
    ///             TargetSubscriptionId = exampleSubscription.Apply(exampleSubscription =&gt; exampleSubscription.Id),
    ///             VersionId = examplePublishedVersion.Apply(examplePublishedVersion =&gt; examplePublishedVersion.Id),
    ///             Location = exampleResourceGroup.Location,
    ///             LockMode = "AllResourcesDoNotDelete",
    ///             LockExcludePrincipals = 
    ///             {
    ///                 current.Apply(current =&gt; current.ObjectId),
    ///             },
    ///             Identity = new Azure.Blueprint.Inputs.AssignmentIdentityArgs
    ///             {
    ///                 Type = "UserAssigned",
    ///                 IdentityIds = 
    ///                 {
    ///                     exampleUserAssignedIdentity.Id,
    ///                 },
    ///             },
    ///             ResourceGroups = @"    {
    ///       ""ResourceGroup"": {
    ///         ""name"": ""exampleRG-bp""
    ///       }
    ///     }
    /// ",
    ///             ParameterValues = @"    {
    ///       ""allowedlocationsforresourcegroups_listOfAllowedLocations"": {
    ///         ""value"": [""westus"", ""westus2"", ""eastus"", ""centralus"", ""centraluseuap"", ""southcentralus"", ""northcentralus"", ""westcentralus"", ""eastus2"", ""eastus2euap"", ""brazilsouth"", ""brazilus"", ""northeurope"", ""westeurope"", ""eastasia"", ""southeastasia"", ""japanwest"", ""japaneast"", ""koreacentral"", ""koreasouth"", ""indiasouth"", ""indiawest"", ""indiacentral"", ""australiaeast"", ""australiasoutheast"", ""canadacentral"", ""canadaeast"", ""uknorth"", ""uksouth2"", ""uksouth"", ""ukwest"", ""francecentral"", ""francesouth"", ""australiacentral"", ""australiacentral2"", ""uaecentral"", ""uaenorth"", ""southafricanorth"", ""southafricawest"", ""switzerlandnorth"", ""switzerlandwest"", ""germanynorth"", ""germanywestcentral"", ""norwayeast"", ""norwaywest""]
    ///       }
    ///     }
    /// ",
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 @operator,
    ///                 owner,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Azure Blueprint Assignments can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:blueprint/assignment:Assignment example "/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Blueprint/blueprintAssignments/assignSimpleBlueprint"
    /// ```
    /// </summary>
    [AzureResourceType("azure:blueprint/assignment:Assignment")]
    public partial class Assignment : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the blueprint assigned
        /// </summary>
        [Output("blueprintName")]
        public Output<string> BlueprintName { get; private set; } = null!;

        /// <summary>
        /// The Description on the Blueprint
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the blueprint
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        [Output("identity")]
        public Output<Outputs.AssignmentIdentity?> Identity { get; private set; } = null!;

        /// <summary>
        /// The Azure location of the Assignment.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// a list of up to 5 Principal IDs that are permitted to bypass the locks applied by the Blueprint.
        /// </summary>
        [Output("lockExcludePrincipals")]
        public Output<ImmutableArray<string>> LockExcludePrincipals { get; private set; } = null!;

        /// <summary>
        /// The locking mode of the Blueprint Assignment.  One of `None` (Default), `AllResourcesReadOnly`, or `AlResourcesDoNotDelete`.
        /// </summary>
        [Output("lockMode")]
        public Output<string?> LockMode { get; private set; } = null!;

        /// <summary>
        /// The name of the Blueprint Assignment
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// a JSON string to supply Blueprint Assignment parameter values.
        /// </summary>
        [Output("parameterValues")]
        public Output<string?> ParameterValues { get; private set; } = null!;

        /// <summary>
        /// a JSON string to supply the Blueprint Resource Group information.
        /// </summary>
        [Output("resourceGroups")]
        public Output<string?> ResourceGroups { get; private set; } = null!;

        /// <summary>
        /// The Subscription ID the Blueprint Published Version is to be applied to.
        /// </summary>
        [Output("targetSubscriptionId")]
        public Output<string> TargetSubscriptionId { get; private set; } = null!;

        /// <summary>
        /// The Identity type for the Managed Service Identity. Currently only `UserAssigned` is supported.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The ID of the Published Version of the blueprint to be assigned.
        /// </summary>
        [Output("versionId")]
        public Output<string> VersionId { get; private set; } = null!;


        /// <summary>
        /// Create a Assignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Assignment(string name, AssignmentArgs args, CustomResourceOptions? options = null)
            : base("azure:blueprint/assignment:Assignment", name, args ?? new AssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Assignment(string name, Input<string> id, AssignmentState? state = null, CustomResourceOptions? options = null)
            : base("azure:blueprint/assignment:Assignment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Assignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Assignment Get(string name, Input<string> id, AssignmentState? state = null, CustomResourceOptions? options = null)
        {
            return new Assignment(name, id, state, options);
        }
    }

    public sealed class AssignmentArgs : Pulumi.ResourceArgs
    {
        [Input("identity")]
        public Input<Inputs.AssignmentIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The Azure location of the Assignment.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("lockExcludePrincipals")]
        private InputList<string>? _lockExcludePrincipals;

        /// <summary>
        /// a list of up to 5 Principal IDs that are permitted to bypass the locks applied by the Blueprint.
        /// </summary>
        public InputList<string> LockExcludePrincipals
        {
            get => _lockExcludePrincipals ?? (_lockExcludePrincipals = new InputList<string>());
            set => _lockExcludePrincipals = value;
        }

        /// <summary>
        /// The locking mode of the Blueprint Assignment.  One of `None` (Default), `AllResourcesReadOnly`, or `AlResourcesDoNotDelete`.
        /// </summary>
        [Input("lockMode")]
        public Input<string>? LockMode { get; set; }

        /// <summary>
        /// The name of the Blueprint Assignment
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// a JSON string to supply Blueprint Assignment parameter values.
        /// </summary>
        [Input("parameterValues")]
        public Input<string>? ParameterValues { get; set; }

        /// <summary>
        /// a JSON string to supply the Blueprint Resource Group information.
        /// </summary>
        [Input("resourceGroups")]
        public Input<string>? ResourceGroups { get; set; }

        /// <summary>
        /// The Subscription ID the Blueprint Published Version is to be applied to.
        /// </summary>
        [Input("targetSubscriptionId", required: true)]
        public Input<string> TargetSubscriptionId { get; set; } = null!;

        /// <summary>
        /// The ID of the Published Version of the blueprint to be assigned.
        /// </summary>
        [Input("versionId", required: true)]
        public Input<string> VersionId { get; set; } = null!;

        public AssignmentArgs()
        {
        }
    }

    public sealed class AssignmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the blueprint assigned
        /// </summary>
        [Input("blueprintName")]
        public Input<string>? BlueprintName { get; set; }

        /// <summary>
        /// The Description on the Blueprint
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the blueprint
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("identity")]
        public Input<Inputs.AssignmentIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// The Azure location of the Assignment.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("lockExcludePrincipals")]
        private InputList<string>? _lockExcludePrincipals;

        /// <summary>
        /// a list of up to 5 Principal IDs that are permitted to bypass the locks applied by the Blueprint.
        /// </summary>
        public InputList<string> LockExcludePrincipals
        {
            get => _lockExcludePrincipals ?? (_lockExcludePrincipals = new InputList<string>());
            set => _lockExcludePrincipals = value;
        }

        /// <summary>
        /// The locking mode of the Blueprint Assignment.  One of `None` (Default), `AllResourcesReadOnly`, or `AlResourcesDoNotDelete`.
        /// </summary>
        [Input("lockMode")]
        public Input<string>? LockMode { get; set; }

        /// <summary>
        /// The name of the Blueprint Assignment
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// a JSON string to supply Blueprint Assignment parameter values.
        /// </summary>
        [Input("parameterValues")]
        public Input<string>? ParameterValues { get; set; }

        /// <summary>
        /// a JSON string to supply the Blueprint Resource Group information.
        /// </summary>
        [Input("resourceGroups")]
        public Input<string>? ResourceGroups { get; set; }

        /// <summary>
        /// The Subscription ID the Blueprint Published Version is to be applied to.
        /// </summary>
        [Input("targetSubscriptionId")]
        public Input<string>? TargetSubscriptionId { get; set; }

        /// <summary>
        /// The Identity type for the Managed Service Identity. Currently only `UserAssigned` is supported.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The ID of the Published Version of the blueprint to be assigned.
        /// </summary>
        [Input("versionId")]
        public Input<string>? VersionId { get; set; }

        public AssignmentState()
        {
        }
    }
}
