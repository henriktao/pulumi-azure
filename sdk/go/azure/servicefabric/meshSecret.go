// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicefabric

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Service Fabric Mesh Secret can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:servicefabric/meshSecret:MeshSecret secret1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceFabricMesh/secrets/secret1
// ```
type MeshSecret struct {
	pulumi.CustomResourceState

	// The type of the content stored in the secret value. Changing this forces a new resource to be created.
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	// A description of this Service Fabric Mesh Secret.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the Azure Region where the Service Fabric Mesh Secret should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Service Fabric Mesh Secret. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group in which the Service Fabric Mesh Secret exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewMeshSecret registers a new resource with the given unique name, arguments, and options.
func NewMeshSecret(ctx *pulumi.Context,
	name string, args *MeshSecretArgs, opts ...pulumi.ResourceOption) (*MeshSecret, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource MeshSecret
	err := ctx.RegisterResource("azure:servicefabric/meshSecret:MeshSecret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMeshSecret gets an existing MeshSecret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMeshSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MeshSecretState, opts ...pulumi.ResourceOption) (*MeshSecret, error) {
	var resource MeshSecret
	err := ctx.ReadResource("azure:servicefabric/meshSecret:MeshSecret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MeshSecret resources.
type meshSecretState struct {
	// The type of the content stored in the secret value. Changing this forces a new resource to be created.
	ContentType *string `pulumi:"contentType"`
	// A description of this Service Fabric Mesh Secret.
	Description *string `pulumi:"description"`
	// Specifies the Azure Region where the Service Fabric Mesh Secret should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Service Fabric Mesh Secret. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the Service Fabric Mesh Secret exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type MeshSecretState struct {
	// The type of the content stored in the secret value. Changing this forces a new resource to be created.
	ContentType pulumi.StringPtrInput
	// A description of this Service Fabric Mesh Secret.
	Description pulumi.StringPtrInput
	// Specifies the Azure Region where the Service Fabric Mesh Secret should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Service Fabric Mesh Secret. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the Service Fabric Mesh Secret exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (MeshSecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*meshSecretState)(nil)).Elem()
}

type meshSecretArgs struct {
	// The type of the content stored in the secret value. Changing this forces a new resource to be created.
	ContentType *string `pulumi:"contentType"`
	// A description of this Service Fabric Mesh Secret.
	Description *string `pulumi:"description"`
	// Specifies the Azure Region where the Service Fabric Mesh Secret should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Service Fabric Mesh Secret. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the Service Fabric Mesh Secret exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MeshSecret resource.
type MeshSecretArgs struct {
	// The type of the content stored in the secret value. Changing this forces a new resource to be created.
	ContentType pulumi.StringPtrInput
	// A description of this Service Fabric Mesh Secret.
	Description pulumi.StringPtrInput
	// Specifies the Azure Region where the Service Fabric Mesh Secret should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Service Fabric Mesh Secret. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the Service Fabric Mesh Secret exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (MeshSecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*meshSecretArgs)(nil)).Elem()
}

type MeshSecretInput interface {
	pulumi.Input

	ToMeshSecretOutput() MeshSecretOutput
	ToMeshSecretOutputWithContext(ctx context.Context) MeshSecretOutput
}

func (*MeshSecret) ElementType() reflect.Type {
	return reflect.TypeOf((*MeshSecret)(nil))
}

func (i *MeshSecret) ToMeshSecretOutput() MeshSecretOutput {
	return i.ToMeshSecretOutputWithContext(context.Background())
}

func (i *MeshSecret) ToMeshSecretOutputWithContext(ctx context.Context) MeshSecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeshSecretOutput)
}

func (i *MeshSecret) ToMeshSecretPtrOutput() MeshSecretPtrOutput {
	return i.ToMeshSecretPtrOutputWithContext(context.Background())
}

func (i *MeshSecret) ToMeshSecretPtrOutputWithContext(ctx context.Context) MeshSecretPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeshSecretPtrOutput)
}

type MeshSecretPtrInput interface {
	pulumi.Input

	ToMeshSecretPtrOutput() MeshSecretPtrOutput
	ToMeshSecretPtrOutputWithContext(ctx context.Context) MeshSecretPtrOutput
}

type meshSecretPtrType MeshSecretArgs

func (*meshSecretPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MeshSecret)(nil))
}

func (i *meshSecretPtrType) ToMeshSecretPtrOutput() MeshSecretPtrOutput {
	return i.ToMeshSecretPtrOutputWithContext(context.Background())
}

func (i *meshSecretPtrType) ToMeshSecretPtrOutputWithContext(ctx context.Context) MeshSecretPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeshSecretPtrOutput)
}

// MeshSecretArrayInput is an input type that accepts MeshSecretArray and MeshSecretArrayOutput values.
// You can construct a concrete instance of `MeshSecretArrayInput` via:
//
//          MeshSecretArray{ MeshSecretArgs{...} }
type MeshSecretArrayInput interface {
	pulumi.Input

	ToMeshSecretArrayOutput() MeshSecretArrayOutput
	ToMeshSecretArrayOutputWithContext(context.Context) MeshSecretArrayOutput
}

