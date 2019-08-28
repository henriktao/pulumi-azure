// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Queue within an Azure Storage Account.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_queue.html.markdown.
type Queue struct {
	s *pulumi.ResourceState
}

// NewQueue registers a new resource with the given unique name, arguments, and options.
func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOpt) (*Queue, error) {
	if args == nil || args.StorageAccountName == nil {
		return nil, errors.New("missing required argument 'StorageAccountName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["metadata"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["storageAccountName"] = nil
	} else {
		inputs["metadata"] = args.Metadata
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["storageAccountName"] = args.StorageAccountName
	}
	s, err := ctx.RegisterResource("azure:storage/queue:Queue", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Queue{s: s}, nil
}

// GetQueue gets an existing Queue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.ID, state *QueueState, opts ...pulumi.ResourceOpt) (*Queue, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["metadata"] = state.Metadata
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["storageAccountName"] = state.StorageAccountName
	}
	s, err := ctx.ReadResource("azure:storage/queue:Queue", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Queue{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Queue) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Queue) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// A mapping of MetaData which should be assigned to this Storage Queue.
func (r *Queue) Metadata() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["metadata"])
}

// The name of the Queue which should be created within the Storage Account. Must be unique within the storage account the queue is located.
func (r *Queue) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The name of the resource group in which to create the storage queue.
func (r *Queue) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// Specifies the Storage Account in which the Storage Queue should exist. Changing this forces a new resource to be created.
func (r *Queue) StorageAccountName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["storageAccountName"])
}

// Input properties used for looking up and filtering Queue resources.
type QueueState struct {
	// A mapping of MetaData which should be assigned to this Storage Queue.
	Metadata interface{}
	// The name of the Queue which should be created within the Storage Account. Must be unique within the storage account the queue is located.
	Name interface{}
	// The name of the resource group in which to create the storage queue.
	ResourceGroupName interface{}
	// Specifies the Storage Account in which the Storage Queue should exist. Changing this forces a new resource to be created.
	StorageAccountName interface{}
}

// The set of arguments for constructing a Queue resource.
type QueueArgs struct {
	// A mapping of MetaData which should be assigned to this Storage Queue.
	Metadata interface{}
	// The name of the Queue which should be created within the Storage Account. Must be unique within the storage account the queue is located.
	Name interface{}
	// The name of the resource group in which to create the storage queue.
	ResourceGroupName interface{}
	// Specifies the Storage Account in which the Storage Queue should exist. Changing this forces a new resource to be created.
	StorageAccountName interface{}
}
