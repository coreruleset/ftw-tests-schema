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

<code>rule_id</code>  <i>uint</i>

</div>
<div class="dt">

RuleId is the ID of the rule this test targets.



Examples:


```yaml
# RuleId
rule_id: 123456
```


</div>

<hr />

<div class="dd">

<code>tests</code>  <i>[]<a href="#test">Test</a></i>

</div>
<div class="dt">

Tests is a list of FTW tests



Examples:


```yaml
tests:
    - test_title: 123456-1
      ruleid: 0
      test_id: 0
      desc: Unix RCE using `time`
      stages:
        - description: Get cookie from server
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
            response_contains: HTTP/1.1
            log_contains: nothing
            no_log_contains: everything
            log:
                expect_ids:
                    - 123456
                no_expect_ids:
                    - 123456
                match_regex: id[:\s"]*123456
                no_match_regex: id[:\s"]*123456
            expect_error: true
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

<div class="dd">

<code>tags</code>  <i>[]string</i>

</div>
<div class="dt">

description: |
   Tags is list of strings that can be used for arbitrary grouping of tests.
 examples:
   - name: Tags
     value: ["PHP", "bug-123"]

</div>

<hr />





## Test

Appears in:


- <code><a href="#ftwtest">FTWTest</a>.tests</code>


```yaml
- test_title: 123456-1
  ruleid: 0
  test_id: 0
  desc: Unix RCE using `time`
  stages:
    - description: Get cookie from server
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
        response_contains: HTTP/1.1
        log_contains: nothing
        no_log_contains: everything
        log:
            expect_ids:
                - 123456
            no_expect_ids:
                - 123456
            match_regex: id[:\s"]*123456
            no_match_regex: id[:\s"]*123456
        expect_error: true
```



<hr />

<div class="dd">

<code>test_title</code>  <i>string</i>

</div>
<div class="dt">

TestTitle is the title of this particular types. It is used for inclusion/exclusion of each run by the tool.



Examples:


```yaml
test_title: 123456-1
```


</div>

<hr />

<div class="dd">

<code>test_id</code>  <i>uint</i>

</div>
<div class="dt">

TestId is the ID of the test, in relation to `rule_id`.

When this field is not set, the ID will be inferred from the
position.



Examples:


```yaml
# TestId
test_id: 4
```


</div>

<hr />

<div class="dd">

<code>desc</code>  <i>string</i>

</div>
<div class="dt">

TestDescription is the description for this particular test.

Should be used to describe the internals of the specific things this test is targeting.



Examples:


```yaml
desc: Unix RCE using `time`
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
stages:
    - description: Get cookie from server
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
        response_contains: HTTP/1.1
        log_contains: nothing
        no_log_contains: everything
        log:
            expect_ids:
                - 123456
            no_expect_ids:
                - 123456
            match_regex: id[:\s"]*123456
            no_match_regex: id[:\s"]*123456
        expect_error: true
```


</div>

<hr />

<div class="dd">

<code>tags</code>  <i>[]string</i>

</div>
<div class="dt">

description: |
   Tags is list of strings that can be used for arbitrary grouping of tests.
 examples:
   - name: Tags
     value: ["PHP", "bug-123"]

</div>

<hr />





## Stage

Appears in:


- <code><a href="#test">Test</a>.stages</code>


```yaml
- description: Get cookie from server
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
    response_contains: HTTP/1.1
    log_contains: nothing
    no_log_contains: everything
    log:
        expect_ids:
            - 123456
        no_expect_ids:
            - 123456
        match_regex: id[:\s"]*123456
        no_match_regex: id[:\s"]*123456
    expect_error: true
```



<hr />

<div class="dd">

<code>description</code>  <i>string</i>

</div>
<div class="dt">

Describes the purpose of this stage.



Examples:


```yaml
description: Get cookie from server
```


</div>

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
    response_contains: HTTP/1.1
    log_contains: nothing
    no_log_contains: everything
    log:
        expect_ids:
            - 123456
        no_expect_ids:
            - 123456
        match_regex: id[:\s"]*123456
        no_match_regex: id[:\s"]*123456
    expect_error: true
```


</div>

<hr />





## Input

Appears in:


- <code><a href="#stage">Stage</a>.input</code>


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

<code>follow_redirect</code>  <i>bool</i>

</div>
<div class="dt">

FollowRedirect will expect the previous stage of the same test to have received a
redirect response, it will fail the test otherwise. The redirect location will be used
to send the request for the current stage and any settings for port, protocol, address,
or URI will be ignored.



Examples:


```yaml
# follow_redirect
follow_redirect: true
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

Headers allows you to declare headers that the test should send.



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

<code>ordered_headers</code>  <i>[]<a href="#headertuple">HeaderTuple</a></i>

</div>
<div class="dt">

Headers allows you to declare headers that the test should send.
The headers will be sent in the exact order specified. It is also possible
to specify the identical header multiple times.



Examples:


```yaml
# Headers
ordered_headers:
    - name: Host
      value: localhost
    - name: User-Agent
      value: CRS Tests
    - name: Host
      value: localhost
    - name: Accept
      value: '*/*'
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

<code>encoded_data</code>  <i>string</i>

</div>
<div class="dt">

EncodedData allows you to declare the payload as a base64 encoded string, which
will be decoded into bytes and sent verbatimt to the server. This allows for complex
payloads that include invisible characters or invalid Unicode byte sequences.



Examples:


```yaml
# encoded_data
encoded_data: c29tZXRoaW5nIHdpdGgKbmV3bGluZQo=
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

Defaults: `true`.



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

<code>response</code>  <i><a href="#response">Response</a></i>

</div>
<div class="dt">

Response describes a response from the web server that a WAF is expected to analyse.

Note: This functionality requires a backend that can send the specified request to the
      reverse proxy. Currently, only Albedo (https://github.com/coreruleset/albedo) is supported.

</div>

<hr />

<div class="dd">

<code>virtual_host_mode</code>  <i>bool</i>

</div>
<div class="dt">

VirtualHostMode determines the value of the `Host` header for internal requests (e.g., the
requests used to insert markers into the web server log). This is useful for running tests
against a virtual host, as the log entries for all requests must end up in the same log file,
and often, log files are segregated by virtual host.

If `true`, internal requests will use the same value for the `Host` header as the test request.

If `false`, the value for the `Host` header of internal requests will be `localhost`.

Default: `false`.

</div>

<hr />





## HeaderTuple

Appears in:


- <code><a href="#input">Input</a>.ordered_headers</code>


```yaml
# Headers
- name: Host
  value: localhost
- name: User-Agent
  value: CRS Tests
- name: Host
  value: localhost
- name: Accept
  value: '*/*'
```





## Response

Appears in:


- <code><a href="#input">Input</a>.response</code>





<hr />

<div class="dd">

<code>headers</code>  <i>map[string]string</i>

</div>
<div class="dt">

Headers defines the headers the response will carry.



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

<code>status</code>  <i>int</i>

</div>
<div class="dt">

Status describes the HTTP status code of the response.

Default: `200` if omitted.



Examples:


```yaml
# Status
status: 302
```


</div>

<hr />

<div class="dd">

<code>body</code>  <i>string</i>

</div>
<div class="dt">

Body defines the body of the response as a plain string.



Examples:


```yaml
# Body
body: |
    {"aJsonDocument": ["in the response"]}
```


</div>

<hr />

<div class="dd">

<code>encoded_body</code>  <i>string</i>

</div>
<div class="dt">

EncodedBody defines the body of the response as a base64 encoded string. This is useful if the response
needs to contain non-printable characters.



Examples:


```yaml
# EncodedBody
encoded_body: eyJhSnNvbkRvY3VtZW50IjogWyJpbiB0aGUgcmVzcG9uc2UiXX0=
```


</div>

<hr />

<div class="dd">

<code>log_message</code>  <i>string</i>

</div>
<div class="dt">

LogMessage specifies a message to be printed in the log of the backend server that sends the response.
This can be helpful when debugging, to match resopnses sent by the backend to test executions.



Examples:


```yaml
# LogMessage
log_message: Response splitting test 1
```


</div>

<hr />





## Output

Appears in:


- <code><a href="#stage">Stage</a>.output</code>


```yaml
# Output
status: 200
response_contains: HTTP/1.1
log_contains: nothing
no_log_contains: everything
log:
    expect_ids:
        - 123456
    no_expect_ids:
        - 123456
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
    expect_ids:
        - 123456
    no_expect_ids:
        - 123456
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

<div class="dd">

<code>retry_once</code>  <i>bool</i>

</div>
<div class="dt">

When `RetryOnce` is true, the test run will be retried once upon failures. This options
primary purpose is to work around a race condition in phase 5, where the log entry for
a phase 5 rule may appear after the end marker of the previous test.

</div>

<hr />

<div class="dd">

<code>isolated</code>  <i>bool</i>

</div>
<div class="dt">

Isolated specifies that the test is expected to trigger a single rule only.
If the rule triggers any other rule than the (single) one specified in
expect_ids, the test fill be considered a failure.

Default: `false`



Examples:


```yaml
# Isolated
isolated: true
```


</div>

<hr />





## Log

Appears in:


- <code><a href="#output">Output</a>.log</code>


```yaml
expect_ids:
    - 123456
no_expect_ids:
    - 123456
match_regex: id[:\s"]*123456
no_match_regex: id[:\s"]*123456
```



<hr />

<div class="dd">

<code>expect_ids</code>  <i>[]uint</i>

</div>
<div class="dt">

description: |
   Expect the given IDs to be contained in the log output.
 examples:
   -value: ExampleLog.ExpectIds

</div>

<hr />

<div class="dd">

<code>no_expect_ids</code>  <i>[]uint</i>

</div>
<div class="dt">

Expect the given IDs _not_ to be contained in the log output.



Examples:


```yaml
no_expect_ids:
    - 123456
```


</div>

<hr />

<div class="dd">

<code>match_regex</code>  <i>string</i>

</div>
<div class="dt">

Expect the regular expression to match log content for the current types.



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

Expect the regular expression to _not_ match log content for the current types.



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
      test_ids: [4, 6]
      reason: |-
        nginx returns 400 when `Content-Length` header is sent in a
        `Transfer-Encoding: chunked` request.
      output:
        status: 200
        response_contains: HTTP/1.1
        log_contains: nothing
        no_log_contains: everything
        log:
            expect_ids:
                - 123456
            no_expect_ids:
                - 123456
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
  test_ids: [4, 6]
  reason: |-
    nginx returns 400 when `Content-Length` header is sent in a
    `Transfer-Encoding: chunked` request.
  output:
    status: 200
    response_contains: HTTP/1.1
    log_contains: nothing
    no_log_contains: everything
    log:
        expect_ids:
            - 123456
        no_expect_ids:
            - 123456
        match_regex: id[:\s"]*123456
        no_match_regex: id[:\s"]*123456
    expect_error: true
```



<hr />

<div class="dd">

<code>rule_id</code>  <i>uint</i>

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

<code>test_ids</code>  <i>[]uint</i>

</div>
<div class="dt">

IDs of the tests for rule_id that overrides should be applied to.
If this field is not set, the overrides will be applied to all tests of rule_id.



Examples:


```yaml
test_ids:
    - 4
    - 6
```


</div>

<hr />

<div class="dd">

<code>stage_ids</code>  <i>[]uint</i>

</div>
<div class="dt">

IDs of the stages to which overrides should be applied.
Stage IDs listed will be overridden for all test IDs listed in `TestIds`.
If this field is not set, the overrides will be applied to all stages.

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

<code>retry_once</code>  <i>bool</i>

</div>
<div class="dt">

Whether a stage should be retried once in case of failure.
This option is primarily a workaround for a race condition in phase 5,
where the log entry of a rule may be flushed after the test end marker.



Examples:


```yaml
retry_once: true
```


</div>

<hr />

<div class="dd">

<code>output</code>  <i><a href="#typesoutput">types.Output</a></i>

</div>
<div class="dt">

Specifies overrides on the test output.
This definition *replaces* the output definition of the test.



Examples:


```yaml
output:
    status: 200
    response_contains: HTTP/1.1
    log_contains: nothing
    no_log_contains: everything
    log:
        expect_ids:
            - 123456
        no_expect_ids:
            - 123456
        match_regex: id[:\s"]*123456
        no_match_regex: id[:\s"]*123456
    expect_error: true
```


</div>

<hr />





## types.Output
Output defines the expectations of a test

Appears in:


- <code><a href="#testoverride">TestOverride</a>.output</code>


```yaml
status: 200
response_contains: HTTP/1.1
log_contains: nothing
no_log_contains: everything
log:
    expect_ids:
        - 123456
    no_expect_ids:
        - 123456
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

<code>log</code>  <i>Log</i>

</div>
<div class="dt">

Log is used to configure expectations about the log contents.



Examples:


```yaml
log:
    expect_ids:
        - 123456
    no_expect_ids:
        - 123456
    match_regex: id[:\s"]*123456
    no_match_regex: id[:\s"]*123456
```


</div>

<hr />

<div class="dd">

<code>expect_error</code>  <i>types.bool</i>

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

<div class="dd">

<code>retry_once</code>  <i>types.bool</i>

</div>
<div class="dt">

When `RetryOnce` is true, the test run will be retried once upon failures. This options
primary purpose is to work around a race condition in phase 5, where the log entry for
a phase 5 rule may appear after the end marker of the previous test.

</div>

<hr />

<div class="dd">

<code>isolated</code>  <i>bool</i>

</div>
<div class="dt">

Isolated specifies that the test is expected to trigger a single rule only.
If the rule triggers any other rule than the (single) one specified in
expect_ids, the test fill be considered a failure.

Default: `false`



Examples:


```yaml
# Isolated
isolated: true
```


</div>

<hr />





## types.Log






<hr />

<div class="dd">

<code>expect_ids</code>  <i>[]uint</i>

</div>
<div class="dt">

description: |
   Expect the given IDs to be contained in the log output.
 examples:
   -value: ExampleLog.ExpectIds

</div>

<hr />

<div class="dd">

<code>no_expect_ids</code>  <i>[]uint</i>

</div>
<div class="dt">

Expect the given IDs _not_ to be contained in the log output.



Examples:


```yaml
no_expect_ids:
    - 123456
```


</div>

<hr />

<div class="dd">

<code>match_regex</code>  <i>string</i>

</div>
<div class="dt">

Expect the regular expression to match log content for the current types.



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

Expect the regular expression to _not_ match log content for the current types.



Examples:


```yaml
no_match_regex: id[:\s"]*123456
```


</div>

<hr />




