// Copyright 2023 OWASP CRS
// SPDX-License-Identifier: Apache-2.0

//go:generate dstdocgen -package types -path . -structure FTWTest -output ./test_doc.go

package types

import "fmt"

// Welcome to the FTW YAMLFormat documentation.
// In this document we will explain all the possible options that can be used within the YAML format.
// Generally this is the preferred format for writing tests in as they don't require any programming skills
// in order to understand and change. If you find a bug in this format please open an issue.

// FTWTest is the base type used when unmarshaling YAML tests files
type FTWTest struct {
	// description: |
	//   Meta describes the metadata information of this yaml test file
	Meta FTWTestMeta `yaml:"meta" json:"meta"`

	// description: |
	//   RuleId is the ID of the rule this test targets.
	// examples:
	//   - name: RuleId
	//     value: 123456
	RuleId uint `yaml:"rule_id" json:"rule_id"`

	// description: |
	//   Tests is a list of FTW tests
	// examples:
	//   - value: ExampleTests
	Tests []Test `yaml:"tests" json:"tests"`
}

// FTWTestMeta describes the metadata information of this yaml test file
type FTWTestMeta struct {
	// description: |
	//   Author is the list of authors that added content to this file
	// examples:
	//   - name: Author
	//     value: "\"Felipe Zipitria\""
	Author string `yaml:"author,omitempty" json:"author,omitempty"`

	// description: |
	//   Enabled indicates if the tests are enabled to be run by the engine or not.
	// examples:
	//   - name: Enabled
	//     value: false
	//
	// Deprecated: ignored; use platform specific overrides instead
	Enabled *bool `yaml:"enabled,omitempty" json:"enabled,omitempty"`

	// description: |
	//   Name is the name of the tests contained in this file.
	// examples:
	//   - name: Name
	//     value: "\"test01\""
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// description: |
	//   Description is a textual description of the tests contained in this file.
	// examples:
	//   - name: Description
	//     value: "\"The tests here target SQL injection.\""
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// description: |
	//   Version is the version of the YAML Schema.
	// examples:
	//   - name: Version
	//     value: "\"v1\""
	Version string `yaml:"version,omitempty" json:"version,omitempty"`

	// description: |
	//   Tags is list of strings that can be used for arbitrary grouping of tests.
	// examples:
	//   - name: Tags
	//     value: ["PHP", "bug-123"]
	Tags []string `yaml:"tags,omitempty" json:"tags,omitempty"`
}

// Test is an individual test case. One test can have multiple stages.
type Test struct {
	// description: |
	//   TestTitle is the title of this particular types. It is used for inclusion/exclusion of each run by the tool.
	// examples:
	//   - value: ExampleTest.TestTitle
	//
	// Deprecated: use `rule_id` and `test_id`
	TestTitle string `yaml:"test_title,omitempty" json:"test_title,omitempty"`

	// description: |
	//   RuleId is the ID of the rule this test targets.
	//
	//   This field is for internal use and not exposed via YAML.
	// examples:
	//   - name: RuleId
	//     value: 123456
	RuleId uint `json:"-"`

	// description: |
	//   TestId is the ID of the test, in relation to `rule_id`.
	//
	//   When this field is not set, the ID will be inferred from the
	//   position.
	// examples:
	//   - name: TestId
	//     value: 4
	TestId uint `yaml:"test_id" json:"test_id"`

	// description: |
	//   TestDescription is the description for this particular test.
	//
	//   Should be used to describe the internals of the specific things this test is targeting.
	// examples:
	//   - value: ExampleTest.TestDescription
	TestDescription string `yaml:"desc,omitempty" json:"desc,omitempty"`

	// description: |
	//   Stages is the list of all the stages to perform this test.
	// examples:
	//   - value: ExampleStages
	Stages []Stage `yaml:"stages" json:"stages"`

	// description: |
	//   Tags is list of strings that can be used for arbitrary grouping of tests.
	// examples:
	//   - name: Tags
	//     value: ["PHP", "bug-123"]
	Tags []string `yaml:"tags,omitempty" json:"tags,omitempty"`
}

// IdString prints the human readable ID of a test in the format
// <rule ID>-<test ID>. This format is also used when matching
// the include / exclude regular expressions.
func (t *Test) IdString() string {
	return fmt.Sprintf("%d-%d", t.RuleId, t.TestId)
}

