// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./activeDirectoryAdministrator";
export * from "./database";
export * from "./elasticPool";
export * from "./failoverGroup";
export * from "./firewallRule";
export * from "./getDatabase";
export * from "./getServer";
export * from "./managedDatabase";
export * from "./managedInstance";
export * from "./managedInstanceActiveDirectoryAdministrator";
export * from "./managedInstanceFailoverGroup";
export * from "./sqlServer";
export * from "./virtualNetworkRule";

// Import resources to register:
import { ActiveDirectoryAdministrator } from "./activeDirectoryAdministrator";
import { Database } from "./database";
import { ElasticPool } from "./elasticPool";
import { FailoverGroup } from "./failoverGroup";
import { FirewallRule } from "./firewallRule";
import { ManagedDatabase } from "./managedDatabase";
import { ManagedInstance } from "./managedInstance";
import { ManagedInstanceActiveDirectoryAdministrator } from "./managedInstanceActiveDirectoryAdministrator";
import { ManagedInstanceFailoverGroup } from "./managedInstanceFailoverGroup";
import { SqlServer } from "./sqlServer";
import { VirtualNetworkRule } from "./virtualNetworkRule";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure:sql/activeDirectoryAdministrator:ActiveDirectoryAdministrator":
                return new ActiveDirectoryAdministrator(name, <any>undefined, { urn })
            case "azure:sql/database:Database":
                return new Database(name, <any>undefined, { urn })
            case "azure:sql/elasticPool:ElasticPool":
                return new ElasticPool(name, <any>undefined, { urn })
            case "azure:sql/failoverGroup:FailoverGroup":
                return new FailoverGroup(name, <any>undefined, { urn })
            case "azure:sql/firewallRule:FirewallRule":
                return new FirewallRule(name, <any>undefined, { urn })
            case "azure:sql/managedDatabase:ManagedDatabase":
                return new ManagedDatabase(name, <any>undefined, { urn })
            case "azure:sql/managedInstance:ManagedInstance":
                return new ManagedInstance(name, <any>undefined, { urn })
            case "azure:sql/managedInstanceActiveDirectoryAdministrator:ManagedInstanceActiveDirectoryAdministrator":
                return new ManagedInstanceActiveDirectoryAdministrator(name, <any>undefined, { urn })
            case "azure:sql/managedInstanceFailoverGroup:ManagedInstanceFailoverGroup":
                return new ManagedInstanceFailoverGroup(name, <any>undefined, { urn })
            case "azure:sql/sqlServer:SqlServer":
                return new SqlServer(name, <any>undefined, { urn })
            case "azure:sql/virtualNetworkRule:VirtualNetworkRule":
                return new VirtualNetworkRule(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure", "sql/activeDirectoryAdministrator", _module)
pulumi.runtime.registerResourceModule("azure", "sql/database", _module)
pulumi.runtime.registerResourceModule("azure", "sql/elasticPool", _module)
pulumi.runtime.registerResourceModule("azure", "sql/failoverGroup", _module)
pulumi.runtime.registerResourceModule("azure", "sql/firewallRule", _module)
pulumi.runtime.registerResourceModule("azure", "sql/managedDatabase", _module)
pulumi.runtime.registerResourceModule("azure", "sql/managedInstance", _module)
pulumi.runtime.registerResourceModule("azure", "sql/managedInstanceActiveDirectoryAdministrator", _module)
pulumi.runtime.registerResourceModule("azure", "sql/managedInstanceFailoverGroup", _module)
pulumi.runtime.registerResourceModule("azure", "sql/sqlServer", _module)
pulumi.runtime.registerResourceModule("azure", "sql/virtualNetworkRule", _module)
