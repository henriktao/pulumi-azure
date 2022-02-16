// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Orchestrated Virtual Machine Scale Set.
//
// ## Disclaimers
//
// > **NOTE:** As of the **v2.86.0** (November 19, 2021) release of the provider this resource will only create Virtual Machine Scale Sets with the **Flexible** Orchestration Mode.
//
// > **NOTE:** All arguments including the administrator login and password will be stored in the raw state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/compute"
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
// 		_, err = compute.NewOrchestratedVirtualMachineScaleSet(ctx, "exampleOrchestratedVirtualMachineScaleSet", &compute.OrchestratedVirtualMachineScaleSetArgs{
// 			Location:                 exampleResourceGroup.Location,
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			PlatformFaultDomainCount: pulumi.Int(1),
// 			Zones: pulumi.String{
// 				"1",
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
// An Orchestrated Virtual Machine Scale Set can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/Microsoft.Compute/virtualMachineScaleSets/scaleset1
// ```
type OrchestratedVirtualMachineScaleSet struct {
	pulumi.CustomResourceState

	AutomaticInstanceRepair OrchestratedVirtualMachineScaleSetAutomaticInstanceRepairOutput `pulumi:"automaticInstanceRepair"`
	BootDiagnostics         OrchestratedVirtualMachineScaleSetBootDiagnosticsPtrOutput      `pulumi:"bootDiagnostics"`
	DataDisks               OrchestratedVirtualMachineScaleSetDataDiskArrayOutput           `pulumi:"dataDisks"`
	EncryptionAtHostEnabled pulumi.BoolPtrOutput                                            `pulumi:"encryptionAtHostEnabled"`
	EvictionPolicy          pulumi.StringPtrOutput                                          `pulumi:"evictionPolicy"`
	Extensions              OrchestratedVirtualMachineScaleSetExtensionArrayOutput          `pulumi:"extensions"`
	// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M).
	ExtensionsTimeBudget pulumi.StringPtrOutput                              `pulumi:"extensionsTimeBudget"`
	Identity             OrchestratedVirtualMachineScaleSetIdentityPtrOutput `pulumi:"identity"`
	// The number of Virtual Machines in the Orcestrated Virtual Machine Scale Set.
	Instances   pulumi.IntOutput       `pulumi:"instances"`
	LicenseType pulumi.StringPtrOutput `pulumi:"licenseType"`
	// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location    pulumi.StringOutput     `pulumi:"location"`
	MaxBidPrice pulumi.Float64PtrOutput `pulumi:"maxBidPrice"`
	// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name              pulumi.StringOutput                                           `pulumi:"name"`
	NetworkInterfaces OrchestratedVirtualMachineScaleSetNetworkInterfaceArrayOutput `pulumi:"networkInterfaces"`
	OsDisk            OrchestratedVirtualMachineScaleSetOsDiskPtrOutput             `pulumi:"osDisk"`
	OsProfile         OrchestratedVirtualMachineScaleSetOsProfilePtrOutput          `pulumi:"osProfile"`
	Plan              OrchestratedVirtualMachineScaleSetPlanPtrOutput               `pulumi:"plan"`
	// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	PlatformFaultDomainCount pulumi.IntOutput       `pulumi:"platformFaultDomainCount"`
	Priority                 pulumi.StringPtrOutput `pulumi:"priority"`
	// The ID of the Proximity Placement Group which the Orchestrated Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId pulumi.StringPtrOutput `pulumi:"proximityPlacementGroupId"`
	// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput    `pulumi:"resourceGroupName"`
	SkuName           pulumi.StringPtrOutput `pulumi:"skuName"`
	SourceImageId     pulumi.StringPtrOutput `pulumi:"sourceImageId"`
	// A `sourceImageReference` block as defined below.
	SourceImageReference OrchestratedVirtualMachineScaleSetSourceImageReferencePtrOutput `pulumi:"sourceImageReference"`
	// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
	Tags                    pulumi.StringMapOutput                                          `pulumi:"tags"`
	TerminationNotification OrchestratedVirtualMachineScaleSetTerminationNotificationOutput `pulumi:"terminationNotification"`
	// The Unique ID for the Orchestrated Virtual Machine Scale Set.
	UniqueId    pulumi.StringOutput  `pulumi:"uniqueId"`
	ZoneBalance pulumi.BoolPtrOutput `pulumi:"zoneBalance"`
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones pulumi.StringPtrOutput `pulumi:"zones"`
}

