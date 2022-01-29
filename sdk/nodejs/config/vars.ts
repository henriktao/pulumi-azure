// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("azure");

export declare const auxiliaryTenantIds: string[] | undefined;
Object.defineProperty(exports, "auxiliaryTenantIds", {
    get() {
        return __config.getObject<string[]>("auxiliaryTenantIds");
    },
    enumerable: true,
});

/**
 * The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
 * Certificate
 */
export declare const clientCertificatePassword: string | undefined;
Object.defineProperty(exports, "clientCertificatePassword", {
    get() {
        return __config.get("clientCertificatePassword");
    },
    enumerable: true,
});

/**
 * The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
 * Principal using a Client Certificate.
 */
export declare const clientCertificatePath: string | undefined;
Object.defineProperty(exports, "clientCertificatePath", {
    get() {
        return __config.get("clientCertificatePath");
    },
    enumerable: true,
});

/**
 * The Client ID which should be used.
 */
export declare const clientId: string | undefined;
Object.defineProperty(exports, "clientId", {
    get() {
        return __config.get("clientId");
    },
    enumerable: true,
});

/**
 * The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
 */
export declare const clientSecret: string | undefined;
Object.defineProperty(exports, "clientSecret", {
    get() {
        return __config.get("clientSecret");
    },
    enumerable: true,
});

/**
 * This will disable the x-ms-correlation-request-id header.
 */
export declare const disableCorrelationRequestId: boolean | undefined;
Object.defineProperty(exports, "disableCorrelationRequestId", {
    get() {
        return __config.getObject<boolean>("disableCorrelationRequestId");
    },
    enumerable: true,
});

/**
 * This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
 */
export declare const disableTerraformPartnerId: boolean | undefined;
Object.defineProperty(exports, "disableTerraformPartnerId", {
    get() {
        return __config.getObject<boolean>("disableTerraformPartnerId");
    },
    enumerable: true,
});

/**
 * The Cloud Environment which should be used. Possible values are public, usgovernment, and china. Defaults to public.
 */
export declare const environment: string;
Object.defineProperty(exports, "environment", {
    get() {
        return __config.get("environment") ?? (utilities.getEnv("AZURE_ENVIRONMENT", "ARM_ENVIRONMENT") || "public");
    },
    enumerable: true,
});

export declare const features: outputs.config.Features | undefined;
Object.defineProperty(exports, "features", {
    get() {
        return __config.getObject<outputs.config.Features>("features");
    },
    enumerable: true,
});

export declare const location: string | undefined;
Object.defineProperty(exports, "location", {
    get() {
        return __config.get("location") ?? utilities.getEnv("ARM_LOCATION");
    },
    enumerable: true,
});

/**
 * The Hostname which should be used for the Azure Metadata Service.
 */
export declare const metadataHost: string | undefined;
Object.defineProperty(exports, "metadataHost", {
    get() {
        return __config.get("metadataHost") ?? utilities.getEnv("ARM_METADATA_HOSTNAME");
    },
    enumerable: true,
});

/**
 * Deprecated - replaced by `metadata_host`.
 */
export declare const metadataUrl: string | undefined;
Object.defineProperty(exports, "metadataUrl", {
    get() {
        return __config.get("metadataUrl") ?? utilities.getEnv("ARM_METADATA_URL");
    },
    enumerable: true,
});

/**
 * The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected
 * automatically.
 */
export declare const msiEndpoint: string | undefined;
Object.defineProperty(exports, "msiEndpoint", {
    get() {
        return __config.get("msiEndpoint");
    },
    enumerable: true,
});

/**
 * A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
 */
export declare const partnerId: string | undefined;
Object.defineProperty(exports, "partnerId", {
    get() {
        return __config.get("partnerId");
    },
    enumerable: true,
});

/**
 * [DEPRECATED] This will cause the AzureRM Provider to skip verifying the credentials being used are valid.
 */
export declare const skipCredentialsValidation: boolean | undefined;
Object.defineProperty(exports, "skipCredentialsValidation", {
    get() {
        return __config.getObject<boolean>("skipCredentialsValidation");
    },
    enumerable: true,
});

/**
 * Should the AzureRM Provider skip registering all of the Resource Providers that it supports, if they're not already
 * registered?
 */
export declare const skipProviderRegistration: boolean;
Object.defineProperty(exports, "skipProviderRegistration", {
    get() {
        return __config.getObject<boolean>("skipProviderRegistration") ?? (<any>utilities.getEnvBoolean("ARM_SKIP_PROVIDER_REGISTRATION") || false);
    },
    enumerable: true,
});

/**
 * Should the AzureRM Provider use AzureAD to access the Storage Data Plane API's?
 */
export declare const storageUseAzuread: boolean;
Object.defineProperty(exports, "storageUseAzuread", {
    get() {
        return __config.getObject<boolean>("storageUseAzuread") ?? (<any>utilities.getEnvBoolean("ARM_STORAGE_USE_AZUREAD") || false);
    },
    enumerable: true,
});

/**
 * The Subscription ID which should be used.
 */
export declare const subscriptionId: string;
Object.defineProperty(exports, "subscriptionId", {
    get() {
        return __config.get("subscriptionId") ?? (utilities.getEnv("ARM_SUBSCRIPTION_ID") || "");
    },
    enumerable: true,
});

/**
 * The Tenant ID which should be used.
 */
export declare const tenantId: string | undefined;
Object.defineProperty(exports, "tenantId", {
    get() {
        return __config.get("tenantId");
    },
    enumerable: true,
});

/**
 * Should Terraform obtain MSAL auth tokens and no longer use Azure Active Directory Graph?
 */
export declare const useMsal: boolean | undefined;
Object.defineProperty(exports, "useMsal", {
    get() {
        return __config.getObject<boolean>("useMsal");
    },
    enumerable: true,
});

/**
 * Allowed Managed Service Identity be used for Authentication.
 */
export declare const useMsi: boolean | undefined;
Object.defineProperty(exports, "useMsi", {
    get() {
        return __config.getObject<boolean>("useMsi");
    },
    enumerable: true,
});

