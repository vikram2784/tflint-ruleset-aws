// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53ResolverQueryLogConfigAssociationInvalidResourceIDRule checks the pattern is valid
type AwsRoute53ResolverQueryLogConfigAssociationInvalidResourceIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsRoute53ResolverQueryLogConfigAssociationInvalidResourceIDRule returns new rule with default attributes
func NewAwsRoute53ResolverQueryLogConfigAssociationInvalidResourceIDRule() *AwsRoute53ResolverQueryLogConfigAssociationInvalidResourceIDRule {
	return &AwsRoute53ResolverQueryLogConfigAssociationInvalidResourceIDRule{
		resourceType:  "aws_route53_resolver_query_log_config_association",
		attributeName: "resource_id",
		max:           64,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsRoute53ResolverQueryLogConfigAssociationInvalidResourceIDRule) Name() string {
	return "aws_route53_resolver_query_log_config_association_invalid_resource_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53ResolverQueryLogConfigAssociationInvalidResourceIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53ResolverQueryLogConfigAssociationInvalidResourceIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53ResolverQueryLogConfigAssociationInvalidResourceIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53ResolverQueryLogConfigAssociationInvalidResourceIDRule) Check(runner tflint.Runner) error {
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
					"resource_id must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"resource_id must be 1 characters or higher",
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