// Stage is a list of stages
type Stage struct {
	// description: |
	//   Describes the purpose of this stage.
	// examples:
	//   - value: ExampleStage.Description
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// description: |
	//   Input is the data that is passed to the test
	// examples:
	//   - name: Input
	//     value: ExampleInput
	Input Input `yaml:"input" json:"input"`

	// description: |
	//   Output is the data that is returned from the test
	// examples:
	//   - name: Output
	//     value: ExampleOutput
	Output Output `yaml:"output" json:"output"`
}

// StageData is the data that is passed to the test, and the data that is returned from the test
//
// Deprecated: use the other fields of `stage`
type StageData struct {
	// description: |
	//   Input is the data that is passed to the test
	// examples:
	//   - name: Input
	//     value: ExampleInput
	Input Input `yaml:"input" json:"input"`

	// description: |
	//   Output is the data that is returned from the test
	// examples:
	//   - name: Output
	//     value: ExampleOutput
	Output Output `yaml:"output" json:"output"`
}

// Input represents the input request in a stage
type Input struct {
	// description: |
	//   DestAddr is the IP of the destination host that the test will send the message to.
	// examples:
	//   - name: DestAddr
	//     value: "\"127.0.0.1\""
	DestAddr *string `yaml:"dest_addr,omitempty" json:"dest_addr,omitempty" koanf:"dest_addr,omitempty"`

	// description: |
	//   Port allows you to declare which port on the destination host the test should connect to.
	// examples:
	//   - name: Port
	//     value: 80
	Port *int `yaml:"port,omitempty" json:"port,omitempty" koanf:"port,omitempty"`

	// description: |
	//   Protocol allows you to declare which protocol the test should use when sending the request.
	// examples:
	//   - name: Protocol
	//     value: "\"http\""
	Protocol *string `yaml:"protocol,omitempty" json:"protocol,omitempty" koanf:"protocol,omitempty"`

	// description: |
	//   URI allows you to declare the URI the test should use as part of the request line.
	// examples:
	//   - name: URI
	//     value: '/get?hello=world"
	URI *string `yaml:"uri,omitempty" json:"uri,omitempty" koanf:"uri,omitempty"`

	// description: |
	//   URIs is a list of URIs that the test should use as part of the request line.
	//   If URIs is set, URI will be ignored. URIs is useful for tests that need to send multiple requests.
	//   Caveat: you cannot change the method or headers between requests.
	// examples:
	//   - name: URIs
	//     value: ['/get?hello=world', '/request/this?var=value']
	URIs []string `yaml:"uris,omitempty" json:"uris,omitempty" koanf:"uris,omitempty"`

	// description: |
	//   FollowRedirect will expect the previous stage of the same test to have received a
	//   redirect response, it will fail the test otherwise. The redirect location will be used
	//   to send the request for the current stage and any settings for port, protocol, address,
	//   or URI will be ignored.
	// examples:
	//   - name: follow_redirect
	//     value: true
	FollowRedirect *bool `yaml:"follow_redirect,omitempty" json:"follow_redirect,omitempty" koanf:"follow_redirect,omitempty"`

	// description: |
	//   Version allows you to declare the HTTP version the test should use as part of the request line.
	// examples:
	//   - name: Version
	//     value: "\"1.1\""
	Version *string `yaml:"version,omitempty" json:"version,omitempty" koanf:"version,omitempty"`

	// description: |
	//   Method allows you to declare the HTTP method the test should use as part of the request line.
	// examples:
	//   - name: Method
	//     value: "\"GET\""
	Method *string `yaml:"method,omitempty" json:"method,omitempty" koanf:"method,omitempty"`

	// description: |
	//   Headers allows you to declare headers that the test should send.
	// examples:
	//   - name: Headers
	//     value: ExampleHeaders
	//
	// Deprecated: use OrderedHeaders instead
	Headers map[string]string `yaml:"headers,omitempty" json:"headers,omitempty" koanf:"headers,omitempty"`

	// description: |
	//   Headers allows you to declare headers that the test should send.
	//   The headers will be sent in the exact order specified. It is also possible
	//   to specify the identical header multiple times.
	// examples:
	//   - name: Headers
	//     value: ExampleOrderedHeaders
	OrderedHeaders []HeaderTuple `yaml:"ordered_headers,omitempty" json:"ordered_headers,omitempty" koanf:"ordered_headers,omitempty"`

	// description: |
	//   Data allows you to declare the payload that the test should in the request body.
	// examples:
	//   - name: Data
	//     value: "\"Bibitti bopi\""
	Data *string `yaml:"data,omitempty" json:"data,omitempty" koanf:"data,omitempty"`

	// description: |
	//   EncodedData allows you to declare the payload as a base64 encoded string, which
	//   will be decoded into bytes and sent verbatimt to the server. This allows for complex
	//   payloads that include invisible characters or invalid Unicode byte sequences.
	// examples:
	//   - name: encoded_data
	//     value: ExampleEncodedData
	EncodedData *string `yaml:"encoded_data,omitempty" json:"encoded_data,omitempty" koanf:"encoded_data,omitempty"`

	// description: |
	//   SaveCookie allows you to automatically provide cookies if there are multiple stages and save cookie is set
	// examples:
	//   - name: SaveCookie
	//     value: 80
	SaveCookie *bool `yaml:"save_cookie,omitempty" json:"save_cookie,omitempty" koanf:"save_cookie,omitempty"`

	// description: |
	//   StopMagic is deprecated.
	// examples:
	//   - name: StopMagic
	//     value: false
	//
	// Deprecated: use AutocompleteHeaders instead
	StopMagic *bool `yaml:"stop_magic,omitempty" json:"stop_magic,omitempty" koanf:"stop_magic,omitempty"`

	// description: |
	//   AutocompleteHeaders allows the test framework to automatically fill the request with Content-Type and Connection headers.
	//
	//   Defaults: `true`.
	// examples:
	//   - name: StopMagic
	//     value: false
	AutocompleteHeaders *bool `yaml:"autocomplete_headers,omitempty" json:"autocomplete_headers,omitempty" koanf:"autocomplete_headers,omitempty"`

	// description: |
	//   EncodedRequest will take a base64 encoded string that will be decoded and sent through as the request.
	//
	//   It will override all other settings
	// examples:
	//   - name: EncodedRequest
	//     value: "\"a\""
	EncodedRequest string `yaml:"encoded_request,omitempty" json:"encoded_request,omitempty" koanf:"encoded_request,omitempty"`

	// description: |
	//   Response describes a response from the web server that a WAF is expected to analyse.
	//
	//   Note: This functionality requires a backend that can send the specified request to the
	//         reverse proxy. Currently, only Albedo (https://github.com/coreruleset/albedo) is supported.
	// example:
	//   - name: Response
	//     value: ExampleResponse
	Response Response `yaml:"response,omitempty" json:"response,omitempty" koanf:"response,omitempty"`

	// description: |
	//   VirtualHostMode determines the value of the `Host` header for internal requests (e.g., the
	//   requests used to insert markers into the web server log). This is useful for running tests
	//   against a virtual host, as the log entries for all requests must end up in the same log file,
	//   and often, log files are segregated by virtual host.
	//
	//   If `true`, internal requests will use the same value for the `Host` header as the test request.
	//
	//   If `false`, the value for the `Host` header of internal requests will be `localhost`.
	//
	//   Default: `false`.
	// example:
	//   - name: VirtualHostMode
	//     value: true
	VirtualHostMode bool `yaml:"virtual_host_mode,omitempty" json:"virtual_host_mode,omitempty" koanf:"virtual_host_mode,omitempty"`
}

