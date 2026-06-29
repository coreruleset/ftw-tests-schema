# FTW Tests YAML Schema

[![Lint and Generate](https://github.com/coreruleset/ftw-tests-schema/actions/workflows/lint.yml/badge.svg)](https://github.com/coreruleset/ftw-tests-schema/actions/workflows/lint.yml)

This repository contains the YAML schema used by the [FTW](https://github.com/coreruleset/go-ftw) tool for WAF testing. The schema is maintained as Go types so it can be imported directly as a library.

## v3 format (current)

v3 introduces Go template support to make test files significantly simpler. Instead of repeating the full HTTP request structure in every test, you define a file-level `template` once and only write the `payload` per test.

```yaml
meta:
  author: "you"
  description: "SQL injection tests for rule 942100."
rule_id: 942100
template:
  method: GET
  uri: "/?q={{.Payload}}"
  headers:
    - name: Host
      value: localhost
    - name: User-Agent
      value: "CRS Tests"
tests:
  - id: 1
    description: "UNION-based injection"
    payload: "1 UNION SELECT 1,2,3--"
    output:
      log:
        expect_ids: [942100]
  - id: 2
    description: "Benign query (should not trigger)"
    payload: "hello world"
    output:
      log:
        no_expect_ids: [942100]
```

### Per-test template override

A test can override individual fields from the file-level template via its own `template` section. Only fields that are set in the per-test template take precedence; the rest fall through to the file-level template.

```yaml
tests:
  - id: 3
    description: "Same payload injected via POST body instead"
    template:
      method: POST
      uri: /submit
      data: "q={{.Payload}}"
    payload: "1 UNION SELECT 1,2,3--"
    output:
      log:
        expect_ids: [942100]
```

### Template variables

The following template variables are available inside `uri`, `data`, and header values:

| Variable | Description |
|----------|-------------|
| `{{.Payload}}` | The raw `payload` string from the test |

### Field name aliases

In v3, `id` and `description` are the canonical field names. The v2 names are accepted as deprecated aliases:

| v3 canonical | Deprecated v2 alias |
|--------------|---------------------|
| `id` | `test_id` |
| `description` | `desc` |

## v2 format

v2 is the prior stable format. Every test specifies its HTTP request in a `stages` list, with one `input`/`output` pair per stage. Multi-stage tests (e.g. fetch a cookie, then use it) are expressed as multiple entries in the same `stages` list.

```yaml
meta:
  author: "you"
  description: "SQL injection tests for rule 942100."
rule_id: 942100
tests:
  - test_id: 1
    desc: "UNION-based injection via GET parameter"
    stages:
      - input:
          dest_addr: "127.0.0.1"
          port: 80
          protocol: http
          method: GET
          uri: "/?q=1 UNION SELECT 1,2,3--"
          headers:
            Host: localhost
            User-Agent: "CRS Tests"
        output:
          log:
            expect_ids: [942100]
  - test_id: 2
    desc: "Benign query (should not trigger)"
    stages:
      - input:
          dest_addr: "127.0.0.1"
          port: 80
          protocol: http
          method: GET
          uri: "/?q=hello world"
          headers:
            Host: localhost
            User-Agent: "CRS Tests"
        output:
          log:
            no_expect_ids: [942100]
```

### Multi-stage example

```yaml
  - test_id: 3
    desc: "Obtain session cookie, then send the attack"
    stages:
      - input:
          dest_addr: "127.0.0.1"
          port: 80
          method: GET
          uri: /login
          headers:
            Host: localhost
          save_cookie: true
        output:
          status: 200
      - input:
          dest_addr: "127.0.0.1"
          port: 80
          method: POST
          uri: /search
          headers:
            Host: localhost
            Content-Type: application/x-www-form-urlencoded
          data: "q=1 UNION SELECT 1,2,3--"
        output:
          log:
            expect_ids: [942100]
```

## Backward compatibility

v3 is fully backward-compatible with v2 test files. The explicit `stages` format continues to work unchanged:

```yaml
rule_id: 123456
tests:
  - test_id: 1
    desc: "Complex multi-stage test"
    stages:
      - input:
          dest_addr: "127.0.0.1"
          port: 80
          method: POST
          headers:
            Host: localhost
          data: "some data"
        output:
          log:
            expect_ids: [123456]
```

Deprecated fields (`stop_magic`, `log_contains`, `no_log_contains`, `headers` map, `enabled`) are still parsed so existing test files produce correct results without modification.

## Using as a Go library

```go
import "github.com/coreruleset/ftw-tests-schema/v3/types"
```

```go
var test types.FTWTest
if err := yaml.Unmarshal(data, &test); err != nil {
    return err
}
```

## Development

### Prerequisites

```bash
go install github.com/projectdiscovery/yamldoc-go/cmd/docgen/dstdocgen@latest
```

### Common tasks

| Command | Description |
|---------|-------------|
| `mage generate` | Regenerate `test_doc.go` and `overrides_doc.go` from struct comments |
| `mage test` | Run the test suite |
| `mage format` | Format code, tidy modules, apply license headers |
| `mage jsonSchemas` | Write JSON schemas for the current version to `spec/` |
| `mage markdown` | Print markdown documentation to stdout |

### Generating JSON schemas for a new version

Create the version directory first, then run the schema generator:

```bash
mkdir spec/v3.x.x
mage jsonSchemas
```

## Schema versions

| Version | Spec | Comment |
|---------|------|---------|
| v3.0.0 | [spec/v3.0.0](spec/v3.0.0) | Go template support (`{{.Payload}}`), file-level `template`, per-test `payload` and `output`; `id`/`description` canonical (`test_id`/`desc` deprecated); module path → `v3` |
| v2.3.0 | [spec/v2.3.0](spec/v2.3.0) | Dependency updates; no schema changes |
| v2.2.0 | [spec/v2.2.0](spec/v2.2.0) | Added `ordered_headers` (`[]HeaderTuple`) to preserve header order and allow duplicate header names; `headers` map deprecated |
| v2.1.1 | [spec/v2.1.1](spec/v2.1.1) | Added `virtual_host_mode` to `Input` for tests running against a virtual host |
| v2.1.0 | [spec/v2.1.0](spec/v2.1.0) | Added `Response` object inside `Input` for Albedo-backed response-reflection tests |
| v2.0.0 | [spec/v2.0.0](spec/v2.0.0) | Major restructure: `rule_id` moved to file level, `tags` added, platform overrides schema introduced, JSON Schema files added |
| v1.1.0 | [spec/v1.1.0](spec/v1.1.0) | Added `follow_redirect`, `encoded_data`, `retry_once`, `isolated`; log IDs changed from singular to arrays (`expect_ids`/`no_expect_ids`) |
| v1.0.0 | [spec/v1.0.0](spec/v1.0.0) | Initial release |
