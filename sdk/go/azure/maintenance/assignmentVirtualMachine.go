// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package maintenance

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a maintenance assignment to virtual machine.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/maintenance"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnet, err := network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.2.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleNetworkInterface, err := network.NewNetworkInterface(ctx, "exampleNetworkInterface", &network.NetworkInterfaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			IpConfigurations: network.NetworkInterfaceIpConfigurationArray{
// 				&network.NetworkInterfaceIpConfigurationArgs{
// 					Name:                       pulumi.String("internal"),
// 					SubnetId:                   exampleSubnet.ID(),
// 					PrivateIpAddressAllocation: pulumi.String("Dynamic"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleLinuxVirtualMachine, err := compute.NewLinuxVirtualMachine(ctx, "exampleLinuxVirtualMachine", &compute.LinuxVirtualMachineArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Size:              pulumi.String("Standard_F2"),
// 			AdminUsername:     pulumi.String("adminuser"),
// 			NetworkInterfaceIds: pulumi.StringArray{
// 				exampleNetworkInterface.ID(),
// 			},
// 			AdminSshKeys: compute.LinuxVirtualMachineAdminSshKeyArray{
// 				&compute.LinuxVirtualMachineAdminSshKeyArgs{
// 					Username:  pulumi.String("adminuser"),
// 					PublicKey: readFileOrPanic("~/.ssh/id_rsa.pub"),
// 				},
// 			},
// 			OsDisk: &compute.LinuxVirtualMachineOsDiskArgs{
// 				Caching:            pulumi.String("ReadWrite"),
// 				StorageAccountType: pulumi.String("Standard_LRS"),
// 			},
// 			SourceImageReference: &compute.LinuxVirtualMachineSourceImageReferenceArgs{
// 				Publisher: pulumi.String("Canonical"),
// 				Offer:     pulumi.String("UbuntuServer"),
// 				Sku:       pulumi.String("16.04-LTS"),
// 				Version:   pulumi.String("latest"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleConfiguration, err := maintenance.NewConfiguration(ctx, "exampleConfiguration", &maintenance.ConfigurationArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Scope:             pulumi.String("All"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = maintenance.NewAssignmentVirtualMachine(ctx, "exampleAssignmentVirtualMachine", &maintenance.AssignmentVirtualMachineArgs{
// 			Location:                   exampleResourceGroup.Location,
// 			MaintenanceConfigurationId: exampleConfiguration.ID(),
// 			VirtualMachineId:           exampleLinuxVirtualMachine.ID(),
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
// Maintenance Assignment can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:maintenance/assignmentVirtualMachine:AssignmentVirtualMachine example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resGroup1/providers/microsoft.compute/virtualMachines/vm1/providers/Microsoft.Maintenance/configurationAssignments/assign1
// ```
type AssignmentVirtualMachine struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
	MaintenanceConfigurationId pulumi.StringOutput `pulumi:"maintenanceConfigurationId"`
	// Specifies the Virtual Machine ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
	VirtualMachineId pulumi.StringOutput `pulumi:"virtualMachineId"`
}

// NewAssignmentVirtualMachine registers a new resource with the given unique name, arguments, and options.
func NewAssignmentVirtualMachine(ctx *pulumi.Context,
	name string, args *AssignmentVirtualMachineArgs, opts ...pulumi.ResourceOption) (*AssignmentVirtualMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MaintenanceConfigurationId == nil {
		return nil, errors.New("invalid value for required argument 'MaintenanceConfigurationId'")
	}
	if args.VirtualMachineId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineId'")
	}
	var resource AssignmentVirtualMachine
	err := ctx.RegisterResource("azure:maintenance/assignmentVirtualMachine:AssignmentVirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssignmentVirtualMachine gets an existing AssignmentVirtualMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssignmentVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssignmentVirtualMachineState, opts ...pulumi.ResourceOption) (*AssignmentVirtualMachine, error) {
	var resource AssignmentVirtualMachine
	err := ctx.ReadResource("azure:maintenance/assignmentVirtualMachine:AssignmentVirtualMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AssignmentVirtualMachine resources.
type assignmentVirtualMachineState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
	MaintenanceConfigurationId *string `pulumi:"maintenanceConfigurationId"`
	// Specifies the Virtual Machine ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
	VirtualMachineId *string `pulumi:"virtualMachineId"`
}

type AssignmentVirtualMachineState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
	MaintenanceConfigurationId pulumi.StringPtrInput
	// Specifies the Virtual Machine ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
	VirtualMachineId pulumi.StringPtrInput
}

func (AssignmentVirtualMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentVirtualMachineState)(nil)).Elem()
}

type assignmentVirtualMachineArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
	MaintenanceConfigurationId string `pulumi:"maintenanceConfigurationId"`
	// Specifies the Virtual Machine ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
	VirtualMachineId string `pulumi:"virtualMachineId"`
}

// The set of arguments for constructing a AssignmentVirtualMachine resource.
type AssignmentVirtualMachineArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
	MaintenanceConfigurationId pulumi.StringInput
	// Specifies the Virtual Machine ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
	VirtualMachineId pulumi.StringInput
}

func (AssignmentVirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentVirtualMachineArgs)(nil)).Elem()
}

type AssignmentVirtualMachineInput interface {
	pulumi.Input

	ToAssignmentVirtualMachineOutput() AssignmentVirtualMachineOutput
	ToAssignmentVirtualMachineOutputWithContext(ctx context.Context) AssignmentVirtualMachineOutput
}

func (*AssignmentVirtualMachine) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentVirtualMachine)(nil)).Elem()
}

