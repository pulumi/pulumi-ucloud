// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an VPC Peering Connection for establishing a connection between multiple VPC.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-ucloud/sdk/go/ucloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		foo, err := vpc.NewVPC(ctx, "foo", &vpc.VPCArgs{
// 			Tag: pulumi.String("tf-example"),
// 			CidrBlocks: pulumi.StringArray{
// 				pulumi.String("192.168.0.0/16"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		bar, err := vpc.NewVPC(ctx, "bar", &vpc.VPCArgs{
// 			Tag: pulumi.String("tf-example"),
// 			CidrBlocks: pulumi.StringArray{
// 				pulumi.String("10.10.0.0/16"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vpc.NewVPCPeeringConnection(ctx, "connection", &vpc.VPCPeeringConnectionArgs{
// 			VpcId:     foo.ID(),
// 			PeerVpcId: bar.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type VPCPeeringConnection struct {
	pulumi.CustomResourceState

	// The ID of accepter project of the specific VPC Peering Connection to retrieve.
	PeerProjectId pulumi.StringOutput `pulumi:"peerProjectId"`
	// The short ID of accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerVpcId pulumi.StringOutput `pulumi:"peerVpcId"`
	// The short of ID of the requester VPC of the specific VPC Peering Connection to retrieve.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewVPCPeeringConnection registers a new resource with the given unique name, arguments, and options.
func NewVPCPeeringConnection(ctx *pulumi.Context,
	name string, args *VPCPeeringConnectionArgs, opts ...pulumi.ResourceOption) (*VPCPeeringConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PeerVpcId == nil {
		return nil, errors.New("invalid value for required argument 'PeerVpcId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource VPCPeeringConnection
	err := ctx.RegisterResource("ucloud:vpc/vPCPeeringConnection:VPCPeeringConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVPCPeeringConnection gets an existing VPCPeeringConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVPCPeeringConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VPCPeeringConnectionState, opts ...pulumi.ResourceOption) (*VPCPeeringConnection, error) {
	var resource VPCPeeringConnection
	err := ctx.ReadResource("ucloud:vpc/vPCPeeringConnection:VPCPeeringConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VPCPeeringConnection resources.
type vpcpeeringConnectionState struct {
	// The ID of accepter project of the specific VPC Peering Connection to retrieve.
	PeerProjectId *string `pulumi:"peerProjectId"`
	// The short ID of accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerVpcId *string `pulumi:"peerVpcId"`
	// The short of ID of the requester VPC of the specific VPC Peering Connection to retrieve.
	VpcId *string `pulumi:"vpcId"`
}

type VPCPeeringConnectionState struct {
	// The ID of accepter project of the specific VPC Peering Connection to retrieve.
	PeerProjectId pulumi.StringPtrInput
	// The short ID of accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerVpcId pulumi.StringPtrInput
	// The short of ID of the requester VPC of the specific VPC Peering Connection to retrieve.
	VpcId pulumi.StringPtrInput
}

func (VPCPeeringConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcpeeringConnectionState)(nil)).Elem()
}

type vpcpeeringConnectionArgs struct {
	// The ID of accepter project of the specific VPC Peering Connection to retrieve.
	PeerProjectId *string `pulumi:"peerProjectId"`
	// The short ID of accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerVpcId string `pulumi:"peerVpcId"`
	// The short of ID of the requester VPC of the specific VPC Peering Connection to retrieve.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a VPCPeeringConnection resource.
type VPCPeeringConnectionArgs struct {
	// The ID of accepter project of the specific VPC Peering Connection to retrieve.
	PeerProjectId pulumi.StringPtrInput
	// The short ID of accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerVpcId pulumi.StringInput
	// The short of ID of the requester VPC of the specific VPC Peering Connection to retrieve.
	VpcId pulumi.StringInput
}

func (VPCPeeringConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcpeeringConnectionArgs)(nil)).Elem()
}

type VPCPeeringConnectionInput interface {
	pulumi.Input

	ToVPCPeeringConnectionOutput() VPCPeeringConnectionOutput
	ToVPCPeeringConnectionOutputWithContext(ctx context.Context) VPCPeeringConnectionOutput
}

func (*VPCPeeringConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*VPCPeeringConnection)(nil))
}

