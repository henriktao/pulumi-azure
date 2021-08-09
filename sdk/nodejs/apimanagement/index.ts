// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./api";
export * from "./apiDiagnostic";
export * from "./apiOperation";
export * from "./apiOperationPolicy";
export * from "./apiOperationTag";
export * from "./apiPolicy";
export * from "./apiRelease";
export * from "./apiSchema";
export * from "./apiVersionSet";
export * from "./authorizationServer";
export * from "./backend";
export * from "./certificate";
export * from "./customDomain";
export * from "./diagnostic";
export * from "./emailTemplate";
export * from "./gateway";
export * from "./gatewayApi";
export * from "./getApi";
export * from "./getApiVersionSet";
export * from "./getGateway";
export * from "./getGroup";
export * from "./getProduct";
export * from "./getService";
export * from "./getUser";
export * from "./group";
export * from "./groupUser";
export * from "./identityProviderAad";
export * from "./identityProviderAadb2c";
export * from "./identityProviderFacebook";
export * from "./identityProviderGoogle";
export * from "./identityProviderMicrosoft";
export * from "./identityProviderTwitter";
export * from "./logger";
export * from "./namedValue";
export * from "./openIdConnectProvider";
export * from "./policy";
export * from "./product";
export * from "./productApi";
export * from "./productGroup";
export * from "./productPolicy";
export * from "./property";
export * from "./redisCache";
export * from "./service";
export * from "./subscription";
export * from "./tag";
export * from "./user";

