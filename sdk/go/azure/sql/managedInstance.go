// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a SQL Azure Managed Instance.
//
// > **Note:** All arguments including the administrator login and password will be stored in the raw state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/sql"
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
// 		exampleNetworkSecurityGroup, err := network.NewNetworkSecurityGroup(ctx, "exampleNetworkSecurityGroup", &network.NetworkSecurityGroupArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewNetworkSecurityRule(ctx, "allowManagementInbound", &network.NetworkSecurityRuleArgs{
// 			Priority:        pulumi.Int(106),
// 			Direction:       pulumi.String("Inbound"),
// 			Access:          pulumi.String("Allow"),
// 			Protocol:        pulumi.String("Tcp"),
// 			SourcePortRange: pulumi.String("*"),
// 			DestinationPortRanges: pulumi.StringArray{
// 				pulumi.String("9000"),
// 				pulumi.String("9003"),
// 				pulumi.String("1438"),
// 				pulumi.String("1440"),
// 				pulumi.String("1452"),
// 			},
// 			SourceAddressPrefix:      pulumi.String("*"),
// 			DestinationAddressPrefix: pulumi.String("*"),
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			NetworkSecurityGroupName: exampleNetworkSecurityGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewNetworkSecurityRule(ctx, "allowMisubnetInbound", &network.NetworkSecurityRuleArgs{
// 			Priority:                 pulumi.Int(200),
// 			Direction:                pulumi.String("Inbound"),
// 			Access:                   pulumi.String("Allow"),
// 			Protocol:                 pulumi.String("*"),
// 			SourcePortRange:          pulumi.String("*"),
// 			DestinationPortRange:     pulumi.String("*"),
// 			SourceAddressPrefix:      pulumi.String("10.0.0.0/24"),
// 			DestinationAddressPrefix: pulumi.String("*"),
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			NetworkSecurityGroupName: exampleNetworkSecurityGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewNetworkSecurityRule(ctx, "allowHealthProbeInbound", &network.NetworkSecurityRuleArgs{
// 			Priority:                 pulumi.Int(300),
// 			Direction:                pulumi.String("Inbound"),
// 			Access:                   pulumi.String("Allow"),
// 			Protocol:                 pulumi.String("*"),
// 			SourcePortRange:          pulumi.String("*"),
// 			DestinationPortRange:     pulumi.String("*"),
// 			SourceAddressPrefix:      pulumi.String("AzureLoadBalancer"),
// 			DestinationAddressPrefix: pulumi.String("*"),
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			NetworkSecurityGroupName: exampleNetworkSecurityGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewNetworkSecurityRule(ctx, "allowTdsInbound", &network.NetworkSecurityRuleArgs{
// 			Priority:                 pulumi.Int(1000),
// 			Direction:                pulumi.String("Inbound"),
// 			Access:                   pulumi.String("Allow"),
// 			Protocol:                 pulumi.String("Tcp"),
// 			SourcePortRange:          pulumi.String("*"),
// 			DestinationPortRange:     pulumi.String("1433"),
// 			SourceAddressPrefix:      pulumi.String("VirtualNetwork"),
// 			DestinationAddressPrefix: pulumi.String("*"),
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			NetworkSecurityGroupName: exampleNetworkSecurityGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewNetworkSecurityRule(ctx, "denyAllInbound", &network.NetworkSecurityRuleArgs{
// 			Priority:                 pulumi.Int(4096),
// 			Direction:                pulumi.String("Inbound"),
// 			Access:                   pulumi.String("Deny"),
// 			Protocol:                 pulumi.String("*"),
// 			SourcePortRange:          pulumi.String("*"),
// 			DestinationPortRange:     pulumi.String("*"),
// 			SourceAddressPrefix:      pulumi.String("*"),
// 			DestinationAddressPrefix: pulumi.String("*"),
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			NetworkSecurityGroupName: exampleNetworkSecurityGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewNetworkSecurityRule(ctx, "allowManagementOutbound", &network.NetworkSecurityRuleArgs{
// 			Priority:        pulumi.Int(102),
// 			Direction:       pulumi.String("Outbound"),
// 			Access:          pulumi.String("Allow"),
// 			Protocol:        pulumi.String("Tcp"),
// 			SourcePortRange: pulumi.String("*"),
// 			DestinationPortRanges: pulumi.StringArray{
// 				pulumi.String("80"),
// 				pulumi.String("443"),
// 				pulumi.String("12000"),
// 			},
// 			SourceAddressPrefix:      pulumi.String("*"),
// 			DestinationAddressPrefix: pulumi.String("*"),
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			NetworkSecurityGroupName: exampleNetworkSecurityGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewNetworkSecurityRule(ctx, "allowMisubnetOutbound", &network.NetworkSecurityRuleArgs{
// 			Priority:                 pulumi.Int(200),
// 			Direction:                pulumi.String("Outbound"),
// 			Access:                   pulumi.String("Allow"),
// 			Protocol:                 pulumi.String("*"),
// 			SourcePortRange:          pulumi.String("*"),
// 			DestinationPortRange:     pulumi.String("*"),
// 			SourceAddressPrefix:      pulumi.String("10.0.0.0/24"),
// 			DestinationAddressPrefix: pulumi.String("*"),
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			NetworkSecurityGroupName: exampleNetworkSecurityGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewNetworkSecurityRule(ctx, "denyAllOutbound", &network.NetworkSecurityRuleArgs{
// 			Priority:                 pulumi.Int(4096),
// 			Direction:                pulumi.String("Outbound"),
// 			Access:                   pulumi.String("Deny"),
// 			Protocol:                 pulumi.String("*"),
// 			SourcePortRange:          pulumi.String("*"),
// 			DestinationPortRange:     pulumi.String("*"),
// 			SourceAddressPrefix:      pulumi.String("*"),
// 			DestinationAddressPrefix: pulumi.String("*"),
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			NetworkSecurityGroupName: exampleNetworkSecurityGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 			Location: exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnet, err := network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefix:      pulumi.String("10.0.0.0/24"),
// 			Delegations: network.SubnetDelegationArray{
// 				&network.SubnetDelegationArgs{
// 					Name: pulumi.String("managedinstancedelegation"),
// 					ServiceDelegation: &network.SubnetDelegationServiceDelegationArgs{
// 						Name: pulumi.String("Microsoft.Sql/managedInstances"),
// 						Actions: pulumi.StringArray{
// 							pulumi.String("Microsoft.Network/virtualNetworks/subnets/join/action"),
// 							pulumi.String("Microsoft.Network/virtualNetworks/subnets/prepareNetworkPolicies/action"),
// 							pulumi.String("Microsoft.Network/virtualNetworks/subnets/unprepareNetworkPolicies/action"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnetNetworkSecurityGroupAssociation, err := network.NewSubnetNetworkSecurityGroupAssociation(ctx, "exampleSubnetNetworkSecurityGroupAssociation", &network.SubnetNetworkSecurityGroupAssociationArgs{
// 			SubnetId:               exampleSubnet.ID(),
// 			NetworkSecurityGroupId: exampleNetworkSecurityGroup.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleRouteTable, err := network.NewRouteTable(ctx, "exampleRouteTable", &network.RouteTableArgs{
// 			Location:                   exampleResourceGroup.Location,
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			DisableBgpRoutePropagation: pulumi.Bool(false),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleSubnet,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnetRouteTableAssociation, err := network.NewSubnetRouteTableAssociation(ctx, "exampleSubnetRouteTableAssociation", &network.SubnetRouteTableAssociationArgs{
// 			SubnetId:     exampleSubnet.ID(),
// 			RouteTableId: exampleRouteTable.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sql.NewManagedInstance(ctx, "exampleManagedInstance", &sql.ManagedInstanceArgs{
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			Location:                   exampleResourceGroup.Location,
// 			AdministratorLogin:         pulumi.String("mradministrator"),
// 			AdministratorLoginPassword: pulumi.String("thisIsDog11"),
// 			LicenseType:                pulumi.String("BasePrice"),
// 			SubnetId:                   exampleSubnet.ID(),
// 			SkuName:                    pulumi.String("GP_Gen5"),
// 			Vcores:                     pulumi.Int(4),
// 			StorageSizeInGb:            pulumi.Int(32),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleSubnetNetworkSecurityGroupAssociation,
// 			exampleSubnetRouteTableAssociation,
// 		}))
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
// SQL Servers can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:sql/managedInstance:ManagedInstance example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/managedInstances/myserver
// ```
type ManagedInstance struct {
	pulumi.CustomResourceState

	// The administrator login name for the new server. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringOutput `pulumi:"administratorLogin"`
	// The password associated with the `administratorLogin` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
	AdministratorLoginPassword pulumi.StringOutput `pulumi:"administratorLoginPassword"`
	// Specifies how the SQL Managed Instance will be collated. Default value is `SQL_Latin1_General_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation pulumi.StringPtrOutput `pulumi:"collation"`
	// The ID of the Managed Instance which will share the DNS zone. This is a prerequisite for creating a failover group, although creation of a failover group is not yet possible in `azurerm`. Setting this after creation forces a new resource to be created.
	DnsZonePartnerId pulumi.StringPtrOutput `pulumi:"dnsZonePartnerId"`
	// The fully qualified domain name of the Azure Managed SQL Instance
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// An `identity` block as defined below.
	Identity ManagedInstanceIdentityPtrOutput `pulumi:"identity"`
	// What type of license the Managed Instance will use. Valid values include can be `PriceIncluded` or `BasePrice`.
	LicenseType pulumi.StringOutput `pulumi:"licenseType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The Minimum TLS Version. Default value is `1.2` Valid values include `1.0`, `1.1`, `1.2`.
	MinimumTlsVersion pulumi.StringPtrOutput `pulumi:"minimumTlsVersion"`
	// The name of the SQL Managed Instance. This needs to be globally unique within Azure.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies how the SQL Managed Instance will be accessed. Default value is `Default`. Valid values include `Default`, `Proxy`, and `Redirect`.
	ProxyOverride pulumi.StringPtrOutput `pulumi:"proxyOverride"`
	// Is the public data endpoint enabled? Default value is `false`.
	PublicDataEndpointEnabled pulumi.BoolPtrOutput `pulumi:"publicDataEndpointEnabled"`
	// The name of the resource group in which to create the SQL Server.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the SKU Name for the SQL Managed Instance. Valid values include `GP_Gen4`, `GP_Gen5`, `BC_Gen4`, `BC_Gen5`. Changing this forces a new resource to be created.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// Maximum storage space for your instance. It should be a multiple of 32GB.
	StorageSizeInGb pulumi.IntOutput `pulumi:"storageSizeInGb"`
	// The subnet resource id that the SQL Managed Instance will be associated with.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The TimeZone ID that the SQL Managed Instance will be operating in. Default value is `UTC`. Changing this forces a new resource to be created.
	TimezoneId pulumi.StringPtrOutput `pulumi:"timezoneId"`
	// Number of cores that should be assigned to your instance. Values can be `8`, `16`, or `24` if `skuName` is `GP_Gen4`, or `8`, `16`, `24`, `32`, or `40` if `skuName` is `GP_Gen5`.
	Vcores pulumi.IntOutput `pulumi:"vcores"`
}

// NewManagedInstance registers a new resource with the given unique name, arguments, and options.
func NewManagedInstance(ctx *pulumi.Context,
	name string, args *ManagedInstanceArgs, opts ...pulumi.ResourceOption) (*ManagedInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdministratorLogin == nil {
		return nil, errors.New("invalid value for required argument 'AdministratorLogin'")
	}
	if args.AdministratorLoginPassword == nil {
		return nil, errors.New("invalid value for required argument 'AdministratorLoginPassword'")
	}
	if args.LicenseType == nil {
		return nil, errors.New("invalid value for required argument 'LicenseType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SkuName == nil {
		return nil, errors.New("invalid value for required argument 'SkuName'")
	}
	if args.StorageSizeInGb == nil {
		return nil, errors.New("invalid value for required argument 'StorageSizeInGb'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.Vcores == nil {
		return nil, errors.New("invalid value for required argument 'Vcores'")
	}
	var resource ManagedInstance
	err := ctx.RegisterResource("azure:sql/managedInstance:ManagedInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedInstance gets an existing ManagedInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedInstanceState, opts ...pulumi.ResourceOption) (*ManagedInstance, error) {
	var resource ManagedInstance
	err := ctx.ReadResource("azure:sql/managedInstance:ManagedInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedInstance resources.
type managedInstanceState struct {
	// The administrator login name for the new server. Changing this forces a new resource to be created.
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The password associated with the `administratorLogin` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
	AdministratorLoginPassword *string `pulumi:"administratorLoginPassword"`
	// Specifies how the SQL Managed Instance will be collated. Default value is `SQL_Latin1_General_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation *string `pulumi:"collation"`
	// The ID of the Managed Instance which will share the DNS zone. This is a prerequisite for creating a failover group, although creation of a failover group is not yet possible in `azurerm`. Setting this after creation forces a new resource to be created.
	DnsZonePartnerId *string `pulumi:"dnsZonePartnerId"`
	// The fully qualified domain name of the Azure Managed SQL Instance
	Fqdn *string `pulumi:"fqdn"`
	// An `identity` block as defined below.
	Identity *ManagedInstanceIdentity `pulumi:"identity"`
	// What type of license the Managed Instance will use. Valid values include can be `PriceIncluded` or `BasePrice`.
	LicenseType *string `pulumi:"licenseType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The Minimum TLS Version. Default value is `1.2` Valid values include `1.0`, `1.1`, `1.2`.
	MinimumTlsVersion *string `pulumi:"minimumTlsVersion"`
	// The name of the SQL Managed Instance. This needs to be globally unique within Azure.
	Name *string `pulumi:"name"`
	// Specifies how the SQL Managed Instance will be accessed. Default value is `Default`. Valid values include `Default`, `Proxy`, and `Redirect`.
	ProxyOverride *string `pulumi:"proxyOverride"`
	// Is the public data endpoint enabled? Default value is `false`.
	PublicDataEndpointEnabled *bool `pulumi:"publicDataEndpointEnabled"`
	// The name of the resource group in which to create the SQL Server.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the SKU Name for the SQL Managed Instance. Valid values include `GP_Gen4`, `GP_Gen5`, `BC_Gen4`, `BC_Gen5`. Changing this forces a new resource to be created.
	SkuName *string `pulumi:"skuName"`
	// Maximum storage space for your instance. It should be a multiple of 32GB.
	StorageSizeInGb *int `pulumi:"storageSizeInGb"`
	// The subnet resource id that the SQL Managed Instance will be associated with.
	SubnetId *string `pulumi:"subnetId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The TimeZone ID that the SQL Managed Instance will be operating in. Default value is `UTC`. Changing this forces a new resource to be created.
	TimezoneId *string `pulumi:"timezoneId"`
	// Number of cores that should be assigned to your instance. Values can be `8`, `16`, or `24` if `skuName` is `GP_Gen4`, or `8`, `16`, `24`, `32`, or `40` if `skuName` is `GP_Gen5`.
	Vcores *int `pulumi:"vcores"`
}

type ManagedInstanceState struct {
	// The administrator login name for the new server. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringPtrInput
	// The password associated with the `administratorLogin` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
	AdministratorLoginPassword pulumi.StringPtrInput
	// Specifies how the SQL Managed Instance will be collated. Default value is `SQL_Latin1_General_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation pulumi.StringPtrInput
	// The ID of the Managed Instance which will share the DNS zone. This is a prerequisite for creating a failover group, although creation of a failover group is not yet possible in `azurerm`. Setting this after creation forces a new resource to be created.
	DnsZonePartnerId pulumi.StringPtrInput
	// The fully qualified domain name of the Azure Managed SQL Instance
	Fqdn pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity ManagedInstanceIdentityPtrInput
	// What type of license the Managed Instance will use. Valid values include can be `PriceIncluded` or `BasePrice`.
	LicenseType pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The Minimum TLS Version. Default value is `1.2` Valid values include `1.0`, `1.1`, `1.2`.
	MinimumTlsVersion pulumi.StringPtrInput
	// The name of the SQL Managed Instance. This needs to be globally unique within Azure.
	Name pulumi.StringPtrInput
	// Specifies how the SQL Managed Instance will be accessed. Default value is `Default`. Valid values include `Default`, `Proxy`, and `Redirect`.
	ProxyOverride pulumi.StringPtrInput
	// Is the public data endpoint enabled? Default value is `false`.
	PublicDataEndpointEnabled pulumi.BoolPtrInput
	// The name of the resource group in which to create the SQL Server.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the SKU Name for the SQL Managed Instance. Valid values include `GP_Gen4`, `GP_Gen5`, `BC_Gen4`, `BC_Gen5`. Changing this forces a new resource to be created.
	SkuName pulumi.StringPtrInput
	// Maximum storage space for your instance. It should be a multiple of 32GB.
	StorageSizeInGb pulumi.IntPtrInput
	// The subnet resource id that the SQL Managed Instance will be associated with.
	SubnetId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The TimeZone ID that the SQL Managed Instance will be operating in. Default value is `UTC`. Changing this forces a new resource to be created.
	TimezoneId pulumi.StringPtrInput
	// Number of cores that should be assigned to your instance. Values can be `8`, `16`, or `24` if `skuName` is `GP_Gen4`, or `8`, `16`, `24`, `32`, or `40` if `skuName` is `GP_Gen5`.
	Vcores pulumi.IntPtrInput
}

func (ManagedInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceState)(nil)).Elem()
}

type managedInstanceArgs struct {
	// The administrator login name for the new server. Changing this forces a new resource to be created.
	AdministratorLogin string `pulumi:"administratorLogin"`
	// The password associated with the `administratorLogin` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
	AdministratorLoginPassword string `pulumi:"administratorLoginPassword"`
	// Specifies how the SQL Managed Instance will be collated. Default value is `SQL_Latin1_General_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation *string `pulumi:"collation"`
	// The ID of the Managed Instance which will share the DNS zone. This is a prerequisite for creating a failover group, although creation of a failover group is not yet possible in `azurerm`. Setting this after creation forces a new resource to be created.
	DnsZonePartnerId *string `pulumi:"dnsZonePartnerId"`
	// An `identity` block as defined below.
	Identity *ManagedInstanceIdentity `pulumi:"identity"`
	// What type of license the Managed Instance will use. Valid values include can be `PriceIncluded` or `BasePrice`.
	LicenseType string `pulumi:"licenseType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The Minimum TLS Version. Default value is `1.2` Valid values include `1.0`, `1.1`, `1.2`.
	MinimumTlsVersion *string `pulumi:"minimumTlsVersion"`
	// The name of the SQL Managed Instance. This needs to be globally unique within Azure.
	Name *string `pulumi:"name"`
	// Specifies how the SQL Managed Instance will be accessed. Default value is `Default`. Valid values include `Default`, `Proxy`, and `Redirect`.
	ProxyOverride *string `pulumi:"proxyOverride"`
	// Is the public data endpoint enabled? Default value is `false`.
	PublicDataEndpointEnabled *bool `pulumi:"publicDataEndpointEnabled"`
	// The name of the resource group in which to create the SQL Server.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the SKU Name for the SQL Managed Instance. Valid values include `GP_Gen4`, `GP_Gen5`, `BC_Gen4`, `BC_Gen5`. Changing this forces a new resource to be created.
	SkuName string `pulumi:"skuName"`
	// Maximum storage space for your instance. It should be a multiple of 32GB.
	StorageSizeInGb int `pulumi:"storageSizeInGb"`
	// The subnet resource id that the SQL Managed Instance will be associated with.
	SubnetId string `pulumi:"subnetId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The TimeZone ID that the SQL Managed Instance will be operating in. Default value is `UTC`. Changing this forces a new resource to be created.
	TimezoneId *string `pulumi:"timezoneId"`
	// Number of cores that should be assigned to your instance. Values can be `8`, `16`, or `24` if `skuName` is `GP_Gen4`, or `8`, `16`, `24`, `32`, or `40` if `skuName` is `GP_Gen5`.
	Vcores int `pulumi:"vcores"`
}

// The set of arguments for constructing a ManagedInstance resource.
type ManagedInstanceArgs struct {
	// The administrator login name for the new server. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringInput
	// The password associated with the `administratorLogin` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
	AdministratorLoginPassword pulumi.StringInput
	// Specifies how the SQL Managed Instance will be collated. Default value is `SQL_Latin1_General_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation pulumi.StringPtrInput
	// The ID of the Managed Instance which will share the DNS zone. This is a prerequisite for creating a failover group, although creation of a failover group is not yet possible in `azurerm`. Setting this after creation forces a new resource to be created.
	DnsZonePartnerId pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity ManagedInstanceIdentityPtrInput
	// What type of license the Managed Instance will use. Valid values include can be `PriceIncluded` or `BasePrice`.
	LicenseType pulumi.StringInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The Minimum TLS Version. Default value is `1.2` Valid values include `1.0`, `1.1`, `1.2`.
	MinimumTlsVersion pulumi.StringPtrInput
	// The name of the SQL Managed Instance. This needs to be globally unique within Azure.
	Name pulumi.StringPtrInput
	// Specifies how the SQL Managed Instance will be accessed. Default value is `Default`. Valid values include `Default`, `Proxy`, and `Redirect`.
	ProxyOverride pulumi.StringPtrInput
	// Is the public data endpoint enabled? Default value is `false`.
	PublicDataEndpointEnabled pulumi.BoolPtrInput
	// The name of the resource group in which to create the SQL Server.
	ResourceGroupName pulumi.StringInput
	// Specifies the SKU Name for the SQL Managed Instance. Valid values include `GP_Gen4`, `GP_Gen5`, `BC_Gen4`, `BC_Gen5`. Changing this forces a new resource to be created.
	SkuName pulumi.StringInput
	// Maximum storage space for your instance. It should be a multiple of 32GB.
	StorageSizeInGb pulumi.IntInput
	// The subnet resource id that the SQL Managed Instance will be associated with.
	SubnetId pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The TimeZone ID that the SQL Managed Instance will be operating in. Default value is `UTC`. Changing this forces a new resource to be created.
	TimezoneId pulumi.StringPtrInput
	// Number of cores that should be assigned to your instance. Values can be `8`, `16`, or `24` if `skuName` is `GP_Gen4`, or `8`, `16`, `24`, `32`, or `40` if `skuName` is `GP_Gen5`.
	Vcores pulumi.IntInput
}

func (ManagedInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceArgs)(nil)).Elem()
}

type ManagedInstanceInput interface {
	pulumi.Input

	ToManagedInstanceOutput() ManagedInstanceOutput
	ToManagedInstanceOutputWithContext(ctx context.Context) ManagedInstanceOutput
}

func (*ManagedInstance) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstance)(nil))
}

func (i *ManagedInstance) ToManagedInstanceOutput() ManagedInstanceOutput {
	return i.ToManagedInstanceOutputWithContext(context.Background())
}

func (i *ManagedInstance) ToManagedInstanceOutputWithContext(ctx context.Context) ManagedInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceOutput)
}

func (i *ManagedInstance) ToManagedInstancePtrOutput() ManagedInstancePtrOutput {
	return i.ToManagedInstancePtrOutputWithContext(context.Background())
}

func (i *ManagedInstance) ToManagedInstancePtrOutputWithContext(ctx context.Context) ManagedInstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePtrOutput)
}

type ManagedInstancePtrInput interface {
	pulumi.Input

	ToManagedInstancePtrOutput() ManagedInstancePtrOutput
	ToManagedInstancePtrOutputWithContext(ctx context.Context) ManagedInstancePtrOutput
}

type managedInstancePtrType ManagedInstanceArgs

func (*managedInstancePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstance)(nil))
}

func (i *managedInstancePtrType) ToManagedInstancePtrOutput() ManagedInstancePtrOutput {
	return i.ToManagedInstancePtrOutputWithContext(context.Background())
}

func (i *managedInstancePtrType) ToManagedInstancePtrOutputWithContext(ctx context.Context) ManagedInstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePtrOutput)
}

// ManagedInstanceArrayInput is an input type that accepts ManagedInstanceArray and ManagedInstanceArrayOutput values.
// You can construct a concrete instance of `ManagedInstanceArrayInput` via:
//
//          ManagedInstanceArray{ ManagedInstanceArgs{...} }
type ManagedInstanceArrayInput interface {
	pulumi.Input

	ToManagedInstanceArrayOutput() ManagedInstanceArrayOutput
	ToManagedInstanceArrayOutputWithContext(context.Context) ManagedInstanceArrayOutput
}

type ManagedInstanceArray []ManagedInstanceInput

func (ManagedInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedInstance)(nil)).Elem()
}

