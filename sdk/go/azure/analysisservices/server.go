// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package analysisservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Analysis Services Server.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/analysisservices"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		rg, err := core.NewResourceGroup(ctx, "rg", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = analysisservices.NewServer(ctx, "server", &analysisservices.ServerArgs{
// 			Location:          pulumi.String("northeurope"),
// 			ResourceGroupName: rg.Name,
// 			Sku:               pulumi.String("S0"),
// 			AdminUsers: pulumi.StringArray{
// 				pulumi.String("myuser@domain.tld"),
// 			},
// 			EnablePowerBiService: pulumi.Bool(true),
// 			Ipv4FirewallRules: analysisservices.ServerIpv4FirewallRuleArray{
// 				&analysisservices.ServerIpv4FirewallRuleArgs{
// 					Name:       pulumi.String("myRule1"),
// 					RangeStart: pulumi.String("210.117.252.0"),
// 					RangeEnd:   pulumi.String("210.117.252.255"),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"abc": pulumi.String("123"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// > **NOTE:** The server resource will automatically be started and stopped during an update if it is in `paused` state.
//
// ## Import
//
// Analysis Services Server can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:analysisservices/server:Server server /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AnalysisServices/servers/server1
// ```
type Server struct {
	pulumi.CustomResourceState

	// List of email addresses of admin users.
	AdminUsers pulumi.StringArrayOutput `pulumi:"adminUsers"`
	// URI and SAS token for a blob container to store backups.
	BackupBlobContainerUri pulumi.StringPtrOutput `pulumi:"backupBlobContainerUri"`
	// Indicates if the Power BI service is allowed to access or not.
	EnablePowerBiService pulumi.BoolPtrOutput `pulumi:"enablePowerBiService"`
	// One or more `ipv4FirewallRule` block(s) as defined below.
	Ipv4FirewallRules ServerIpv4FirewallRuleArrayOutput `pulumi:"ipv4FirewallRules"`
	// The Azure location where the Analysis Services Server exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the firewall rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// Controls how the read-write server is used in the query pool. If this value is set to `All` then read-write servers are also used for queries. Otherwise with `ReadOnly` these servers do not participate in query operations.
	QuerypoolConnectionMode pulumi.StringOutput `pulumi:"querypoolConnectionMode"`
	// The name of the Resource Group in which the Analysis Services Server should be exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The full name of the Analysis Services Server.
	ServerFullName pulumi.StringOutput `pulumi:"serverFullName"`
	// SKU for the Analysis Services Server. Possible values are: `D1`, `B1`, `B2`, `S0`, `S1`, `S2`, `S4`, `S8`, `S9`, `S8v2` and `S9v2`.
	Sku  pulumi.StringOutput    `pulumi:"sku"`
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewServer registers a new resource with the given unique name, arguments, and options.
func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOption) (*Server, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	var resource Server
	err := ctx.RegisterResource("azure:analysisservices/server:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServer gets an existing Server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("azure:analysisservices/server:Server", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Server resources.
type serverState struct {
	// List of email addresses of admin users.
	AdminUsers []string `pulumi:"adminUsers"`
	// URI and SAS token for a blob container to store backups.
	BackupBlobContainerUri *string `pulumi:"backupBlobContainerUri"`
	// Indicates if the Power BI service is allowed to access or not.
	EnablePowerBiService *bool `pulumi:"enablePowerBiService"`
	// One or more `ipv4FirewallRule` block(s) as defined below.
	Ipv4FirewallRules []ServerIpv4FirewallRule `pulumi:"ipv4FirewallRules"`
	// The Azure location where the Analysis Services Server exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the firewall rule.
	Name *string `pulumi:"name"`
	// Controls how the read-write server is used in the query pool. If this value is set to `All` then read-write servers are also used for queries. Otherwise with `ReadOnly` these servers do not participate in query operations.
	QuerypoolConnectionMode *string `pulumi:"querypoolConnectionMode"`
	// The name of the Resource Group in which the Analysis Services Server should be exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The full name of the Analysis Services Server.
	ServerFullName *string `pulumi:"serverFullName"`
	// SKU for the Analysis Services Server. Possible values are: `D1`, `B1`, `B2`, `S0`, `S1`, `S2`, `S4`, `S8`, `S9`, `S8v2` and `S9v2`.
	Sku  *string           `pulumi:"sku"`
	Tags map[string]string `pulumi:"tags"`
}

type ServerState struct {
	// List of email addresses of admin users.
	AdminUsers pulumi.StringArrayInput
	// URI and SAS token for a blob container to store backups.
	BackupBlobContainerUri pulumi.StringPtrInput
	// Indicates if the Power BI service is allowed to access or not.
	EnablePowerBiService pulumi.BoolPtrInput
	// One or more `ipv4FirewallRule` block(s) as defined below.
	Ipv4FirewallRules ServerIpv4FirewallRuleArrayInput
	// The Azure location where the Analysis Services Server exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the firewall rule.
	Name pulumi.StringPtrInput
	// Controls how the read-write server is used in the query pool. If this value is set to `All` then read-write servers are also used for queries. Otherwise with `ReadOnly` these servers do not participate in query operations.
	QuerypoolConnectionMode pulumi.StringPtrInput
	// The name of the Resource Group in which the Analysis Services Server should be exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The full name of the Analysis Services Server.
	ServerFullName pulumi.StringPtrInput
	// SKU for the Analysis Services Server. Possible values are: `D1`, `B1`, `B2`, `S0`, `S1`, `S2`, `S4`, `S8`, `S9`, `S8v2` and `S9v2`.
	Sku  pulumi.StringPtrInput
	Tags pulumi.StringMapInput
}

func (ServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverState)(nil)).Elem()
}

