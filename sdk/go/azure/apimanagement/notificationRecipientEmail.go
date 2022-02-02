// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a API Management Notification Recipient Email.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/apimanagement"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
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
// 		exampleService, err := apimanagement.NewService(ctx, "exampleService", &apimanagement.ServiceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			PublisherName:     pulumi.String("My Company"),
// 			PublisherEmail:    pulumi.String("company@terraform.io"),
// 			SkuName:           pulumi.String("Developer_1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apimanagement.NewNotificationRecipientEmail(ctx, "exampleNotificationRecipientEmail", &apimanagement.NotificationRecipientEmailArgs{
// 			ApiManagementId:  exampleService.ID(),
// 			NotificationType: pulumi.String("AccountClosedPublisher"),
// 			Email:            pulumi.String("foo@bar.com"),
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
// API Management Notification Recipient Emails can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/notificationRecipientEmail:NotificationRecipientEmail example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/notifications/notificationName1/recipientEmails/email1
// ```
type NotificationRecipientEmail struct {
	pulumi.CustomResourceState

	// The ID of the API Management Service from which to create this Notification Recipient Email. Changing this forces a new API Management Notification Recipient Email to be created.
	ApiManagementId pulumi.StringOutput `pulumi:"apiManagementId"`
	// The recipient email address. Changing this forces a new API Management Notification Recipient Email to be created.
	Email pulumi.StringOutput `pulumi:"email"`
	// The Notification Name to be received. Changing this forces a new API Management Notification Recipient Email to be created. Possible values are `AccountClosedPublisher`, `BCC`, `NewApplicationNotificationMessage`, `NewIssuePublisherNotificationMessage`, `PurchasePublisherNotificationMessage`, `QuotaLimitApproachingPublisherNotificationMessage`, and `RequestPublisherNotificationMessage`.
	NotificationType pulumi.StringOutput `pulumi:"notificationType"`
}

