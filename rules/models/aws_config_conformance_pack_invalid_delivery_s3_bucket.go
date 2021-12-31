// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsConfigConformancePackInvalidDeliveryS3BucketRule checks the pattern is valid
type AwsConfigConformancePackInvalidDeliveryS3BucketRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsConfigConformancePackInvalidDeliveryS3BucketRule returns new rule with default attributes
func NewAwsConfigConformancePackInvalidDeliveryS3BucketRule() *AwsConfigConformancePackInvalidDeliveryS3BucketRule {
	return &AwsConfigConformancePackInvalidDeliveryS3BucketRule{
		resourceType:  "aws_config_conformance_pack",
		attributeName: "delivery_s3_bucket",
		max:           63,
	}
}

// Name returns the rule name
func (r *AwsConfigConformancePackInvalidDeliveryS3BucketRule) Name() string {
	return "aws_config_conformance_pack_invalid_delivery_s3_bucket"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigConformancePackInvalidDeliveryS3BucketRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigConformancePackInvalidDeliveryS3BucketRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigConformancePackInvalidDeliveryS3BucketRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigConformancePackInvalidDeliveryS3BucketRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"delivery_s3_bucket must be 63 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
