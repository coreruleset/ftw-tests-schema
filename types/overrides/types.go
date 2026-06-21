// Copyright 2023 OWASP CRS
// SPDX-License-Identifier: Apache-2.0

//go:generate dstdocgen -package overrides -path . -structure FTWOverrides -output ./overrides_doc.go

package overrides

import (
	"github.com/coreruleset/ftw-tests-schema/v3/types"
)

// FTWOverrides describes platform specific overrides for tests
type FTWOverrides struct {
	// description: |
	//   The version field designates the version of the schema that validates this file
	// examples:
	//   - value: "\"v0.1.0\""
	Version string `yaml:"version" json:"version"`

	// description: |
	//   Meta describes the metadata information
	// examples:
	//   - value: MetaExample
	Meta FTWOverridesMeta `yaml:"meta" json:"meta"`

	// description: |
	//   List of test override specifications
	// examples:
	//   - value: TestOverridesExample
	TestOverrides []TestOverride `yaml:"test_overrides" json:"test_overrides"`
}

// FTWOverridesMeta describes the metadata information of this yaml file
type FTWOverridesMeta struct {
	// description: |
	//   The name of the WAF engine the tests are expected to run against
	// examples:
	//   - value: "\"coraza\""
	Engine string `yaml:"engine" json:"engine"`

	// description: |
	//   The name of the platform (e.g., web server) the tests are expected to run against
	// examples:
	//   - value: "\"nginx\""
	Platform string `yaml:"platform" json:"platform"`

	// description: |
	//   Custom annotations; can be used to add additional meta information
	// examples:
	//   - value: AnnotationsExample
	Annotations map[string]string `yaml:"annotations" json:"annotations"`
}

// TestOverride describes overrides for a single test
type TestOverride struct {
	// description: |
	//   ID of the rule this test targets.
	// examples:
	//   - value: TestOverridesExample[0].RuleId
	RuleId uint `yaml:"rule_id" json:"rule_id"`

	// description: |
	//   IDs of the tests for rule_id that overrides should be applied to.
	//   If this field is not set, the overrides will be applied to all tests of rule_id.
	// examples:
	//   - value: TestOverridesExample[0].TestIds
	TestIds []uint `yaml:"test_ids,flow,omitempty" json:"test_ids,omitempty"`

	// description: |
	//   IDs of the stages to which overrides should be applied.
	//   Stage IDs listed will be overridden for all test IDs listed in `TestIds`.
	//   If this field is not set, the overrides will be applied to all stages.
	StageIds []uint `yaml:"stage_ids,omitempty" json:"stage_ids,omitempty"`

	// description: |
	//   Describes why this override is necessary.
	// examples:
	//   - value: ReasonExample
	Reason string `yaml:"reason" json:"reason"`

	// description: |
	//   Whether a stage should be retried once in case of failure.
	//   This option is primarily a workaround for a race condition in phase 5,
	//   where the log entry of a rule may be flushed after the test end marker.
	// examples:
	//   - value: true
	RetryOnce *bool `yaml:"retry_once,omitempty" json:"retry_once,omitempty"`

	// description: |
	//   Specifies overrides on the test output.
	//   This definition *replaces* the output definition of the test.
	// examples:
	//   - value: ExampleOutput
	Output types.Output `yaml:"output" json:"output"`
}
