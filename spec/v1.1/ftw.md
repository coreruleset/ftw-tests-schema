



## FTWTest
Welcome to the FTW YAMLFormat documentation.
 In this document we will explain all the possible options that can be used within the YAML format.
 Generally this is the preferred format for writing tests in as they don't require any programming skills
 in order to understand and change. If you find a bug in this format please open an issue.


 FTWTest is the base type used when unmarshaling YAML tests files






<hr />

<div class="dd">

<code>meta</code>  <i><a href="#ftwtestmeta">FTWTestMeta</a></i>

</div>
<div class="dt">

Meta describes the metadata information of this yaml test file

</div>

<hr />

<div class="dd">

<code>tests</code>  <i>[]<a href="#test">Test</a></i>

</div>
<div class="dt">

Tests is a list of FTW tests



Examples:


```yaml
# Tests
tests: the tests
```


</div>

<hr />





## FTWTestMeta

Appears in:


- <code><a href="#ftwtest">FTWTest</a>.meta</code>





<hr />

<div class="dd">

<code>author</code>  <i>string</i>

</div>
<div class="dt">

Author is the list of authors that added content to this file



Examples:


```yaml
# Author
author: Felipe Zipitria
```


</div>

<hr />

<div class="dd">

<code>enabled</code>  <i>bool</i>

</div>
<div class="dt">

Enabled indicates if the tests are enabled to be run by the engine or not.



Examples:


```yaml
# Enabled
enabled: false
```


</div>

<hr />

<div class="dd">

<code>name</code>  <i>string</i>

</div>
<div class="dt">

Name is the name of the tests contained in this file.



Examples:


```yaml
# Name
name: test01
```


</div>

<hr />

<div class="dd">

<code>description</code>  <i>string</i>

</div>
<div class="dt">

Description is a textual description of the tests contained in this file.



Examples:


```yaml
# Description
description: The tests here target SQL injection.
```


</div>

<hr />

<div class="dd">

<code>version</code>  <i>string</i>

</div>
<div class="dt">

Version is the version of the YAML Schema.



Examples:


```yaml
# Version
version: v1
```


</div>

<hr />





## Test

Appears in:


- <code><a href="#ftwtest">FTWTest</a>.tests</code>


```yaml
# Tests
the tests
```



<hr />

<div class="dd">

<code>test_title</code>  <i>string</i>

</div>
<div class="dt">

TestTitle the title of this particular test. It is used for inclusion/exclusion of each run by the tool.



Examples:


```yaml
# TestTitle
test_title: 920100-1
```


</div>

<hr />

<div class="dd">

<code>desc</code>  <i>string</i>

</div>
<div class="dt">

TestDescription is the description for this particular test. Should be used to describe the internals of
the specific things this test is targeting.



Examples:


```yaml
# TestDescription
desc: This test targets something
```


</div>

<hr />

<div class="dd">

<code>stages</code>  <i>[]<a href="#stage">Stage</a></i>

</div>
<div class="dt">

Stages is the list of all the stages to perform this test.



Examples:


```yaml
# Stages
stages:
    - stage:
        input:
            dest_addr: 192.168.0.1
            port: 8080
            protocol: http
            uri: /test
            version: HTTP/1.1
            method: REPORT
            headers:
                Accept: '*/*'
                Host: localhost
                User-Agent: CRS Tests
            save_cookie: false
            stop_magic: true
            autocomplete_headers: false
            encoded_request: TXkgRGF0YQo=
        output:
            status: 200
            log_contains: nothing
            log:
                expect_id: 123456
                expect_id: 123456
                match_regex: id[:\s"]*123456
                no_match_regex: id[:\s"]*123456
            expect_error: true
```


</div>

<hr />





## Stage

Appears in:


- <code><a href="#test">Test</a>.stages</code>


```yaml
# Stages
- stage:
    input:
        dest_addr: 192.168.0.1
        port: 8080
        protocol: http
        uri: /test
        version: HTTP/1.1
        method: REPORT
        headers:
            Accept: '*/*'
            Host: localhost
            User-Agent: CRS Tests
        save_cookie: false
        stop_magic: true
        autocomplete_headers: false
        encoded_request: TXkgRGF0YQo=
    output:
        status: 200
        log_contains: nothing
        log:
            expect_id: 123456
            expect_id: 123456
            match_regex: id[:\s"]*123456
            no_match_regex: id[:\s"]*123456
        expect_error: true
```



