// Copyright 2023 OWASP CRS
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/coreruleset/ftw-tests-schema/v3/internal/helpers"
)

var testYaml = `---
meta:
  author: "ftw-tests-schema"
  enabled: true
  name: "testYaml"
  description: "Simple YAML to test that the schema is working."
rule_id: 123456
tests:
  - test_id: 1
    desc: "Test that the schema is working."
    stages:
      - input:
          dest_addr: "192.168.0.1"
          port: 8080
          method: "REPORT"
          headers:
            User-Agent: "CRS Tests"
            Host: "localhost"
            Accept: "*/*"
          encoded_request: "TXkgRGF0YQo="
          uri: "/test"
          protocol: "http"
          autocomplete_headers: false
          stop_magic: true
          save_cookie: false
        output:
          status: 200
          response_contains: ""
          log_contains: "nothing"
          no_log_contains: "everything"
      - input:
          dest_addr: "127.0.0.1"
          port: 80
          method: "OPTIONS"
          headers:
            User-Agent: "FTW Schema Tests"
            Host: "localhost"
          response:
            status: 502
            headers:
              Content-Type: "application/problem+json"
              x-mr-burns: excellent
            body: |
              {"aJsonDocument": ["in the response"]}
            encoded_body: eyJhSnNvbkRvY3VtZW50IjogWyJpbiB0aGUgcmVzcG9uc2UiXX0=
            log_message: Response splitting test 1
        output:
          status: 200
`

var ftwTest = &FTWTest{
	RuleId: 123456,
	Meta: FTWTestMeta{
		Author:      "ftw-tests-schema",
		Enabled:     helpers.BoolPtr(true),
		Name:        "testYaml",
		Description: "Simple YAML to test that the schema is working.",
	},
	Tests: []Test{
		{
			TestId:          1,
			TestDescription: "Test that the schema is working.",
			Stages: []Stage{
				{
					Description: ExampleStage.Description,
					Input:       ExampleInput,
					Output:      ExampleOutput,
				},
				{
					Input: Input{
						DestAddr: helpers.StrPtr("127.0.0.1"),
						Port:     helpers.IntPtr(80),
						Method:   helpers.StrPtr("OPTIONS"),
						Headers: map[string]string{
							"User-Agent": "FTW Schema Tests",
							"Host":       "localhost",
						},
						Response: ExampleRespone,
					},
					Output: Output{
						Status: 200,
					},
				},
			},
		},
	},
}

// testV3Yaml exercises the v3 simplified format:
//   - file-level template with {{.Payload}} in the URI
//   - per-test `payload` field
//   - `id` and `description` as canonical v3 field names
//   - `test_id` and `desc` as deprecated v2 aliases (backward compat)
//   - test-level `output` (no explicit stages)
var testV3Yaml = `---
meta:
  author: "ftw-tests-schema"
  description: "v3 template format test."
rule_id: 942100
template:
  method: "GET"
  uri: "/?q={{.Payload}}"
  protocol: "http"
  headers:
    - name: Host
      value: localhost
    - name: User-Agent
      value: "CRS Tests"
tests:
  - id: 1
    description: "UNION-based SQL injection"
    payload: "1 UNION SELECT 1,2,3--"
    output:
      log:
        expect_ids: [942100]
  - id: 2
    description: "Benign query (no trigger)"
    payload: "hello world"
    output:
      log:
        no_expect_ids: [942100]
  - test_id: 3
    desc: "Backward-compat field names still work"
    payload: "another payload"
    output:
      status: 200
`

