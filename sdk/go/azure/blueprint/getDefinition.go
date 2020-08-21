// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package blueprint

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Azure Blueprint Definition
//
// > **NOTE:** Azure Blueprints are in Preview and potentially subject to breaking change without notice.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/blueprint"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/management"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt0 := current.TenantId
// 		root, err := management.LookupGroup(ctx, &management.LookupGroupArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = blueprint.GetDefinition(ctx, &blueprint.GetDefinitionArgs{
// 			Name:    "exampleManagementGroupBP",
// 			ScopeId: root.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDefinition(ctx *pulumi.Context, args *GetDefinitionArgs, opts ...pulumi.InvokeOption) (*GetDefinitionResult, error) {
	var rv GetDefinitionResult
	err := ctx.Invoke("azure:blueprint/getDefinition:getDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDefinition.
type GetDefinitionArgs struct {
	// The name of the Blueprint.
	Name string `pulumi:"name"`
	// The ID of the Subscription or Management Group, as the scope at which the blueprint definition is stored.
	ScopeId string `pulumi:"scopeId"`
}

// A collection of values returned by getDefinition.
type GetDefinitionResult struct {
	// The description of the Blueprint Definition.
	Description string `pulumi:"description"`
	// The display name of the Blueprint Definition.
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The timestamp of when this last modification was saved to the Blueprint Definition.
	LastModified string `pulumi:"lastModified"`
	Name         string `pulumi:"name"`
	ScopeId      string `pulumi:"scopeId"`
	// The target scope.
	TargetScope string `pulumi:"targetScope"`
	// The timestamp of when this Blueprint Definition was created.
	TimeCreated string `pulumi:"timeCreated"`
	// A list of versions published for this Blueprint Definition.
	Versions []string `pulumi:"versions"`
}