func (i *VPCPeeringConnection) ToVPCPeeringConnectionOutput() VPCPeeringConnectionOutput {
	return i.ToVPCPeeringConnectionOutputWithContext(context.Background())
}

func (i *VPCPeeringConnection) ToVPCPeeringConnectionOutputWithContext(ctx context.Context) VPCPeeringConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPCPeeringConnectionOutput)
}

func (i *VPCPeeringConnection) ToVPCPeeringConnectionPtrOutput() VPCPeeringConnectionPtrOutput {
	return i.ToVPCPeeringConnectionPtrOutputWithContext(context.Background())
}

func (i *VPCPeeringConnection) ToVPCPeeringConnectionPtrOutputWithContext(ctx context.Context) VPCPeeringConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPCPeeringConnectionPtrOutput)
}

type VPCPeeringConnectionPtrInput interface {
	pulumi.Input

	ToVPCPeeringConnectionPtrOutput() VPCPeeringConnectionPtrOutput
	ToVPCPeeringConnectionPtrOutputWithContext(ctx context.Context) VPCPeeringConnectionPtrOutput
}

type vpcpeeringConnectionPtrType VPCPeeringConnectionArgs

func (*vpcpeeringConnectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VPCPeeringConnection)(nil))
}

func (i *vpcpeeringConnectionPtrType) ToVPCPeeringConnectionPtrOutput() VPCPeeringConnectionPtrOutput {
	return i.ToVPCPeeringConnectionPtrOutputWithContext(context.Background())
}

func (i *vpcpeeringConnectionPtrType) ToVPCPeeringConnectionPtrOutputWithContext(ctx context.Context) VPCPeeringConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPCPeeringConnectionPtrOutput)
}

// VPCPeeringConnectionArrayInput is an input type that accepts VPCPeeringConnectionArray and VPCPeeringConnectionArrayOutput values.
// You can construct a concrete instance of `VPCPeeringConnectionArrayInput` via:
//
//          VPCPeeringConnectionArray{ VPCPeeringConnectionArgs{...} }
type VPCPeeringConnectionArrayInput interface {
	pulumi.Input

	ToVPCPeeringConnectionArrayOutput() VPCPeeringConnectionArrayOutput
	ToVPCPeeringConnectionArrayOutputWithContext(context.Context) VPCPeeringConnectionArrayOutput
}

type VPCPeeringConnectionArray []VPCPeeringConnectionInput

func (VPCPeeringConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*VPCPeeringConnection)(nil))
}

func (i VPCPeeringConnectionArray) ToVPCPeeringConnectionArrayOutput() VPCPeeringConnectionArrayOutput {
	return i.ToVPCPeeringConnectionArrayOutputWithContext(context.Background())
}

func (i VPCPeeringConnectionArray) ToVPCPeeringConnectionArrayOutputWithContext(ctx context.Context) VPCPeeringConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPCPeeringConnectionArrayOutput)
}

// VPCPeeringConnectionMapInput is an input type that accepts VPCPeeringConnectionMap and VPCPeeringConnectionMapOutput values.
// You can construct a concrete instance of `VPCPeeringConnectionMapInput` via:
//
//          VPCPeeringConnectionMap{ "key": VPCPeeringConnectionArgs{...} }
type VPCPeeringConnectionMapInput interface {
	pulumi.Input

	ToVPCPeeringConnectionMapOutput() VPCPeeringConnectionMapOutput
	ToVPCPeeringConnectionMapOutputWithContext(context.Context) VPCPeeringConnectionMapOutput
}

type VPCPeeringConnectionMap map[string]VPCPeeringConnectionInput

func (VPCPeeringConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*VPCPeeringConnection)(nil))
}