<hr />

<div class="dd">

<code>stage</code>  <i><a href="#stagedata">StageData</a></i>

</div>
<div class="dt">

StageData is an individual test stage



Examples:


```yaml
# StageData
stage:
    input:
        dest_addr: 192.168.0.1
        port: 8080
        protocol: http
        uri: /test
        version: HTTP/1.1
        method: REPORT
        headers:
            Accept: '*/*'
            Host: localhost
            User-Agent: CRS Tests
        save_cookie: false
        stop_magic: true
        autocomplete_headers: false
        encoded_request: TXkgRGF0YQo=
    output:
        status: 200
        log_contains: nothing
        log:
            expect_id: 123456
            expect_id: 123456
            match_regex: id[:\s"]*123456
            no_match_regex: id[:\s"]*123456
        expect_error: true
```


</div>

<hr />





## StageData

Appears in:


- <code><a href="#stage">Stage</a>.stage</code>


```yaml
# StageData
input:
    dest_addr: 192.168.0.1
    port: 8080
    protocol: http
    uri: /test
    version: HTTP/1.1
    method: REPORT
    headers:
        Accept: '*/*'
        Host: localhost
        User-Agent: CRS Tests
    save_cookie: false
    stop_magic: true
    autocomplete_headers: false
    encoded_request: TXkgRGF0YQo=
output:
    status: 200
    log_contains: nothing
    log:
        expect_id: 123456
        expect_id: 123456
        match_regex: id[:\s"]*123456
        no_match_regex: id[:\s"]*123456
    expect_error: true
```



<hr />

<div class="dd">

<code>input</code>  <i><a href="#input">Input</a></i>

</div>
<div class="dt">

Input is the data that is passed to the test



Examples:


```yaml
# Input
input:
    dest_addr: 192.168.0.1
    port: 8080
    protocol: http
    uri: /test
    version: HTTP/1.1
    method: REPORT
    headers:
        Accept: '*/*'
        Host: localhost
        User-Agent: CRS Tests
    save_cookie: false
    stop_magic: true
    autocomplete_headers: false
    encoded_request: TXkgRGF0YQo=
```


</div>

<hr />

<div class="dd">

<code>output</code>  <i><a href="#output">Output</a></i>

</div>
<div class="dt">

Output is the data that is returned from the test



Examples:


```yaml
# Output
output:
    status: 200
    log_contains: nothing
    log:
        expect_id: 123456
        expect_id: 123456
        match_regex: id[:\s"]*123456
        no_match_regex: id[:\s"]*123456
    expect_error: true
```


</div>

<hr />





## Input

Appears in:


- <code><a href="#stagedata">StageData</a>.input</code>


```yaml
# Input
dest_addr: 192.168.0.1
port: 8080
protocol: http
uri: /test
version: HTTP/1.1
method: REPORT
headers:
    Accept: '*/*'
    Host: localhost
    User-Agent: CRS Tests
save_cookie: false
stop_magic: true
autocomplete_headers: false
encoded_request: TXkgRGF0YQo=
```



<hr />

<div class="dd">

<code>dest_addr</code>  <i>string</i>

</div>
<div class="dt">

DestAddr is the IP of the destination host that the test will send the message to.



Examples:


```yaml
# DestAddr
dest_addr: 127.0.0.1
```


</div>

<hr />

<div class="dd">

<code>port</code>  <i>int</i>

</div>
<div class="dt">

Port allows you to declare which port on the destination host the test should connect to.



Examples:


```yaml
# Port
port: 80
```


</div>

<hr />

<div class="dd">

<code>protocol</code>  <i>string</i>

</div>
<div class="dt">

Protocol allows you to declare which protocol the test should use when sending the request.



Examples:


```yaml
# Protocol
protocol: http
```


</div>

<hr />

<div class="dd">

<code>uri</code>  <i>string</i>

</div>
<div class="dt">

URI allows you to declare the URI the test should use as part of the request line.



Examples:


```yaml
# URI
uri: /get?hello=world
```


</div>

<hr />

