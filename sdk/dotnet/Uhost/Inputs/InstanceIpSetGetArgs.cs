// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ucloud.Uhost.Inputs
{

    public sealed class InstanceIpSetGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of Elastic IP routes. Possible values are: `International` as international BGP IP, `BGP` as china BGP IP and `Private` as private IP.
        /// </summary>
        [Input("internetType")]
        public Input<string>? InternetType { get; set; }

        /// <summary>
        /// Elastic IP address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        public InstanceIpSetGetArgs()
        {
        }
    }
}