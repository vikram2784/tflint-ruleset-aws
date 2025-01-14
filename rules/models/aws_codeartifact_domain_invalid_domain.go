// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodeartifactDomainInvalidDomainRule checks the pattern is valid
type AwsCodeartifactDomainInvalidDomainRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCodeartifactDomainInvalidDomainRule returns new rule with default attributes
func NewAwsCodeartifactDomainInvalidDomainRule() *AwsCodeartifactDomainInvalidDomainRule {
	return &AwsCodeartifactDomainInvalidDomainRule{
		resourceType:  "aws_codeartifact_domain",
		attributeName: "domain",
		max:           50,
		min:           2,
		pattern:       regexp.MustCompile(`^[a-z][a-z0-9\-]{0,48}[a-z0-9]$`),
	}
}

// Name returns the rule name
func (r *AwsCodeartifactDomainInvalidDomainRule) Name() string {
	return "aws_codeartifact_domain_invalid_domain"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodeartifactDomainInvalidDomainRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodeartifactDomainInvalidDomainRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodeartifactDomainInvalidDomainRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodeartifactDomainInvalidDomainRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"domain must be 50 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"domain must be 2 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-z][a-z0-9\-]{0,48}[a-z0-9]$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
