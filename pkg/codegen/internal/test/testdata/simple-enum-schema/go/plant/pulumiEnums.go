// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package plant

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContainerBrightness float64

const (
	ContainerBrightnessZeroPointOne = ContainerBrightness(0.1)
	ContainerBrightnessOne          = ContainerBrightness(1)
)

func (ContainerBrightness) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerBrightness)(nil)).Elem()
}

func (e ContainerBrightness) ToContainerBrightnessOutput() ContainerBrightnessOutput {
	return pulumi.ToOutput(e).(ContainerBrightnessOutput)
}

func (e ContainerBrightness) ToContainerBrightnessOutputWithContext(ctx context.Context) ContainerBrightnessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContainerBrightnessOutput)
}

func (e ContainerBrightness) ToContainerBrightnessPtrOutput() ContainerBrightnessPtrOutput {
	return e.ToContainerBrightnessPtrOutputWithContext(context.Background())
}

func (e ContainerBrightness) ToContainerBrightnessPtrOutputWithContext(ctx context.Context) ContainerBrightnessPtrOutput {
	return ContainerBrightness(e).ToContainerBrightnessOutputWithContext(ctx).ToContainerBrightnessPtrOutputWithContext(ctx)
}

func (e ContainerBrightness) ToFloat64Output() pulumi.Float64Output {
	return pulumi.ToOutput(pulumi.Float64(e)).(pulumi.Float64Output)
}

func (e ContainerBrightness) ToFloat64OutputWithContext(ctx context.Context) pulumi.Float64Output {
	return pulumi.ToOutputWithContext(ctx, pulumi.Float64(e)).(pulumi.Float64Output)
}

func (e ContainerBrightness) ToFloat64PtrOutput() pulumi.Float64PtrOutput {
	return pulumi.Float64(e).ToFloat64PtrOutputWithContext(context.Background())
}

func (e ContainerBrightness) ToFloat64PtrOutputWithContext(ctx context.Context) pulumi.Float64PtrOutput {
	return pulumi.Float64(e).ToFloat64OutputWithContext(ctx).ToFloat64PtrOutputWithContext(ctx)
}

type ContainerBrightnessOutput struct{ *pulumi.OutputState }

func (ContainerBrightnessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerBrightness)(nil)).Elem()
}

func (o ContainerBrightnessOutput) ToContainerBrightnessOutput() ContainerBrightnessOutput {
	return o
}

func (o ContainerBrightnessOutput) ToContainerBrightnessOutputWithContext(ctx context.Context) ContainerBrightnessOutput {
	return o
}

func (o ContainerBrightnessOutput) ToContainerBrightnessPtrOutput() ContainerBrightnessPtrOutput {
	return o.ToContainerBrightnessPtrOutputWithContext(context.Background())
}

func (o ContainerBrightnessOutput) ToContainerBrightnessPtrOutputWithContext(ctx context.Context) ContainerBrightnessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerBrightness) *ContainerBrightness {
		return &v
	}).(ContainerBrightnessPtrOutput)
}

type ContainerBrightnessPtrOutput struct{ *pulumi.OutputState }

func (ContainerBrightnessPtrOutput) ElementType() reflect.Type {
	return containerBrightnessPtrType
}

func (o ContainerBrightnessPtrOutput) ToContainerBrightnessPtrOutput() ContainerBrightnessPtrOutput {
	return o
}

func (o ContainerBrightnessPtrOutput) ToContainerBrightnessPtrOutputWithContext(ctx context.Context) ContainerBrightnessPtrOutput {
	return o
}

func (o ContainerBrightnessPtrOutput) Elem() ContainerBrightnessOutput {
	return o.ApplyT(func(v *ContainerBrightness) ContainerBrightness {
		var ret ContainerBrightness
		if v != nil {
			ret = *v
		}
		return ret
	}).(ContainerBrightnessOutput)
}

// ContainerBrightnessInput is an input type that accepts ContainerBrightnessArgs and ContainerBrightnessOutput values.
// You can construct a concrete instance of `ContainerBrightnessInput` via:
//
//          ContainerBrightnessArgs{...}
type ContainerBrightnessInput interface {
	pulumi.Input

	ToContainerBrightnessOutput() ContainerBrightnessOutput
	ToContainerBrightnessOutputWithContext(context.Context) ContainerBrightnessOutput
}

var containerBrightnessPtrType = reflect.TypeOf((**ContainerBrightness)(nil)).Elem()

type ContainerBrightnessPtrInput interface {
	pulumi.Input

	ToContainerBrightnessPtrOutput() ContainerBrightnessPtrOutput
	ToContainerBrightnessPtrOutputWithContext(context.Context) ContainerBrightnessPtrOutput
}

type containerBrightnessPtr float64

