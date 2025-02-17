// Copyright 2023 OWASP CRS
// SPDX-License-Identifier: Apache-2.0

package types

import "github.com/coreruleset/ftw-tests-schema/v2/internal/helpers"

var (
	ExampleTests = []Test{
		ExampleTest,
	}
	ExampleTest = Test{
		TestTitle:       "123456-1",
		TestDescription: "Unix RCE using `time`",
		Stages:          ExampleStages,
	}
	ExampleStage = Stage{
		Description: "Get cookie from server",
		Input:       ExampleInput,
		Output:      ExampleOutput,
	}
	ExampleStages = []Stage{
		ExampleStage,
	}
	ExampleHeaders = map[string]string{
		"User-Agent": "CRS Tests",
		"Host":       "localhost",
		"Accept":     "*/*",
	}
	ExampleOrderedHeaders = []HeaderTuple{
		{"Host", "localhost"},
		{"User-Agent", "CRS Tests"},
		{"Host", "localhost"},
		{"Accept", "*/*"},
	}
	ExampleInput = Input{
		DestAddr:            helpers.StrPtr("192.168.0.1"),
		Port:                helpers.IntPtr(8080),
		Protocol:            helpers.StrPtr("http"),
		URI:                 helpers.StrPtr("/test"),
		Version:             helpers.StrPtr("HTTP/1.1"),
		Headers:             ExampleHeaders,
		Method:              helpers.StrPtr("REPORT"),
		Data:                nil,
		EncodedRequest:      "TXkgRGF0YQo=",
		SaveCookie:          helpers.BoolPtr(false),
		StopMagic:           helpers.BoolPtr(true),
		AutocompleteHeaders: helpers.BoolPtr(false),
	}
	ExampleOutput = Output{
		Status:           200,
		ResponseContains: "HTTP/1.1",
		LogContains:      "nothing",
		NoLogContains:    "everything",
		Log:              ExampleLog,
		ExpectError:      helpers.BoolPtr(true),
	}
	ExampleLog = Log{
		ExpectIds:    []uint{123456},
		NoExpectIds:  []uint{123456},
		MatchRegex:   `id[:\s"]*123456`,
		NoMatchRegex: `id[:\s"]*123456`,
	}
	AnnotationsExample = map[string]string{
		"os":      "Debian Bullseye",
		"purpose": "L7ASR test suite",
	}
	ReasonExample = "nginx returns 400 when `Content-Length` header is sent in a\n" +
		"`Transfer-Encoding: chunked` request."
	ExampleEncodedData         = "c29tZXRoaW5nIHdpdGgKbmV3bGluZQo="
	ExampleResponseBody        = "{\"aJsonDocument\": [\"in the response\"]}\n"
	ExampleEncodedResponseBody = "eyJhSnNvbkRvY3VtZW50IjogWyJpbiB0aGUgcmVzcG9uc2UiXX0="
	ExampleRespone             = Response{
		Status: 502,
		Headers: map[string]string{
			"Content-Type": "application/problem+json",
			"x-mr-burns":   "excellent",
		},
		Body:        ExampleResponseBody,
		EncodedBody: ExampleEncodedResponseBody,
		LogMessage:  "Response splitting test 1",
	}
)
