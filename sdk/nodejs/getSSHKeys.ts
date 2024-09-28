// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSSHKeys(opts?: pulumi.InvokeOptions): Promise<GetSSHKeysResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("pgedge:index/getSSHKeys:getSSHKeys", {
    }, opts);
}

/**
 * A collection of values returned by getSSHKeys.
 */
export interface GetSSHKeysResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly sshKeys: outputs.GetSSHKeysSshKey[];
}
export function getSSHKeysOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetSSHKeysResult> {
    return pulumi.output(getSSHKeys(opts))
}