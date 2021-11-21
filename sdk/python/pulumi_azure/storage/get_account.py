# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetAccountResult',
    'AwaitableGetAccountResult',
    'get_account',
    'get_account_output',
]

@pulumi.output_type
class GetAccountResult:
    """
    A collection of values returned by getAccount.
    """
    def __init__(__self__, access_tier=None, account_kind=None, account_replication_type=None, account_tier=None, allow_blob_public_access=None, custom_domains=None, enable_https_traffic_only=None, id=None, is_hns_enabled=None, location=None, min_tls_version=None, name=None, primary_access_key=None, primary_blob_connection_string=None, primary_blob_endpoint=None, primary_blob_host=None, primary_connection_string=None, primary_dfs_endpoint=None, primary_dfs_host=None, primary_file_endpoint=None, primary_file_host=None, primary_location=None, primary_queue_endpoint=None, primary_queue_host=None, primary_table_endpoint=None, primary_table_host=None, primary_web_endpoint=None, primary_web_host=None, queue_encryption_key_type=None, resource_group_name=None, secondary_access_key=None, secondary_blob_connection_string=None, secondary_blob_endpoint=None, secondary_blob_host=None, secondary_connection_string=None, secondary_dfs_endpoint=None, secondary_dfs_host=None, secondary_file_endpoint=None, secondary_file_host=None, secondary_location=None, secondary_queue_endpoint=None, secondary_queue_host=None, secondary_table_endpoint=None, secondary_table_host=None, secondary_web_endpoint=None, secondary_web_host=None, table_encryption_key_type=None, tags=None):
        if access_tier and not isinstance(access_tier, str):
            raise TypeError("Expected argument 'access_tier' to be a str")
        pulumi.set(__self__, "access_tier", access_tier)
        if account_kind and not isinstance(account_kind, str):
            raise TypeError("Expected argument 'account_kind' to be a str")
        pulumi.set(__self__, "account_kind", account_kind)
        if account_replication_type and not isinstance(account_replication_type, str):
            raise TypeError("Expected argument 'account_replication_type' to be a str")
        pulumi.set(__self__, "account_replication_type", account_replication_type)
        if account_tier and not isinstance(account_tier, str):
            raise TypeError("Expected argument 'account_tier' to be a str")
        pulumi.set(__self__, "account_tier", account_tier)
        if allow_blob_public_access and not isinstance(allow_blob_public_access, bool):
            raise TypeError("Expected argument 'allow_blob_public_access' to be a bool")
        pulumi.set(__self__, "allow_blob_public_access", allow_blob_public_access)
        if custom_domains and not isinstance(custom_domains, list):
            raise TypeError("Expected argument 'custom_domains' to be a list")
        pulumi.set(__self__, "custom_domains", custom_domains)
        if enable_https_traffic_only and not isinstance(enable_https_traffic_only, bool):
            raise TypeError("Expected argument 'enable_https_traffic_only' to be a bool")
        pulumi.set(__self__, "enable_https_traffic_only", enable_https_traffic_only)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_hns_enabled and not isinstance(is_hns_enabled, bool):
            raise TypeError("Expected argument 'is_hns_enabled' to be a bool")
        pulumi.set(__self__, "is_hns_enabled", is_hns_enabled)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if min_tls_version and not isinstance(min_tls_version, str):
            raise TypeError("Expected argument 'min_tls_version' to be a str")
        pulumi.set(__self__, "min_tls_version", min_tls_version)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if primary_access_key and not isinstance(primary_access_key, str):
            raise TypeError("Expected argument 'primary_access_key' to be a str")
        pulumi.set(__self__, "primary_access_key", primary_access_key)
        if primary_blob_connection_string and not isinstance(primary_blob_connection_string, str):
            raise TypeError("Expected argument 'primary_blob_connection_string' to be a str")
        pulumi.set(__self__, "primary_blob_connection_string", primary_blob_connection_string)
        if primary_blob_endpoint and not isinstance(primary_blob_endpoint, str):
            raise TypeError("Expected argument 'primary_blob_endpoint' to be a str")
        pulumi.set(__self__, "primary_blob_endpoint", primary_blob_endpoint)
        if primary_blob_host and not isinstance(primary_blob_host, str):
            raise TypeError("Expected argument 'primary_blob_host' to be a str")
        pulumi.set(__self__, "primary_blob_host", primary_blob_host)
        if primary_connection_string and not isinstance(primary_connection_string, str):
            raise TypeError("Expected argument 'primary_connection_string' to be a str")
        pulumi.set(__self__, "primary_connection_string", primary_connection_string)
        if primary_dfs_endpoint and not isinstance(primary_dfs_endpoint, str):
            raise TypeError("Expected argument 'primary_dfs_endpoint' to be a str")
        pulumi.set(__self__, "primary_dfs_endpoint", primary_dfs_endpoint)
        if primary_dfs_host and not isinstance(primary_dfs_host, str):
            raise TypeError("Expected argument 'primary_dfs_host' to be a str")
        pulumi.set(__self__, "primary_dfs_host", primary_dfs_host)
        if primary_file_endpoint and not isinstance(primary_file_endpoint, str):
            raise TypeError("Expected argument 'primary_file_endpoint' to be a str")
        pulumi.set(__self__, "primary_file_endpoint", primary_file_endpoint)
        if primary_file_host and not isinstance(primary_file_host, str):
            raise TypeError("Expected argument 'primary_file_host' to be a str")
        pulumi.set(__self__, "primary_file_host", primary_file_host)
        if primary_location and not isinstance(primary_location, str):
            raise TypeError("Expected argument 'primary_location' to be a str")
        pulumi.set(__self__, "primary_location", primary_location)
        if primary_queue_endpoint and not isinstance(primary_queue_endpoint, str):
            raise TypeError("Expected argument 'primary_queue_endpoint' to be a str")
        pulumi.set(__self__, "primary_queue_endpoint", primary_queue_endpoint)
        if primary_queue_host and not isinstance(primary_queue_host, str):
            raise TypeError("Expected argument 'primary_queue_host' to be a str")
        pulumi.set(__self__, "primary_queue_host", primary_queue_host)
        if primary_table_endpoint and not isinstance(primary_table_endpoint, str):
            raise TypeError("Expected argument 'primary_table_endpoint' to be a str")
        pulumi.set(__self__, "primary_table_endpoint", primary_table_endpoint)
        if primary_table_host and not isinstance(primary_table_host, str):
            raise TypeError("Expected argument 'primary_table_host' to be a str")
        pulumi.set(__self__, "primary_table_host", primary_table_host)
        if primary_web_endpoint and not isinstance(primary_web_endpoint, str):
            raise TypeError("Expected argument 'primary_web_endpoint' to be a str")
        pulumi.set(__self__, "primary_web_endpoint", primary_web_endpoint)
        if primary_web_host and not isinstance(primary_web_host, str):
            raise TypeError("Expected argument 'primary_web_host' to be a str")
        pulumi.set(__self__, "primary_web_host", primary_web_host)
        if queue_encryption_key_type and not isinstance(queue_encryption_key_type, str):
            raise TypeError("Expected argument 'queue_encryption_key_type' to be a str")
        pulumi.set(__self__, "queue_encryption_key_type", queue_encryption_key_type)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if secondary_access_key and not isinstance(secondary_access_key, str):
            raise TypeError("Expected argument 'secondary_access_key' to be a str")
        pulumi.set(__self__, "secondary_access_key", secondary_access_key)
        if secondary_blob_connection_string and not isinstance(secondary_blob_connection_string, str):
            raise TypeError("Expected argument 'secondary_blob_connection_string' to be a str")
        pulumi.set(__self__, "secondary_blob_connection_string", secondary_blob_connection_string)
        if secondary_blob_endpoint and not isinstance(secondary_blob_endpoint, str):
            raise TypeError("Expected argument 'secondary_blob_endpoint' to be a str")
        pulumi.set(__self__, "secondary_blob_endpoint", secondary_blob_endpoint)
        if secondary_blob_host and not isinstance(secondary_blob_host, str):
            raise TypeError("Expected argument 'secondary_blob_host' to be a str")
        pulumi.set(__self__, "secondary_blob_host", secondary_blob_host)
        if secondary_connection_string and not isinstance(secondary_connection_string, str):
            raise TypeError("Expected argument 'secondary_connection_string' to be a str")
        pulumi.set(__self__, "secondary_connection_string", secondary_connection_string)
        if secondary_dfs_endpoint and not isinstance(secondary_dfs_endpoint, str):
            raise TypeError("Expected argument 'secondary_dfs_endpoint' to be a str")
        pulumi.set(__self__, "secondary_dfs_endpoint", secondary_dfs_endpoint)
        if secondary_dfs_host and not isinstance(secondary_dfs_host, str):
            raise TypeError("Expected argument 'secondary_dfs_host' to be a str")
        pulumi.set(__self__, "secondary_dfs_host", secondary_dfs_host)
        if secondary_file_endpoint and not isinstance(secondary_file_endpoint, str):
            raise TypeError("Expected argument 'secondary_file_endpoint' to be a str")
        pulumi.set(__self__, "secondary_file_endpoint", secondary_file_endpoint)
        if secondary_file_host and not isinstance(secondary_file_host, str):
            raise TypeError("Expected argument 'secondary_file_host' to be a str")
        pulumi.set(__self__, "secondary_file_host", secondary_file_host)
        if secondary_location and not isinstance(secondary_location, str):
            raise TypeError("Expected argument 'secondary_location' to be a str")
        pulumi.set(__self__, "secondary_location", secondary_location)
        if secondary_queue_endpoint and not isinstance(secondary_queue_endpoint, str):
            raise TypeError("Expected argument 'secondary_queue_endpoint' to be a str")
        pulumi.set(__self__, "secondary_queue_endpoint", secondary_queue_endpoint)
        if secondary_queue_host and not isinstance(secondary_queue_host, str):
            raise TypeError("Expected argument 'secondary_queue_host' to be a str")
        pulumi.set(__self__, "secondary_queue_host", secondary_queue_host)
        if secondary_table_endpoint and not isinstance(secondary_table_endpoint, str):
            raise TypeError("Expected argument 'secondary_table_endpoint' to be a str")
        pulumi.set(__self__, "secondary_table_endpoint", secondary_table_endpoint)
        if secondary_table_host and not isinstance(secondary_table_host, str):
            raise TypeError("Expected argument 'secondary_table_host' to be a str")
        pulumi.set(__self__, "secondary_table_host", secondary_table_host)
        if secondary_web_endpoint and not isinstance(secondary_web_endpoint, str):
            raise TypeError("Expected argument 'secondary_web_endpoint' to be a str")
        pulumi.set(__self__, "secondary_web_endpoint", secondary_web_endpoint)
        if secondary_web_host and not isinstance(secondary_web_host, str):
            raise TypeError("Expected argument 'secondary_web_host' to be a str")
        pulumi.set(__self__, "secondary_web_host", secondary_web_host)
        if table_encryption_key_type and not isinstance(table_encryption_key_type, str):
            raise TypeError("Expected argument 'table_encryption_key_type' to be a str")
        pulumi.set(__self__, "table_encryption_key_type", table_encryption_key_type)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accessTier")
    def access_tier(self) -> str:
        """
        The access tier for `BlobStorage` accounts.
        """
        return pulumi.get(self, "access_tier")

    @property
    @pulumi.getter(name="accountKind")
    def account_kind(self) -> str:
        """
        The Kind of account.
        """
        return pulumi.get(self, "account_kind")

    @property
    @pulumi.getter(name="accountReplicationType")
    def account_replication_type(self) -> str:
        """
        The type of replication used for this storage account.
        """
        return pulumi.get(self, "account_replication_type")

    @property
    @pulumi.getter(name="accountTier")
    def account_tier(self) -> str:
        """
        The Tier of this storage account.
        """
        return pulumi.get(self, "account_tier")

    @property
    @pulumi.getter(name="allowBlobPublicAccess")
    def allow_blob_public_access(self) -> bool:
        """
        Is public access allowed to all blobs or containers in the storage account?
        """
        return pulumi.get(self, "allow_blob_public_access")

    @property
    @pulumi.getter(name="customDomains")
    def custom_domains(self) -> Sequence['outputs.GetAccountCustomDomainResult']:
        """
        A `custom_domain` block as documented below.
        """
        return pulumi.get(self, "custom_domains")

    @property
    @pulumi.getter(name="enableHttpsTrafficOnly")
    def enable_https_traffic_only(self) -> bool:
        """
        Is traffic only allowed via HTTPS? See [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
        for more information.
        """
        return pulumi.get(self, "enable_https_traffic_only")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isHnsEnabled")
    def is_hns_enabled(self) -> bool:
        """
        Is Hierarchical Namespace enabled?
        """
        return pulumi.get(self, "is_hns_enabled")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The Azure location where the Storage Account exists
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="minTlsVersion")
    def min_tls_version(self) -> Optional[str]:
        """
        The minimum supported TLS version for this storage account.
        """
        return pulumi.get(self, "min_tls_version")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The Custom Domain Name used for the Storage Account.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="primaryAccessKey")
    def primary_access_key(self) -> str:
        """
        The primary access key for the Storage Account.
        """
        return pulumi.get(self, "primary_access_key")

    @property
    @pulumi.getter(name="primaryBlobConnectionString")
    def primary_blob_connection_string(self) -> str:
        """
        The connection string associated with the primary blob location
        """
        return pulumi.get(self, "primary_blob_connection_string")

    @property
    @pulumi.getter(name="primaryBlobEndpoint")
    def primary_blob_endpoint(self) -> str:
        """
        The endpoint URL for blob storage in the primary location.
        """
        return pulumi.get(self, "primary_blob_endpoint")

    @property
    @pulumi.getter(name="primaryBlobHost")
    def primary_blob_host(self) -> str:
        """
        The hostname with port if applicable for blob storage in the primary location.
        """
        return pulumi.get(self, "primary_blob_host")

    @property
    @pulumi.getter(name="primaryConnectionString")
    def primary_connection_string(self) -> str:
        """
        The connection string associated with the primary location
        """
        return pulumi.get(self, "primary_connection_string")

    @property
    @pulumi.getter(name="primaryDfsEndpoint")
    def primary_dfs_endpoint(self) -> str:
        """
        The endpoint URL for DFS storage in the primary location.
        """
        return pulumi.get(self, "primary_dfs_endpoint")

    @property
    @pulumi.getter(name="primaryDfsHost")
    def primary_dfs_host(self) -> str:
        """
        The hostname with port if applicable for DFS storage in the primary location.
        """
        return pulumi.get(self, "primary_dfs_host")

    @property
    @pulumi.getter(name="primaryFileEndpoint")
    def primary_file_endpoint(self) -> str:
        """
        The endpoint URL for file storage in the primary location.
        """
        return pulumi.get(self, "primary_file_endpoint")

    @property
    @pulumi.getter(name="primaryFileHost")
    def primary_file_host(self) -> str:
        """
        The hostname with port if applicable for file storage in the primary location.
        """
        return pulumi.get(self, "primary_file_host")

    @property
    @pulumi.getter(name="primaryLocation")
    def primary_location(self) -> str:
        """
        The primary location of the Storage Account.
        """
        return pulumi.get(self, "primary_location")

    @property
    @pulumi.getter(name="primaryQueueEndpoint")
    def primary_queue_endpoint(self) -> str:
        """
        The endpoint URL for queue storage in the primary location.
        """
        return pulumi.get(self, "primary_queue_endpoint")

    @property
    @pulumi.getter(name="primaryQueueHost")
    def primary_queue_host(self) -> str:
        """
        The hostname with port if applicable for queue storage in the primary location.
        """
        return pulumi.get(self, "primary_queue_host")

    @property
    @pulumi.getter(name="primaryTableEndpoint")
    def primary_table_endpoint(self) -> str:
        """
        The endpoint URL for table storage in the primary location.
        """
        return pulumi.get(self, "primary_table_endpoint")

    @property
    @pulumi.getter(name="primaryTableHost")
    def primary_table_host(self) -> str:
        """
        The hostname with port if applicable for table storage in the primary location.
        """
        return pulumi.get(self, "primary_table_host")

    @property
    @pulumi.getter(name="primaryWebEndpoint")
    def primary_web_endpoint(self) -> str:
        """
        The endpoint URL for web storage in the primary location.
        """
        return pulumi.get(self, "primary_web_endpoint")

    @property
    @pulumi.getter(name="primaryWebHost")
    def primary_web_host(self) -> str:
        """
        The hostname with port if applicable for web storage in the primary location.
        """
        return pulumi.get(self, "primary_web_host")

    @property
    @pulumi.getter(name="queueEncryptionKeyType")
    def queue_encryption_key_type(self) -> str:
        """
        The encryption key type of the queue.
        """
        return pulumi.get(self, "queue_encryption_key_type")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="secondaryAccessKey")
    def secondary_access_key(self) -> str:
        """
        The secondary access key for the Storage Account.
        """
        return pulumi.get(self, "secondary_access_key")

    @property
    @pulumi.getter(name="secondaryBlobConnectionString")
    def secondary_blob_connection_string(self) -> str:
        """
        The connection string associated with the secondary blob location
        """
        return pulumi.get(self, "secondary_blob_connection_string")

    @property
    @pulumi.getter(name="secondaryBlobEndpoint")
    def secondary_blob_endpoint(self) -> str:
        """
        The endpoint URL for blob storage in the secondary location.
        """
        return pulumi.get(self, "secondary_blob_endpoint")

    @property
    @pulumi.getter(name="secondaryBlobHost")
    def secondary_blob_host(self) -> str:
        """
        The hostname with port if applicable for blob storage in the secondary location.
        """
        return pulumi.get(self, "secondary_blob_host")

    @property
    @pulumi.getter(name="secondaryConnectionString")
    def secondary_connection_string(self) -> str:
        """
        The connection string associated with the secondary location
        """
        return pulumi.get(self, "secondary_connection_string")

    @property
    @pulumi.getter(name="secondaryDfsEndpoint")
    def secondary_dfs_endpoint(self) -> str:
        """
        The endpoint URL for DFS storage in the secondary location.
        """
        return pulumi.get(self, "secondary_dfs_endpoint")

    @property
    @pulumi.getter(name="secondaryDfsHost")
    def secondary_dfs_host(self) -> str:
        """
        The hostname with port if applicable for DFS storage in the secondary location.
        """
        return pulumi.get(self, "secondary_dfs_host")

    @property
    @pulumi.getter(name="secondaryFileEndpoint")
    def secondary_file_endpoint(self) -> str:
        """
        The endpoint URL for file storage in the secondary location.
        """
        return pulumi.get(self, "secondary_file_endpoint")

    @property
    @pulumi.getter(name="secondaryFileHost")
    def secondary_file_host(self) -> str:
        """
        The hostname with port if applicable for file storage in the secondary location.
        """
        return pulumi.get(self, "secondary_file_host")

    @property
    @pulumi.getter(name="secondaryLocation")
    def secondary_location(self) -> str:
        """
        The secondary location of the Storage Account.
        """
        return pulumi.get(self, "secondary_location")

    @property
    @pulumi.getter(name="secondaryQueueEndpoint")
    def secondary_queue_endpoint(self) -> str:
        """
        The endpoint URL for queue storage in the secondary location.
        """
        return pulumi.get(self, "secondary_queue_endpoint")

    @property
    @pulumi.getter(name="secondaryQueueHost")
    def secondary_queue_host(self) -> str:
        """
        The hostname with port if applicable for queue storage in the secondary location.
        """
        return pulumi.get(self, "secondary_queue_host")

    @property
    @pulumi.getter(name="secondaryTableEndpoint")
    def secondary_table_endpoint(self) -> str:
        """
        The endpoint URL for table storage in the secondary location.
        """
        return pulumi.get(self, "secondary_table_endpoint")

    @property
    @pulumi.getter(name="secondaryTableHost")
    def secondary_table_host(self) -> str:
        """
        The hostname with port if applicable for table storage in the secondary location.
        """
        return pulumi.get(self, "secondary_table_host")

    @property
    @pulumi.getter(name="secondaryWebEndpoint")
    def secondary_web_endpoint(self) -> str:
        """
        The endpoint URL for web storage in the secondary location.
        """
        return pulumi.get(self, "secondary_web_endpoint")

    @property
    @pulumi.getter(name="secondaryWebHost")
    def secondary_web_host(self) -> str:
        """
        The hostname with port if applicable for web storage in the secondary location.
        """
        return pulumi.get(self, "secondary_web_host")

    @property
    @pulumi.getter(name="tableEncryptionKeyType")
    def table_encryption_key_type(self) -> str:
        """
        The encryption key type of the table.
        """
        return pulumi.get(self, "table_encryption_key_type")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A mapping of tags to assigned to the resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetAccountResult(GetAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountResult(
            access_tier=self.access_tier,
            account_kind=self.account_kind,
            account_replication_type=self.account_replication_type,
            account_tier=self.account_tier,
            allow_blob_public_access=self.allow_blob_public_access,
            custom_domains=self.custom_domains,
            enable_https_traffic_only=self.enable_https_traffic_only,
            id=self.id,
            is_hns_enabled=self.is_hns_enabled,
            location=self.location,
            min_tls_version=self.min_tls_version,
            name=self.name,
            primary_access_key=self.primary_access_key,
            primary_blob_connection_string=self.primary_blob_connection_string,
            primary_blob_endpoint=self.primary_blob_endpoint,
            primary_blob_host=self.primary_blob_host,
            primary_connection_string=self.primary_connection_string,
            primary_dfs_endpoint=self.primary_dfs_endpoint,
            primary_dfs_host=self.primary_dfs_host,
            primary_file_endpoint=self.primary_file_endpoint,
            primary_file_host=self.primary_file_host,
            primary_location=self.primary_location,
            primary_queue_endpoint=self.primary_queue_endpoint,
            primary_queue_host=self.primary_queue_host,
            primary_table_endpoint=self.primary_table_endpoint,
            primary_table_host=self.primary_table_host,
            primary_web_endpoint=self.primary_web_endpoint,
            primary_web_host=self.primary_web_host,
            queue_encryption_key_type=self.queue_encryption_key_type,
            resource_group_name=self.resource_group_name,
            secondary_access_key=self.secondary_access_key,
            secondary_blob_connection_string=self.secondary_blob_connection_string,
            secondary_blob_endpoint=self.secondary_blob_endpoint,
            secondary_blob_host=self.secondary_blob_host,
            secondary_connection_string=self.secondary_connection_string,
            secondary_dfs_endpoint=self.secondary_dfs_endpoint,
            secondary_dfs_host=self.secondary_dfs_host,
            secondary_file_endpoint=self.secondary_file_endpoint,
            secondary_file_host=self.secondary_file_host,
            secondary_location=self.secondary_location,
            secondary_queue_endpoint=self.secondary_queue_endpoint,
            secondary_queue_host=self.secondary_queue_host,
            secondary_table_endpoint=self.secondary_table_endpoint,
            secondary_table_host=self.secondary_table_host,
            secondary_web_endpoint=self.secondary_web_endpoint,
            secondary_web_host=self.secondary_web_host,
            table_encryption_key_type=self.table_encryption_key_type,
            tags=self.tags)


def get_account(min_tls_version: Optional[str] = None,
                name: Optional[str] = None,
                resource_group_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccountResult:
    """
    Use this data source to access information about an existing Storage Account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.storage.get_account(name="packerimages",
        resource_group_name="packer-storage")
    pulumi.export("storageAccountTier", example.account_tier)
    ```


    :param str min_tls_version: The minimum supported TLS version for this storage account.
    :param str name: Specifies the name of the Storage Account
    :param str resource_group_name: Specifies the name of the resource group the Storage Account is located in.
    """
    __args__ = dict()
    __args__['minTlsVersion'] = min_tls_version
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:storage/getAccount:getAccount', __args__, opts=opts, typ=GetAccountResult).value

    return AwaitableGetAccountResult(
        access_tier=__ret__.access_tier,
        account_kind=__ret__.account_kind,
        account_replication_type=__ret__.account_replication_type,
        account_tier=__ret__.account_tier,
        allow_blob_public_access=__ret__.allow_blob_public_access,
        custom_domains=__ret__.custom_domains,
        enable_https_traffic_only=__ret__.enable_https_traffic_only,
        id=__ret__.id,
        is_hns_enabled=__ret__.is_hns_enabled,
        location=__ret__.location,
        min_tls_version=__ret__.min_tls_version,
        name=__ret__.name,
        primary_access_key=__ret__.primary_access_key,
        primary_blob_connection_string=__ret__.primary_blob_connection_string,
        primary_blob_endpoint=__ret__.primary_blob_endpoint,
        primary_blob_host=__ret__.primary_blob_host,
        primary_connection_string=__ret__.primary_connection_string,
        primary_dfs_endpoint=__ret__.primary_dfs_endpoint,
        primary_dfs_host=__ret__.primary_dfs_host,
        primary_file_endpoint=__ret__.primary_file_endpoint,
        primary_file_host=__ret__.primary_file_host,
        primary_location=__ret__.primary_location,
        primary_queue_endpoint=__ret__.primary_queue_endpoint,
        primary_queue_host=__ret__.primary_queue_host,
        primary_table_endpoint=__ret__.primary_table_endpoint,
        primary_table_host=__ret__.primary_table_host,
        primary_web_endpoint=__ret__.primary_web_endpoint,
        primary_web_host=__ret__.primary_web_host,
        queue_encryption_key_type=__ret__.queue_encryption_key_type,
        resource_group_name=__ret__.resource_group_name,
        secondary_access_key=__ret__.secondary_access_key,
        secondary_blob_connection_string=__ret__.secondary_blob_connection_string,
        secondary_blob_endpoint=__ret__.secondary_blob_endpoint,
        secondary_blob_host=__ret__.secondary_blob_host,
        secondary_connection_string=__ret__.secondary_connection_string,
        secondary_dfs_endpoint=__ret__.secondary_dfs_endpoint,
        secondary_dfs_host=__ret__.secondary_dfs_host,
        secondary_file_endpoint=__ret__.secondary_file_endpoint,
        secondary_file_host=__ret__.secondary_file_host,
        secondary_location=__ret__.secondary_location,
        secondary_queue_endpoint=__ret__.secondary_queue_endpoint,
        secondary_queue_host=__ret__.secondary_queue_host,
        secondary_table_endpoint=__ret__.secondary_table_endpoint,
        secondary_table_host=__ret__.secondary_table_host,
        secondary_web_endpoint=__ret__.secondary_web_endpoint,
        secondary_web_host=__ret__.secondary_web_host,
        table_encryption_key_type=__ret__.table_encryption_key_type,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_account)
def get_account_output(min_tls_version: Optional[pulumi.Input[Optional[str]]] = None,
                       name: Optional[pulumi.Input[str]] = None,
                       resource_group_name: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccountResult]:
    """
    Use this data source to access information about an existing Storage Account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.storage.get_account(name="packerimages",
        resource_group_name="packer-storage")
    pulumi.export("storageAccountTier", example.account_tier)
    ```


    :param str min_tls_version: The minimum supported TLS version for this storage account.
    :param str name: Specifies the name of the Storage Account
    :param str resource_group_name: Specifies the name of the resource group the Storage Account is located in.
    """
    ...
