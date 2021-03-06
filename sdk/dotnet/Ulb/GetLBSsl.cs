// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ucloud.Ulb
{
    public static class GetLBSsl
    {
        /// <summary>
        /// This data source provides a list of Load Balancer SSL certificate resources according to their Load Balancer SSL certificate resource ID and name.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Ucloud = Pulumi.Ucloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Ucloud.Ulb.GetLBSsl.InvokeAsync());
        ///         this.First = example.Apply(example =&gt; example.LbSsls[0].Id);
        ///     }
        /// 
        ///     [Output("first")]
        ///     public Output&lt;string&gt; First { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetLBSslResult> InvokeAsync(GetLBSslArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLBSslResult>("ucloud:ulb/getLBSsl:getLBSsl", args ?? new GetLBSslArgs(), options.WithVersion());
    }


    public sealed class GetLBSslArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of LB SSL certificate resource IDs, all the LB SSL certificate resources in the current region will be retrieved if the ID is `[]`.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter resulting LB SSL by name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetLBSslArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLBSslResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// It is a nested type which documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLBSslLbSslResult> LbSsls;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// Total number of LB SSL certificate resources that satisfy the condition.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetLBSslResult(
            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetLBSslLbSslResult> lbSsls,

            string? nameRegex,

            string? outputFile,

            int totalCount)
        {
            Id = id;
            Ids = ids;
            LbSsls = lbSsls;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