// NewOrchestratedVirtualMachineScaleSet registers a new resource with the given unique name, arguments, and options.
func NewOrchestratedVirtualMachineScaleSet(ctx *pulumi.Context,
	name string, args *OrchestratedVirtualMachineScaleSetArgs, opts ...pulumi.ResourceOption) (*OrchestratedVirtualMachineScaleSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PlatformFaultDomainCount == nil {
		return nil, errors.New("invalid value for required argument 'PlatformFaultDomainCount'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource OrchestratedVirtualMachineScaleSet
	err := ctx.RegisterResource("azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrchestratedVirtualMachineScaleSet gets an existing OrchestratedVirtualMachineScaleSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrchestratedVirtualMachineScaleSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrchestratedVirtualMachineScaleSetState, opts ...pulumi.ResourceOption) (*OrchestratedVirtualMachineScaleSet, error) {
	var resource OrchestratedVirtualMachineScaleSet
	err := ctx.ReadResource("azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrchestratedVirtualMachineScaleSet resources.
type orchestratedVirtualMachineScaleSetState struct {
	AutomaticInstanceRepair *OrchestratedVirtualMachineScaleSetAutomaticInstanceRepair `pulumi:"automaticInstanceRepair"`
	BootDiagnostics         *OrchestratedVirtualMachineScaleSetBootDiagnostics         `pulumi:"bootDiagnostics"`
	DataDisks               []OrchestratedVirtualMachineScaleSetDataDisk               `pulumi:"dataDisks"`
	EncryptionAtHostEnabled *bool                                                      `pulumi:"encryptionAtHostEnabled"`
	EvictionPolicy          *string                                                    `pulumi:"evictionPolicy"`
	Extensions              []OrchestratedVirtualMachineScaleSetExtension              `pulumi:"extensions"`
	// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M).
	ExtensionsTimeBudget *string                                     `pulumi:"extensionsTimeBudget"`
	Identity             *OrchestratedVirtualMachineScaleSetIdentity `pulumi:"identity"`
	// The number of Virtual Machines in the Orcestrated Virtual Machine Scale Set.
	Instances   *int    `pulumi:"instances"`
	LicenseType *string `pulumi:"licenseType"`
	// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location    *string  `pulumi:"location"`
	MaxBidPrice *float64 `pulumi:"maxBidPrice"`
	// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name              *string                                              `pulumi:"name"`
	NetworkInterfaces []OrchestratedVirtualMachineScaleSetNetworkInterface `pulumi:"networkInterfaces"`
	OsDisk            *OrchestratedVirtualMachineScaleSetOsDisk            `pulumi:"osDisk"`
	OsProfile         *OrchestratedVirtualMachineScaleSetOsProfile         `pulumi:"osProfile"`
	Plan              *OrchestratedVirtualMachineScaleSetPlan              `pulumi:"plan"`
	// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	PlatformFaultDomainCount *int    `pulumi:"platformFaultDomainCount"`
	Priority                 *string `pulumi:"priority"`
	// The ID of the Proximity Placement Group which the Orchestrated Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId *string `pulumi:"proximityPlacementGroupId"`
	// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	SkuName           *string `pulumi:"skuName"`
	SourceImageId     *string `pulumi:"sourceImageId"`
	// A `sourceImageReference` block as defined below.
	SourceImageReference *OrchestratedVirtualMachineScaleSetSourceImageReference `pulumi:"sourceImageReference"`
	// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
	Tags                    map[string]string                                          `pulumi:"tags"`
	TerminationNotification *OrchestratedVirtualMachineScaleSetTerminationNotification `pulumi:"terminationNotification"`
	// The Unique ID for the Orchestrated Virtual Machine Scale Set.
	UniqueId    *string `pulumi:"uniqueId"`
	ZoneBalance *bool   `pulumi:"zoneBalance"`
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones *string `pulumi:"zones"`
}

type OrchestratedVirtualMachineScaleSetState struct {
	AutomaticInstanceRepair OrchestratedVirtualMachineScaleSetAutomaticInstanceRepairPtrInput
	BootDiagnostics         OrchestratedVirtualMachineScaleSetBootDiagnosticsPtrInput
	DataDisks               OrchestratedVirtualMachineScaleSetDataDiskArrayInput
	EncryptionAtHostEnabled pulumi.BoolPtrInput
	EvictionPolicy          pulumi.StringPtrInput
	Extensions              OrchestratedVirtualMachineScaleSetExtensionArrayInput
	// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M).
	ExtensionsTimeBudget pulumi.StringPtrInput
	Identity             OrchestratedVirtualMachineScaleSetIdentityPtrInput
	// The number of Virtual Machines in the Orcestrated Virtual Machine Scale Set.
	Instances   pulumi.IntPtrInput
	LicenseType pulumi.StringPtrInput
	// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location    pulumi.StringPtrInput
	MaxBidPrice pulumi.Float64PtrInput
	// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name              pulumi.StringPtrInput
	NetworkInterfaces OrchestratedVirtualMachineScaleSetNetworkInterfaceArrayInput
	OsDisk            OrchestratedVirtualMachineScaleSetOsDiskPtrInput
	OsProfile         OrchestratedVirtualMachineScaleSetOsProfilePtrInput
	Plan              OrchestratedVirtualMachineScaleSetPlanPtrInput
	// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	PlatformFaultDomainCount pulumi.IntPtrInput
	Priority                 pulumi.StringPtrInput
	// The ID of the Proximity Placement Group which the Orchestrated Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId pulumi.StringPtrInput
	// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	SkuName           pulumi.StringPtrInput
	SourceImageId     pulumi.StringPtrInput
	// A `sourceImageReference` block as defined below.
	SourceImageReference OrchestratedVirtualMachineScaleSetSourceImageReferencePtrInput
	// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
	Tags                    pulumi.StringMapInput
	TerminationNotification OrchestratedVirtualMachineScaleSetTerminationNotificationPtrInput
	// The Unique ID for the Orchestrated Virtual Machine Scale Set.
	UniqueId    pulumi.StringPtrInput
	ZoneBalance pulumi.BoolPtrInput
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones pulumi.StringPtrInput
}

func (OrchestratedVirtualMachineScaleSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*orchestratedVirtualMachineScaleSetState)(nil)).Elem()
}

type orchestratedVirtualMachineScaleSetArgs struct {
	AutomaticInstanceRepair *OrchestratedVirtualMachineScaleSetAutomaticInstanceRepair `pulumi:"automaticInstanceRepair"`
	BootDiagnostics         *OrchestratedVirtualMachineScaleSetBootDiagnostics         `pulumi:"bootDiagnostics"`
	DataDisks               []OrchestratedVirtualMachineScaleSetDataDisk               `pulumi:"dataDisks"`
	EncryptionAtHostEnabled *bool                                                      `pulumi:"encryptionAtHostEnabled"`
	EvictionPolicy          *string                                                    `pulumi:"evictionPolicy"`
	Extensions              []OrchestratedVirtualMachineScaleSetExtension              `pulumi:"extensions"`
	// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M).
	ExtensionsTimeBudget *string                                     `pulumi:"extensionsTimeBudget"`
	Identity             *OrchestratedVirtualMachineScaleSetIdentity `pulumi:"identity"`
	// The number of Virtual Machines in the Orcestrated Virtual Machine Scale Set.
	Instances   *int    `pulumi:"instances"`
	LicenseType *string `pulumi:"licenseType"`
	// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location    *string  `pulumi:"location"`
	MaxBidPrice *float64 `pulumi:"maxBidPrice"`
	// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name              *string                                              `pulumi:"name"`
	NetworkInterfaces []OrchestratedVirtualMachineScaleSetNetworkInterface `pulumi:"networkInterfaces"`
	OsDisk            *OrchestratedVirtualMachineScaleSetOsDisk            `pulumi:"osDisk"`
	OsProfile         *OrchestratedVirtualMachineScaleSetOsProfile         `pulumi:"osProfile"`
	Plan              *OrchestratedVirtualMachineScaleSetPlan              `pulumi:"plan"`
	// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	PlatformFaultDomainCount int     `pulumi:"platformFaultDomainCount"`
	Priority                 *string `pulumi:"priority"`
	// The ID of the Proximity Placement Group which the Orchestrated Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId *string `pulumi:"proximityPlacementGroupId"`
	// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SkuName           *string `pulumi:"skuName"`
	SourceImageId     *string `pulumi:"sourceImageId"`
	// A `sourceImageReference` block as defined below.
	SourceImageReference *OrchestratedVirtualMachineScaleSetSourceImageReference `pulumi:"sourceImageReference"`
	// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
	Tags                    map[string]string                                          `pulumi:"tags"`
	TerminationNotification *OrchestratedVirtualMachineScaleSetTerminationNotification `pulumi:"terminationNotification"`
	ZoneBalance             *bool                                                      `pulumi:"zoneBalance"`
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones *string `pulumi:"zones"`
}

// The set of arguments for constructing a OrchestratedVirtualMachineScaleSet resource.
type OrchestratedVirtualMachineScaleSetArgs struct {
	AutomaticInstanceRepair OrchestratedVirtualMachineScaleSetAutomaticInstanceRepairPtrInput
	BootDiagnostics         OrchestratedVirtualMachineScaleSetBootDiagnosticsPtrInput
	DataDisks               OrchestratedVirtualMachineScaleSetDataDiskArrayInput
	EncryptionAtHostEnabled pulumi.BoolPtrInput
	EvictionPolicy          pulumi.StringPtrInput
	Extensions              OrchestratedVirtualMachineScaleSetExtensionArrayInput
	// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M).
	ExtensionsTimeBudget pulumi.StringPtrInput
	Identity             OrchestratedVirtualMachineScaleSetIdentityPtrInput
	// The number of Virtual Machines in the Orcestrated Virtual Machine Scale Set.
	Instances   pulumi.IntPtrInput
	LicenseType pulumi.StringPtrInput
	// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location    pulumi.StringPtrInput
	MaxBidPrice pulumi.Float64PtrInput
	// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name              pulumi.StringPtrInput
	NetworkInterfaces OrchestratedVirtualMachineScaleSetNetworkInterfaceArrayInput
	OsDisk            OrchestratedVirtualMachineScaleSetOsDiskPtrInput
	OsProfile         OrchestratedVirtualMachineScaleSetOsProfilePtrInput
	Plan              OrchestratedVirtualMachineScaleSetPlanPtrInput
	// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	PlatformFaultDomainCount pulumi.IntInput
	Priority                 pulumi.StringPtrInput
	// The ID of the Proximity Placement Group which the Orchestrated Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId pulumi.StringPtrInput
	// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	SkuName           pulumi.StringPtrInput
	SourceImageId     pulumi.StringPtrInput
	// A `sourceImageReference` block as defined below.
	SourceImageReference OrchestratedVirtualMachineScaleSetSourceImageReferencePtrInput
	// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
	Tags                    pulumi.StringMapInput
	TerminationNotification OrchestratedVirtualMachineScaleSetTerminationNotificationPtrInput
	ZoneBalance             pulumi.BoolPtrInput
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones pulumi.StringPtrInput
}

func (OrchestratedVirtualMachineScaleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orchestratedVirtualMachineScaleSetArgs)(nil)).Elem()
}

type OrchestratedVirtualMachineScaleSetInput interface {
	pulumi.Input

	ToOrchestratedVirtualMachineScaleSetOutput() OrchestratedVirtualMachineScaleSetOutput
	ToOrchestratedVirtualMachineScaleSetOutputWithContext(ctx context.Context) OrchestratedVirtualMachineScaleSetOutput
}

func (*OrchestratedVirtualMachineScaleSet) ElementType() reflect.Type {
	return reflect.TypeOf((**OrchestratedVirtualMachineScaleSet)(nil)).Elem()
}

func (i *OrchestratedVirtualMachineScaleSet) ToOrchestratedVirtualMachineScaleSetOutput() OrchestratedVirtualMachineScaleSetOutput {
	return i.ToOrchestratedVirtualMachineScaleSetOutputWithContext(context.Background())
}

func (i *OrchestratedVirtualMachineScaleSet) ToOrchestratedVirtualMachineScaleSetOutputWithContext(ctx context.Context) OrchestratedVirtualMachineScaleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrchestratedVirtualMachineScaleSetOutput)
}