// Import resources to register:
import { Api } from "./api";
import { ApiDiagnostic } from "./apiDiagnostic";
import { ApiOperation } from "./apiOperation";
import { ApiOperationPolicy } from "./apiOperationPolicy";
import { ApiOperationTag } from "./apiOperationTag";
import { ApiPolicy } from "./apiPolicy";
import { ApiRelease } from "./apiRelease";
import { ApiSchema } from "./apiSchema";
import { ApiVersionSet } from "./apiVersionSet";
import { AuthorizationServer } from "./authorizationServer";
import { Backend } from "./backend";
import { Certificate } from "./certificate";
import { CustomDomain } from "./customDomain";
import { Diagnostic } from "./diagnostic";
import { EmailTemplate } from "./emailTemplate";
import { Gateway } from "./gateway";
import { GatewayApi } from "./gatewayApi";
import { Group } from "./group";
import { GroupUser } from "./groupUser";
import { IdentityProviderAad } from "./identityProviderAad";
import { IdentityProviderAadb2c } from "./identityProviderAadb2c";
import { IdentityProviderFacebook } from "./identityProviderFacebook";
import { IdentityProviderGoogle } from "./identityProviderGoogle";
import { IdentityProviderMicrosoft } from "./identityProviderMicrosoft";
import { IdentityProviderTwitter } from "./identityProviderTwitter";
import { Logger } from "./logger";
import { NamedValue } from "./namedValue";
import { OpenIdConnectProvider } from "./openIdConnectProvider";
import { Policy } from "./policy";
import { Product } from "./product";
import { ProductApi } from "./productApi";
import { ProductGroup } from "./productGroup";
import { ProductPolicy } from "./productPolicy";
import { Property } from "./property";
import { RedisCache } from "./redisCache";
import { Service } from "./service";
import { Subscription } from "./subscription";
import { Tag } from "./tag";
import { User } from "./user";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure:apimanagement/api:Api":
                return new Api(name, <any>undefined, { urn })
            case "azure:apimanagement/apiDiagnostic:ApiDiagnostic":
                return new ApiDiagnostic(name, <any>undefined, { urn })
            case "azure:apimanagement/apiOperation:ApiOperation":
                return new ApiOperation(name, <any>undefined, { urn })
            case "azure:apimanagement/apiOperationPolicy:ApiOperationPolicy":
                return new ApiOperationPolicy(name, <any>undefined, { urn })
            case "azure:apimanagement/apiOperationTag:ApiOperationTag":
                return new ApiOperationTag(name, <any>undefined, { urn })
            case "azure:apimanagement/apiPolicy:ApiPolicy":
                return new ApiPolicy(name, <any>undefined, { urn })
            case "azure:apimanagement/apiRelease:ApiRelease":
                return new ApiRelease(name, <any>undefined, { urn })
            case "azure:apimanagement/apiSchema:ApiSchema":
                return new ApiSchema(name, <any>undefined, { urn })
            case "azure:apimanagement/apiVersionSet:ApiVersionSet":
                return new ApiVersionSet(name, <any>undefined, { urn })
            case "azure:apimanagement/authorizationServer:AuthorizationServer":
                return new AuthorizationServer(name, <any>undefined, { urn })
            case "azure:apimanagement/backend:Backend":
                return new Backend(name, <any>undefined, { urn })
            case "azure:apimanagement/certificate:Certificate":
                return new Certificate(name, <any>undefined, { urn })
            case "azure:apimanagement/customDomain:CustomDomain":
                return new CustomDomain(name, <any>undefined, { urn })
            case "azure:apimanagement/diagnostic:Diagnostic":
                return new Diagnostic(name, <any>undefined, { urn })
            case "azure:apimanagement/emailTemplate:EmailTemplate":
                return new EmailTemplate(name, <any>undefined, { urn })
            case "azure:apimanagement/gateway:Gateway":
                return new Gateway(name, <any>undefined, { urn })
            case "azure:apimanagement/gatewayApi:GatewayApi":
                return new GatewayApi(name, <any>undefined, { urn })
            case "azure:apimanagement/group:Group":
                return new Group(name, <any>undefined, { urn })
            case "azure:apimanagement/groupUser:GroupUser":
                return new GroupUser(name, <any>undefined, { urn })
            case "azure:apimanagement/identityProviderAad:IdentityProviderAad":
                return new IdentityProviderAad(name, <any>undefined, { urn })
            case "azure:apimanagement/identityProviderAadb2c:IdentityProviderAadb2c":
                return new IdentityProviderAadb2c(name, <any>undefined, { urn })
            case "azure:apimanagement/identityProviderFacebook:IdentityProviderFacebook":
                return new IdentityProviderFacebook(name, <any>undefined, { urn })
            case "azure:apimanagement/identityProviderGoogle:IdentityProviderGoogle":
                return new IdentityProviderGoogle(name, <any>undefined, { urn })
            case "azure:apimanagement/identityProviderMicrosoft:IdentityProviderMicrosoft":
                return new IdentityProviderMicrosoft(name, <any>undefined, { urn })
            case "azure:apimanagement/identityProviderTwitter:IdentityProviderTwitter":
                return new IdentityProviderTwitter(name, <any>undefined, { urn })
            case "azure:apimanagement/logger:Logger":
                return new Logger(name, <any>undefined, { urn })
            case "azure:apimanagement/namedValue:NamedValue":
                return new NamedValue(name, <any>undefined, { urn })
            case "azure:apimanagement/openIdConnectProvider:OpenIdConnectProvider":
                return new OpenIdConnectProvider(name, <any>undefined, { urn })
            case "azure:apimanagement/policy:Policy":
                return new Policy(name, <any>undefined, { urn })
            case "azure:apimanagement/product:Product":
                return new Product(name, <any>undefined, { urn })
            case "azure:apimanagement/productApi:ProductApi":
                return new ProductApi(name, <any>undefined, { urn })
            case "azure:apimanagement/productGroup:ProductGroup":
                return new ProductGroup(name, <any>undefined, { urn })
            case "azure:apimanagement/productPolicy:ProductPolicy":
                return new ProductPolicy(name, <any>undefined, { urn })
            case "azure:apimanagement/property:Property":
                return new Property(name, <any>undefined, { urn })
            case "azure:apimanagement/redisCache:RedisCache":
                return new RedisCache(name, <any>undefined, { urn })
            case "azure:apimanagement/service:Service":
                return new Service(name, <any>undefined, { urn })
            case "azure:apimanagement/subscription:Subscription":
                return new Subscription(name, <any>undefined, { urn })
            case "azure:apimanagement/tag:Tag":
                return new Tag(name, <any>undefined, { urn })
            case "azure:apimanagement/user:User":
                return new User(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure", "apimanagement/api", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/apiDiagnostic", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/apiOperation", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/apiOperationPolicy", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/apiOperationTag", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/apiPolicy", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/apiRelease", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/apiSchema", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/apiVersionSet", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/authorizationServer", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/backend", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/certificate", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/customDomain", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/diagnostic", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/emailTemplate", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/gateway", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/gatewayApi", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/group", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/groupUser", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/identityProviderAad", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/identityProviderAadb2c", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/identityProviderFacebook", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/identityProviderGoogle", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/identityProviderMicrosoft", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/identityProviderTwitter", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/logger", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/namedValue", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/openIdConnectProvider", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/policy", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/product", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/productApi", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/productGroup", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/productPolicy", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/property", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/redisCache", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/service", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/subscription", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/tag", _module)
pulumi.runtime.registerResourceModule("azure", "apimanagement/user", _module)