// NewNotificationRecipientEmail registers a new resource with the given unique name, arguments, and options.
func NewNotificationRecipientEmail(ctx *pulumi.Context,
	name string, args *NotificationRecipientEmailArgs, opts ...pulumi.ResourceOption) (*NotificationRecipientEmail, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiManagementId == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementId'")
	}
	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.NotificationType == nil {
		return nil, errors.New("invalid value for required argument 'NotificationType'")
	}
	var resource NotificationRecipientEmail
	err := ctx.RegisterResource("azure:apimanagement/notificationRecipientEmail:NotificationRecipientEmail", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotificationRecipientEmail gets an existing NotificationRecipientEmail resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotificationRecipientEmail(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationRecipientEmailState, opts ...pulumi.ResourceOption) (*NotificationRecipientEmail, error) {
	var resource NotificationRecipientEmail
	err := ctx.ReadResource("azure:apimanagement/notificationRecipientEmail:NotificationRecipientEmail", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotificationRecipientEmail resources.
type notificationRecipientEmailState struct {
	// The ID of the API Management Service from which to create this Notification Recipient Email. Changing this forces a new API Management Notification Recipient Email to be created.
	ApiManagementId *string `pulumi:"apiManagementId"`
	// The recipient email address. Changing this forces a new API Management Notification Recipient Email to be created.
	Email *string `pulumi:"email"`
	// The Notification Name to be received. Changing this forces a new API Management Notification Recipient Email to be created. Possible values are `AccountClosedPublisher`, `BCC`, `NewApplicationNotificationMessage`, `NewIssuePublisherNotificationMessage`, `PurchasePublisherNotificationMessage`, `QuotaLimitApproachingPublisherNotificationMessage`, and `RequestPublisherNotificationMessage`.
	NotificationType *string `pulumi:"notificationType"`
}

type NotificationRecipientEmailState struct {
	// The ID of the API Management Service from which to create this Notification Recipient Email. Changing this forces a new API Management Notification Recipient Email to be created.
	ApiManagementId pulumi.StringPtrInput
	// The recipient email address. Changing this forces a new API Management Notification Recipient Email to be created.
	Email pulumi.StringPtrInput
	// The Notification Name to be received. Changing this forces a new API Management Notification Recipient Email to be created. Possible values are `AccountClosedPublisher`, `BCC`, `NewApplicationNotificationMessage`, `NewIssuePublisherNotificationMessage`, `PurchasePublisherNotificationMessage`, `QuotaLimitApproachingPublisherNotificationMessage`, and `RequestPublisherNotificationMessage`.
	NotificationType pulumi.StringPtrInput
}

func (NotificationRecipientEmailState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationRecipientEmailState)(nil)).Elem()
}

type notificationRecipientEmailArgs struct {
	// The ID of the API Management Service from which to create this Notification Recipient Email. Changing this forces a new API Management Notification Recipient Email to be created.
	ApiManagementId string `pulumi:"apiManagementId"`
	// The recipient email address. Changing this forces a new API Management Notification Recipient Email to be created.
	Email string `pulumi:"email"`
	// The Notification Name to be received. Changing this forces a new API Management Notification Recipient Email to be created. Possible values are `AccountClosedPublisher`, `BCC`, `NewApplicationNotificationMessage`, `NewIssuePublisherNotificationMessage`, `PurchasePublisherNotificationMessage`, `QuotaLimitApproachingPublisherNotificationMessage`, and `RequestPublisherNotificationMessage`.
	NotificationType string `pulumi:"notificationType"`
}

// The set of arguments for constructing a NotificationRecipientEmail resource.
type NotificationRecipientEmailArgs struct {
	// The ID of the API Management Service from which to create this Notification Recipient Email. Changing this forces a new API Management Notification Recipient Email to be created.
	ApiManagementId pulumi.StringInput
	// The recipient email address. Changing this forces a new API Management Notification Recipient Email to be created.
	Email pulumi.StringInput
	// The Notification Name to be received. Changing this forces a new API Management Notification Recipient Email to be created. Possible values are `AccountClosedPublisher`, `BCC`, `NewApplicationNotificationMessage`, `NewIssuePublisherNotificationMessage`, `PurchasePublisherNotificationMessage`, `QuotaLimitApproachingPublisherNotificationMessage`, and `RequestPublisherNotificationMessage`.
	NotificationType pulumi.StringInput
}

func (NotificationRecipientEmailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationRecipientEmailArgs)(nil)).Elem()
}

type NotificationRecipientEmailInput interface {
	pulumi.Input

	ToNotificationRecipientEmailOutput() NotificationRecipientEmailOutput
	ToNotificationRecipientEmailOutputWithContext(ctx context.Context) NotificationRecipientEmailOutput
}

func (*NotificationRecipientEmail) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationRecipientEmail)(nil)).Elem()
}

func (i *NotificationRecipientEmail) ToNotificationRecipientEmailOutput() NotificationRecipientEmailOutput {
	return i.ToNotificationRecipientEmailOutputWithContext(context.Background())
}

func (i *NotificationRecipientEmail) ToNotificationRecipientEmailOutputWithContext(ctx context.Context) NotificationRecipientEmailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationRecipientEmailOutput)
}

// NotificationRecipientEmailArrayInput is an input type that accepts NotificationRecipientEmailArray and NotificationRecipientEmailArrayOutput values.
// You can construct a concrete instance of `NotificationRecipientEmailArrayInput` via:
//
//          NotificationRecipientEmailArray{ NotificationRecipientEmailArgs{...} }
type NotificationRecipientEmailArrayInput interface {
	pulumi.Input

	ToNotificationRecipientEmailArrayOutput() NotificationRecipientEmailArrayOutput
	ToNotificationRecipientEmailArrayOutputWithContext(context.Context) NotificationRecipientEmailArrayOutput
}

type NotificationRecipientEmailArray []NotificationRecipientEmailInput

func (NotificationRecipientEmailArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotificationRecipientEmail)(nil)).Elem()
}

func (i NotificationRecipientEmailArray) ToNotificationRecipientEmailArrayOutput() NotificationRecipientEmailArrayOutput {
	return i.ToNotificationRecipientEmailArrayOutputWithContext(context.Background())
}