type Response struct {
	// description: |
	//   Headers defines the headers the response will carry.
	// examples:
	//   - name: Headers
	//     value: ExampleHeaders
	Headers map[string]string `yaml:"headers,omitempty" json:"headers,omitempty" koanf:"headers,omitempty"`

	// description: |
	//   Status describes the HTTP status code of the response.
	//
	//   Default: `200` if omitted.
	// examples:
	//   - name: Status
	//     value: 302
	Status int `yaml:"status,omitempty" json:"status,omitempty" koanf:"status,omitempty"`

	// description: |
	//   Body defines the body of the response as a plain string.
	// examples:
	//   - name: Body
	//     value: ExampleResponseBody
	Body string `yaml:"body,omitempty" json:"body,omitempty" koanf:"body,omitempty"`

	// description: |
	//   EncodedBody defines the body of the response as a base64 encoded string. This is useful if the response
	//   needs to contain non-printable characters.
	// examples:
	//   - name: EncodedBody
	//     value: ExampleEncodedResponseBody
	EncodedBody string `yaml:"encoded_body,omitempty" json:"encoded_body,omitempty" koanf:"encoded_body,omitempty"`

	// description: |
	//   LogMessage specifies a message to be printed in the log of the backend server that sends the response.
	//   This can be helpful when debugging, to match responses sent by the backend to test executions.
	// examples:
	//   - name: LogMessage
	//     value: "\"Response splitting test 1\""
	LogMessage string `yaml:"log_message,omitempty" json:"log_message,omitempty" koanf:"log_message,omitempty"`
}

