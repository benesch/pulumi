// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Configstation.Config/sandwich.Outputs
{

    [OutputType]
    public sealed class Sandwich
    {
        public readonly string? Bread;
        public readonly ImmutableArray<string> Veggies;

        [OutputConstructor]
        private Sandwich(
            string? bread,

            ImmutableArray<string> veggies)
        {
            Bread = bread;
            Veggies = veggies;
        }
    }
}
