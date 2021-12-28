// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsMacie2ClassificationJobInvalidJobTypeRule checks the pattern is valid
type AwsMacie2ClassificationJobInvalidJobTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsMacie2ClassificationJobInvalidJobTypeRule returns new rule with default attributes
func NewAwsMacie2ClassificationJobInvalidJobTypeRule() *AwsMacie2ClassificationJobInvalidJobTypeRule {
	return &AwsMacie2ClassificationJobInvalidJobTypeRule{
		resourceType:  "aws_macie2_classification_job",
		attributeName: "job_type",
		enum: []string{
			"ONE_TIME",
			"SCHEDULED",
		},
	}
}

// Name returns the rule name
func (r *AwsMacie2ClassificationJobInvalidJobTypeRule) Name() string {
	return "aws_macie2_classification_job_invalid_job_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsMacie2ClassificationJobInvalidJobTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsMacie2ClassificationJobInvalidJobTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsMacie2ClassificationJobInvalidJobTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsMacie2ClassificationJobInvalidJobTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as job_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}