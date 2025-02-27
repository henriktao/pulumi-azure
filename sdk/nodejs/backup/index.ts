// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./containerStorageAccount";
export * from "./getPolicyFileshare";
export * from "./getPolicyVM";
export * from "./policyFileShare";
export * from "./policyVM";
export * from "./protectedFileShare";
export * from "./protectedVM";

// Import resources to register:
import { ContainerStorageAccount } from "./containerStorageAccount";
import { PolicyFileShare } from "./policyFileShare";
import { PolicyVM } from "./policyVM";
import { ProtectedFileShare } from "./protectedFileShare";
import { ProtectedVM } from "./protectedVM";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure:backup/containerStorageAccount:ContainerStorageAccount":
                return new ContainerStorageAccount(name, <any>undefined, { urn })
            case "azure:backup/policyFileShare:PolicyFileShare":
                return new PolicyFileShare(name, <any>undefined, { urn })
            case "azure:backup/policyVM:PolicyVM":
                return new PolicyVM(name, <any>undefined, { urn })
            case "azure:backup/protectedFileShare:ProtectedFileShare":
                return new ProtectedFileShare(name, <any>undefined, { urn })
            case "azure:backup/protectedVM:ProtectedVM":
                return new ProtectedVM(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure", "backup/containerStorageAccount", _module)
pulumi.runtime.registerResourceModule("azure", "backup/policyFileShare", _module)
pulumi.runtime.registerResourceModule("azure", "backup/policyVM", _module)
pulumi.runtime.registerResourceModule("azure", "backup/protectedFileShare", _module)
pulumi.runtime.registerResourceModule("azure", "backup/protectedVM", _module)
