// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides a list of Load Balancer Rule resources according to their Load Balancer Rule ID.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ucloud from "@pulumi/ucloud";
 *
 * const example = ucloud.ulb.getLBRules({
 *     loadBalancerId: "ulb-xxx",
 *     listenerId: "vserver-xxx",
 * });
 * export const first = example.then(example => example.lbRules[0].id);
 * ```
 */
export function getLBRules(args: GetLBRulesArgs, opts?: pulumi.InvokeOptions): Promise<GetLBRulesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ucloud:ulb/getLBRules:getLBRules", {
        "ids": args.ids,
        "listenerId": args.listenerId,
        "loadBalancerId": args.loadBalancerId,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getLBRules.
 */
export interface GetLBRulesArgs {
    /**
     * A list of LB Rule IDs, all the LB Rules belong to the Load Balancer listener will be retrieved if the ID is `[]`.
     */
    ids?: string[];
    /**
     * The ID of a listener server.
     */
    listenerId: string;
    /**
     * The ID of a load balancer.
     */
    loadBalancerId: string;
    outputFile?: string;
}

/**
 * A collection of values returned by getLBRules.
 */
export interface GetLBRulesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    /**
     * It is a nested type which documented below.
     */
    readonly lbRules: outputs.ulb.GetLBRulesLbRule[];
    readonly listenerId: string;
    readonly loadBalancerId: string;
    readonly outputFile?: string;
    /**
     * Total number of LB Rules that satisfy the condition.
     */
    readonly totalCount: number;
}
