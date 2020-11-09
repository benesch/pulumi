// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ArgFunction(ctx *pulumi.Context, args *ArgFunctionArgs, opts ...pulumi.InvokeOption) (*ArgFunctionResult, error) {
	var rv ArgFunctionResult
	err := ctx.Invoke("example::argFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ArgFunctionArgs struct {
	Arg1 *Resource `pulumi:"arg1"`
}

type ArgFunctionResult struct {
	Result *Resource `pulumi:"result"`
}