// OrchestratedVirtualMachineScaleSetArrayInput is an input type that accepts OrchestratedVirtualMachineScaleSetArray and OrchestratedVirtualMachineScaleSetArrayOutput values.
// You can construct a concrete instance of `OrchestratedVirtualMachineScaleSetArrayInput` via:
//
//          OrchestratedVirtualMachineScaleSetArray{ OrchestratedVirtualMachineScaleSetArgs{...} }
type OrchestratedVirtualMachineScaleSetArrayInput interface {
	pulumi.Input

	ToOrchestratedVirtualMachineScaleSetArrayOutput() OrchestratedVirtualMachineScaleSetArrayOutput
	ToOrchestratedVirtualMachineScaleSetArrayOutputWithContext(context.Context) OrchestratedVirtualMachineScaleSetArrayOutput
}

type OrchestratedVirtualMachineScaleSetArray []OrchestratedVirtualMachineScaleSetInput

func (OrchestratedVirtualMachineScaleSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrchestratedVirtualMachineScaleSet)(nil)).Elem()
}

func (i OrchestratedVirtualMachineScaleSetArray) ToOrchestratedVirtualMachineScaleSetArrayOutput() OrchestratedVirtualMachineScaleSetArrayOutput {
	return i.ToOrchestratedVirtualMachineScaleSetArrayOutputWithContext(context.Background())
}

