// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./activeSlot";
export * from "./appService";
export * from "./certificate";
export * from "./certificateBinding";
export * from "./certificateOrder";
export * from "./customHostnameBinding";
export * from "./environment";
export * from "./environmentV3";
export * from "./functionApp";
export * from "./functionAppSlot";
export * from "./getAppService";
export * from "./getAppServiceEnvironment";
export * from "./getAppServicePlan";
export * from "./getCertificate";
export * from "./getCertificateOrder";
export * from "./getEnvironmentV3";
export * from "./getFunctionApp";
export * from "./getFunctionAppHostKeys";
export * from "./hybridConnection";
export * from "./kind";
export * from "./managedCertificate";
export * from "./plan";
export * from "./publicCertificate";
export * from "./slot";
export * from "./slotCustomHostnameBinding";
export * from "./slotVirtualNetworkSwiftConnection";
export * from "./sourceCodeToken";
export * from "./staticSite";
export * from "./staticSiteCustomDomain";
export * from "./virtualNetworkSwiftConnection";
export * from "./zMixins";
export * from "./zMixins_durable";
export * from "./zMixins_http";
export * from "./zMixins_timer";

// Import resources to register:
import { ActiveSlot } from "./activeSlot";
import { AppService } from "./appService";
import { Certificate } from "./certificate";
import { CertificateBinding } from "./certificateBinding";
import { CertificateOrder } from "./certificateOrder";
import { CustomHostnameBinding } from "./customHostnameBinding";
import { Environment } from "./environment";
import { EnvironmentV3 } from "./environmentV3";
import { FunctionApp } from "./functionApp";
import { FunctionAppSlot } from "./functionAppSlot";
import { HybridConnection } from "./hybridConnection";
import { ManagedCertificate } from "./managedCertificate";
import { Plan } from "./plan";
import { PublicCertificate } from "./publicCertificate";
import { Slot } from "./slot";
import { SlotCustomHostnameBinding } from "./slotCustomHostnameBinding";
import { SlotVirtualNetworkSwiftConnection } from "./slotVirtualNetworkSwiftConnection";
import { SourceCodeToken } from "./sourceCodeToken";
import { StaticSite } from "./staticSite";
import { StaticSiteCustomDomain } from "./staticSiteCustomDomain";
import { VirtualNetworkSwiftConnection } from "./virtualNetworkSwiftConnection";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure:appservice/activeSlot:ActiveSlot":
                return new ActiveSlot(name, <any>undefined, { urn })
            case "azure:appservice/appService:AppService":
                return new AppService(name, <any>undefined, { urn })
            case "azure:appservice/certificate:Certificate":
                return new Certificate(name, <any>undefined, { urn })
            case "azure:appservice/certificateBinding:CertificateBinding":
                return new CertificateBinding(name, <any>undefined, { urn })
            case "azure:appservice/certificateOrder:CertificateOrder":
                return new CertificateOrder(name, <any>undefined, { urn })
            case "azure:appservice/customHostnameBinding:CustomHostnameBinding":
                return new CustomHostnameBinding(name, <any>undefined, { urn })
            case "azure:appservice/environment:Environment":
                return new Environment(name, <any>undefined, { urn })
            case "azure:appservice/environmentV3:EnvironmentV3":
                return new EnvironmentV3(name, <any>undefined, { urn })
            case "azure:appservice/functionApp:FunctionApp":
                return new FunctionApp(name, <any>undefined, { urn })
            case "azure:appservice/functionAppSlot:FunctionAppSlot":
                return new FunctionAppSlot(name, <any>undefined, { urn })
            case "azure:appservice/hybridConnection:HybridConnection":
                return new HybridConnection(name, <any>undefined, { urn })
            case "azure:appservice/managedCertificate:ManagedCertificate":
                return new ManagedCertificate(name, <any>undefined, { urn })
            case "azure:appservice/plan:Plan":
                return new Plan(name, <any>undefined, { urn })
            case "azure:appservice/publicCertificate:PublicCertificate":
                return new PublicCertificate(name, <any>undefined, { urn })
            case "azure:appservice/slot:Slot":
                return new Slot(name, <any>undefined, { urn })
            case "azure:appservice/slotCustomHostnameBinding:SlotCustomHostnameBinding":
                return new SlotCustomHostnameBinding(name, <any>undefined, { urn })
            case "azure:appservice/slotVirtualNetworkSwiftConnection:SlotVirtualNetworkSwiftConnection":
                return new SlotVirtualNetworkSwiftConnection(name, <any>undefined, { urn })
            case "azure:appservice/sourceCodeToken:SourceCodeToken":
                return new SourceCodeToken(name, <any>undefined, { urn })
            case "azure:appservice/staticSite:StaticSite":
                return new StaticSite(name, <any>undefined, { urn })
            case "azure:appservice/staticSiteCustomDomain:StaticSiteCustomDomain":
                return new StaticSiteCustomDomain(name, <any>undefined, { urn })
            case "azure:appservice/virtualNetworkSwiftConnection:VirtualNetworkSwiftConnection":
                return new VirtualNetworkSwiftConnection(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure", "appservice/activeSlot", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/appService", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/certificate", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/certificateBinding", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/certificateOrder", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/customHostnameBinding", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/environment", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/environmentV3", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/functionApp", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/functionAppSlot", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/hybridConnection", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/managedCertificate", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/plan", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/publicCertificate", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/slot", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/slotCustomHostnameBinding", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/slotVirtualNetworkSwiftConnection", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/sourceCodeToken", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/staticSite", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/staticSiteCustomDomain", _module)
pulumi.runtime.registerResourceModule("azure", "appservice/virtualNetworkSwiftConnection", _module)