func (i VPCPeeringConnectionMap) ToVPCPeeringConnectionMapOutput() VPCPeeringConnectionMapOutput {
	return i.ToVPCPeeringConnectionMapOutputWithContext(context.Background())
}

func (i VPCPeeringConnectionMap) ToVPCPeeringConnectionMapOutputWithContext(ctx context.Context) VPCPeeringConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPCPeeringConnectionMapOutput)
}

type VPCPeeringConnectionOutput struct {
	*pulumi.OutputState
}

func (VPCPeeringConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VPCPeeringConnection)(nil))
}

func (o VPCPeeringConnectionOutput) ToVPCPeeringConnectionOutput() VPCPeeringConnectionOutput {
	return o
}

func (o VPCPeeringConnectionOutput) ToVPCPeeringConnectionOutputWithContext(ctx context.Context) VPCPeeringConnectionOutput {
	return o
}

func (o VPCPeeringConnectionOutput) ToVPCPeeringConnectionPtrOutput() VPCPeeringConnectionPtrOutput {
	return o.ToVPCPeeringConnectionPtrOutputWithContext(context.Background())
}

func (o VPCPeeringConnectionOutput) ToVPCPeeringConnectionPtrOutputWithContext(ctx context.Context) VPCPeeringConnectionPtrOutput {
	return o.ApplyT(func(v VPCPeeringConnection) *VPCPeeringConnection {
		return &v
	}).(VPCPeeringConnectionPtrOutput)
}

type VPCPeeringConnectionPtrOutput struct {
	*pulumi.OutputState
}

func (VPCPeeringConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VPCPeeringConnection)(nil))
}

func (o VPCPeeringConnectionPtrOutput) ToVPCPeeringConnectionPtrOutput() VPCPeeringConnectionPtrOutput {
	return o
}

func (o VPCPeeringConnectionPtrOutput) ToVPCPeeringConnectionPtrOutputWithContext(ctx context.Context) VPCPeeringConnectionPtrOutput {
	return o
}

type VPCPeeringConnectionArrayOutput struct{ *pulumi.OutputState }

func (VPCPeeringConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VPCPeeringConnection)(nil))
}

func (o VPCPeeringConnectionArrayOutput) ToVPCPeeringConnectionArrayOutput() VPCPeeringConnectionArrayOutput {
	return o
}

func (o VPCPeeringConnectionArrayOutput) ToVPCPeeringConnectionArrayOutputWithContext(ctx context.Context) VPCPeeringConnectionArrayOutput {
	return o
}

func (o VPCPeeringConnectionArrayOutput) Index(i pulumi.IntInput) VPCPeeringConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VPCPeeringConnection {
		return vs[0].([]VPCPeeringConnection)[vs[1].(int)]
	}).(VPCPeeringConnectionOutput)
}

type VPCPeeringConnectionMapOutput struct{ *pulumi.OutputState }

func (VPCPeeringConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VPCPeeringConnection)(nil))
}

func (o VPCPeeringConnectionMapOutput) ToVPCPeeringConnectionMapOutput() VPCPeeringConnectionMapOutput {
	return o
}

func (o VPCPeeringConnectionMapOutput) ToVPCPeeringConnectionMapOutputWithContext(ctx context.Context) VPCPeeringConnectionMapOutput {
	return o
}

func (o VPCPeeringConnectionMapOutput) MapIndex(k pulumi.StringInput) VPCPeeringConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VPCPeeringConnection {
		return vs[0].(map[string]VPCPeeringConnection)[vs[1].(string)]
	}).(VPCPeeringConnectionOutput)
}

func init() {
	pulumi.RegisterOutputType(VPCPeeringConnectionOutput{})
	pulumi.RegisterOutputType(VPCPeeringConnectionPtrOutput{})
	pulumi.RegisterOutputType(VPCPeeringConnectionArrayOutput{})
	pulumi.RegisterOutputType(VPCPeeringConnectionMapOutput{})
}
