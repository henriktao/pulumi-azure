// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package notificationhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Authorization Rule associated with a Notification Hub within a Notification Hub Namespace.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/notificationhub"
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
// 		exampleNamespace, err := notificationhub.NewNamespace(ctx, "exampleNamespace", &notificationhub.NamespaceArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			NamespaceType:     pulumi.String("NotificationHub"),
// 			SkuName:           pulumi.String("Free"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleHub, err := notificationhub.NewHub(ctx, "exampleHub", &notificationhub.HubArgs{
// 			NamespaceName:     exampleNamespace.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = notificationhub.NewAuthorizationRule(ctx, "exampleAuthorizationRule", &notificationhub.AuthorizationRuleArgs{
// 			NotificationHubName: exampleHub.Name,
// 			NamespaceName:       exampleNamespace.Name,
// 			ResourceGroupName:   exampleResourceGroup.Name,
// 			Manage:              pulumi.Bool(true),
// 			Send:                pulumi.Bool(true),
// 			Listen:              pulumi.Bool(true),
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
// Notification Hub Authorization Rule can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:notificationhub/authorizationRule:AuthorizationRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.NotificationHubs/namespaces/namespace1/notificationHubs/hub1/authorizationRules/rule1
// ```
type AuthorizationRule struct {
	pulumi.CustomResourceState

	// Does this Authorization Rule have Listen access to the Notification Hub? Defaults to `false`.
	Listen pulumi.BoolPtrOutput `pulumi:"listen"`
	// Does this Authorization Rule have Manage access to the Notification Hub? Defaults to `false`.
	Manage pulumi.BoolPtrOutput `pulumi:"manage"`
	// The name to use for this Authorization Rule. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Notification Hub Namespace in which the Notification Hub exists. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// The name of the Notification Hub for which the Authorization Rule should be created. Changing this forces a new resource to be created.
	NotificationHubName pulumi.StringOutput `pulumi:"notificationHubName"`
	// The Primary Access Key associated with this Authorization Rule.
	PrimaryAccessKey pulumi.StringOutput `pulumi:"primaryAccessKey"`
	// The name of the Resource Group in which the Notification Hub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Secondary Access Key associated with this Authorization Rule.
	SecondaryAccessKey pulumi.StringOutput `pulumi:"secondaryAccessKey"`
	// Does this Authorization Rule have Send access to the Notification Hub? Defaults to `false`.
	Send pulumi.BoolPtrOutput `pulumi:"send"`
}

// NewAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewAuthorizationRule(ctx *pulumi.Context,
	name string, args *AuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*AuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.NotificationHubName == nil {
		return nil, errors.New("invalid value for required argument 'NotificationHubName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource AuthorizationRule
	err := ctx.RegisterResource("azure:notificationhub/authorizationRule:AuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthorizationRule gets an existing AuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizationRuleState, opts ...pulumi.ResourceOption) (*AuthorizationRule, error) {
	var resource AuthorizationRule
	err := ctx.ReadResource("azure:notificationhub/authorizationRule:AuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthorizationRule resources.
type authorizationRuleState struct {
	// Does this Authorization Rule have Listen access to the Notification Hub? Defaults to `false`.
	Listen *bool `pulumi:"listen"`
	// Does this Authorization Rule have Manage access to the Notification Hub? Defaults to `false`.
	Manage *bool `pulumi:"manage"`
	// The name to use for this Authorization Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Notification Hub Namespace in which the Notification Hub exists. Changing this forces a new resource to be created.
	NamespaceName *string `pulumi:"namespaceName"`
	// The name of the Notification Hub for which the Authorization Rule should be created. Changing this forces a new resource to be created.
	NotificationHubName *string `pulumi:"notificationHubName"`
	// The Primary Access Key associated with this Authorization Rule.
	PrimaryAccessKey *string `pulumi:"primaryAccessKey"`
	// The name of the Resource Group in which the Notification Hub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Secondary Access Key associated with this Authorization Rule.
	SecondaryAccessKey *string `pulumi:"secondaryAccessKey"`
	// Does this Authorization Rule have Send access to the Notification Hub? Defaults to `false`.
	Send *bool `pulumi:"send"`
}

type AuthorizationRuleState struct {
	// Does this Authorization Rule have Listen access to the Notification Hub? Defaults to `false`.
	Listen pulumi.BoolPtrInput
	// Does this Authorization Rule have Manage access to the Notification Hub? Defaults to `false`.
	Manage pulumi.BoolPtrInput
	// The name to use for this Authorization Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Notification Hub Namespace in which the Notification Hub exists. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringPtrInput
	// The name of the Notification Hub for which the Authorization Rule should be created. Changing this forces a new resource to be created.
	NotificationHubName pulumi.StringPtrInput
	// The Primary Access Key associated with this Authorization Rule.
	PrimaryAccessKey pulumi.StringPtrInput
	// The name of the Resource Group in which the Notification Hub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Secondary Access Key associated with this Authorization Rule.
	SecondaryAccessKey pulumi.StringPtrInput
	// Does this Authorization Rule have Send access to the Notification Hub? Defaults to `false`.
	Send pulumi.BoolPtrInput
}

func (AuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationRuleState)(nil)).Elem()
}

type authorizationRuleArgs struct {
	// Does this Authorization Rule have Listen access to the Notification Hub? Defaults to `false`.
	Listen *bool `pulumi:"listen"`
	// Does this Authorization Rule have Manage access to the Notification Hub? Defaults to `false`.
	Manage *bool `pulumi:"manage"`
	// The name to use for this Authorization Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Notification Hub Namespace in which the Notification Hub exists. Changing this forces a new resource to be created.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the Notification Hub for which the Authorization Rule should be created. Changing this forces a new resource to be created.
	NotificationHubName string `pulumi:"notificationHubName"`
	// The name of the Resource Group in which the Notification Hub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Does this Authorization Rule have Send access to the Notification Hub? Defaults to `false`.
	Send *bool `pulumi:"send"`
}

// The set of arguments for constructing a AuthorizationRule resource.
type AuthorizationRuleArgs struct {
	// Does this Authorization Rule have Listen access to the Notification Hub? Defaults to `false`.
	Listen pulumi.BoolPtrInput
	// Does this Authorization Rule have Manage access to the Notification Hub? Defaults to `false`.
	Manage pulumi.BoolPtrInput
	// The name to use for this Authorization Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Notification Hub Namespace in which the Notification Hub exists. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput
	// The name of the Notification Hub for which the Authorization Rule should be created. Changing this forces a new resource to be created.
	NotificationHubName pulumi.StringInput
	// The name of the Resource Group in which the Notification Hub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Does this Authorization Rule have Send access to the Notification Hub? Defaults to `false`.
	Send pulumi.BoolPtrInput
}

func (AuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationRuleArgs)(nil)).Elem()
}

type AuthorizationRuleInput interface {
	pulumi.Input

	ToAuthorizationRuleOutput() AuthorizationRuleOutput
	ToAuthorizationRuleOutputWithContext(ctx context.Context) AuthorizationRuleOutput
}

func (*AuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationRule)(nil)).Elem()
}

func (i *AuthorizationRule) ToAuthorizationRuleOutput() AuthorizationRuleOutput {
	return i.ToAuthorizationRuleOutputWithContext(context.Background())
}

func (i *AuthorizationRule) ToAuthorizationRuleOutputWithContext(ctx context.Context) AuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationRuleOutput)
}

// AuthorizationRuleArrayInput is an input type that accepts AuthorizationRuleArray and AuthorizationRuleArrayOutput values.
// You can construct a concrete instance of `AuthorizationRuleArrayInput` via:
//
//          AuthorizationRuleArray{ AuthorizationRuleArgs{...} }
type AuthorizationRuleArrayInput interface {
	pulumi.Input

	ToAuthorizationRuleArrayOutput() AuthorizationRuleArrayOutput
	ToAuthorizationRuleArrayOutputWithContext(context.Context) AuthorizationRuleArrayOutput
}