type MeshSecretArray []MeshSecretInput

func (MeshSecretArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MeshSecret)(nil)).Elem()
}

func (i MeshSecretArray) ToMeshSecretArrayOutput() MeshSecretArrayOutput {
	return i.ToMeshSecretArrayOutputWithContext(context.Background())
}

func (i MeshSecretArray) ToMeshSecretArrayOutputWithContext(ctx context.Context) MeshSecretArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeshSecretArrayOutput)
}

// MeshSecretMapInput is an input type that accepts MeshSecretMap and MeshSecretMapOutput values.
// You can construct a concrete instance of `MeshSecretMapInput` via:
//
//          MeshSecretMap{ "key": MeshSecretArgs{...} }
type MeshSecretMapInput interface {
	pulumi.Input

	ToMeshSecretMapOutput() MeshSecretMapOutput
	ToMeshSecretMapOutputWithContext(context.Context) MeshSecretMapOutput
}

type MeshSecretMap map[string]MeshSecretInput

func (MeshSecretMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MeshSecret)(nil)).Elem()
}

func (i MeshSecretMap) ToMeshSecretMapOutput() MeshSecretMapOutput {
	return i.ToMeshSecretMapOutputWithContext(context.Background())
}

func (i MeshSecretMap) ToMeshSecretMapOutputWithContext(ctx context.Context) MeshSecretMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeshSecretMapOutput)
}

type MeshSecretOutput struct{ *pulumi.OutputState }

func (MeshSecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MeshSecret)(nil))
}

func (o MeshSecretOutput) ToMeshSecretOutput() MeshSecretOutput {
	return o
}

func (o MeshSecretOutput) ToMeshSecretOutputWithContext(ctx context.Context) MeshSecretOutput {
	return o
}

func (o MeshSecretOutput) ToMeshSecretPtrOutput() MeshSecretPtrOutput {
	return o.ToMeshSecretPtrOutputWithContext(context.Background())
}

func (o MeshSecretOutput) ToMeshSecretPtrOutputWithContext(ctx context.Context) MeshSecretPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MeshSecret) *MeshSecret {
		return &v
	}).(MeshSecretPtrOutput)
}

type MeshSecretPtrOutput struct{ *pulumi.OutputState }

func (MeshSecretPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MeshSecret)(nil))
}

func (o MeshSecretPtrOutput) ToMeshSecretPtrOutput() MeshSecretPtrOutput {
	return o
}

func (o MeshSecretPtrOutput) ToMeshSecretPtrOutputWithContext(ctx context.Context) MeshSecretPtrOutput {
	return o
}

func (o MeshSecretPtrOutput) Elem() MeshSecretOutput {
	return o.ApplyT(func(v *MeshSecret) MeshSecret {
		if v != nil {
			return *v
		}
		var ret MeshSecret
		return ret
	}).(MeshSecretOutput)
}

type MeshSecretArrayOutput struct{ *pulumi.OutputState }

func (MeshSecretArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MeshSecret)(nil))
}

func (o MeshSecretArrayOutput) ToMeshSecretArrayOutput() MeshSecretArrayOutput {
	return o
}

func (o MeshSecretArrayOutput) ToMeshSecretArrayOutputWithContext(ctx context.Context) MeshSecretArrayOutput {
	return o
}

func (o MeshSecretArrayOutput) Index(i pulumi.IntInput) MeshSecretOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MeshSecret {
		return vs[0].([]MeshSecret)[vs[1].(int)]
	}).(MeshSecretOutput)
}

type MeshSecretMapOutput struct{ *pulumi.OutputState }

func (MeshSecretMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MeshSecret)(nil))
}

func (o MeshSecretMapOutput) ToMeshSecretMapOutput() MeshSecretMapOutput {
	return o
}

func (o MeshSecretMapOutput) ToMeshSecretMapOutputWithContext(ctx context.Context) MeshSecretMapOutput {
	return o
}

func (o MeshSecretMapOutput) MapIndex(k pulumi.StringInput) MeshSecretOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MeshSecret {
		return vs[0].(map[string]MeshSecret)[vs[1].(string)]
	}).(MeshSecretOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MeshSecretInput)(nil)).Elem(), &MeshSecret{})
	pulumi.RegisterInputType(reflect.TypeOf((*MeshSecretPtrInput)(nil)).Elem(), &MeshSecret{})
	pulumi.RegisterInputType(reflect.TypeOf((*MeshSecretArrayInput)(nil)).Elem(), MeshSecretArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MeshSecretMapInput)(nil)).Elem(), MeshSecretMap{})
	pulumi.RegisterOutputType(MeshSecretOutput{})
	pulumi.RegisterOutputType(MeshSecretPtrOutput{})
	pulumi.RegisterOutputType(MeshSecretArrayOutput{})
	pulumi.RegisterOutputType(MeshSecretMapOutput{})
}