func (i ManagedInstanceArray) ToManagedInstanceArrayOutput() ManagedInstanceArrayOutput {
	return i.ToManagedInstanceArrayOutputWithContext(context.Background())
}

func (i ManagedInstanceArray) ToManagedInstanceArrayOutputWithContext(ctx context.Context) ManagedInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceArrayOutput)
}

// ManagedInstanceMapInput is an input type that accepts ManagedInstanceMap and ManagedInstanceMapOutput values.
// You can construct a concrete instance of `ManagedInstanceMapInput` via:
//
//          ManagedInstanceMap{ "key": ManagedInstanceArgs{...} }
type ManagedInstanceMapInput interface {
	pulumi.Input

	ToManagedInstanceMapOutput() ManagedInstanceMapOutput
	ToManagedInstanceMapOutputWithContext(context.Context) ManagedInstanceMapOutput
}

type ManagedInstanceMap map[string]ManagedInstanceInput

func (ManagedInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedInstance)(nil)).Elem()
}

func (i ManagedInstanceMap) ToManagedInstanceMapOutput() ManagedInstanceMapOutput {
	return i.ToManagedInstanceMapOutputWithContext(context.Background())
}

func (i ManagedInstanceMap) ToManagedInstanceMapOutputWithContext(ctx context.Context) ManagedInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceMapOutput)
}

