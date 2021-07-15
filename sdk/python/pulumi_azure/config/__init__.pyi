# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

auxiliaryTenantIds: Optional[str]

clientCertificatePassword: Optional[str]
"""
The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
Certificate
"""

clientCertificatePath: Optional[str]
"""
The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
Principal using a Client Certificate.
"""

clientId: Optional[str]
"""
The Client ID which should be used.
"""

clientSecret: Optional[str]
"""
The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
"""

disableCorrelationRequestId: Optional[str]
"""
This will disable the x-ms-correlation-request-id header.
"""

disableTerraformPartnerId: Optional[str]
"""
This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
"""

environment: str
"""
The Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to
public.
"""

features: Optional[str]

location: Optional[str]

metadataHost: str
"""
The Hostname which should be used for the Azure Metadata Service.
"""

metadataUrl: Optional[str]
"""
Deprecated - replaced by `metadata_host`.
"""

msiEndpoint: Optional[str]
"""
The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected
automatically.
"""

partnerId: Optional[str]
"""
A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
"""

skipCredentialsValidation: Optional[str]
"""
[DEPRECATED] This will cause the AzureRM Provider to skip verifying the credentials being used are valid.
"""

skipProviderRegistration: Optional[str]
"""
Should the AzureRM Provider skip registering all of the Resource Providers that it supports, if they're not already
registered?
"""

storageUseAzuread: Optional[str]
"""
Should the AzureRM Provider use AzureAD to access the Storage Data Plane API's?
"""

subscriptionId: Optional[str]
"""
The Subscription ID which should be used.
"""

tenantId: Optional[str]
"""
The Tenant ID which should be used.
"""

useMsi: Optional[str]
"""
Allowed Managed Service Identity be used for Authentication.
"""