type serverArgs struct {
	// List of email addresses of admin users.
	AdminUsers []string `pulumi:"adminUsers"`
	// URI and SAS token for a blob container to store backups.
	BackupBlobContainerUri *string `pulumi:"backupBlobContainerUri"`
	// Indicates if the Power BI service is allowed to access or not.
	EnablePowerBiService *bool `pulumi:"enablePowerBiService"`
	// One or more `ipv4FirewallRule` block(s) as defined below.
	Ipv4FirewallRules []ServerIpv4FirewallRule `pulumi:"ipv4FirewallRules"`
	// The Azure location where the Analysis Services Server exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the firewall rule.
	Name *string `pulumi:"name"`
	// Controls how the read-write server is used in the query pool. If this value is set to `All` then read-write servers are also used for queries. Otherwise with `ReadOnly` these servers do not participate in query operations.
	QuerypoolConnectionMode *string `pulumi:"querypoolConnectionMode"`
	// The name of the Resource Group in which the Analysis Services Server should be exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SKU for the Analysis Services Server. Possible values are: `D1`, `B1`, `B2`, `S0`, `S1`, `S2`, `S4`, `S8`, `S9`, `S8v2` and `S9v2`.
	Sku  string            `pulumi:"sku"`
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
	// List of email addresses of admin users.
	AdminUsers pulumi.StringArrayInput
	// URI and SAS token for a blob container to store backups.
	BackupBlobContainerUri pulumi.StringPtrInput
	// Indicates if the Power BI service is allowed to access or not.
	EnablePowerBiService pulumi.BoolPtrInput
	// One or more `ipv4FirewallRule` block(s) as defined below.
	Ipv4FirewallRules ServerIpv4FirewallRuleArrayInput
	// The Azure location where the Analysis Services Server exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the firewall rule.
	Name pulumi.StringPtrInput
	// Controls how the read-write server is used in the query pool. If this value is set to `All` then read-write servers are also used for queries. Otherwise with `ReadOnly` these servers do not participate in query operations.
	QuerypoolConnectionMode pulumi.StringPtrInput
	// The name of the Resource Group in which the Analysis Services Server should be exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// SKU for the Analysis Services Server. Possible values are: `D1`, `B1`, `B2`, `S0`, `S1`, `S2`, `S4`, `S8`, `S9`, `S8v2` and `S9v2`.
	Sku  pulumi.StringInput
	Tags pulumi.StringMapInput
}

