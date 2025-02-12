// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsXrayEncryptionConfigInvalidTypeRule checks the pattern is valid
type AwsXrayEncryptionConfigInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsXrayEncryptionConfigInvalidTypeRule returns new rule with default attributes
func NewAwsXrayEncryptionConfigInvalidTypeRule() *AwsXrayEncryptionConfigInvalidTypeRule {
	return &AwsXrayEncryptionConfigInvalidTypeRule{
		resourceType:  "aws_xray_encryption_config",
		attributeName: "type",
		enum: []string{
			"NONE",
			"KMS",
		},
	}
}

// Name returns the rule name
func (r *AwsXrayEncryptionConfigInvalidTypeRule) Name() string {
	return "aws_xray_encryption_config_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsXrayEncryptionConfigInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsXrayEncryptionConfigInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsXrayEncryptionConfigInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsXrayEncryptionConfigInvalidTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
