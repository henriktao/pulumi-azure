// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datalake

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Azure Data Lake Store Firewall Rule.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/datalake"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleStore, err := datalake.NewStore(ctx, "exampleStore", &datalake.StoreArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datalake.NewStoreFirewallRule(ctx, "exampleStoreFirewallRule", &datalake.StoreFirewallRuleArgs{
// 			AccountName:       exampleStore.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			StartIpAddress:    pulumi.String("1.2.3.4"),
// 			EndIpAddress:      pulumi.String("2.3.4.5"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Data Lake Store Firewall Rules can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datalake/storeFirewallRule:StoreFirewallRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DataLakeStore/accounts/mydatalakeaccount/firewallRules/rule1
// ```
type StoreFirewallRule struct {
	pulumi.CustomResourceState

	// Specifies the name of the Data Lake Store for which the Firewall Rule should take effect.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// The End IP Address for the firewall rule.
	EndIpAddress pulumi.StringOutput `pulumi:"endIpAddress"`
	// Specifies the name of the Data Lake Store. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Data Lake Store.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Start IP address for the firewall rule.
	StartIpAddress pulumi.StringOutput `pulumi:"startIpAddress"`
}

// NewStoreFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewStoreFirewallRule(ctx *pulumi.Context,
	name string, args *StoreFirewallRuleArgs, opts ...pulumi.ResourceOption) (*StoreFirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.EndIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'EndIpAddress'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StartIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'StartIpAddress'")
	}
	var resource StoreFirewallRule
	err := ctx.RegisterResource("azure:datalake/storeFirewallRule:StoreFirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStoreFirewallRule gets an existing StoreFirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStoreFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StoreFirewallRuleState, opts ...pulumi.ResourceOption) (*StoreFirewallRule, error) {
	var resource StoreFirewallRule
	err := ctx.ReadResource("azure:datalake/storeFirewallRule:StoreFirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StoreFirewallRule resources.
type storeFirewallRuleState struct {
	// Specifies the name of the Data Lake Store for which the Firewall Rule should take effect.
	AccountName *string `pulumi:"accountName"`
	// The End IP Address for the firewall rule.
	EndIpAddress *string `pulumi:"endIpAddress"`
	// Specifies the name of the Data Lake Store. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Data Lake Store.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Start IP address for the firewall rule.
	StartIpAddress *string `pulumi:"startIpAddress"`
}

type StoreFirewallRuleState struct {
	// Specifies the name of the Data Lake Store for which the Firewall Rule should take effect.
	AccountName pulumi.StringPtrInput
	// The End IP Address for the firewall rule.
	EndIpAddress pulumi.StringPtrInput
	// Specifies the name of the Data Lake Store. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Data Lake Store.
	ResourceGroupName pulumi.StringPtrInput
	// The Start IP address for the firewall rule.
	StartIpAddress pulumi.StringPtrInput
}

func (StoreFirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*storeFirewallRuleState)(nil)).Elem()
}

type storeFirewallRuleArgs struct {
	// Specifies the name of the Data Lake Store for which the Firewall Rule should take effect.
	AccountName string `pulumi:"accountName"`
	// The End IP Address for the firewall rule.
	EndIpAddress string `pulumi:"endIpAddress"`
	// Specifies the name of the Data Lake Store. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Data Lake Store.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Start IP address for the firewall rule.
	StartIpAddress string `pulumi:"startIpAddress"`
}

// The set of arguments for constructing a StoreFirewallRule resource.
type StoreFirewallRuleArgs struct {
	// Specifies the name of the Data Lake Store for which the Firewall Rule should take effect.
	AccountName pulumi.StringInput
	// The End IP Address for the firewall rule.
	EndIpAddress pulumi.StringInput
	// Specifies the name of the Data Lake Store. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Data Lake Store.
	ResourceGroupName pulumi.StringInput
	// The Start IP address for the firewall rule.
	StartIpAddress pulumi.StringInput
}