func (ServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverArgs)(nil)).Elem()
}

type ServerInput interface {
	pulumi.Input

	ToServerOutput() ServerOutput
	ToServerOutputWithContext(ctx context.Context) ServerOutput
}

func (*Server) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (i *Server) ToServerOutput() ServerOutput {
	return i.ToServerOutputWithContext(context.Background())
}

func (i *Server) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerOutput)
}

// ServerArrayInput is an input type that accepts ServerArray and ServerArrayOutput values.
// You can construct a concrete instance of `ServerArrayInput` via:
//
//          ServerArray{ ServerArgs{...} }
type ServerArrayInput interface {
	pulumi.Input

	ToServerArrayOutput() ServerArrayOutput
	ToServerArrayOutputWithContext(context.Context) ServerArrayOutput
}

type ServerArray []ServerInput

func (ServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Server)(nil)).Elem()
}

func (i ServerArray) ToServerArrayOutput() ServerArrayOutput {
	return i.ToServerArrayOutputWithContext(context.Background())
}

func (i ServerArray) ToServerArrayOutputWithContext(ctx context.Context) ServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerArrayOutput)
}

// ServerMapInput is an input type that accepts ServerMap and ServerMapOutput values.
// You can construct a concrete instance of `ServerMapInput` via:
//
//          ServerMap{ "key": ServerArgs{...} }
type ServerMapInput interface {
	pulumi.Input

	ToServerMapOutput() ServerMapOutput
	ToServerMapOutputWithContext(context.Context) ServerMapOutput
}

type ServerMap map[string]ServerInput

func (ServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Server)(nil)).Elem()
}

func (i ServerMap) ToServerMapOutput() ServerMapOutput {
	return i.ToServerMapOutputWithContext(context.Background())
}

func (i ServerMap) ToServerMapOutputWithContext(ctx context.Context) ServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerMapOutput)
}

type ServerOutput struct{ *pulumi.OutputState }

func (ServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (o ServerOutput) ToServerOutput() ServerOutput {
	return o
}

func (o ServerOutput) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return o
}

type ServerArrayOutput struct{ *pulumi.OutputState }

func (ServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Server)(nil)).Elem()
}

func (o ServerArrayOutput) ToServerArrayOutput() ServerArrayOutput {
	return o
}

func (o ServerArrayOutput) ToServerArrayOutputWithContext(ctx context.Context) ServerArrayOutput {
	return o
}

func (o ServerArrayOutput) Index(i pulumi.IntInput) ServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Server {
		return vs[0].([]*Server)[vs[1].(int)]
	}).(ServerOutput)
}

type ServerMapOutput struct{ *pulumi.OutputState }

func (ServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Server)(nil)).Elem()
}

func (o ServerMapOutput) ToServerMapOutput() ServerMapOutput {
	return o
}

func (o ServerMapOutput) ToServerMapOutputWithContext(ctx context.Context) ServerMapOutput {
	return o
}

func (o ServerMapOutput) MapIndex(k pulumi.StringInput) ServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Server {
		return vs[0].(map[string]*Server)[vs[1].(string)]
	}).(ServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerInput)(nil)).Elem(), &Server{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerArrayInput)(nil)).Elem(), ServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerMapInput)(nil)).Elem(), ServerMap{})
	pulumi.RegisterOutputType(ServerOutput{})
	pulumi.RegisterOutputType(ServerArrayOutput{})
	pulumi.RegisterOutputType(ServerMapOutput{})
}