func (i OrchestratedVirtualMachineScaleSetArray) ToOrchestratedVirtualMachineScaleSetArrayOutputWithContext(ctx context.Context) OrchestratedVirtualMachineScaleSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrchestratedVirtualMachineScaleSetArrayOutput)
}

// OrchestratedVirtualMachineScaleSetMapInput is an input type that accepts OrchestratedVirtualMachineScaleSetMap and OrchestratedVirtualMachineScaleSetMapOutput values.
// You can construct a concrete instance of `OrchestratedVirtualMachineScaleSetMapInput` via:
//
//          OrchestratedVirtualMachineScaleSetMap{ "key": OrchestratedVirtualMachineScaleSetArgs{...} }
type OrchestratedVirtualMachineScaleSetMapInput interface {
	pulumi.Input

	ToOrchestratedVirtualMachineScaleSetMapOutput() OrchestratedVirtualMachineScaleSetMapOutput
	ToOrchestratedVirtualMachineScaleSetMapOutputWithContext(context.Context) OrchestratedVirtualMachineScaleSetMapOutput
}

type OrchestratedVirtualMachineScaleSetMap map[string]OrchestratedVirtualMachineScaleSetInput

func (OrchestratedVirtualMachineScaleSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrchestratedVirtualMachineScaleSet)(nil)).Elem()
}

