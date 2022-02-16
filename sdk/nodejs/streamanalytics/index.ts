// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./cluster";
export * from "./functionJavaScriptUDF";
export * from "./getJob";
export * from "./job";
export * from "./managedPrivateEndpoint";
export * from "./outputBlob";
export * from "./outputEventHub";
export * from "./outputFunction";
export * from "./outputMssql";
export * from "./outputServiceBusQueue";
export * from "./outputServicebusTopic";
export * from "./outputSynapse";
export * from "./outputTable";
export * from "./referenceInputBlob";
export * from "./referenceInputMssql";
export * from "./streamInputBlob";
export * from "./streamInputEventHub";
export * from "./streamInputIotHub";

// Import resources to register:
import { Cluster } from "./cluster";
import { FunctionJavaScriptUDF } from "./functionJavaScriptUDF";
import { Job } from "./job";
import { ManagedPrivateEndpoint } from "./managedPrivateEndpoint";
import { OutputBlob } from "./outputBlob";
import { OutputEventHub } from "./outputEventHub";
import { OutputFunction } from "./outputFunction";
import { OutputMssql } from "./outputMssql";
import { OutputServiceBusQueue } from "./outputServiceBusQueue";
import { OutputServicebusTopic } from "./outputServicebusTopic";
import { OutputSynapse } from "./outputSynapse";
import { OutputTable } from "./outputTable";
import { ReferenceInputBlob } from "./referenceInputBlob";
import { ReferenceInputMssql } from "./referenceInputMssql";
import { StreamInputBlob } from "./streamInputBlob";
import { StreamInputEventHub } from "./streamInputEventHub";
import { StreamInputIotHub } from "./streamInputIotHub";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure:streamanalytics/cluster:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "azure:streamanalytics/functionJavaScriptUDF:FunctionJavaScriptUDF":
                return new FunctionJavaScriptUDF(name, <any>undefined, { urn })
            case "azure:streamanalytics/job:Job":
                return new Job(name, <any>undefined, { urn })
            case "azure:streamanalytics/managedPrivateEndpoint:ManagedPrivateEndpoint":
                return new ManagedPrivateEndpoint(name, <any>undefined, { urn })
            case "azure:streamanalytics/outputBlob:OutputBlob":
                return new OutputBlob(name, <any>undefined, { urn })
            case "azure:streamanalytics/outputEventHub:OutputEventHub":
                return new OutputEventHub(name, <any>undefined, { urn })
            case "azure:streamanalytics/outputFunction:OutputFunction":
                return new OutputFunction(name, <any>undefined, { urn })
            case "azure:streamanalytics/outputMssql:OutputMssql":
                return new OutputMssql(name, <any>undefined, { urn })
            case "azure:streamanalytics/outputServiceBusQueue:OutputServiceBusQueue":
                return new OutputServiceBusQueue(name, <any>undefined, { urn })
            case "azure:streamanalytics/outputServicebusTopic:OutputServicebusTopic":
                return new OutputServicebusTopic(name, <any>undefined, { urn })
            case "azure:streamanalytics/outputSynapse:OutputSynapse":
                return new OutputSynapse(name, <any>undefined, { urn })
            case "azure:streamanalytics/outputTable:OutputTable":
                return new OutputTable(name, <any>undefined, { urn })
            case "azure:streamanalytics/referenceInputBlob:ReferenceInputBlob":
                return new ReferenceInputBlob(name, <any>undefined, { urn })
            case "azure:streamanalytics/referenceInputMssql:ReferenceInputMssql":
                return new ReferenceInputMssql(name, <any>undefined, { urn })
            case "azure:streamanalytics/streamInputBlob:StreamInputBlob":
                return new StreamInputBlob(name, <any>undefined, { urn })
            case "azure:streamanalytics/streamInputEventHub:StreamInputEventHub":
                return new StreamInputEventHub(name, <any>undefined, { urn })
            case "azure:streamanalytics/streamInputIotHub:StreamInputIotHub":
                return new StreamInputIotHub(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure", "streamanalytics/cluster", _module)
pulumi.runtime.registerResourceModule("azure", "streamanalytics/functionJavaScriptUDF", _module)
pulumi.runtime.registerResourceModule("azure", "streamanalytics/job", _module)
pulumi.runtime.registerResourceModule("azure", "streamanalytics/managedPrivateEndpoint", _module)
pulumi.runtime.registerResourceModule("azure", "streamanalytics/outputBlob", _module)
pulumi.runtime.registerResourceModule("azure", "streamanalytics/outputEventHub", _module)
pulumi.runtime.registerResourceModule("azure", "streamanalytics/outputFunction", _module)
pulumi.runtime.registerResourceModule("azure", "streamanalytics/outputMssql", _module)
pulumi.runtime.registerResourceModule("azure", "streamanalytics/outputServiceBusQueue", _module)
pulumi.runtime.registerResourceModule("azure", "streamanalytics/outputServicebusTopic", _module)
pulumi.runtime.registerResourceModule("azure", "streamanalytics/outputSynapse", _module)
pulumi.runtime.registerResourceModule("azure", "streamanalytics/outputTable", _module)
pulumi.runtime.registerResourceModule("azure", "streamanalytics/referenceInputBlob", _module)
pulumi.runtime.registerResourceModule("azure", "streamanalytics/referenceInputMssql", _module)
pulumi.runtime.registerResourceModule("azure", "streamanalytics/streamInputBlob", _module)
pulumi.runtime.registerResourceModule("azure", "streamanalytics/streamInputEventHub", _module)
pulumi.runtime.registerResourceModule("azure", "streamanalytics/streamInputIotHub", _module)
