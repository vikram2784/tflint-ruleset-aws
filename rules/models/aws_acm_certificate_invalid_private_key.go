// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAcmCertificateInvalidPrivateKeyRule checks the pattern is valid
type AwsAcmCertificateInvalidPrivateKeyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsAcmCertificateInvalidPrivateKeyRule returns new rule with default attributes
func NewAwsAcmCertificateInvalidPrivateKeyRule() *AwsAcmCertificateInvalidPrivateKeyRule {
	return &AwsAcmCertificateInvalidPrivateKeyRule{
		resourceType:  "aws_acm_certificate",
		attributeName: "private_key",
		max:           524288,
		min:           1,
		pattern:       regexp.MustCompile(`^-{5}BEGIN PRIVATE KEY-{5}\x{000D}?\x{000A}([A-Za-z0-9/+]{64}\x{000D}?\x{000A})*[A-Za-z0-9/+]{1,64}={0,2}\x{000D}?\x{000A}-{5}END PRIVATE KEY-{5}(\x{000D}?\x{000A})?$`),
	}
}

// Name returns the rule name
func (r *AwsAcmCertificateInvalidPrivateKeyRule) Name() string {
	return "aws_acm_certificate_invalid_private_key"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAcmCertificateInvalidPrivateKeyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAcmCertificateInvalidPrivateKeyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAcmCertificateInvalidPrivateKeyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAcmCertificateInvalidPrivateKeyRule) Check(runner tflint.Runner) error {
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
					"private_key must be 524288 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"private_key must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`private_key does not match valid pattern ^-{5}BEGIN PRIVATE KEY-{5}\x{000D}?\x{000A}([A-Za-z0-9/+]{64}\x{000D}?\x{000A})*[A-Za-z0-9/+]{1,64}={0,2}\x{000D}?\x{000A}-{5}END PRIVATE KEY-{5}(\x{000D}?\x{000A})?$`,
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
