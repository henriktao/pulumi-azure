# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .attached_database_configuration import *
from .cluster import *
from .cluster_customer_managed_key import *
from .cluster_principal_assignment import *
from .database import *
from .database_principal import *
from .database_principal_assignment import *
from .event_grid_data_connection import *
from .eventhub_data_connection import *
from .get_cluster import *
from .iot_hub_data_connection import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure:kusto/attachedDatabaseConfiguration:AttachedDatabaseConfiguration":
                return AttachedDatabaseConfiguration(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:kusto/cluster:Cluster":
                return Cluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:kusto/clusterCustomerManagedKey:ClusterCustomerManagedKey":
                return ClusterCustomerManagedKey(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:kusto/clusterPrincipalAssignment:ClusterPrincipalAssignment":
                return ClusterPrincipalAssignment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:kusto/database:Database":
                return Database(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:kusto/databasePrincipal:DatabasePrincipal":
                return DatabasePrincipal(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:kusto/databasePrincipalAssignment:DatabasePrincipalAssignment":
                return DatabasePrincipalAssignment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:kusto/eventGridDataConnection:EventGridDataConnection":
                return EventGridDataConnection(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:kusto/eventhubDataConnection:EventhubDataConnection":
                return EventhubDataConnection(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:kusto/iotHubDataConnection:IotHubDataConnection":
                return IotHubDataConnection(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure", "kusto/attachedDatabaseConfiguration", _module_instance)
    pulumi.runtime.register_resource_module("azure", "kusto/cluster", _module_instance)
    pulumi.runtime.register_resource_module("azure", "kusto/clusterCustomerManagedKey", _module_instance)
    pulumi.runtime.register_resource_module("azure", "kusto/clusterPrincipalAssignment", _module_instance)
    pulumi.runtime.register_resource_module("azure", "kusto/database", _module_instance)
    pulumi.runtime.register_resource_module("azure", "kusto/databasePrincipal", _module_instance)
    pulumi.runtime.register_resource_module("azure", "kusto/databasePrincipalAssignment", _module_instance)
    pulumi.runtime.register_resource_module("azure", "kusto/eventGridDataConnection", _module_instance)
    pulumi.runtime.register_resource_module("azure", "kusto/eventhubDataConnection", _module_instance)
    pulumi.runtime.register_resource_module("azure", "kusto/iotHubDataConnection", _module_instance)

_register_module()
