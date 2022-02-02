// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * SQL Instance Failover Groups can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:sql/managedInstanceFailoverGroup:ManagedInstanceFailoverGroup example /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Sql/locations/Location/instanceFailoverGroups/failoverGroup1
 * ```
 */
export class ManagedInstanceFailoverGroup extends pulumi.CustomResource {
    /**
     * Get an existing ManagedInstanceFailoverGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ManagedInstanceFailoverGroupState, opts?: pulumi.CustomResourceOptions): ManagedInstanceFailoverGroup {
        return new ManagedInstanceFailoverGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:sql/managedInstanceFailoverGroup:ManagedInstanceFailoverGroup';

    /**
     * Returns true if the given object is an instance of ManagedInstanceFailoverGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagedInstanceFailoverGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagedInstanceFailoverGroup.__pulumiType;
    }

    /**
     * The Azure Region where the SQL Instance Failover Group exists.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the SQL Managed Instance which will be replicated using a SQL Instance Failover Group. Changing this forces a new SQL Instance Failover Group to be created.
     */
    public readonly managedInstanceName!: pulumi.Output<string>;
    /**
     * The name which should be used for this SQL Instance Failover Group. Changing this forces a new SQL Instance Failover Group to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the SQL Managed Instance which will be replicated to.
     */
    public readonly partnerManagedInstanceId!: pulumi.Output<string>;
    /**
     * A `partnerRegion` block as defined below.
     */
    public /*out*/ readonly partnerRegions!: pulumi.Output<outputs.sql.ManagedInstanceFailoverGroupPartnerRegion[]>;
    /**
     * A `readWriteEndpointFailoverPolicy` block as defined below.
     */
    public readonly readWriteEndpointFailoverPolicy!: pulumi.Output<outputs.sql.ManagedInstanceFailoverGroupReadWriteEndpointFailoverPolicy>;
    /**
     * Failover policy for the read-only endpoint. Defaults to `false`.
     */
    public readonly readonlyEndpointFailoverPolicyEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the Resource Group where the SQL Instance Failover Group should exist. Changing this forces a new SQL Instance Failover Group to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The partner replication role of the SQL Instance Failover Group.
     */
    public /*out*/ readonly role!: pulumi.Output<string>;

    /**
     * Create a ManagedInstanceFailoverGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedInstanceFailoverGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManagedInstanceFailoverGroupArgs | ManagedInstanceFailoverGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ManagedInstanceFailoverGroupState | undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["managedInstanceName"] = state ? state.managedInstanceName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["partnerManagedInstanceId"] = state ? state.partnerManagedInstanceId : undefined;
            resourceInputs["partnerRegions"] = state ? state.partnerRegions : undefined;
            resourceInputs["readWriteEndpointFailoverPolicy"] = state ? state.readWriteEndpointFailoverPolicy : undefined;
            resourceInputs["readonlyEndpointFailoverPolicyEnabled"] = state ? state.readonlyEndpointFailoverPolicyEnabled : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as ManagedInstanceFailoverGroupArgs | undefined;
            if ((!args || args.managedInstanceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'managedInstanceName'");
            }
            if ((!args || args.partnerManagedInstanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'partnerManagedInstanceId'");
            }
            if ((!args || args.readWriteEndpointFailoverPolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'readWriteEndpointFailoverPolicy'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["managedInstanceName"] = args ? args.managedInstanceName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["partnerManagedInstanceId"] = args ? args.partnerManagedInstanceId : undefined;
            resourceInputs["readWriteEndpointFailoverPolicy"] = args ? args.readWriteEndpointFailoverPolicy : undefined;
            resourceInputs["readonlyEndpointFailoverPolicyEnabled"] = args ? args.readonlyEndpointFailoverPolicyEnabled : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["partnerRegions"] = undefined /*out*/;
            resourceInputs["role"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ManagedInstanceFailoverGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ManagedInstanceFailoverGroup resources.
 */
export interface ManagedInstanceFailoverGroupState {
    /**
     * The Azure Region where the SQL Instance Failover Group exists.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the SQL Managed Instance which will be replicated using a SQL Instance Failover Group. Changing this forces a new SQL Instance Failover Group to be created.
     */
    managedInstanceName?: pulumi.Input<string>;
    /**
     * The name which should be used for this SQL Instance Failover Group. Changing this forces a new SQL Instance Failover Group to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the SQL Managed Instance which will be replicated to.
     */
    partnerManagedInstanceId?: pulumi.Input<string>;
    /**
     * A `partnerRegion` block as defined below.
     */
    partnerRegions?: pulumi.Input<pulumi.Input<inputs.sql.ManagedInstanceFailoverGroupPartnerRegion>[]>;
    /**
     * A `readWriteEndpointFailoverPolicy` block as defined below.
     */
    readWriteEndpointFailoverPolicy?: pulumi.Input<inputs.sql.ManagedInstanceFailoverGroupReadWriteEndpointFailoverPolicy>;
    /**
     * Failover policy for the read-only endpoint. Defaults to `false`.
     */
    readonlyEndpointFailoverPolicyEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the Resource Group where the SQL Instance Failover Group should exist. Changing this forces a new SQL Instance Failover Group to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The partner replication role of the SQL Instance Failover Group.
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ManagedInstanceFailoverGroup resource.
 */
export interface ManagedInstanceFailoverGroupArgs {
    /**
     * The Azure Region where the SQL Instance Failover Group exists.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the SQL Managed Instance which will be replicated using a SQL Instance Failover Group. Changing this forces a new SQL Instance Failover Group to be created.
     */
    managedInstanceName: pulumi.Input<string>;
    /**
     * The name which should be used for this SQL Instance Failover Group. Changing this forces a new SQL Instance Failover Group to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the SQL Managed Instance which will be replicated to.
     */
    partnerManagedInstanceId: pulumi.Input<string>;
    /**
     * A `readWriteEndpointFailoverPolicy` block as defined below.
     */
    readWriteEndpointFailoverPolicy: pulumi.Input<inputs.sql.ManagedInstanceFailoverGroupReadWriteEndpointFailoverPolicy>;
    /**
     * Failover policy for the read-only endpoint. Defaults to `false`.
     */
    readonlyEndpointFailoverPolicyEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the Resource Group where the SQL Instance Failover Group should exist. Changing this forces a new SQL Instance Failover Group to be created.
     */
    resourceGroupName: pulumi.Input<string>;
}