<div class="dd">

<code>version</code>  <i>string</i>

</div>
<div class="dt">

Version allows you to declare the HTTP version the test should use as part of the request line.



Examples:


```yaml
# Version
version: "1.1"
```


</div>

<hr />

<div class="dd">

<code>method</code>  <i>string</i>

</div>
<div class="dt">

Method allows you to declare the HTTP method the test should use as part of the request line.



Examples:


```yaml
# Method
method: GET
```


</div>

<hr />

<div class="dd">

<code>headers</code>  <i>map[string]string</i>

</div>
<div class="dt">

Method allows you to declare headers that the test should send.



Examples:


```yaml
# Headers
headers:
    Accept: '*/*'
    Host: localhost
    User-Agent: CRS Tests
```


</div>

<hr />

<div class="dd">

<code>data</code>  <i>string</i>

</div>
<div class="dt">

Data allows you to declare the payload that the test should in the request body.



Examples:


```yaml
# Data
data: Bibitti bopi
```


</div>

<hr />

<div class="dd">

<code>save_cookie</code>  <i>bool</i>

</div>
<div class="dt">

SaveCookie allows you to automatically provide cookies if there are multiple stages and save cookie is set



Examples:


```yaml
# SaveCookie
save_cookie: 80
```


</div>

<hr />

<div class="dd">

<code>stop_magic</code>  <i>bool</i>

</div>
<div class="dt">

StopMagic is deprecated.



Examples:


```yaml
# StopMagic
stop_magic: false
```


</div>

<hr />

<div class="dd">

<code>autocomplete_headers</code>  <i>bool</i>

</div>
<div class="dt">

AutocompleteHeaders allows the test framework to automatically fill the request with Content-Type and Connection headers.
Defaults to true.



Examples:


```yaml
# StopMagic
autocomplete_headers: false
```


</div>

<hr />

<div class="dd">

<code>encoded_request</code>  <i>string</i>

</div>
<div class="dt">

EncodedRequest will take a base64 encoded string that will be decoded and sent through as the request.
It will override all other settings



Examples:


```yaml
# EncodedRequest
encoded_request: a
```


</div>

<hr />

<div class="dd">

<code>raw_request</code>  <i>string</i>

</div>
<div class="dt">

RAWRequest is deprecated.



Examples:


```yaml
# RAWRequest
raw_request: TXkgRGF0YQo=
```


</div>

<hr />





## Output

Appears in:


- <code><a href="#stagedata">StageData</a>.output</code>


```yaml
# Output
status: 200
log_contains: nothing
log:
    expect_id: 123456
    expect_id: 123456
    match_regex: id[:\s"]*123456
    no_match_regex: id[:\s"]*123456
expect_error: true
```



<hr />

<div class="dd">

<code>status</code>  <i>int</i>

</div>
<div class="dt">

Status describes the HTTP status code expected in the response.



Examples:


```yaml
# Status
status: 200
```


</div>

<hr />

<div class="dd">

<code>response_contains</code>  <i>string</i>

</div>
<div class="dt">

ResponseContains describes the text that should be contained in the HTTP response.



Examples:


```yaml
# ResponseContains
response_contains: Hello, World
```


</div>

<hr />

<div class="dd">

<code>log_contains</code>  <i>string</i>

</div>
<div class="dt">

LogContains describes the text that should be contained in the WAF logs.



Examples:


```yaml
# LogContains
log_contains: id 920100
```


</div>

<hr />

<div class="dd">

<code>no_log_contains</code>  <i>string</i>

</div>
<div class="dt">

NoLogContains describes the text that should not be contained in the WAF logs.



Examples:


```yaml
# NoLogContains
no_log_contains: id 920100
```


</div>

<hr />

<div class="dd">

<code>log</code>  <i><a href="#log">Log</a></i>

</div>
<div class="dt">

Log is used to configure expectations about the log contents.



Examples:


```yaml
log:
    expect_id: 123456
    expect_id: 123456
    match_regex: id[:\s"]*123456
    no_match_regex: id[:\s"]*123456
```


</div>

<hr />

<div class="dd">

<code>expect_error</code>  <i>bool</i>

</div>
<div class="dt">

When `ExpectError` is true, we don't expect an answer from the WAF, just an error.



