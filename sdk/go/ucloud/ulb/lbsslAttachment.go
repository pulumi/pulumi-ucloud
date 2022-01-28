// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ulb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Load Balancer SSL attachment resource for attaching SSL certificate to Load Balancer Listener.
type LBSslAttachment struct {
	pulumi.CustomResourceState

	// The ID of listener servers.
	ListenerId     pulumi.StringOutput `pulumi:"listenerId"`
	LoadBalancerId pulumi.StringOutput `pulumi:"loadBalancerId"`
	// The ID of SSL certificate.
	SslId pulumi.StringOutput `pulumi:"sslId"`
}

// NewLBSslAttachment registers a new resource with the given unique name, arguments, and options.
func NewLBSslAttachment(ctx *pulumi.Context,
	name string, args *LBSslAttachmentArgs, opts ...pulumi.ResourceOption) (*LBSslAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ListenerId == nil {
		return nil, errors.New("invalid value for required argument 'ListenerId'")
	}
	if args.LoadBalancerId == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerId'")
	}
	if args.SslId == nil {
		return nil, errors.New("invalid value for required argument 'SslId'")
	}
	var resource LBSslAttachment
	err := ctx.RegisterResource("ucloud:ulb/lBSslAttachment:LBSslAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLBSslAttachment gets an existing LBSslAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLBSslAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LBSslAttachmentState, opts ...pulumi.ResourceOption) (*LBSslAttachment, error) {
	var resource LBSslAttachment
	err := ctx.ReadResource("ucloud:ulb/lBSslAttachment:LBSslAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LBSslAttachment resources.
type lbsslAttachmentState struct {
	// The ID of listener servers.
	ListenerId     *string `pulumi:"listenerId"`
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	// The ID of SSL certificate.
	SslId *string `pulumi:"sslId"`
}

type LBSslAttachmentState struct {
	// The ID of listener servers.
	ListenerId     pulumi.StringPtrInput
	LoadBalancerId pulumi.StringPtrInput
	// The ID of SSL certificate.
	SslId pulumi.StringPtrInput
}

func (LBSslAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*lbsslAttachmentState)(nil)).Elem()
}

type lbsslAttachmentArgs struct {
	// The ID of listener servers.
	ListenerId     string `pulumi:"listenerId"`
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// The ID of SSL certificate.
	SslId string `pulumi:"sslId"`
}

// The set of arguments for constructing a LBSslAttachment resource.
type LBSslAttachmentArgs struct {
	// The ID of listener servers.
	ListenerId     pulumi.StringInput
	LoadBalancerId pulumi.StringInput
	// The ID of SSL certificate.
	SslId pulumi.StringInput
}

func (LBSslAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lbsslAttachmentArgs)(nil)).Elem()
}

type LBSslAttachmentInput interface {
	pulumi.Input

	ToLBSslAttachmentOutput() LBSslAttachmentOutput
	ToLBSslAttachmentOutputWithContext(ctx context.Context) LBSslAttachmentOutput
}

func (*LBSslAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*LBSslAttachment)(nil))
}

func (i *LBSslAttachment) ToLBSslAttachmentOutput() LBSslAttachmentOutput {
	return i.ToLBSslAttachmentOutputWithContext(context.Background())
}

func (i *LBSslAttachment) ToLBSslAttachmentOutputWithContext(ctx context.Context) LBSslAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBSslAttachmentOutput)
}

func (i *LBSslAttachment) ToLBSslAttachmentPtrOutput() LBSslAttachmentPtrOutput {
	return i.ToLBSslAttachmentPtrOutputWithContext(context.Background())
}

func (i *LBSslAttachment) ToLBSslAttachmentPtrOutputWithContext(ctx context.Context) LBSslAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBSslAttachmentPtrOutput)
}

type LBSslAttachmentPtrInput interface {
	pulumi.Input

	ToLBSslAttachmentPtrOutput() LBSslAttachmentPtrOutput
	ToLBSslAttachmentPtrOutputWithContext(ctx context.Context) LBSslAttachmentPtrOutput
}

type lbsslAttachmentPtrType LBSslAttachmentArgs

func (*lbsslAttachmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LBSslAttachment)(nil))
}

func (i *lbsslAttachmentPtrType) ToLBSslAttachmentPtrOutput() LBSslAttachmentPtrOutput {
	return i.ToLBSslAttachmentPtrOutputWithContext(context.Background())
}

func (i *lbsslAttachmentPtrType) ToLBSslAttachmentPtrOutputWithContext(ctx context.Context) LBSslAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBSslAttachmentPtrOutput)
}

// LBSslAttachmentArrayInput is an input type that accepts LBSslAttachmentArray and LBSslAttachmentArrayOutput values.
// You can construct a concrete instance of `LBSslAttachmentArrayInput` via:
//
//          LBSslAttachmentArray{ LBSslAttachmentArgs{...} }
type LBSslAttachmentArrayInput interface {
	pulumi.Input

	ToLBSslAttachmentArrayOutput() LBSslAttachmentArrayOutput
	ToLBSslAttachmentArrayOutputWithContext(context.Context) LBSslAttachmentArrayOutput
}