// Output defines the expectations of a test
type Output struct {
	// description: |
	//   Status describes the HTTP status code expected in the response.
	// examples:
	//   - name: Status
	//     value: 200
	Status int `yaml:"status,omitempty" json:"status,omitempty"`

	// description: |
	//   ResponseContains describes the text that should be contained in the HTTP response.
	// examples:
	//   - name: ResponseContains
	//     value: "\"Hello, World\""
	ResponseContains string `yaml:"response_contains,omitempty" json:"response_contains,omitempty"`

	// description: |
	//   LogContains describes the text that should be contained in the WAF logs.
	// examples:
	//   - name: LogContains
	//     value: "\"id 920100\""
	//
	// Deprecated: use Log instead
	LogContains string `yaml:"log_contains,omitempty" json:"log_contains,omitempty"`

	// description: |
	//   NoLogContains describes the text that should not be contained in the WAF logs.
	// examples:
	//   - name: NoLogContains
	//     value: "\"id 920100\""
	//
	// Deprecated: use Log instead
	NoLogContains string `yaml:"no_log_contains,omitempty" json:"no_log_contains,omitempty"`

	// description: |
	//   Log is used to configure expectations about the log contents.
	// examples:
	//   - value: ExampleLog
	Log Log `yaml:"log,omitempty" json:"log,omitempty"`

	// description: |
	//   When `ExpectError` is true, we don't expect an answer from the WAF, just an error.
	// examples:
	//   - name: ExpectError
	//     value: false
	ExpectError *bool `yaml:"expect_error,omitempty" json:"expect_error,omitempty"`

	// description: |
	//   When `RetryOnce` is true, the test run will be retried once upon failures. This options
	//   primary purpose is to work around a race condition in phase 5, where the log entry for
	//   a phase 5 rule may appear after the end marker of the previous test.
	RetryOnce *bool `yaml:"retry_once,omitempty" json:"retry_once,omitempty"`

	// description: |
	//   Isolated specifies that the test is expected to trigger a single rule only.
	//   If the rule triggers any other rule than the (single) one specified in
	//   expect_ids, the test fill be considered a failure.
	//
	//   Default: `false`
	// examples:
	//   - name: Isolated
	//     value: true
	Isolated bool `yaml:"isolated,omitempty" json:"isolated,omitempty" koanf:"isolated, omitempty"`
}

// Log is used to configure expectations about the log contents.
type Log struct {
	// description: |
	//   Expect the given IDs to be contained in the log output.
	// examples:
	//   -value: ExampleLog.ExpectIds
	ExpectIds []uint `yaml:"expect_ids,omitempty" json:"expect_ids,omitempty"`

	// description: |
	//   Expect the given IDs _not_ to be contained in the log output.
	// examples:
	//   - value: ExampleLog.NoExpectIds
	NoExpectIds []uint `yaml:"no_expect_ids,omitempty" json:"no_expect_ids,omitempty"`

	// description: |
	//   Expect the regular expression to match log content for the current types.
	// examples:
	//   - value: ExampleLog.MatchRegex
	MatchRegex string `yaml:"match_regex,omitempty" json:"match_regex,omitempty"`

	// description: |
	//   Expect the regular expression to _not_ match log content for the current types.
	// examples:
	//   - value: ExampleLog.NoMatchRegex
	NoMatchRegex string `yaml:"no_match_regex,omitempty" json:"no_match_regex,omitempty"`
}

// Header represents the (name, value) tuple of an HTTP header
type HeaderTuple struct {
	Name  string
	Value string
}