Examples:


```yaml
# ExpectError
expect_error: false
```


</div>

<hr />





## Log

Appears in:


- <code><a href="#output">Output</a>.log</code>


```yaml
expect_id: 123456
expect_id: 123456
match_regex: id[:\s"]*123456
no_match_regex: id[:\s"]*123456
```



<hr />

<div class="dd">

<code>expect_id</code>  <i>int</i>

</div>
<div class="dt">

description: |
   Expect the given ID to be contained in the log output.
 examples:
   - exampleLog.ExpectId

</div>

<hr />

<div class="dd">

<code>expect_id</code>  <i>int</i>

</div>
<div class="dt">

description: |
   Expect the given ID _not_ to be contained in the log output.
 examples:
   - exampleLog.NoExpectId

</div>

<hr />

<div class="dd">

<code>match_regex</code>  <i>string</i>

</div>
<div class="dt">

Expect the regular expression to match log content for the current test.



Examples:


```yaml
match_regex: id[:\s"]*123456
```


</div>

<hr />

<div class="dd">

<code>no_match_regex</code>  <i>string</i>

</div>
<div class="dt">

Expect the regular expression to _not_ match log content for the current test.



Examples:


```yaml
no_match_regex: id[:\s"]*123456
```


</div>

<hr />








## FTWOverrides
FTWOverrides describes platform specific overrides for tests






<hr />

<div class="dd">

<code>version</code>  <i>string</i>

</div>
<div class="dt">

The version field designates the version of the schema that validates this file



Examples:


```yaml
version: v0.1.0
```


</div>

<hr />

<div class="dd">

<code>meta</code>  <i><a href="#ftwoverridesmeta">FTWOverridesMeta</a></i>

</div>
<div class="dt">

Meta describes the metadata information



Examples:


```yaml
meta:
    engine: libmodsecurity3
    platform: nginx
    annotations:
        os: Debian Bullseye
        purpose: L7ASR test suite
```


</div>

<hr />

<div class="dd">

<code>test_overrides</code>  <i>[]<a href="#testoverride">TestOverride</a></i>

</div>
<div class="dt">

List of test override specifications



Examples:


```yaml
test_overrides:
    - rule_id: 920100
      test_ids:
        - 4
        - 6
      reason: |-
        nginx returns 400 when `Content-Length` header is sent in a
        `Transfer-Encoding: chunked` request.
      expect_failure: true
      output:
        status: 200
        log_contains: nothing
        log:
            expect_id: 123456
            expect_id: 123456
            match_regex: id[:\s"]*123456
            no_match_regex: id[:\s"]*123456
        expect_error: true
```


</div>

<hr />





## FTWOverridesMeta

Appears in:


- <code><a href="#ftwoverrides">FTWOverrides</a>.meta</code>


```yaml
engine: libmodsecurity3
platform: nginx
annotations:
    os: Debian Bullseye
    purpose: L7ASR test suite
```



<hr />

<div class="dd">

<code>engine</code>  <i>string</i>

</div>
<div class="dt">

The name of the WAF engine the tests are expected to run against



Examples:


```yaml
engine: coraza
```


</div>

<hr />

<div class="dd">

<code>platform</code>  <i>string</i>

</div>
<div class="dt">

The name of the platform (e.g., web server) the tests are expected to run against



Examples:


```yaml
platform: nginx
```


</div>

<hr />

<div class="dd">

<code>annotations</code>  <i>map[string]string</i>

</div>
<div class="dt">

Custom annotations; can be used to add additional meta information



Examples:


```yaml
annotations:
    os: Debian Bullseye
    purpose: L7ASR test suite
```


</div>

<hr />





## TestOverride

Appears in:


- <code><a href="#ftwoverrides">FTWOverrides</a>.test_overrides</code>


```yaml
- rule_id: 920100
  test_ids:
    - 4
    - 6
  reason: |-
    nginx returns 400 when `Content-Length` header is sent in a
    `Transfer-Encoding: chunked` request.
  expect_failure: true
  output:
    status: 200
    log_contains: nothing
    log:
        expect_id: 123456
        expect_id: 123456
        match_regex: id[:\s"]*123456
        no_match_regex: id[:\s"]*123456
    expect_error: true
```