func (i OrchestratedVirtualMachineScaleSetMap) ToOrchestratedVirtualMachineScaleSetMapOutput() OrchestratedVirtualMachineScaleSetMapOutput {
	return i.ToOrchestratedVirtualMachineScaleSetMapOutputWithContext(context.Background())
}

func (i OrchestratedVirtualMachineScaleSetMap) ToOrchestratedVirtualMachineScaleSetMapOutputWithContext(ctx context.Context) OrchestratedVirtualMachineScaleSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrchestratedVirtualMachineScaleSetMapOutput)
}

type OrchestratedVirtualMachineScaleSetOutput struct{ *pulumi.OutputState }

func (OrchestratedVirtualMachineScaleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrchestratedVirtualMachineScaleSet)(nil)).Elem()
}

func (o OrchestratedVirtualMachineScaleSetOutput) ToOrchestratedVirtualMachineScaleSetOutput() OrchestratedVirtualMachineScaleSetOutput {
	return o
}

func (o OrchestratedVirtualMachineScaleSetOutput) ToOrchestratedVirtualMachineScaleSetOutputWithContext(ctx context.Context) OrchestratedVirtualMachineScaleSetOutput {
	return o
}

type OrchestratedVirtualMachineScaleSetArrayOutput struct{ *pulumi.OutputState }

func (OrchestratedVirtualMachineScaleSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrchestratedVirtualMachineScaleSet)(nil)).Elem()
}

func (o OrchestratedVirtualMachineScaleSetArrayOutput) ToOrchestratedVirtualMachineScaleSetArrayOutput() OrchestratedVirtualMachineScaleSetArrayOutput {
	return o
}

func (o OrchestratedVirtualMachineScaleSetArrayOutput) ToOrchestratedVirtualMachineScaleSetArrayOutputWithContext(ctx context.Context) OrchestratedVirtualMachineScaleSetArrayOutput {
	return o
}

func (o OrchestratedVirtualMachineScaleSetArrayOutput) Index(i pulumi.IntInput) OrchestratedVirtualMachineScaleSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrchestratedVirtualMachineScaleSet {
		return vs[0].([]*OrchestratedVirtualMachineScaleSet)[vs[1].(int)]
	}).(OrchestratedVirtualMachineScaleSetOutput)
}

type OrchestratedVirtualMachineScaleSetMapOutput struct{ *pulumi.OutputState }

func (OrchestratedVirtualMachineScaleSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrchestratedVirtualMachineScaleSet)(nil)).Elem()
}

func (o OrchestratedVirtualMachineScaleSetMapOutput) ToOrchestratedVirtualMachineScaleSetMapOutput() OrchestratedVirtualMachineScaleSetMapOutput {
	return o
}

func (o OrchestratedVirtualMachineScaleSetMapOutput) ToOrchestratedVirtualMachineScaleSetMapOutputWithContext(ctx context.Context) OrchestratedVirtualMachineScaleSetMapOutput {
	return o
}

func (o OrchestratedVirtualMachineScaleSetMapOutput) MapIndex(k pulumi.StringInput) OrchestratedVirtualMachineScaleSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrchestratedVirtualMachineScaleSet {
		return vs[0].(map[string]*OrchestratedVirtualMachineScaleSet)[vs[1].(string)]
	}).(OrchestratedVirtualMachineScaleSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrchestratedVirtualMachineScaleSetInput)(nil)).Elem(), &OrchestratedVirtualMachineScaleSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrchestratedVirtualMachineScaleSetArrayInput)(nil)).Elem(), OrchestratedVirtualMachineScaleSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrchestratedVirtualMachineScaleSetMapInput)(nil)).Elem(), OrchestratedVirtualMachineScaleSetMap{})
	pulumi.RegisterOutputType(OrchestratedVirtualMachineScaleSetOutput{})
	pulumi.RegisterOutputType(OrchestratedVirtualMachineScaleSetArrayOutput{})
	pulumi.RegisterOutputType(OrchestratedVirtualMachineScaleSetMapOutput{})
}
