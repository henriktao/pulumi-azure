# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SynapseSparkArgs', 'SynapseSpark']

@pulumi.input_type
class SynapseSparkArgs:
    def __init__(__self__, *,
                 machine_learning_workspace_id: pulumi.Input[str],
                 synapse_spark_pool_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input['SynapseSparkIdentityArgs']] = None,
                 local_auth_enabled: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SynapseSpark resource.
        :param pulumi.Input[str] machine_learning_workspace_id: The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[str] synapse_spark_pool_id: The ID of the linked Synapse Spark Pool. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[str] description: The description of the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input['SynapseSparkIdentityArgs'] identity: A `identity` block as defined below. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[bool] local_auth_enabled: Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[str] location: The Azure Region where the Machine Learning Synapse Spark should exist. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[str] name: The name which should be used for this Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        pulumi.set(__self__, "machine_learning_workspace_id", machine_learning_workspace_id)
        pulumi.set(__self__, "synapse_spark_pool_id", synapse_spark_pool_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if local_auth_enabled is not None:
            pulumi.set(__self__, "local_auth_enabled", local_auth_enabled)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="machineLearningWorkspaceId")
    def machine_learning_workspace_id(self) -> pulumi.Input[str]:
        """
        The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "machine_learning_workspace_id")

    @machine_learning_workspace_id.setter
    def machine_learning_workspace_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "machine_learning_workspace_id", value)

    @property
    @pulumi.getter(name="synapseSparkPoolId")
    def synapse_spark_pool_id(self) -> pulumi.Input[str]:
        """
        The ID of the linked Synapse Spark Pool. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "synapse_spark_pool_id")

    @synapse_spark_pool_id.setter
    def synapse_spark_pool_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "synapse_spark_pool_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['SynapseSparkIdentityArgs']]:
        """
        A `identity` block as defined below. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['SynapseSparkIdentityArgs']]):
        pulumi.set(self, "identity", value)

    @property
    @pulumi.getter(name="localAuthEnabled")
    def local_auth_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "local_auth_enabled")

    @local_auth_enabled.setter
    def local_auth_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "local_auth_enabled", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure Region where the Machine Learning Synapse Spark should exist. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags which should be assigned to the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _SynapseSparkState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input['SynapseSparkIdentityArgs']] = None,
                 local_auth_enabled: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 machine_learning_workspace_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 synapse_spark_pool_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering SynapseSpark resources.
        :param pulumi.Input[str] description: The description of the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input['SynapseSparkIdentityArgs'] identity: A `identity` block as defined below. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[bool] local_auth_enabled: Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[str] location: The Azure Region where the Machine Learning Synapse Spark should exist. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[str] machine_learning_workspace_id: The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[str] name: The name which should be used for this Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[str] synapse_spark_pool_id: The ID of the linked Synapse Spark Pool. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if local_auth_enabled is not None:
            pulumi.set(__self__, "local_auth_enabled", local_auth_enabled)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if machine_learning_workspace_id is not None:
            pulumi.set(__self__, "machine_learning_workspace_id", machine_learning_workspace_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if synapse_spark_pool_id is not None:
            pulumi.set(__self__, "synapse_spark_pool_id", synapse_spark_pool_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['SynapseSparkIdentityArgs']]:
        """
        A `identity` block as defined below. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['SynapseSparkIdentityArgs']]):
        pulumi.set(self, "identity", value)

    @property
    @pulumi.getter(name="localAuthEnabled")
    def local_auth_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "local_auth_enabled")

    @local_auth_enabled.setter
    def local_auth_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "local_auth_enabled", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure Region where the Machine Learning Synapse Spark should exist. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="machineLearningWorkspaceId")
    def machine_learning_workspace_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "machine_learning_workspace_id")

    @machine_learning_workspace_id.setter
    def machine_learning_workspace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "machine_learning_workspace_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="synapseSparkPoolId")
    def synapse_spark_pool_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the linked Synapse Spark Pool. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "synapse_spark_pool_id")

    @synapse_spark_pool_id.setter
    def synapse_spark_pool_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "synapse_spark_pool_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags which should be assigned to the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class SynapseSpark(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['SynapseSparkIdentityArgs']]] = None,
                 local_auth_enabled: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 machine_learning_workspace_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 synapse_spark_pool_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manages a Machine Learning Synapse Spark.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup",
            location="west europe",
            tags={
                "stage": "example",
            })
        example_insights = azure.appinsights.Insights("exampleInsights",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            application_type="web")
        example_key_vault = azure.keyvault.KeyVault("exampleKeyVault",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tenant_id=current.tenant_id,
            sku_name="standard",
            purge_protection_enabled=True)
        example_account = azure.storage.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            account_tier="Standard",
            account_replication_type="LRS")
        example_workspace = azure.machinelearning.Workspace("exampleWorkspace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            application_insights_id=example_insights.id,
            key_vault_id=example_key_vault.id,
            storage_account_id=example_account.id,
            identity=azure.machinelearning.WorkspaceIdentityArgs(
                type="SystemAssigned",
            ))
        example_data_lake_gen2_filesystem = azure.storage.DataLakeGen2Filesystem("exampleDataLakeGen2Filesystem", storage_account_id=example_account.id)
        example_synapse_workspace_workspace = azure.synapse.Workspace("exampleSynapse/workspaceWorkspace",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            storage_data_lake_gen2_filesystem_id=example_data_lake_gen2_filesystem.id,
            sql_administrator_login="sqladminuser",
            sql_administrator_login_password="H@Sh1CoR3!")
        example_spark_pool = azure.synapse.SparkPool("exampleSparkPool",
            synapse_workspace_id=example_synapse / workspace_workspace["id"],
            node_size_family="MemoryOptimized",
            node_size="Small",
            node_count=3)
        example_synapse_spark = azure.machinelearning.SynapseSpark("exampleSynapseSpark",
            machine_learning_workspace_id=example_workspace.id,
            location=example_resource_group.location,
            synapse_spark_pool_id=example_spark_pool.id,
            identity=azure.machinelearning.SynapseSparkIdentityArgs(
                type="SystemAssigned",
            ))
        ```

        ## Import

        Machine Learning Synapse Sparks can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:machinelearning/synapseSpark:SynapseSpark example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/compute1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[pulumi.InputType['SynapseSparkIdentityArgs']] identity: A `identity` block as defined below. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[bool] local_auth_enabled: Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[str] location: The Azure Region where the Machine Learning Synapse Spark should exist. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[str] machine_learning_workspace_id: The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[str] name: The name which should be used for this Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[str] synapse_spark_pool_id: The ID of the linked Synapse Spark Pool. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SynapseSparkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Machine Learning Synapse Spark.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup",
            location="west europe",
            tags={
                "stage": "example",
            })
        example_insights = azure.appinsights.Insights("exampleInsights",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            application_type="web")
        example_key_vault = azure.keyvault.KeyVault("exampleKeyVault",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tenant_id=current.tenant_id,
            sku_name="standard",
            purge_protection_enabled=True)
        example_account = azure.storage.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            account_tier="Standard",
            account_replication_type="LRS")
        example_workspace = azure.machinelearning.Workspace("exampleWorkspace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            application_insights_id=example_insights.id,
            key_vault_id=example_key_vault.id,
            storage_account_id=example_account.id,
            identity=azure.machinelearning.WorkspaceIdentityArgs(
                type="SystemAssigned",
            ))
        example_data_lake_gen2_filesystem = azure.storage.DataLakeGen2Filesystem("exampleDataLakeGen2Filesystem", storage_account_id=example_account.id)
        example_synapse_workspace_workspace = azure.synapse.Workspace("exampleSynapse/workspaceWorkspace",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            storage_data_lake_gen2_filesystem_id=example_data_lake_gen2_filesystem.id,
            sql_administrator_login="sqladminuser",
            sql_administrator_login_password="H@Sh1CoR3!")
        example_spark_pool = azure.synapse.SparkPool("exampleSparkPool",
            synapse_workspace_id=example_synapse / workspace_workspace["id"],
            node_size_family="MemoryOptimized",
            node_size="Small",
            node_count=3)
        example_synapse_spark = azure.machinelearning.SynapseSpark("exampleSynapseSpark",
            machine_learning_workspace_id=example_workspace.id,
            location=example_resource_group.location,
            synapse_spark_pool_id=example_spark_pool.id,
            identity=azure.machinelearning.SynapseSparkIdentityArgs(
                type="SystemAssigned",
            ))
        ```

        ## Import

        Machine Learning Synapse Sparks can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:machinelearning/synapseSpark:SynapseSpark example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/compute1
        ```

        :param str resource_name: The name of the resource.
        :param SynapseSparkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SynapseSparkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['SynapseSparkIdentityArgs']]] = None,
                 local_auth_enabled: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 machine_learning_workspace_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 synapse_spark_pool_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SynapseSparkArgs.__new__(SynapseSparkArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["identity"] = identity
            __props__.__dict__["local_auth_enabled"] = local_auth_enabled
            __props__.__dict__["location"] = location
            if machine_learning_workspace_id is None and not opts.urn:
                raise TypeError("Missing required property 'machine_learning_workspace_id'")
            __props__.__dict__["machine_learning_workspace_id"] = machine_learning_workspace_id
            __props__.__dict__["name"] = name
            if synapse_spark_pool_id is None and not opts.urn:
                raise TypeError("Missing required property 'synapse_spark_pool_id'")
            __props__.__dict__["synapse_spark_pool_id"] = synapse_spark_pool_id
            __props__.__dict__["tags"] = tags
        super(SynapseSpark, __self__).__init__(
            'azure:machinelearning/synapseSpark:SynapseSpark',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            identity: Optional[pulumi.Input[pulumi.InputType['SynapseSparkIdentityArgs']]] = None,
            local_auth_enabled: Optional[pulumi.Input[bool]] = None,
            location: Optional[pulumi.Input[str]] = None,
            machine_learning_workspace_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            synapse_spark_pool_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'SynapseSpark':
        """
        Get an existing SynapseSpark resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[pulumi.InputType['SynapseSparkIdentityArgs']] identity: A `identity` block as defined below. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[bool] local_auth_enabled: Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[str] location: The Azure Region where the Machine Learning Synapse Spark should exist. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[str] machine_learning_workspace_id: The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[str] name: The name which should be used for this Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[str] synapse_spark_pool_id: The ID of the linked Synapse Spark Pool. Changing this forces a new Machine Learning Synapse Spark to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SynapseSparkState.__new__(_SynapseSparkState)

        __props__.__dict__["description"] = description
        __props__.__dict__["identity"] = identity
        __props__.__dict__["local_auth_enabled"] = local_auth_enabled
        __props__.__dict__["location"] = location
        __props__.__dict__["machine_learning_workspace_id"] = machine_learning_workspace_id
        __props__.__dict__["name"] = name
        __props__.__dict__["synapse_spark_pool_id"] = synapse_spark_pool_id
        __props__.__dict__["tags"] = tags
        return SynapseSpark(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.SynapseSparkIdentity']]:
        """
        A `identity` block as defined below. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="localAuthEnabled")
    def local_auth_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "local_auth_enabled")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The Azure Region where the Machine Learning Synapse Spark should exist. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="machineLearningWorkspaceId")
    def machine_learning_workspace_id(self) -> pulumi.Output[str]:
        """
        The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "machine_learning_workspace_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="synapseSparkPoolId")
    def synapse_spark_pool_id(self) -> pulumi.Output[str]:
        """
        The ID of the linked Synapse Spark Pool. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "synapse_spark_pool_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags which should be assigned to the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
        """
        return pulumi.get(self, "tags")

