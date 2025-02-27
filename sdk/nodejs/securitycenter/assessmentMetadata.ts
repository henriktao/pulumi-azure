// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages the Security Center Assessment Metadata for Azure Security Center.
 *
 * > **NOTE:** This resource has been deprecated in favour of the `azure.securitycenter.AssessmentPolicy` resource and will be removed in the next major version of the AzureRM Provider. The new resource shares the same fields as this one, and information on migrating across can be found in this guide.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = new azure.securitycenter.AssessmentMetadata("example", {
 *     description: "Test Description",
 *     displayName: "Test Display Name",
 *     severity: "Medium",
 * });
 * ```
 *
 * ## Import
 *
 * Security Assessments Metadata can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:securitycenter/assessmentMetadata:AssessmentMetadata example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Security/assessmentMetadata/metadata1
 * ```
 */
export class AssessmentMetadata extends pulumi.CustomResource {
    /**
     * Get an existing AssessmentMetadata resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AssessmentMetadataState, opts?: pulumi.CustomResourceOptions): AssessmentMetadata {
        return new AssessmentMetadata(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:securitycenter/assessmentMetadata:AssessmentMetadata';

    /**
     * Returns true if the given object is an instance of AssessmentMetadata.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AssessmentMetadata {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AssessmentMetadata.__pulumiType;
    }

    /**
     * A list of the categories of resource that is at risk when the Security Center Assessment is unhealthy. Possible values are `Unknown`, `Compute`, `Data`, `IdentityAndAccess`, `IoT` and `Networking`.
     */
    public readonly categories!: pulumi.Output<string[]>;
    /**
     * The description of the Security Center Assessment.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The user-friendly display name of the Security Center Assessment.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The implementation effort which is used to remediate the Security Center Assessment. Possible values are `Low`, `Moderate` and `High`.
     */
    public readonly implementationEffort!: pulumi.Output<string | undefined>;
    /**
     * The GUID as the name of the Security Center Assessment Metadata.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The description which is used to mitigate the security issue.
     */
    public readonly remediationDescription!: pulumi.Output<string | undefined>;
    /**
     * The severity level of the Security Center Assessment. Possible values are `Low`, `Medium` and `High`. Defaults to `Medium`.
     */
    public readonly severity!: pulumi.Output<string | undefined>;
    /**
     * A list of the threat impacts for the Security Center Assessment. Possible values are `AccountBreach`, `DataExfiltration`, `DataSpillage`, `DenialOfService`, `ElevationOfPrivilege`, `MaliciousInsider`, `MissingCoverage` and `ThreatResistance`.
     */
    public readonly threats!: pulumi.Output<string[] | undefined>;
    /**
     * The user impact of the Security Center Assessment. Possible values are `Low`, `Moderate` and `High`.
     */
    public readonly userImpact!: pulumi.Output<string | undefined>;

    /**
     * Create a AssessmentMetadata resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AssessmentMetadataArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AssessmentMetadataArgs | AssessmentMetadataState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AssessmentMetadataState | undefined;
            resourceInputs["categories"] = state ? state.categories : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["implementationEffort"] = state ? state.implementationEffort : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["remediationDescription"] = state ? state.remediationDescription : undefined;
            resourceInputs["severity"] = state ? state.severity : undefined;
            resourceInputs["threats"] = state ? state.threats : undefined;
            resourceInputs["userImpact"] = state ? state.userImpact : undefined;
        } else {
            const args = argsOrState as AssessmentMetadataArgs | undefined;
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["categories"] = args ? args.categories : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["implementationEffort"] = args ? args.implementationEffort : undefined;
            resourceInputs["remediationDescription"] = args ? args.remediationDescription : undefined;
            resourceInputs["severity"] = args ? args.severity : undefined;
            resourceInputs["threats"] = args ? args.threats : undefined;
            resourceInputs["userImpact"] = args ? args.userImpact : undefined;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AssessmentMetadata.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AssessmentMetadata resources.
 */
export interface AssessmentMetadataState {
    /**
     * A list of the categories of resource that is at risk when the Security Center Assessment is unhealthy. Possible values are `Unknown`, `Compute`, `Data`, `IdentityAndAccess`, `IoT` and `Networking`.
     */
    categories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of the Security Center Assessment.
     */
    description?: pulumi.Input<string>;
    /**
     * The user-friendly display name of the Security Center Assessment.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The implementation effort which is used to remediate the Security Center Assessment. Possible values are `Low`, `Moderate` and `High`.
     */
    implementationEffort?: pulumi.Input<string>;
    /**
     * The GUID as the name of the Security Center Assessment Metadata.
     */
    name?: pulumi.Input<string>;
    /**
     * The description which is used to mitigate the security issue.
     */
    remediationDescription?: pulumi.Input<string>;
    /**
     * The severity level of the Security Center Assessment. Possible values are `Low`, `Medium` and `High`. Defaults to `Medium`.
     */
    severity?: pulumi.Input<string>;
    /**
     * A list of the threat impacts for the Security Center Assessment. Possible values are `AccountBreach`, `DataExfiltration`, `DataSpillage`, `DenialOfService`, `ElevationOfPrivilege`, `MaliciousInsider`, `MissingCoverage` and `ThreatResistance`.
     */
    threats?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The user impact of the Security Center Assessment. Possible values are `Low`, `Moderate` and `High`.
     */
    userImpact?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AssessmentMetadata resource.
 */
export interface AssessmentMetadataArgs {
    /**
     * A list of the categories of resource that is at risk when the Security Center Assessment is unhealthy. Possible values are `Unknown`, `Compute`, `Data`, `IdentityAndAccess`, `IoT` and `Networking`.
     */
    categories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of the Security Center Assessment.
     */
    description: pulumi.Input<string>;
    /**
     * The user-friendly display name of the Security Center Assessment.
     */
    displayName: pulumi.Input<string>;
    /**
     * The implementation effort which is used to remediate the Security Center Assessment. Possible values are `Low`, `Moderate` and `High`.
     */
    implementationEffort?: pulumi.Input<string>;
    /**
     * The description which is used to mitigate the security issue.
     */
    remediationDescription?: pulumi.Input<string>;
    /**
     * The severity level of the Security Center Assessment. Possible values are `Low`, `Medium` and `High`. Defaults to `Medium`.
     */
    severity?: pulumi.Input<string>;
    /**
     * A list of the threat impacts for the Security Center Assessment. Possible values are `AccountBreach`, `DataExfiltration`, `DataSpillage`, `DenialOfService`, `ElevationOfPrivilege`, `MaliciousInsider`, `MissingCoverage` and `ThreatResistance`.
     */
    threats?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The user impact of the Security Center Assessment. Possible values are `Low`, `Moderate` and `High`.
     */
    userImpact?: pulumi.Input<string>;
}