func TestUnmarshalFTWTest(t *testing.T) {
	var ftw FTWTest
	assertions := assert.New(t)

	err := yaml.Unmarshal([]byte(testYaml), &ftw)
	assertions.NoError(err)

	assertions.Equal(ftwTest.RuleId, ftw.RuleId)
	assertions.Equal(ftwTest.Meta.Author, ftw.Meta.Author)
	assertions.Equal(ftwTest.Meta.Enabled, ftw.Meta.Enabled)
	assertions.Equal(ftwTest.Meta.Name, ftw.Meta.Name)
	assertions.Equal(ftwTest.Meta.Description, ftw.Meta.Description)
	assertions.Len(ftwTest.Tests, len(ftw.Tests))

	for i, test := range ftw.Tests {
		expectedTest := ftwTest.Tests[i]
		assertions.Equal(expectedTest.TestTitle, test.TestTitle)
		assertions.Equal(expectedTest.TestId, test.TestId)
		assertions.Len(test.Stages, len(expectedTest.Stages))

		for j, stage := range test.Stages {
			expectedStage := expectedTest.Stages[j]
			assertions.Equal(expectedStage.Input.DestAddr, stage.Input.DestAddr)
			assertions.Equal(expectedStage.Input.Port, stage.Input.Port)
			assertions.Equal(expectedStage.Input.Method, stage.Input.Method)
			assertions.Len(stage.Input.Headers, len(expectedStage.Input.Headers))

			for name, value := range stage.Input.Headers {
				assertions.Contains(expectedStage.Input.Headers, name)
				assertions.Equal(expectedStage.Input.Headers[name], value)
			}

			response := stage.Input.Response
			if j == 1 {
				assertions.NotNil(response)
				assertions.Equal(502, response.Status)
				assertions.Equal("{\"aJsonDocument\": [\"in the response\"]}\n", response.Body)
				assertions.Equal("eyJhSnNvbkRvY3VtZW50IjogWyJpbiB0aGUgcmVzcG9uc2UiXX0=", response.EncodedBody)
				assertions.Equal("Response splitting test 1", response.LogMessage)
				assertions.Len(response.Headers, len(expectedStage.Input.Headers))

				for name, value := range stage.Input.Headers {
					assertions.Contains(expectedStage.Input.Headers, name)
					assertions.Equal(expectedStage.Input.Headers[name], value)
				}
			} else {
				assertions.Equal(response, Response{})
			}

			assertions.Equal(expectedStage.Output.NoLogContains, stage.Output.NoLogContains)
			assertions.Equal(expectedStage.Output.Status, stage.Output.Status)
		}
	}
}

func TestUnmarshalV3Format(t *testing.T) {
	var ftw FTWTest
	assertions := assert.New(t)

	err := yaml.Unmarshal([]byte(testV3Yaml), &ftw)
	assertions.NoError(err)

	assertions.Equal(uint(942100), ftw.RuleId)
	assertions.NotNil(ftw.Template, "file-level template should be set")
	assertions.Equal("GET", *ftw.Template.Method)
	assertions.Equal("/?q={{.Payload}}", *ftw.Template.URI)
	assertions.Len(ftw.Template.Headers, 2)

	assertions.Len(ftw.Tests, 3)

	// `id` alias
	assertions.Equal(uint(1), ftw.Tests[0].TestId)
	assertions.Equal("UNION-based SQL injection", ftw.Tests[0].TestDescription)
	assertions.NotNil(ftw.Tests[0].Payload)
	assertions.Equal("1 UNION SELECT 1,2,3--", *ftw.Tests[0].Payload)
	assertions.NotNil(ftw.Tests[0].Output)
	assertions.Equal([]uint{942100}, ftw.Tests[0].Output.Log.ExpectIds)
	assertions.Empty(ftw.Tests[0].Stages, "payload-based test should have no stages")

	// second test
	assertions.Equal(uint(2), ftw.Tests[1].TestId)
	assertions.Equal([]uint{942100}, ftw.Tests[1].Output.Log.NoExpectIds)

	// backward-compat field names
	assertions.Equal(uint(3), ftw.Tests[2].TestId)
	assertions.Equal("Backward-compat field names still work", ftw.Tests[2].TestDescription)
}

func TestIdString(t *testing.T) {
	test := Test{
		RuleId: 942100,
		TestId: 3,
	}
	assert.Equal(t, "942100-3", test.IdString())
}
