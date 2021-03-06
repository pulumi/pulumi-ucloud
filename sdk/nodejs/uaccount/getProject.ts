// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source providers a list of projects owned by user according to finance permission and name.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ucloud from "@pulumi/ucloud";
 *
 * const example = ucloud.uaccount.getProject({
 *     isFinance: false,
 * });
 * export const first = example.then(example => example.projects[0].id);
 * ```
 */
export function getProject(args?: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ucloud:uaccount/getProject:getProject", {
        "isFinance": args.isFinance,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectArgs {
    /**
     * To identify if the current account is granted with financial permission.
     */
    isFinance?: boolean;
    /**
     * A regex string to filter resulting projects by name.
     */
    nameRegex?: string;
    outputFile?: string;
}

/**
 * A collection of values returned by getProject.
 */
export interface GetProjectResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isFinance?: boolean;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * It is a nested type which documented below.
     */
    readonly projects: outputs.uaccount.GetProjectProject[];
    /**
     * Total number of projects that satisfy the condition.
     */
    readonly totalCount: number;
}
