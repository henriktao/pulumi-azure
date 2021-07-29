// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package domainservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about an Active Directory Domain Service.
//
// > **Supported Modes:** At present this data source only supports **User Forest** mode and _not_ **Resource Forest** mode. [Read more](https://docs.microsoft.com/en-us/azure/active-directory-domain-services/concepts-resource-forest) about the different operation modes for this service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/domainservices"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := domainservices.LookupService(ctx, &domainservices.LookupServiceArgs{
// 			Name:              "example-aadds",
// 			ResourceGroupName: "example-aadds-rg",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure:domainservices/getService:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getService.
type LookupServiceArgs struct {
	// The display name for your managed Active Directory Domain Service resource. Changing this forces a new resource to be created.
	Name string `pulumi:"name"`
	// The name of the Resource Group in which the Domain Service should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getService.
type LookupServiceResult struct {
	// A unique ID for the managed domain deployment.
	DeploymentId string `pulumi:"deploymentId"`
	// The forest type used by the managed domain. One of `ResourceTrusting`, for a _Resource Forest_, or blank, for a _User Forest_.
	DomainConfigurationType string `pulumi:"domainConfigurationType"`
	// The Active Directory domain of the Domain Service. See [official documentation](https://docs.microsoft.com/en-us/azure/active-directory-domain-services/tutorial-create-instance#create-a-managed-domain) for constraints and recommendations.
	DomainName string `pulumi:"domainName"`
	// Whether group-based filtered sync (also called scoped synchronisation) is enabled.
	FilteredSyncEnabled bool `pulumi:"filteredSyncEnabled"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Azure location in which the replica set resides.
	Location string `pulumi:"location"`
	Name     string `pulumi:"name"`
	// A `notifications` block as defined below.
	Notifications     []GetServiceNotification `pulumi:"notifications"`
	ReplicaSets       []GetServiceReplicaSet   `pulumi:"replicaSets"`
	ResourceGroupName string                   `pulumi:"resourceGroupName"`
	// A `secureLdap` block as defined below.
	SecureLdaps []GetServiceSecureLdap `pulumi:"secureLdaps"`
	// A `security` block as defined below.
	Securities []GetServiceSecurity `pulumi:"securities"`
	// The SKU of the Domain Service resource. One of `Standard`, `Enterprise` or `Premium`.
	Sku       string `pulumi:"sku"`
	SyncOwner string `pulumi:"syncOwner"`
	// A mapping of tags assigned to the resource.
	Tags     map[string]string `pulumi:"tags"`
	TenantId string            `pulumi:"tenantId"`
	Version  int               `pulumi:"version"`
}