func ContainerBrightnessPtr(v float64) ContainerBrightnessPtrInput {
	return (*containerBrightnessPtr)(&v)
}

func (*containerBrightnessPtr) ElementType() reflect.Type {
	return containerBrightnessPtrType
}

func (in *containerBrightnessPtr) ToContainerBrightnessPtrOutput() ContainerBrightnessPtrOutput {
	return pulumi.ToOutput(in).(ContainerBrightnessPtrOutput)
}

func (in *containerBrightnessPtr) ToContainerBrightnessPtrOutputWithContext(ctx context.Context) ContainerBrightnessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContainerBrightnessPtrOutput)
}

// plant container colors
type ContainerColor string

const (
	ContainerColorRed    = ContainerColor("red")
	ContainerColorBlue   = ContainerColor("blue")
	ContainerColorYellow = ContainerColor("yellow")
)

func (ContainerColor) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerColor)(nil)).Elem()
}

func (e ContainerColor) ToContainerColorOutput() ContainerColorOutput {
	return pulumi.ToOutput(e).(ContainerColorOutput)
}

func (e ContainerColor) ToContainerColorOutputWithContext(ctx context.Context) ContainerColorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContainerColorOutput)
}

func (e ContainerColor) ToContainerColorPtrOutput() ContainerColorPtrOutput {
	return e.ToContainerColorPtrOutputWithContext(context.Background())
}

func (e ContainerColor) ToContainerColorPtrOutputWithContext(ctx context.Context) ContainerColorPtrOutput {
	return ContainerColor(e).ToContainerColorOutputWithContext(ctx).ToContainerColorPtrOutputWithContext(ctx)
}

func (e ContainerColor) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerColor) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerColor) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContainerColor) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContainerColorOutput struct{ *pulumi.OutputState }

func (ContainerColorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerColor)(nil)).Elem()
}

func (o ContainerColorOutput) ToContainerColorOutput() ContainerColorOutput {
	return o
}

func (o ContainerColorOutput) ToContainerColorOutputWithContext(ctx context.Context) ContainerColorOutput {
	return o
}

func (o ContainerColorOutput) ToContainerColorPtrOutput() ContainerColorPtrOutput {
	return o.ToContainerColorPtrOutputWithContext(context.Background())
}

func (o ContainerColorOutput) ToContainerColorPtrOutputWithContext(ctx context.Context) ContainerColorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerColor) *ContainerColor {
		return &v
	}).(ContainerColorPtrOutput)
}

type ContainerColorPtrOutput struct{ *pulumi.OutputState }

func (ContainerColorPtrOutput) ElementType() reflect.Type {
	return containerColorPtrType
}

func (o ContainerColorPtrOutput) ToContainerColorPtrOutput() ContainerColorPtrOutput {
	return o
}

func (o ContainerColorPtrOutput) ToContainerColorPtrOutputWithContext(ctx context.Context) ContainerColorPtrOutput {
	return o
}

func (o ContainerColorPtrOutput) Elem() ContainerColorOutput {
	return o.ApplyT(func(v *ContainerColor) ContainerColor {
		var ret ContainerColor
		if v != nil {
			ret = *v
		}
		return ret
	}).(ContainerColorOutput)
}

// ContainerColorInput is an input type that accepts ContainerColorArgs and ContainerColorOutput values.
// You can construct a concrete instance of `ContainerColorInput` via:
//
//          ContainerColorArgs{...}
type ContainerColorInput interface {
	pulumi.Input

	ToContainerColorOutput() ContainerColorOutput
	ToContainerColorOutputWithContext(context.Context) ContainerColorOutput
}

var containerColorPtrType = reflect.TypeOf((**ContainerColor)(nil)).Elem()

type ContainerColorPtrInput interface {
	pulumi.Input

	ToContainerColorPtrOutput() ContainerColorPtrOutput
	ToContainerColorPtrOutputWithContext(context.Context) ContainerColorPtrOutput
}

type containerColorPtr string

func ContainerColorPtr(v string) ContainerColorPtrInput {
	return (*containerColorPtr)(&v)
}

func (*containerColorPtr) ElementType() reflect.Type {
	return containerColorPtrType
}

func (in *containerColorPtr) ToContainerColorPtrOutput() ContainerColorPtrOutput {
	return pulumi.ToOutput(in).(ContainerColorPtrOutput)
}

func (in *containerColorPtr) ToContainerColorPtrOutputWithContext(ctx context.Context) ContainerColorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContainerColorPtrOutput)
}

// plant container sizes
type ContainerSize int

const (
	ContainerSizeFourInch = ContainerSize(4)
	ContainerSizeSixInch  = ContainerSize(6)
	// Deprecated: Eight inch pots are no longer supported.
	ContainerSizeEightInch = ContainerSize(8)
)