func (i NotificationRecipientEmailArray) ToNotificationRecipientEmailArrayOutputWithContext(ctx context.Context) NotificationRecipientEmailArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationRecipientEmailArrayOutput)
}

// NotificationRecipientEmailMapInput is an input type that accepts NotificationRecipientEmailMap and NotificationRecipientEmailMapOutput values.
// You can construct a concrete instance of `NotificationRecipientEmailMapInput` via:
//
//          NotificationRecipientEmailMap{ "key": NotificationRecipientEmailArgs{...} }
type NotificationRecipientEmailMapInput interface {
	pulumi.Input

	ToNotificationRecipientEmailMapOutput() NotificationRecipientEmailMapOutput
	ToNotificationRecipientEmailMapOutputWithContext(context.Context) NotificationRecipientEmailMapOutput
}

type NotificationRecipientEmailMap map[string]NotificationRecipientEmailInput

func (NotificationRecipientEmailMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotificationRecipientEmail)(nil)).Elem()
}

func (i NotificationRecipientEmailMap) ToNotificationRecipientEmailMapOutput() NotificationRecipientEmailMapOutput {
	return i.ToNotificationRecipientEmailMapOutputWithContext(context.Background())
}

func (i NotificationRecipientEmailMap) ToNotificationRecipientEmailMapOutputWithContext(ctx context.Context) NotificationRecipientEmailMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationRecipientEmailMapOutput)
}

type NotificationRecipientEmailOutput struct{ *pulumi.OutputState }

func (NotificationRecipientEmailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationRecipientEmail)(nil)).Elem()
}

func (o NotificationRecipientEmailOutput) ToNotificationRecipientEmailOutput() NotificationRecipientEmailOutput {
	return o
}

func (o NotificationRecipientEmailOutput) ToNotificationRecipientEmailOutputWithContext(ctx context.Context) NotificationRecipientEmailOutput {
	return o
}

type NotificationRecipientEmailArrayOutput struct{ *pulumi.OutputState }

func (NotificationRecipientEmailArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotificationRecipientEmail)(nil)).Elem()
}

func (o NotificationRecipientEmailArrayOutput) ToNotificationRecipientEmailArrayOutput() NotificationRecipientEmailArrayOutput {
	return o
}

func (o NotificationRecipientEmailArrayOutput) ToNotificationRecipientEmailArrayOutputWithContext(ctx context.Context) NotificationRecipientEmailArrayOutput {
	return o
}

func (o NotificationRecipientEmailArrayOutput) Index(i pulumi.IntInput) NotificationRecipientEmailOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NotificationRecipientEmail {
		return vs[0].([]*NotificationRecipientEmail)[vs[1].(int)]
	}).(NotificationRecipientEmailOutput)
}

type NotificationRecipientEmailMapOutput struct{ *pulumi.OutputState }

func (NotificationRecipientEmailMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotificationRecipientEmail)(nil)).Elem()
}

func (o NotificationRecipientEmailMapOutput) ToNotificationRecipientEmailMapOutput() NotificationRecipientEmailMapOutput {
	return o
}

func (o NotificationRecipientEmailMapOutput) ToNotificationRecipientEmailMapOutputWithContext(ctx context.Context) NotificationRecipientEmailMapOutput {
	return o
}

func (o NotificationRecipientEmailMapOutput) MapIndex(k pulumi.StringInput) NotificationRecipientEmailOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NotificationRecipientEmail {
		return vs[0].(map[string]*NotificationRecipientEmail)[vs[1].(string)]
	}).(NotificationRecipientEmailOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationRecipientEmailInput)(nil)).Elem(), &NotificationRecipientEmail{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationRecipientEmailArrayInput)(nil)).Elem(), NotificationRecipientEmailArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationRecipientEmailMapInput)(nil)).Elem(), NotificationRecipientEmailMap{})
	pulumi.RegisterOutputType(NotificationRecipientEmailOutput{})
	pulumi.RegisterOutputType(NotificationRecipientEmailArrayOutput{})
	pulumi.RegisterOutputType(NotificationRecipientEmailMapOutput{})
}
