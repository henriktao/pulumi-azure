// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package role

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

func GetRoleDefinition(ctx *pulumi.Context, args *GetRoleDefinitionArgs, opts ...pulumi.InvokeOption) (*GetRoleDefinitionResult, error) {
	var rv GetRoleDefinitionResult
	err := ctx.Invoke("azure:role/getRoleDefinition:getRoleDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRoleDefinition.
type GetRoleDefinitionArgs struct {
	Name *string `pulumi:"name"`
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	Scope *string `pulumi:"scope"`
}


// A collection of values returned by getRoleDefinition.
type GetRoleDefinitionResult struct {
	AssignableScopes []string `pulumi:"assignableScopes"`
	Description string `pulumi:"description"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Name string `pulumi:"name"`
	Permissions []GetRoleDefinitionPermission `pulumi:"permissions"`
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
	Scope *string `pulumi:"scope"`
	Type string `pulumi:"type"`
}