func (StoreFirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storeFirewallRuleArgs)(nil)).Elem()
}

type StoreFirewallRuleInput interface {
	pulumi.Input

	ToStoreFirewallRuleOutput() StoreFirewallRuleOutput
	ToStoreFirewallRuleOutputWithContext(ctx context.Context) StoreFirewallRuleOutput
}

func (*StoreFirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((*StoreFirewallRule)(nil))
}

func (i *StoreFirewallRule) ToStoreFirewallRuleOutput() StoreFirewallRuleOutput {
	return i.ToStoreFirewallRuleOutputWithContext(context.Background())
}

func (i *StoreFirewallRule) ToStoreFirewallRuleOutputWithContext(ctx context.Context) StoreFirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreFirewallRuleOutput)
}

func (i *StoreFirewallRule) ToStoreFirewallRulePtrOutput() StoreFirewallRulePtrOutput {
	return i.ToStoreFirewallRulePtrOutputWithContext(context.Background())
}

func (i *StoreFirewallRule) ToStoreFirewallRulePtrOutputWithContext(ctx context.Context) StoreFirewallRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreFirewallRulePtrOutput)
}

type StoreFirewallRulePtrInput interface {
	pulumi.Input

	ToStoreFirewallRulePtrOutput() StoreFirewallRulePtrOutput
	ToStoreFirewallRulePtrOutputWithContext(ctx context.Context) StoreFirewallRulePtrOutput
}

type storeFirewallRulePtrType StoreFirewallRuleArgs

func (*storeFirewallRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StoreFirewallRule)(nil))
}

func (i *storeFirewallRulePtrType) ToStoreFirewallRulePtrOutput() StoreFirewallRulePtrOutput {
	return i.ToStoreFirewallRulePtrOutputWithContext(context.Background())
}

func (i *storeFirewallRulePtrType) ToStoreFirewallRulePtrOutputWithContext(ctx context.Context) StoreFirewallRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreFirewallRulePtrOutput)
}

// StoreFirewallRuleArrayInput is an input type that accepts StoreFirewallRuleArray and StoreFirewallRuleArrayOutput values.
// You can construct a concrete instance of `StoreFirewallRuleArrayInput` via:
//
//          StoreFirewallRuleArray{ StoreFirewallRuleArgs{...} }
type StoreFirewallRuleArrayInput interface {
	pulumi.Input

	ToStoreFirewallRuleArrayOutput() StoreFirewallRuleArrayOutput
	ToStoreFirewallRuleArrayOutputWithContext(context.Context) StoreFirewallRuleArrayOutput
}

type StoreFirewallRuleArray []StoreFirewallRuleInput

func (StoreFirewallRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StoreFirewallRule)(nil)).Elem()
}

func (i StoreFirewallRuleArray) ToStoreFirewallRuleArrayOutput() StoreFirewallRuleArrayOutput {
	return i.ToStoreFirewallRuleArrayOutputWithContext(context.Background())
}

func (i StoreFirewallRuleArray) ToStoreFirewallRuleArrayOutputWithContext(ctx context.Context) StoreFirewallRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreFirewallRuleArrayOutput)
}

// StoreFirewallRuleMapInput is an input type that accepts StoreFirewallRuleMap and StoreFirewallRuleMapOutput values.
// You can construct a concrete instance of `StoreFirewallRuleMapInput` via:
//
//          StoreFirewallRuleMap{ "key": StoreFirewallRuleArgs{...} }
type StoreFirewallRuleMapInput interface {
	pulumi.Input

	ToStoreFirewallRuleMapOutput() StoreFirewallRuleMapOutput
	ToStoreFirewallRuleMapOutputWithContext(context.Context) StoreFirewallRuleMapOutput
}

type StoreFirewallRuleMap map[string]StoreFirewallRuleInput

func (StoreFirewallRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StoreFirewallRule)(nil)).Elem()
}

