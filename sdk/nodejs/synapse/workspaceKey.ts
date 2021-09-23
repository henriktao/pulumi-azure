// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class WorkspaceKey extends pulumi.CustomResource {
    /**
     * Get an existing WorkspaceKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkspaceKeyState, opts?: pulumi.CustomResourceOptions): WorkspaceKey {
        return new WorkspaceKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:synapse/workspaceKey:WorkspaceKey';

    /**
     * Returns true if the given object is an instance of WorkspaceKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkspaceKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkspaceKey.__pulumiType;
    }

    public readonly active!: pulumi.Output<boolean>;
    public readonly cusomterManagedKeyName!: pulumi.Output<string>;
    public readonly customerManagedKeyVersionlessId!: pulumi.Output<string | undefined>;
    public readonly synapseWorkspaceId!: pulumi.Output<string>;

    /**
     * Create a WorkspaceKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkspaceKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkspaceKeyArgs | WorkspaceKeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkspaceKeyState | undefined;
            inputs["active"] = state ? state.active : undefined;
            inputs["cusomterManagedKeyName"] = state ? state.cusomterManagedKeyName : undefined;
            inputs["customerManagedKeyVersionlessId"] = state ? state.customerManagedKeyVersionlessId : undefined;
            inputs["synapseWorkspaceId"] = state ? state.synapseWorkspaceId : undefined;
        } else {
            const args = argsOrState as WorkspaceKeyArgs | undefined;
            if ((!args || args.active === undefined) && !opts.urn) {
                throw new Error("Missing required property 'active'");
            }
            if ((!args || args.cusomterManagedKeyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cusomterManagedKeyName'");
            }
            if ((!args || args.synapseWorkspaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'synapseWorkspaceId'");
            }
            inputs["active"] = args ? args.active : undefined;
            inputs["cusomterManagedKeyName"] = args ? args.cusomterManagedKeyName : undefined;
            inputs["customerManagedKeyVersionlessId"] = args ? args.customerManagedKeyVersionlessId : undefined;
            inputs["synapseWorkspaceId"] = args ? args.synapseWorkspaceId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WorkspaceKey.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WorkspaceKey resources.
 */
export interface WorkspaceKeyState {
    active?: pulumi.Input<boolean>;
    cusomterManagedKeyName?: pulumi.Input<string>;
    customerManagedKeyVersionlessId?: pulumi.Input<string>;
    synapseWorkspaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WorkspaceKey resource.
 */
export interface WorkspaceKeyArgs {
    active: pulumi.Input<boolean>;
    cusomterManagedKeyName: pulumi.Input<string>;
    customerManagedKeyVersionlessId?: pulumi.Input<string>;
    synapseWorkspaceId: pulumi.Input<string>;
}
