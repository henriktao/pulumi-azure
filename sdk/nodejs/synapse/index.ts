// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./firewallRule";
export * from "./getWorkspace";
export * from "./integrationRuntimeAzure";
export * from "./integrationRuntimeSelfHosted";
export * from "./linkedService";
export * from "./managedPrivateEndpoint";
export * from "./privateLinkHub";
export * from "./roleAssignment";
export * from "./sparkPool";
export * from "./sqlPool";
export * from "./sqlPoolExtendedAuditingPolicy";
export * from "./sqlPoolSecurityAlertPolicy";
export * from "./sqlPoolVulnerabilityAssessment";
export * from "./sqlPoolVulnerabilityAssessmentBaseline";
export * from "./sqlPoolWorkloadClassifier";
export * from "./sqlPoolWorkloadGroup";
export * from "./workspace";
export * from "./workspaceAadAdmin";
export * from "./workspaceExtendedAuditingPolicy";
export * from "./workspaceKey";
export * from "./workspaceSecurityAlertPolicy";
export * from "./workspaceSqlAadAdmin";
export * from "./workspaceVulnerabilityAssessment";

// Import resources to register:
import { FirewallRule } from "./firewallRule";
import { IntegrationRuntimeAzure } from "./integrationRuntimeAzure";
import { IntegrationRuntimeSelfHosted } from "./integrationRuntimeSelfHosted";
import { LinkedService } from "./linkedService";
import { ManagedPrivateEndpoint } from "./managedPrivateEndpoint";
import { PrivateLinkHub } from "./privateLinkHub";
import { RoleAssignment } from "./roleAssignment";
import { SparkPool } from "./sparkPool";
import { SqlPool } from "./sqlPool";
import { SqlPoolExtendedAuditingPolicy } from "./sqlPoolExtendedAuditingPolicy";
import { SqlPoolSecurityAlertPolicy } from "./sqlPoolSecurityAlertPolicy";
import { SqlPoolVulnerabilityAssessment } from "./sqlPoolVulnerabilityAssessment";
import { SqlPoolVulnerabilityAssessmentBaseline } from "./sqlPoolVulnerabilityAssessmentBaseline";
import { SqlPoolWorkloadClassifier } from "./sqlPoolWorkloadClassifier";
import { SqlPoolWorkloadGroup } from "./sqlPoolWorkloadGroup";
import { Workspace } from "./workspace";
import { WorkspaceAadAdmin } from "./workspaceAadAdmin";
import { WorkspaceExtendedAuditingPolicy } from "./workspaceExtendedAuditingPolicy";
import { WorkspaceKey } from "./workspaceKey";
import { WorkspaceSecurityAlertPolicy } from "./workspaceSecurityAlertPolicy";
import { WorkspaceSqlAadAdmin } from "./workspaceSqlAadAdmin";
import { WorkspaceVulnerabilityAssessment } from "./workspaceVulnerabilityAssessment";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure:synapse/firewallRule:FirewallRule":
                return new FirewallRule(name, <any>undefined, { urn })
            case "azure:synapse/integrationRuntimeAzure:IntegrationRuntimeAzure":
                return new IntegrationRuntimeAzure(name, <any>undefined, { urn })
            case "azure:synapse/integrationRuntimeSelfHosted:IntegrationRuntimeSelfHosted":
                return new IntegrationRuntimeSelfHosted(name, <any>undefined, { urn })
            case "azure:synapse/linkedService:LinkedService":
                return new LinkedService(name, <any>undefined, { urn })
            case "azure:synapse/managedPrivateEndpoint:ManagedPrivateEndpoint":
                return new ManagedPrivateEndpoint(name, <any>undefined, { urn })
            case "azure:synapse/privateLinkHub:PrivateLinkHub":
                return new PrivateLinkHub(name, <any>undefined, { urn })
            case "azure:synapse/roleAssignment:RoleAssignment":
                return new RoleAssignment(name, <any>undefined, { urn })
            case "azure:synapse/sparkPool:SparkPool":
                return new SparkPool(name, <any>undefined, { urn })
            case "azure:synapse/sqlPool:SqlPool":
                return new SqlPool(name, <any>undefined, { urn })
            case "azure:synapse/sqlPoolExtendedAuditingPolicy:SqlPoolExtendedAuditingPolicy":
                return new SqlPoolExtendedAuditingPolicy(name, <any>undefined, { urn })
            case "azure:synapse/sqlPoolSecurityAlertPolicy:SqlPoolSecurityAlertPolicy":
                return new SqlPoolSecurityAlertPolicy(name, <any>undefined, { urn })
            case "azure:synapse/sqlPoolVulnerabilityAssessment:SqlPoolVulnerabilityAssessment":
                return new SqlPoolVulnerabilityAssessment(name, <any>undefined, { urn })
            case "azure:synapse/sqlPoolVulnerabilityAssessmentBaseline:SqlPoolVulnerabilityAssessmentBaseline":
                return new SqlPoolVulnerabilityAssessmentBaseline(name, <any>undefined, { urn })
            case "azure:synapse/sqlPoolWorkloadClassifier:SqlPoolWorkloadClassifier":
                return new SqlPoolWorkloadClassifier(name, <any>undefined, { urn })
            case "azure:synapse/sqlPoolWorkloadGroup:SqlPoolWorkloadGroup":
                return new SqlPoolWorkloadGroup(name, <any>undefined, { urn })
            case "azure:synapse/workspace:Workspace":
                return new Workspace(name, <any>undefined, { urn })
            case "azure:synapse/workspaceAadAdmin:WorkspaceAadAdmin":
                return new WorkspaceAadAdmin(name, <any>undefined, { urn })
            case "azure:synapse/workspaceExtendedAuditingPolicy:WorkspaceExtendedAuditingPolicy":
                return new WorkspaceExtendedAuditingPolicy(name, <any>undefined, { urn })
            case "azure:synapse/workspaceKey:WorkspaceKey":
                return new WorkspaceKey(name, <any>undefined, { urn })
            case "azure:synapse/workspaceSecurityAlertPolicy:WorkspaceSecurityAlertPolicy":
                return new WorkspaceSecurityAlertPolicy(name, <any>undefined, { urn })
            case "azure:synapse/workspaceSqlAadAdmin:WorkspaceSqlAadAdmin":
                return new WorkspaceSqlAadAdmin(name, <any>undefined, { urn })
            case "azure:synapse/workspaceVulnerabilityAssessment:WorkspaceVulnerabilityAssessment":
                return new WorkspaceVulnerabilityAssessment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure", "synapse/firewallRule", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/integrationRuntimeAzure", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/integrationRuntimeSelfHosted", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/linkedService", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/managedPrivateEndpoint", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/privateLinkHub", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/roleAssignment", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/sparkPool", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/sqlPool", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/sqlPoolExtendedAuditingPolicy", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/sqlPoolSecurityAlertPolicy", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/sqlPoolVulnerabilityAssessment", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/sqlPoolVulnerabilityAssessmentBaseline", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/sqlPoolWorkloadClassifier", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/sqlPoolWorkloadGroup", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/workspace", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/workspaceAadAdmin", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/workspaceExtendedAuditingPolicy", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/workspaceKey", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/workspaceSecurityAlertPolicy", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/workspaceSqlAadAdmin", _module)
pulumi.runtime.registerResourceModule("azure", "synapse/workspaceVulnerabilityAssessment", _module)