func (ContainerSize) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerSize)(nil)).Elem()
}

func (e ContainerSize) ToContainerSizeOutput() ContainerSizeOutput {
	return pulumi.ToOutput(e).(ContainerSizeOutput)
}

func (e ContainerSize) ToContainerSizeOutputWithContext(ctx context.Context) ContainerSizeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContainerSizeOutput)
}

func (e ContainerSize) ToContainerSizePtrOutput() ContainerSizePtrOutput {
	return e.ToContainerSizePtrOutputWithContext(context.Background())
}

func (e ContainerSize) ToContainerSizePtrOutputWithContext(ctx context.Context) ContainerSizePtrOutput {
	return ContainerSize(e).ToContainerSizeOutputWithContext(ctx).ToContainerSizePtrOutputWithContext(ctx)
}

func (e ContainerSize) ToIntOutput() pulumi.IntOutput {
	return pulumi.ToOutput(pulumi.Int(e)).(pulumi.IntOutput)
}

func (e ContainerSize) ToIntOutputWithContext(ctx context.Context) pulumi.IntOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.Int(e)).(pulumi.IntOutput)
}

func (e ContainerSize) ToIntPtrOutput() pulumi.IntPtrOutput {
	return pulumi.Int(e).ToIntPtrOutputWithContext(context.Background())
}

func (e ContainerSize) ToIntPtrOutputWithContext(ctx context.Context) pulumi.IntPtrOutput {
	return pulumi.Int(e).ToIntOutputWithContext(ctx).ToIntPtrOutputWithContext(ctx)
}

type ContainerSizeOutput struct{ *pulumi.OutputState }

func (ContainerSizeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerSize)(nil)).Elem()
}

func (o ContainerSizeOutput) ToContainerSizeOutput() ContainerSizeOutput {
	return o
}

func (o ContainerSizeOutput) ToContainerSizeOutputWithContext(ctx context.Context) ContainerSizeOutput {
	return o
}

func (o ContainerSizeOutput) ToContainerSizePtrOutput() ContainerSizePtrOutput {
	return o.ToContainerSizePtrOutputWithContext(context.Background())
}

func (o ContainerSizeOutput) ToContainerSizePtrOutputWithContext(ctx context.Context) ContainerSizePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerSize) *ContainerSize {
		return &v
	}).(ContainerSizePtrOutput)
}

type ContainerSizePtrOutput struct{ *pulumi.OutputState }

func (ContainerSizePtrOutput) ElementType() reflect.Type {
	return containerSizePtrType
}

func (o ContainerSizePtrOutput) ToContainerSizePtrOutput() ContainerSizePtrOutput {
	return o
}

func (o ContainerSizePtrOutput) ToContainerSizePtrOutputWithContext(ctx context.Context) ContainerSizePtrOutput {
	return o
}

func (o ContainerSizePtrOutput) Elem() ContainerSizeOutput {
	return o.ApplyT(func(v *ContainerSize) ContainerSize {
		var ret ContainerSize
		if v != nil {
			ret = *v
		}
		return ret
	}).(ContainerSizeOutput)
}

// ContainerSizeInput is an input type that accepts ContainerSizeArgs and ContainerSizeOutput values.
// You can construct a concrete instance of `ContainerSizeInput` via:
//
//          ContainerSizeArgs{...}
type ContainerSizeInput interface {
	pulumi.Input

	ToContainerSizeOutput() ContainerSizeOutput
	ToContainerSizeOutputWithContext(context.Context) ContainerSizeOutput
}

var containerSizePtrType = reflect.TypeOf((**ContainerSize)(nil)).Elem()

type ContainerSizePtrInput interface {
	pulumi.Input

	ToContainerSizePtrOutput() ContainerSizePtrOutput
	ToContainerSizePtrOutputWithContext(context.Context) ContainerSizePtrOutput
}

type containerSizePtr int

func ContainerSizePtr(v int) ContainerSizePtrInput {
	return (*containerSizePtr)(&v)
}

func (*containerSizePtr) ElementType() reflect.Type {
	return containerSizePtrType
}

func (in *containerSizePtr) ToContainerSizePtrOutput() ContainerSizePtrOutput {
	return pulumi.ToOutput(in).(ContainerSizePtrOutput)
}

func (in *containerSizePtr) ToContainerSizePtrOutputWithContext(ctx context.Context) ContainerSizePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContainerSizePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerBrightnessOutput{})
	pulumi.RegisterOutputType(ContainerBrightnessPtrOutput{})
	pulumi.RegisterOutputType(ContainerColorOutput{})
	pulumi.RegisterOutputType(ContainerColorPtrOutput{})
	pulumi.RegisterOutputType(ContainerSizeOutput{})
	pulumi.RegisterOutputType(ContainerSizePtrOutput{})
}
