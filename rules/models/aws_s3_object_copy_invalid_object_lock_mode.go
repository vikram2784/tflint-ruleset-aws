// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsS3ObjectCopyInvalidObjectLockModeRule checks the pattern is valid
type AwsS3ObjectCopyInvalidObjectLockModeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsS3ObjectCopyInvalidObjectLockModeRule returns new rule with default attributes
func NewAwsS3ObjectCopyInvalidObjectLockModeRule() *AwsS3ObjectCopyInvalidObjectLockModeRule {
	return &AwsS3ObjectCopyInvalidObjectLockModeRule{
		resourceType:  "aws_s3_object_copy",
		attributeName: "object_lock_mode",
		enum: []string{
			"GOVERNANCE",
			"COMPLIANCE",
		},
	}
}

// Name returns the rule name
func (r *AwsS3ObjectCopyInvalidObjectLockModeRule) Name() string {
	return "aws_s3_object_copy_invalid_object_lock_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsS3ObjectCopyInvalidObjectLockModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsS3ObjectCopyInvalidObjectLockModeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsS3ObjectCopyInvalidObjectLockModeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsS3ObjectCopyInvalidObjectLockModeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as object_lock_mode`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
