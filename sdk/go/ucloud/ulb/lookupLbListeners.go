// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ulb

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides a list of Load Balancer Listener resources according to their Load Balancer Listener ID.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/d/lb_listeners.html.markdown.
func LookupLbListeners(ctx *pulumi.Context, args *LookupLbListenersArgs) (*LookupLbListenersResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["ids"] = args.Ids
		inputs["loadBalancerId"] = args.LoadBalancerId
		inputs["nameRegex"] = args.NameRegex
		inputs["outputFile"] = args.OutputFile
	}
	outputs, err := ctx.Invoke("ucloud:ulb/lookupLbListeners:lookupLbListeners", inputs)
	if err != nil {
		return nil, err
	}
	return &LookupLbListenersResult{
		Ids:            outputs["ids"],
		LbListeners:    outputs["lbListeners"],
		LoadBalancerId: outputs["loadBalancerId"],
		NameRegex:      outputs["nameRegex"],
		OutputFile:     outputs["outputFile"],
		TotalCount:     outputs["totalCount"],
		Id:             outputs["id"],
	}, nil
}

// A collection of arguments for invoking lookupLbListeners.
type LookupLbListenersArgs struct {
	// A list of LB Listener IDs, all the LB Listeners belong to this region will be retrieved if the ID is `""`.
	Ids interface{}
	// The ID of a load balancer.
	LoadBalancerId interface{}
	// A regex string to filter resulting lb listeners by name.
	NameRegex  interface{}
	OutputFile interface{}
}

// A collection of values returned by lookupLbListeners.
type LookupLbListenersResult struct {
	Ids interface{}
	// It is a nested type which documented below.
	LbListeners    interface{}
	LoadBalancerId interface{}
	NameRegex      interface{}
	OutputFile     interface{}
	// Total number of LB listeners that satisfy the condition.
	TotalCount interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}