<hr />

<div class="dd">

<code>rule_id</code>  <i>int</i>

</div>
<div class="dt">

ID of the rule this test targets.



Examples:


```yaml
rule_id: 920100
```


</div>

<hr />

<div class="dd">

<code>test_ids</code>  <i>[]int</i>

</div>
<div class="dt">

description: |
     IDs of the tests for rule_id that overrides should be applied to.
     If this field is not set, the overrides will be applied to all tests of rule_id.
 examples:
     - value: [5, 6]

</div>

<hr />

<div class="dd">

<code>reason</code>  <i>string</i>

</div>
<div class="dt">

Describes why this override is necessary.



Examples:


```yaml
reason: |-
    nginx returns 400 when `Content-Length` header is sent in a
    `Transfer-Encoding: chunked` request.
```


</div>

<hr />

<div class="dd">

<code>expect_failure</code>  <i>bool</i>

</div>
<div class="dt">

Whether this test is expected to fail for this particular configuration.
Default: false



Examples:


```yaml
expect_failure: true
```


</div>

<hr />

<div class="dd">

<code>output</code>  <i><a href="#output">Output</a></i>

</div>
<div class="dt">

Specifies overrides on the test output



Examples:


```yaml
output: 400
```


</div>

<hr />





## Output

Appears in:


- <code><a href="#testoverride">TestOverride</a>.output</code>


```yaml
400
```



<hr />

<div class="dd">

<code>status</code>  <i>int</i>

</div>
<div class="dt">

Status describes the HTTP status code expected in the response.



Examples:


```yaml
# Status
status: 200
```


</div>

<hr />

<div class="dd">

<code>response_contains</code>  <i>string</i>

</div>
<div class="dt">

ResponseContains describes the text that should be contained in the HTTP response.



Examples:


```yaml
# ResponseContains
response_contains: Hello, World
```


</div>

<hr />

<div class="dd">

<code>log_contains</code>  <i>string</i>

</div>
<div class="dt">

LogContains describes the text that should be contained in the WAF logs.



Examples:


```yaml
# LogContains
log_contains: id 920100
```


</div>

<hr />

<div class="dd">

<code>no_log_contains</code>  <i>string</i>

</div>
<div class="dt">

NoLogContains describes the text that should not be contained in the WAF logs.



Examples:


```yaml
# NoLogContains
no_log_contains: id 920100
```


</div>

<hr />

<div class="dd">

<code>log</code>  <i><a href="#log">Log</a></i>

</div>
<div class="dt">

Log is used to configure expectations about the log contents.



Examples:


```yaml
log:
    expect_id: 123456
    expect_id: 123456
    match_regex: id[:\s"]*123456
    no_match_regex: id[:\s"]*123456
```


</div>

<hr />

<div class="dd">

<code>expect_error</code>  <i>bool</i>

</div>
<div class="dt">

When `ExpectError` is true, we don't expect an answer from the WAF, just an error.



Examples:


```yaml
# ExpectError
expect_error: false
```


</div>

<hr />





## Log

Appears in:


- <code><a href="#output">Output</a>.log</code>


```yaml
expect_id: 123456
expect_id: 123456
match_regex: id[:\s"]*123456
no_match_regex: id[:\s"]*123456
```



<hr />

<div class="dd">

<code>expect_id</code>  <i>int</i>

</div>
<div class="dt">

description: |
   Expect the given ID to be contained in the log output.
 examples:
   - exampleLog.ExpectId

</div>

<hr />

<div class="dd">

<code>expect_id</code>  <i>int</i>

</div>
<div class="dt">

description: |
   Expect the given ID _not_ to be contained in the log output.
 examples:
   - exampleLog.NoExpectId

</div>

<hr />

<div class="dd">

<code>match_regex</code>  <i>string</i>

</div>
<div class="dt">

Expect the regular expression to match log content for the current test.



Examples:


```yaml
match_regex: id[:\s"]*123456
```


</div>

<hr />

<div class="dd">

<code>no_match_regex</code>  <i>string</i>

</div>
<div class="dt">

Expect the regular expression to _not_ match log content for the current test.



Examples:


```yaml
no_match_regex: id[:\s"]*123456
```


</div>

<hr />




