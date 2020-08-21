// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Azure Network DDoS Protection Plan.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := network.GetNetworkDdosProtectionPlan(ctx, &network.GetNetworkDdosProtectionPlanArgs{
// 			Name:              azurerm_network_ddos_protection_plan.Example.Name,
// 			ResourceGroupName: azurerm_network_ddos_protection_plan.Example.Resource_group_name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("ddosProtectionPlanId", example.Id)
// 		return nil
// 	})
// }
// ```
func GetNetworkDdosProtectionPlan(ctx *pulumi.Context, args *GetNetworkDdosProtectionPlanArgs, opts ...pulumi.InvokeOption) (*GetNetworkDdosProtectionPlanResult, error) {
	var rv GetNetworkDdosProtectionPlanResult
	err := ctx.Invoke("azure:network/getNetworkDdosProtectionPlan:getNetworkDdosProtectionPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkDdosProtectionPlan.
type GetNetworkDdosProtectionPlanArgs struct {
	// The name of the Network DDoS Protection Plan.
	Name string `pulumi:"name"`
	// The name of the resource group where the Network DDoS Protection Plan exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getNetworkDdosProtectionPlan.
type GetNetworkDdosProtectionPlanResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Specifies the supported Azure location where the resource exists.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A list of ID's of the Virtual Networks associated with this DDoS Protection Plan.
	VirtualNetworkIds []string `pulumi:"virtualNetworkIds"`
}
