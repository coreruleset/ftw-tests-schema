// Copyright 2023 OWASP CRS
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/coreruleset/ftw-tests-schema/v2/internal/helpers"
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

var templateTestYaml = `---
meta:
  author: "ftw-tests-schema"
  name: "templateTest"
  description: "YAML to test template test parsing."
rule_id: 123456
tests:
  - test_id: 1
    desc: "Template test for Unix RCE."
    templates:
      - key: uri
        values:
          - /get
          - /get?foo=bar
      - key: command
        values:
          - php
          - php%20foo
    stages:
      - input:
          dest_addr: "192.168.0.1"
          port: 8080
          protocol: "http"
          uri: "{{ .uri }}"
          method: "POST"
          data: "cmd={{ .command }}"
          headers:
            User-Agent: "CRS Tests"
            Host: "localhost"
            Accept: "*/*"
        output:
          status: 200
`

func TestUnmarshalTemplateTest(t *testing.T) {
	var ftw FTWTest
	assertions := assert.New(t)

	err := yaml.Unmarshal([]byte(templateTestYaml), &ftw)
	assertions.NoError(err)

	assertions.Equal(uint(123456), ftw.RuleId)
	assertions.Len(ftw.Tests, 1)

	test := ftw.Tests[0]
	assertions.Equal(uint(1), test.TestId)
	assertions.Equal("Template test for Unix RCE.", test.TestDescription)
	assertions.Len(test.Templates, 2)
	assertions.Len(test.Stages, 1)

	uriTemplate := test.Templates[0]
	assertions.Equal("uri", uriTemplate.Key)
	assertions.Equal([]string{"/get", "/get?foo=bar"}, uriTemplate.Values)

	commandTemplate := test.Templates[1]
	assertions.Equal("command", commandTemplate.Key)
	assertions.Equal([]string{"php", "php%20foo"}, commandTemplate.Values)

	stage := test.Stages[0]
	assertions.NotNil(stage.Input.URI)
	assertions.Equal("{{ .uri }}", *stage.Input.URI)
	assertions.NotNil(stage.Input.Data)
	assertions.Equal("cmd={{ .command }}", *stage.Input.Data)
}

func TestMarshalTemplateTest(t *testing.T) {
	assertions := assert.New(t)

	data, err := yaml.Marshal(ExampleTemplateTest)
	assertions.NoError(err)

	var roundtrip Test
	err = yaml.Unmarshal(data, &roundtrip)
	assertions.NoError(err)

	assertions.Len(roundtrip.Templates, len(ExampleTemplates))
	for i, tmpl := range roundtrip.Templates {
		assertions.Equal(ExampleTemplates[i].Key, tmpl.Key)
		assertions.Equal(ExampleTemplates[i].Values, tmpl.Values)
	}
}

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
