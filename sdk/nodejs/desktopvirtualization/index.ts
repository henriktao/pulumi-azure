// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./application";
export * from "./applicationGroup";
export * from "./hostPool";
export * from "./scalingPlan";
export * from "./workspace";
export * from "./workspaceApplicationGroupAssociation";

// Import resources to register:
import { Application } from "./application";
import { ApplicationGroup } from "./applicationGroup";
import { HostPool } from "./hostPool";
import { ScalingPlan } from "./scalingPlan";
import { Workspace } from "./workspace";
import { WorkspaceApplicationGroupAssociation } from "./workspaceApplicationGroupAssociation";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure:desktopvirtualization/application:Application":
                return new Application(name, <any>undefined, { urn })
            case "azure:desktopvirtualization/applicationGroup:ApplicationGroup":
                return new ApplicationGroup(name, <any>undefined, { urn })
            case "azure:desktopvirtualization/hostPool:HostPool":
                return new HostPool(name, <any>undefined, { urn })
            case "azure:desktopvirtualization/scalingPlan:ScalingPlan":
                return new ScalingPlan(name, <any>undefined, { urn })
            case "azure:desktopvirtualization/workspace:Workspace":
                return new Workspace(name, <any>undefined, { urn })
            case "azure:desktopvirtualization/workspaceApplicationGroupAssociation:WorkspaceApplicationGroupAssociation":
                return new WorkspaceApplicationGroupAssociation(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure", "desktopvirtualization/application", _module)
pulumi.runtime.registerResourceModule("azure", "desktopvirtualization/applicationGroup", _module)
pulumi.runtime.registerResourceModule("azure", "desktopvirtualization/hostPool", _module)
pulumi.runtime.registerResourceModule("azure", "desktopvirtualization/scalingPlan", _module)
pulumi.runtime.registerResourceModule("azure", "desktopvirtualization/workspace", _module)
pulumi.runtime.registerResourceModule("azure", "desktopvirtualization/workspaceApplicationGroupAssociation", _module)
