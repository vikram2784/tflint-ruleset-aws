// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodeartifactRepositoryInvalidDomainRule checks the pattern is valid
type AwsCodeartifactRepositoryInvalidDomainRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCodeartifactRepositoryInvalidDomainRule returns new rule with default attributes
func NewAwsCodeartifactRepositoryInvalidDomainRule() *AwsCodeartifactRepositoryInvalidDomainRule {
	return &AwsCodeartifactRepositoryInvalidDomainRule{
		resourceType:  "aws_codeartifact_repository",
		attributeName: "domain",
		max:           50,
		min:           2,
		pattern:       regexp.MustCompile(`^[a-z][a-z0-9\-]{0,48}[a-z0-9]$`),
	}
}

// Name returns the rule name
func (r *AwsCodeartifactRepositoryInvalidDomainRule) Name() string {
	return "aws_codeartifact_repository_invalid_domain"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodeartifactRepositoryInvalidDomainRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodeartifactRepositoryInvalidDomainRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodeartifactRepositoryInvalidDomainRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodeartifactRepositoryInvalidDomainRule) Check(runner tflint.Runner) error {
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
					"domain must be 50 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"domain must be 2 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-z][a-z0-9\-]{0,48}[a-z0-9]$`),
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