type ManagedInstanceOutput struct{ *pulumi.OutputState }

func (ManagedInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstance)(nil))
}

func (o ManagedInstanceOutput) ToManagedInstanceOutput() ManagedInstanceOutput {
	return o
}

func (o ManagedInstanceOutput) ToManagedInstanceOutputWithContext(ctx context.Context) ManagedInstanceOutput {
	return o
}

func (o ManagedInstanceOutput) ToManagedInstancePtrOutput() ManagedInstancePtrOutput {
	return o.ToManagedInstancePtrOutputWithContext(context.Background())
}

func (o ManagedInstanceOutput) ToManagedInstancePtrOutputWithContext(ctx context.Context) ManagedInstancePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedInstance) *ManagedInstance {
		return &v
	}).(ManagedInstancePtrOutput)
}

type ManagedInstancePtrOutput struct{ *pulumi.OutputState }

func (ManagedInstancePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstance)(nil))
}

func (o ManagedInstancePtrOutput) ToManagedInstancePtrOutput() ManagedInstancePtrOutput {
	return o
}

func (o ManagedInstancePtrOutput) ToManagedInstancePtrOutputWithContext(ctx context.Context) ManagedInstancePtrOutput {
	return o
}

func (o ManagedInstancePtrOutput) Elem() ManagedInstanceOutput {
	return o.ApplyT(func(v *ManagedInstance) ManagedInstance {
		if v != nil {
			return *v
		}
		var ret ManagedInstance
		return ret
	}).(ManagedInstanceOutput)
}

