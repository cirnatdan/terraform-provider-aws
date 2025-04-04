// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	awstypes "github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tfslices "github.com/hashicorp/terraform-provider-aws/internal/slices"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKDataSource("aws_waf_web_acl", name="Web ACL")
func dataSourceWebACL() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceWebACLRead,

		Schema: map[string]*schema.Schema{
			names.AttrName: {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceWebACLRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).WAFClient(ctx)

	name := d.Get(names.AttrName).(string)
	input := &waf.ListWebACLsInput{}
	output, err := findWebACL(ctx, conn, input, func(v *awstypes.WebACLSummary) bool {
		return aws.ToString(v.Name) == name
	})

	if err != nil {
		return sdkdiag.AppendFromErr(diags, tfresource.SingularDataSourceFindError("WAF Web ACL", err))
	}

	d.SetId(aws.ToString(output.WebACLId))

	return diags
}

func findWebACL(ctx context.Context, conn *waf.Client, input *waf.ListWebACLsInput, filter tfslices.Predicate[*awstypes.WebACLSummary]) (*awstypes.WebACLSummary, error) {
	output, err := findWebACLs(ctx, conn, input, filter)

	if err != nil {
		return nil, err
	}

	return tfresource.AssertSingleValueResult(output)
}

func findWebACLs(ctx context.Context, conn *waf.Client, input *waf.ListWebACLsInput, filter tfslices.Predicate[*awstypes.WebACLSummary]) ([]awstypes.WebACLSummary, error) {
	var output []awstypes.WebACLSummary

	err := listWebACLsPages(ctx, conn, input, func(page *waf.ListWebACLsOutput, lastPage bool) bool {
		if page == nil {
			return !lastPage
		}

		for _, v := range page.WebACLs {
			if filter(&v) {
				output = append(output, v)
			}
		}

		return !lastPage
	})

	if err != nil {
		return nil, err
	}

	return output, nil
}
