// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Azure NAT Gateway.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const examplePublicIp = new azure.network.PublicIp("examplePublicIp", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     allocationMethod: "Static",
 *     sku: "Standard",
 *     zones: ["1"],
 * });
 * const examplePublicIpPrefix = new azure.network.PublicIpPrefix("examplePublicIpPrefix", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     prefixLength: 30,
 *     zones: ["1"],
 * });
 * const exampleNatGateway = new azure.network.NatGateway("exampleNatGateway", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     publicIpAddressIds: [examplePublicIp.id],
 *     publicIpPrefixIds: [examplePublicIpPrefix.id],
 *     skuName: "Standard",
 *     idleTimeoutInMinutes: 10,
 *     zones: ["1"],
 * });
 * ```
 *
 * ## Import
 *
 * NAT Gateway can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:network/natGateway:NatGateway test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/natGateways/gateway1
 * ```
 */
export class NatGateway extends pulumi.CustomResource {
    /**
     * Get an existing NatGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NatGatewayState, opts?: pulumi.CustomResourceOptions): NatGateway {
        return new NatGateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/natGateway:NatGateway';

    /**
     * Returns true if the given object is an instance of NatGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NatGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NatGateway.__pulumiType;
    }

    /**
     * The idle timeout which should be used in minutes. Defaults to `4`.
     */
    public readonly idleTimeoutInMinutes!: pulumi.Output<number | undefined>;
    /**
     * Specifies the supported Azure location where the NAT Gateway should exist. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the NAT Gateway. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of Public IP Address ID's which should be associated with the NAT Gateway resource.
     *
     * @deprecated Inline Public IP Address ID Associations have been deprecated in favour of the `azurerm_nat_gateway_public_ip_association` resource. This field will be removed in the next major version of the Azure Provider.
     */
    public readonly publicIpAddressIds!: pulumi.Output<string[]>;
    /**
     * / **Deprecated in favour of `azure.network.NatGatewayPublicIpPrefixAssociation`**) A list of Public IP Prefix ID's which should be associated with the NAT Gateway resource.
     *
     * @deprecated Inline Public IP Prefix ID Associations have been deprecated in favour of the `azurerm_nat_gateway_public_ip_prefix_association` resource. This field will be removed in the next major version of the Azure Provider.
     */
    public readonly publicIpPrefixIds!: pulumi.Output<string[]>;
    /**
     * Specifies the name of the Resource Group in which the NAT Gateway should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The resource GUID property of the NAT Gateway.
     */
    public /*out*/ readonly resourceGuid!: pulumi.Output<string>;
    /**
     * The SKU which should be used. At this time the only supported value is `Standard`. Defaults to `Standard`.
     */
    public readonly skuName!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A list of availability zones where the NAT Gateway should be provisioned. Changing this forces a new resource to be created.
     */
    public readonly zones!: pulumi.Output<string[] | undefined>;

    /**
     * Create a NatGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NatGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NatGatewayArgs | NatGatewayState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NatGatewayState | undefined;
            resourceInputs["idleTimeoutInMinutes"] = state ? state.idleTimeoutInMinutes : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["publicIpAddressIds"] = state ? state.publicIpAddressIds : undefined;
            resourceInputs["publicIpPrefixIds"] = state ? state.publicIpPrefixIds : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["resourceGuid"] = state ? state.resourceGuid : undefined;
            resourceInputs["skuName"] = state ? state.skuName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["zones"] = state ? state.zones : undefined;
        } else {
            const args = argsOrState as NatGatewayArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["idleTimeoutInMinutes"] = args ? args.idleTimeoutInMinutes : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["publicIpAddressIds"] = args ? args.publicIpAddressIds : undefined;
            resourceInputs["publicIpPrefixIds"] = args ? args.publicIpPrefixIds : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["skuName"] = args ? args.skuName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["zones"] = args ? args.zones : undefined;
            resourceInputs["resourceGuid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NatGateway.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NatGateway resources.
 */
export interface NatGatewayState {
    /**
     * The idle timeout which should be used in minutes. Defaults to `4`.
     */
    idleTimeoutInMinutes?: pulumi.Input<number>;
    /**
     * Specifies the supported Azure location where the NAT Gateway should exist. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the name of the NAT Gateway. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of Public IP Address ID's which should be associated with the NAT Gateway resource.
     *
     * @deprecated Inline Public IP Address ID Associations have been deprecated in favour of the `azurerm_nat_gateway_public_ip_association` resource. This field will be removed in the next major version of the Azure Provider.
     */
    publicIpAddressIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * / **Deprecated in favour of `azure.network.NatGatewayPublicIpPrefixAssociation`**) A list of Public IP Prefix ID's which should be associated with the NAT Gateway resource.
     *
     * @deprecated Inline Public IP Prefix ID Associations have been deprecated in favour of the `azurerm_nat_gateway_public_ip_prefix_association` resource. This field will be removed in the next major version of the Azure Provider.
     */
    publicIpPrefixIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the name of the Resource Group in which the NAT Gateway should exist. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The resource GUID property of the NAT Gateway.
     */
    resourceGuid?: pulumi.Input<string>;
    /**
     * The SKU which should be used. At this time the only supported value is `Standard`. Defaults to `Standard`.
     */
    skuName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of availability zones where the NAT Gateway should be provisioned. Changing this forces a new resource to be created.
     */
    zones?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a NatGateway resource.
 */
export interface NatGatewayArgs {
    /**
     * The idle timeout which should be used in minutes. Defaults to `4`.
     */
    idleTimeoutInMinutes?: pulumi.Input<number>;
    /**
     * Specifies the supported Azure location where the NAT Gateway should exist. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the name of the NAT Gateway. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of Public IP Address ID's which should be associated with the NAT Gateway resource.
     *
     * @deprecated Inline Public IP Address ID Associations have been deprecated in favour of the `azurerm_nat_gateway_public_ip_association` resource. This field will be removed in the next major version of the Azure Provider.
     */
    publicIpAddressIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * / **Deprecated in favour of `azure.network.NatGatewayPublicIpPrefixAssociation`**) A list of Public IP Prefix ID's which should be associated with the NAT Gateway resource.
     *
     * @deprecated Inline Public IP Prefix ID Associations have been deprecated in favour of the `azurerm_nat_gateway_public_ip_prefix_association` resource. This field will be removed in the next major version of the Azure Provider.
     */
    publicIpPrefixIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the name of the Resource Group in which the NAT Gateway should exist. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The SKU which should be used. At this time the only supported value is `Standard`. Defaults to `Standard`.
     */
    skuName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of availability zones where the NAT Gateway should be provisioned. Changing this forces a new resource to be created.
     */
    zones?: pulumi.Input<pulumi.Input<string>[]>;
}