func (i *AssignmentVirtualMachine) ToAssignmentVirtualMachineOutput() AssignmentVirtualMachineOutput {
	return i.ToAssignmentVirtualMachineOutputWithContext(context.Background())
}

func (i *AssignmentVirtualMachine) ToAssignmentVirtualMachineOutputWithContext(ctx context.Context) AssignmentVirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentVirtualMachineOutput)
}

// AssignmentVirtualMachineArrayInput is an input type that accepts AssignmentVirtualMachineArray and AssignmentVirtualMachineArrayOutput values.
// You can construct a concrete instance of `AssignmentVirtualMachineArrayInput` via:
//
//          AssignmentVirtualMachineArray{ AssignmentVirtualMachineArgs{...} }
type AssignmentVirtualMachineArrayInput interface {
	pulumi.Input

	ToAssignmentVirtualMachineArrayOutput() AssignmentVirtualMachineArrayOutput
	ToAssignmentVirtualMachineArrayOutputWithContext(context.Context) AssignmentVirtualMachineArrayOutput
}

type AssignmentVirtualMachineArray []AssignmentVirtualMachineInput

func (AssignmentVirtualMachineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AssignmentVirtualMachine)(nil)).Elem()
}

func (i AssignmentVirtualMachineArray) ToAssignmentVirtualMachineArrayOutput() AssignmentVirtualMachineArrayOutput {
	return i.ToAssignmentVirtualMachineArrayOutputWithContext(context.Background())
}

func (i AssignmentVirtualMachineArray) ToAssignmentVirtualMachineArrayOutputWithContext(ctx context.Context) AssignmentVirtualMachineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentVirtualMachineArrayOutput)
}

// AssignmentVirtualMachineMapInput is an input type that accepts AssignmentVirtualMachineMap and AssignmentVirtualMachineMapOutput values.
// You can construct a concrete instance of `AssignmentVirtualMachineMapInput` via:
//
//          AssignmentVirtualMachineMap{ "key": AssignmentVirtualMachineArgs{...} }
type AssignmentVirtualMachineMapInput interface {
	pulumi.Input

	ToAssignmentVirtualMachineMapOutput() AssignmentVirtualMachineMapOutput
	ToAssignmentVirtualMachineMapOutputWithContext(context.Context) AssignmentVirtualMachineMapOutput
}

type AssignmentVirtualMachineMap map[string]AssignmentVirtualMachineInput

func (AssignmentVirtualMachineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AssignmentVirtualMachine)(nil)).Elem()
}

func (i AssignmentVirtualMachineMap) ToAssignmentVirtualMachineMapOutput() AssignmentVirtualMachineMapOutput {
	return i.ToAssignmentVirtualMachineMapOutputWithContext(context.Background())
}

func (i AssignmentVirtualMachineMap) ToAssignmentVirtualMachineMapOutputWithContext(ctx context.Context) AssignmentVirtualMachineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentVirtualMachineMapOutput)
}

type AssignmentVirtualMachineOutput struct{ *pulumi.OutputState }

func (AssignmentVirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentVirtualMachine)(nil)).Elem()
}

func (o AssignmentVirtualMachineOutput) ToAssignmentVirtualMachineOutput() AssignmentVirtualMachineOutput {
	return o
}

func (o AssignmentVirtualMachineOutput) ToAssignmentVirtualMachineOutputWithContext(ctx context.Context) AssignmentVirtualMachineOutput {
	return o
}

type AssignmentVirtualMachineArrayOutput struct{ *pulumi.OutputState }

func (AssignmentVirtualMachineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AssignmentVirtualMachine)(nil)).Elem()
}

func (o AssignmentVirtualMachineArrayOutput) ToAssignmentVirtualMachineArrayOutput() AssignmentVirtualMachineArrayOutput {
	return o
}

func (o AssignmentVirtualMachineArrayOutput) ToAssignmentVirtualMachineArrayOutputWithContext(ctx context.Context) AssignmentVirtualMachineArrayOutput {
	return o
}

func (o AssignmentVirtualMachineArrayOutput) Index(i pulumi.IntInput) AssignmentVirtualMachineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AssignmentVirtualMachine {
		return vs[0].([]*AssignmentVirtualMachine)[vs[1].(int)]
	}).(AssignmentVirtualMachineOutput)
}

type AssignmentVirtualMachineMapOutput struct{ *pulumi.OutputState }

func (AssignmentVirtualMachineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AssignmentVirtualMachine)(nil)).Elem()
}

func (o AssignmentVirtualMachineMapOutput) ToAssignmentVirtualMachineMapOutput() AssignmentVirtualMachineMapOutput {
	return o
}

func (o AssignmentVirtualMachineMapOutput) ToAssignmentVirtualMachineMapOutputWithContext(ctx context.Context) AssignmentVirtualMachineMapOutput {
	return o
}

func (o AssignmentVirtualMachineMapOutput) MapIndex(k pulumi.StringInput) AssignmentVirtualMachineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AssignmentVirtualMachine {
		return vs[0].(map[string]*AssignmentVirtualMachine)[vs[1].(string)]
	}).(AssignmentVirtualMachineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AssignmentVirtualMachineInput)(nil)).Elem(), &AssignmentVirtualMachine{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssignmentVirtualMachineArrayInput)(nil)).Elem(), AssignmentVirtualMachineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssignmentVirtualMachineMapInput)(nil)).Elem(), AssignmentVirtualMachineMap{})
	pulumi.RegisterOutputType(AssignmentVirtualMachineOutput{})
	pulumi.RegisterOutputType(AssignmentVirtualMachineArrayOutput{})
	pulumi.RegisterOutputType(AssignmentVirtualMachineMapOutput{})
}