func (i StoreFirewallRuleMap) ToStoreFirewallRuleMapOutput() StoreFirewallRuleMapOutput {
	return i.ToStoreFirewallRuleMapOutputWithContext(context.Background())
}

func (i StoreFirewallRuleMap) ToStoreFirewallRuleMapOutputWithContext(ctx context.Context) StoreFirewallRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreFirewallRuleMapOutput)
}

type StoreFirewallRuleOutput struct{ *pulumi.OutputState }

func (StoreFirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StoreFirewallRule)(nil))
}

func (o StoreFirewallRuleOutput) ToStoreFirewallRuleOutput() StoreFirewallRuleOutput {
	return o
}

func (o StoreFirewallRuleOutput) ToStoreFirewallRuleOutputWithContext(ctx context.Context) StoreFirewallRuleOutput {
	return o
}

func (o StoreFirewallRuleOutput) ToStoreFirewallRulePtrOutput() StoreFirewallRulePtrOutput {
	return o.ToStoreFirewallRulePtrOutputWithContext(context.Background())
}

func (o StoreFirewallRuleOutput) ToStoreFirewallRulePtrOutputWithContext(ctx context.Context) StoreFirewallRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StoreFirewallRule) *StoreFirewallRule {
		return &v
	}).(StoreFirewallRulePtrOutput)
}

type StoreFirewallRulePtrOutput struct{ *pulumi.OutputState }

func (StoreFirewallRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StoreFirewallRule)(nil))
}

func (o StoreFirewallRulePtrOutput) ToStoreFirewallRulePtrOutput() StoreFirewallRulePtrOutput {
	return o
}

func (o StoreFirewallRulePtrOutput) ToStoreFirewallRulePtrOutputWithContext(ctx context.Context) StoreFirewallRulePtrOutput {
	return o
}

func (o StoreFirewallRulePtrOutput) Elem() StoreFirewallRuleOutput {
	return o.ApplyT(func(v *StoreFirewallRule) StoreFirewallRule {
		if v != nil {
			return *v
		}
		var ret StoreFirewallRule
		return ret
	}).(StoreFirewallRuleOutput)
}

type StoreFirewallRuleArrayOutput struct{ *pulumi.OutputState }

func (StoreFirewallRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StoreFirewallRule)(nil))
}

func (o StoreFirewallRuleArrayOutput) ToStoreFirewallRuleArrayOutput() StoreFirewallRuleArrayOutput {
	return o
}

func (o StoreFirewallRuleArrayOutput) ToStoreFirewallRuleArrayOutputWithContext(ctx context.Context) StoreFirewallRuleArrayOutput {
	return o
}

func (o StoreFirewallRuleArrayOutput) Index(i pulumi.IntInput) StoreFirewallRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StoreFirewallRule {
		return vs[0].([]StoreFirewallRule)[vs[1].(int)]
	}).(StoreFirewallRuleOutput)
}

type StoreFirewallRuleMapOutput struct{ *pulumi.OutputState }

func (StoreFirewallRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StoreFirewallRule)(nil))
}

func (o StoreFirewallRuleMapOutput) ToStoreFirewallRuleMapOutput() StoreFirewallRuleMapOutput {
	return o
}

func (o StoreFirewallRuleMapOutput) ToStoreFirewallRuleMapOutputWithContext(ctx context.Context) StoreFirewallRuleMapOutput {
	return o
}

func (o StoreFirewallRuleMapOutput) MapIndex(k pulumi.StringInput) StoreFirewallRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StoreFirewallRule {
		return vs[0].(map[string]StoreFirewallRule)[vs[1].(string)]
	}).(StoreFirewallRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(StoreFirewallRuleOutput{})
	pulumi.RegisterOutputType(StoreFirewallRulePtrOutput{})
	pulumi.RegisterOutputType(StoreFirewallRuleArrayOutput{})
	pulumi.RegisterOutputType(StoreFirewallRuleMapOutput{})
}
