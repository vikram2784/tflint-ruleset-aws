// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsImagebuilderImagePipelineInvalidStatusRule checks the pattern is valid
type AwsImagebuilderImagePipelineInvalidStatusRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsImagebuilderImagePipelineInvalidStatusRule returns new rule with default attributes
func NewAwsImagebuilderImagePipelineInvalidStatusRule() *AwsImagebuilderImagePipelineInvalidStatusRule {
	return &AwsImagebuilderImagePipelineInvalidStatusRule{
		resourceType:  "aws_imagebuilder_image_pipeline",
		attributeName: "status",
		enum: []string{
			"DISABLED",
			"ENABLED",
		},
	}
}

// Name returns the rule name
func (r *AwsImagebuilderImagePipelineInvalidStatusRule) Name() string {
	return "aws_imagebuilder_image_pipeline_invalid_status"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsImagebuilderImagePipelineInvalidStatusRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsImagebuilderImagePipelineInvalidStatusRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsImagebuilderImagePipelineInvalidStatusRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsImagebuilderImagePipelineInvalidStatusRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as status`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
