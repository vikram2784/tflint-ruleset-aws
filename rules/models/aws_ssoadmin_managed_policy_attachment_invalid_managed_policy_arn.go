// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsoadminManagedPolicyAttachmentInvalidManagedPolicyArnRule checks the pattern is valid
type AwsSsoadminManagedPolicyAttachmentInvalidManagedPolicyArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSsoadminManagedPolicyAttachmentInvalidManagedPolicyArnRule returns new rule with default attributes
func NewAwsSsoadminManagedPolicyAttachmentInvalidManagedPolicyArnRule() *AwsSsoadminManagedPolicyAttachmentInvalidManagedPolicyArnRule {
	return &AwsSsoadminManagedPolicyAttachmentInvalidManagedPolicyArnRule{
		resourceType:  "aws_ssoadmin_managed_policy_attachment",
		attributeName: "managed_policy_arn",
		max:           2048,
		min:           20,
	}
}

// Name returns the rule name
func (r *AwsSsoadminManagedPolicyAttachmentInvalidManagedPolicyArnRule) Name() string {
	return "aws_ssoadmin_managed_policy_attachment_invalid_managed_policy_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsoadminManagedPolicyAttachmentInvalidManagedPolicyArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsoadminManagedPolicyAttachmentInvalidManagedPolicyArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsoadminManagedPolicyAttachmentInvalidManagedPolicyArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsoadminManagedPolicyAttachmentInvalidManagedPolicyArnRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"managed_policy_arn must be 2048 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"managed_policy_arn must be 20 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