type ManagedInstanceArrayOutput struct{ *pulumi.OutputState }

func (ManagedInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedInstance)(nil))
}

func (o ManagedInstanceArrayOutput) ToManagedInstanceArrayOutput() ManagedInstanceArrayOutput {
	return o
}

func (o ManagedInstanceArrayOutput) ToManagedInstanceArrayOutputWithContext(ctx context.Context) ManagedInstanceArrayOutput {
	return o
}

func (o ManagedInstanceArrayOutput) Index(i pulumi.IntInput) ManagedInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedInstance {
		return vs[0].([]ManagedInstance)[vs[1].(int)]
	}).(ManagedInstanceOutput)
}

type ManagedInstanceMapOutput struct{ *pulumi.OutputState }

func (ManagedInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedInstance)(nil))
}

func (o ManagedInstanceMapOutput) ToManagedInstanceMapOutput() ManagedInstanceMapOutput {
	return o
}

func (o ManagedInstanceMapOutput) ToManagedInstanceMapOutputWithContext(ctx context.Context) ManagedInstanceMapOutput {
	return o
}

func (o ManagedInstanceMapOutput) MapIndex(k pulumi.StringInput) ManagedInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ManagedInstance {
		return vs[0].(map[string]ManagedInstance)[vs[1].(string)]
	}).(ManagedInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstanceInput)(nil)).Elem(), &ManagedInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePtrInput)(nil)).Elem(), &ManagedInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstanceArrayInput)(nil)).Elem(), ManagedInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstanceMapInput)(nil)).Elem(), ManagedInstanceMap{})
	pulumi.RegisterOutputType(ManagedInstanceOutput{})
	pulumi.RegisterOutputType(ManagedInstancePtrOutput{})
	pulumi.RegisterOutputType(ManagedInstanceArrayOutput{})
	pulumi.RegisterOutputType(ManagedInstanceMapOutput{})
}