type LBSslAttachmentArray []LBSslAttachmentInput

func (LBSslAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LBSslAttachment)(nil))
}

func (i LBSslAttachmentArray) ToLBSslAttachmentArrayOutput() LBSslAttachmentArrayOutput {
	return i.ToLBSslAttachmentArrayOutputWithContext(context.Background())
}

func (i LBSslAttachmentArray) ToLBSslAttachmentArrayOutputWithContext(ctx context.Context) LBSslAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBSslAttachmentArrayOutput)
}

// LBSslAttachmentMapInput is an input type that accepts LBSslAttachmentMap and LBSslAttachmentMapOutput values.
// You can construct a concrete instance of `LBSslAttachmentMapInput` via:
//
//          LBSslAttachmentMap{ "key": LBSslAttachmentArgs{...} }
type LBSslAttachmentMapInput interface {
	pulumi.Input

	ToLBSslAttachmentMapOutput() LBSslAttachmentMapOutput
	ToLBSslAttachmentMapOutputWithContext(context.Context) LBSslAttachmentMapOutput
}

type LBSslAttachmentMap map[string]LBSslAttachmentInput

func (LBSslAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LBSslAttachment)(nil))
}

func (i LBSslAttachmentMap) ToLBSslAttachmentMapOutput() LBSslAttachmentMapOutput {
	return i.ToLBSslAttachmentMapOutputWithContext(context.Background())
}

func (i LBSslAttachmentMap) ToLBSslAttachmentMapOutputWithContext(ctx context.Context) LBSslAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBSslAttachmentMapOutput)
}

type LBSslAttachmentOutput struct {
	*pulumi.OutputState
}

func (LBSslAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LBSslAttachment)(nil))
}

func (o LBSslAttachmentOutput) ToLBSslAttachmentOutput() LBSslAttachmentOutput {
	return o
}

func (o LBSslAttachmentOutput) ToLBSslAttachmentOutputWithContext(ctx context.Context) LBSslAttachmentOutput {
	return o
}

func (o LBSslAttachmentOutput) ToLBSslAttachmentPtrOutput() LBSslAttachmentPtrOutput {
	return o.ToLBSslAttachmentPtrOutputWithContext(context.Background())
}

func (o LBSslAttachmentOutput) ToLBSslAttachmentPtrOutputWithContext(ctx context.Context) LBSslAttachmentPtrOutput {
	return o.ApplyT(func(v LBSslAttachment) *LBSslAttachment {
		return &v
	}).(LBSslAttachmentPtrOutput)
}

type LBSslAttachmentPtrOutput struct {
	*pulumi.OutputState
}

func (LBSslAttachmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LBSslAttachment)(nil))
}

func (o LBSslAttachmentPtrOutput) ToLBSslAttachmentPtrOutput() LBSslAttachmentPtrOutput {
	return o
}

func (o LBSslAttachmentPtrOutput) ToLBSslAttachmentPtrOutputWithContext(ctx context.Context) LBSslAttachmentPtrOutput {
	return o
}

type LBSslAttachmentArrayOutput struct{ *pulumi.OutputState }

func (LBSslAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LBSslAttachment)(nil))
}

func (o LBSslAttachmentArrayOutput) ToLBSslAttachmentArrayOutput() LBSslAttachmentArrayOutput {
	return o
}

func (o LBSslAttachmentArrayOutput) ToLBSslAttachmentArrayOutputWithContext(ctx context.Context) LBSslAttachmentArrayOutput {
	return o
}

func (o LBSslAttachmentArrayOutput) Index(i pulumi.IntInput) LBSslAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LBSslAttachment {
		return vs[0].([]LBSslAttachment)[vs[1].(int)]
	}).(LBSslAttachmentOutput)
}

type LBSslAttachmentMapOutput struct{ *pulumi.OutputState }

func (LBSslAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LBSslAttachment)(nil))
}

func (o LBSslAttachmentMapOutput) ToLBSslAttachmentMapOutput() LBSslAttachmentMapOutput {
	return o
}

func (o LBSslAttachmentMapOutput) ToLBSslAttachmentMapOutputWithContext(ctx context.Context) LBSslAttachmentMapOutput {
	return o
}

func (o LBSslAttachmentMapOutput) MapIndex(k pulumi.StringInput) LBSslAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LBSslAttachment {
		return vs[0].(map[string]LBSslAttachment)[vs[1].(string)]
	}).(LBSslAttachmentOutput)
}

func init() {
	pulumi.RegisterOutputType(LBSslAttachmentOutput{})
	pulumi.RegisterOutputType(LBSslAttachmentPtrOutput{})
	pulumi.RegisterOutputType(LBSslAttachmentArrayOutput{})
	pulumi.RegisterOutputType(LBSslAttachmentMapOutput{})
}