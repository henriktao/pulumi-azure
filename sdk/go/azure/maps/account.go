// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package maps

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure Maps Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/maps"
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
// 		_, err = maps.NewAccount(ctx, "exampleAccount", &maps.AccountArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			SkuName:           pulumi.String("S1"),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("Test"),
// 			},
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
// A Maps Account can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:maps/account:Account example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Maps/accounts/my-maps-account
// ```
type Account struct {
	pulumi.CustomResourceState

	// The name of the Azure Maps Account. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The primary key used to authenticate and authorize access to the Maps REST APIs.
	PrimaryAccessKey pulumi.StringOutput `pulumi:"primaryAccessKey"`
	// The name of the Resource Group in which the Azure Maps Account should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The secondary key used to authenticate and authorize access to the Maps REST APIs.
	SecondaryAccessKey pulumi.StringOutput `pulumi:"secondaryAccessKey"`
	// The sku of the Azure Maps Account. Possible values are `S0`, `S1` and `G2`.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// A mapping of tags to assign to the Azure Maps Account.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A unique identifier for the Maps Account.
	XMsClientId pulumi.StringOutput `pulumi:"xMsClientId"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SkuName == nil {
		return nil, errors.New("invalid value for required argument 'SkuName'")
	}
	var resource Account
	err := ctx.RegisterResource("azure:maps/account:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure:maps/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// The name of the Azure Maps Account. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The primary key used to authenticate and authorize access to the Maps REST APIs.
	PrimaryAccessKey *string `pulumi:"primaryAccessKey"`
	// The name of the Resource Group in which the Azure Maps Account should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The secondary key used to authenticate and authorize access to the Maps REST APIs.
	SecondaryAccessKey *string `pulumi:"secondaryAccessKey"`
	// The sku of the Azure Maps Account. Possible values are `S0`, `S1` and `G2`.
	SkuName *string `pulumi:"skuName"`
	// A mapping of tags to assign to the Azure Maps Account.
	Tags map[string]string `pulumi:"tags"`
	// A unique identifier for the Maps Account.
	XMsClientId *string `pulumi:"xMsClientId"`
}

type AccountState struct {
	// The name of the Azure Maps Account. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The primary key used to authenticate and authorize access to the Maps REST APIs.
	PrimaryAccessKey pulumi.StringPtrInput
	// The name of the Resource Group in which the Azure Maps Account should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The secondary key used to authenticate and authorize access to the Maps REST APIs.
	SecondaryAccessKey pulumi.StringPtrInput
	// The sku of the Azure Maps Account. Possible values are `S0`, `S1` and `G2`.
	SkuName pulumi.StringPtrInput
	// A mapping of tags to assign to the Azure Maps Account.
	Tags pulumi.StringMapInput
	// A unique identifier for the Maps Account.
	XMsClientId pulumi.StringPtrInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The name of the Azure Maps Account. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the Azure Maps Account should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku of the Azure Maps Account. Possible values are `S0`, `S1` and `G2`.
	SkuName string `pulumi:"skuName"`
	// A mapping of tags to assign to the Azure Maps Account.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The name of the Azure Maps Account. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the Azure Maps Account should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The sku of the Azure Maps Account. Possible values are `S0`, `S1` and `G2`.
	SkuName pulumi.StringInput
	// A mapping of tags to assign to the Azure Maps Account.
	Tags pulumi.StringMapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((*Account)(nil))
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

func (i *Account) ToAccountPtrOutput() AccountPtrOutput {
	return i.ToAccountPtrOutputWithContext(context.Background())
}

func (i *Account) ToAccountPtrOutputWithContext(ctx context.Context) AccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPtrOutput)
}

type AccountPtrInput interface {
	pulumi.Input

	ToAccountPtrOutput() AccountPtrOutput
	ToAccountPtrOutputWithContext(ctx context.Context) AccountPtrOutput
}

type accountPtrType AccountArgs

func (*accountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil))
}

func (i *accountPtrType) ToAccountPtrOutput() AccountPtrOutput {
	return i.ToAccountPtrOutputWithContext(context.Background())
}

func (i *accountPtrType) ToAccountPtrOutputWithContext(ctx context.Context) AccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPtrOutput)
}

// AccountArrayInput is an input type that accepts AccountArray and AccountArrayOutput values.
// You can construct a concrete instance of `AccountArrayInput` via:
//
//          AccountArray{ AccountArgs{...} }
type AccountArrayInput interface {
	pulumi.Input

	ToAccountArrayOutput() AccountArrayOutput
	ToAccountArrayOutputWithContext(context.Context) AccountArrayOutput
}

type AccountArray []AccountInput

func (AccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Account)(nil)).Elem()
}

func (i AccountArray) ToAccountArrayOutput() AccountArrayOutput {
	return i.ToAccountArrayOutputWithContext(context.Background())
}

func (i AccountArray) ToAccountArrayOutputWithContext(ctx context.Context) AccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountArrayOutput)
}

// AccountMapInput is an input type that accepts AccountMap and AccountMapOutput values.
// You can construct a concrete instance of `AccountMapInput` via:
//
//          AccountMap{ "key": AccountArgs{...} }
type AccountMapInput interface {
	pulumi.Input

	ToAccountMapOutput() AccountMapOutput
	ToAccountMapOutputWithContext(context.Context) AccountMapOutput
}

type AccountMap map[string]AccountInput

func (AccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Account)(nil)).Elem()
}

func (i AccountMap) ToAccountMapOutput() AccountMapOutput {
	return i.ToAccountMapOutputWithContext(context.Background())
}

func (i AccountMap) ToAccountMapOutputWithContext(ctx context.Context) AccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountMapOutput)
}

type AccountOutput struct{ *pulumi.OutputState }

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Account)(nil))
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

func (o AccountOutput) ToAccountPtrOutput() AccountPtrOutput {
	return o.ToAccountPtrOutputWithContext(context.Background())
}

func (o AccountOutput) ToAccountPtrOutputWithContext(ctx context.Context) AccountPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Account) *Account {
		return &v
	}).(AccountPtrOutput)
}

type AccountPtrOutput struct{ *pulumi.OutputState }

func (AccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil))
}

func (o AccountPtrOutput) ToAccountPtrOutput() AccountPtrOutput {
	return o
}

func (o AccountPtrOutput) ToAccountPtrOutputWithContext(ctx context.Context) AccountPtrOutput {
	return o
}

func (o AccountPtrOutput) Elem() AccountOutput {
	return o.ApplyT(func(v *Account) Account {
		if v != nil {
			return *v
		}
		var ret Account
		return ret
	}).(AccountOutput)
}

type AccountArrayOutput struct{ *pulumi.OutputState }

func (AccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Account)(nil))
}

func (o AccountArrayOutput) ToAccountArrayOutput() AccountArrayOutput {
	return o
}

func (o AccountArrayOutput) ToAccountArrayOutputWithContext(ctx context.Context) AccountArrayOutput {
	return o
}

func (o AccountArrayOutput) Index(i pulumi.IntInput) AccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Account {
		return vs[0].([]Account)[vs[1].(int)]
	}).(AccountOutput)
}

type AccountMapOutput struct{ *pulumi.OutputState }

func (AccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Account)(nil))
}

func (o AccountMapOutput) ToAccountMapOutput() AccountMapOutput {
	return o
}

func (o AccountMapOutput) ToAccountMapOutputWithContext(ctx context.Context) AccountMapOutput {
	return o
}

func (o AccountMapOutput) MapIndex(k pulumi.StringInput) AccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Account {
		return vs[0].(map[string]Account)[vs[1].(string)]
	}).(AccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountInput)(nil)).Elem(), &Account{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPtrInput)(nil)).Elem(), &Account{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountArrayInput)(nil)).Elem(), AccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountMapInput)(nil)).Elem(), AccountMap{})
	pulumi.RegisterOutputType(AccountOutput{})
	pulumi.RegisterOutputType(AccountPtrOutput{})
	pulumi.RegisterOutputType(AccountArrayOutput{})
	pulumi.RegisterOutputType(AccountMapOutput{})
}
