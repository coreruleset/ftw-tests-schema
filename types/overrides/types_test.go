// Copyright 2023 OWASP CRS
// SPDX-License-Identifier: Apache-2.0

package overrides

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.yaml.in/yaml/v4"
)

var overridesYaml = `---
version: "v0.0.0"
meta:
  engine: "libmodsecurity3"
  platform: "nginx"
  annotations:
    os: "Debian Bullseye"
    purpose: "L7ASR test suite"
test_overrides:
  - rule_id: 920100
    test_ids: [4, 6]
    reason: |-
      nginx returns 400 when ` + "`" + "Content-Length" + "`" + ` header is sent in a
      ` + "`" + `Transfer-Encoding: chunked` + "`" + ` request.
    output:
      status: 200
      response_contains: "HTTP/1.1"
      log_contains: "nothing"
      no_log_contains: "everything"
      log:
        expect_ids: [123456]
        no_expect_ids: [123456]
        match_regex: 'id[:\s"]*123456'
        no_match_regex: 'id[:\s"]*123456'
      expect_error: true
`

var ftwOverrides = &FTWOverrides{
	Version:       "v0.0.0",
	Meta:          MetaExample,
	TestOverrides: TestOverridesExample,
}

func TestUnmarshalFTWOverrides(t *testing.T) {
	var overrides FTWOverrides

	err := yaml.Unmarshal([]byte(overridesYaml), &overrides)

	if err != nil {
		t.Errorf("Unmarshal: %v", err)
	}

	assert.Equal(t, ftwOverrides.Version, overrides.Version)
}

func TestUnmarshalMeta(t *testing.T) {
	var overrides FTWOverrides
	assert := assert.New(t)

	err := yaml.Unmarshal([]byte(overridesYaml), &overrides)
	if err != nil {
		t.Errorf("Unmarshal: %v", err)
	}

	meta := overrides.Meta
	expectedMeta := ftwOverrides.Meta
	assert.Equal(expectedMeta.Engine, meta.Engine)
	assert.Equal(expectedMeta.Platform, meta.Platform)
	assert.NotNil(meta.Annotations)
	assert.Len(meta.Annotations, len(expectedMeta.Annotations))

	for key, value := range meta.Annotations {
		expectedValue := expectedMeta.Annotations[key]
		assert.Equal(expectedValue, value)
	}
}

func TestUnmarmarshalTestOverrides(t *testing.T) {
	var overrides FTWOverrides
	asserter := assert.New(t)

	err := yaml.Unmarshal([]byte(overridesYaml), &overrides)
	if err != nil {
		t.Errorf("Unmarshal: %v", err)
	}

	testOverrides := overrides.TestOverrides
	expectedTestOverrides := ftwOverrides.TestOverrides
	asserter.NotNil(testOverrides)
	asserter.Len(testOverrides, len(expectedTestOverrides))

	for index, testOverride := range testOverrides {
		expectedTestOverride := expectedTestOverrides[index]
		asserter.Equal(expectedTestOverride.RuleId, testOverride.RuleId)
		asserter.ElementsMatch(expectedTestOverride.TestIds, testOverride.TestIds)
		asserter.Equal(expectedTestOverride.Reason, testOverride.Reason)
		if !assert.ObjectsAreEqual(expectedTestOverride.Output, testOverride.Output) {
			asserter.Failf("Output:", "%v != %v", testOverride.Output, expectedTestOverride.Output)
		}
	}
}
