// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ucloud.Unet.Outputs
{

    [OutputType]
    public sealed class GetEIPEipIpSetResult
    {
        /// <summary>
        /// Type of Elastic IP routes.
        /// </summary>
        public readonly string InternetType;
        /// <summary>
        /// Elastic IP address.
        /// </summary>
        public readonly string Ip;

        [OutputConstructor]
        private GetEIPEipIpSetResult(
            string internetType,

            string ip)
        {
            InternetType = internetType;
            Ip = ip;
        }
    }
}