type AuthorizationRuleArray []AuthorizationRuleInput

func (AuthorizationRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthorizationRule)(nil)).Elem()
}

func (i AuthorizationRuleArray) ToAuthorizationRuleArrayOutput() AuthorizationRuleArrayOutput {
	return i.ToAuthorizationRuleArrayOutputWithContext(context.Background())
}

func (i AuthorizationRuleArray) ToAuthorizationRuleArrayOutputWithContext(ctx context.Context) AuthorizationRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationRuleArrayOutput)
}

// AuthorizationRuleMapInput is an input type that accepts AuthorizationRuleMap and AuthorizationRuleMapOutput values.
// You can construct a concrete instance of `AuthorizationRuleMapInput` via:
//
//          AuthorizationRuleMap{ "key": AuthorizationRuleArgs{...} }
type AuthorizationRuleMapInput interface {
	pulumi.Input

	ToAuthorizationRuleMapOutput() AuthorizationRuleMapOutput
	ToAuthorizationRuleMapOutputWithContext(context.Context) AuthorizationRuleMapOutput
}

type AuthorizationRuleMap map[string]AuthorizationRuleInput

func (AuthorizationRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthorizationRule)(nil)).Elem()
}

func (i AuthorizationRuleMap) ToAuthorizationRuleMapOutput() AuthorizationRuleMapOutput {
	return i.ToAuthorizationRuleMapOutputWithContext(context.Background())
}

func (i AuthorizationRuleMap) ToAuthorizationRuleMapOutputWithContext(ctx context.Context) AuthorizationRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationRuleMapOutput)
}

type AuthorizationRuleOutput struct{ *pulumi.OutputState }

func (AuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationRule)(nil)).Elem()
}

func (o AuthorizationRuleOutput) ToAuthorizationRuleOutput() AuthorizationRuleOutput {
	return o
}

func (o AuthorizationRuleOutput) ToAuthorizationRuleOutputWithContext(ctx context.Context) AuthorizationRuleOutput {
	return o
}

type AuthorizationRuleArrayOutput struct{ *pulumi.OutputState }

func (AuthorizationRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthorizationRule)(nil)).Elem()
}

func (o AuthorizationRuleArrayOutput) ToAuthorizationRuleArrayOutput() AuthorizationRuleArrayOutput {
	return o
}

func (o AuthorizationRuleArrayOutput) ToAuthorizationRuleArrayOutputWithContext(ctx context.Context) AuthorizationRuleArrayOutput {
	return o
}

func (o AuthorizationRuleArrayOutput) Index(i pulumi.IntInput) AuthorizationRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthorizationRule {
		return vs[0].([]*AuthorizationRule)[vs[1].(int)]
	}).(AuthorizationRuleOutput)
}

type AuthorizationRuleMapOutput struct{ *pulumi.OutputState }

func (AuthorizationRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthorizationRule)(nil)).Elem()
}

func (o AuthorizationRuleMapOutput) ToAuthorizationRuleMapOutput() AuthorizationRuleMapOutput {
	return o
}

func (o AuthorizationRuleMapOutput) ToAuthorizationRuleMapOutputWithContext(ctx context.Context) AuthorizationRuleMapOutput {
	return o
}

func (o AuthorizationRuleMapOutput) MapIndex(k pulumi.StringInput) AuthorizationRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthorizationRule {
		return vs[0].(map[string]*AuthorizationRule)[vs[1].(string)]
	}).(AuthorizationRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthorizationRuleInput)(nil)).Elem(), &AuthorizationRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthorizationRuleArrayInput)(nil)).Elem(), AuthorizationRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthorizationRuleMapInput)(nil)).Elem(), AuthorizationRuleMap{})
	pulumi.RegisterOutputType(AuthorizationRuleOutput{})
	pulumi.RegisterOutputType(AuthorizationRuleArrayOutput{})
	pulumi.RegisterOutputType(AuthorizationRuleMapOutput{})
}
