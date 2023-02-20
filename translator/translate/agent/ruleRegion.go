// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package agent

import (
	"fmt"

	"github.com/aws/private-amazon-cloudwatch-agent-staging/translator"
	"github.com/aws/private-amazon-cloudwatch-agent-staging/translator/context"
	"github.com/aws/private-amazon-cloudwatch-agent-staging/translator/util"
)

type Region struct {
}

const (
	RegionKey = "region"
)

// This region will be provided to the corresponding input and output plugins
// This should be applied before interpreting other component.
func (r *Region) ApplyRule(input interface{}) (returnKey string, returnVal interface{}) {
	var region string
	ctx := context.CurrentContext()
	_, inputRegion := translator.DefaultCase(RegionKey, "", input)
	if inputRegion != "" {
		Global_Config.Region = inputRegion.(string)
		return
	}
	region = util.DetectRegion(ctx.Mode(), ctx.Credentials())

	if region == "" {
		translator.AddErrorMessages(GetCurPath()+"ruleRegion/", fmt.Sprintf("Region info is missing for mode: %s",
			ctx.Mode()))
	}

	Global_Config.Region = region
	return
}

func init() {
	r := new(Region)
	RegisterRule(RegionKey, r)
}
