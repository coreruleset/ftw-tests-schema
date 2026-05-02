// Copyright 2023 CRS
// SPDX-License-Identifier: Apache-2.0

//go:build mage
// +build mage

package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"regexp"
	"slices"

	"github.com/invopop/jsonschema"
	"github.com/magefile/mage/sh"

	"github.com/coreruleset/ftw-tests-schema/v2/types"
	"github.com/coreruleset/ftw-tests-schema/v2/types/overrides"
)

var addLicenseVersion = "v1.1.1" // https://github.com/google/addlicense
var gosImportsVer = "v0.3.8"     // https://github.com/rinchsan/gosimports/releases/tag/v0.3.8

// Format formats code in this repository.
func Format() error {
	if err := sh.RunV("go", "mod", "tidy"); err != nil {
		return err
	}

	// addlicense strangely logs skipped files to stderr despite not being erroneous, so use the long sh.Exec form to
	// discard stderr too.
	if _, err := sh.Exec(map[string]string{}, io.Discard, io.Discard, "go", "run", fmt.Sprintf("github.com/google/addlicense@%s", addLicenseVersion),
		"-c", "OWASP CRS",
		"-s=only",
		"-ignore", "**/*.yml",
		"-ignore", "**/*.yaml",
		"-ignore", "v?/**", "."); err != nil {
		return err
	}
	return sh.RunV("go", "run", fmt.Sprintf("github.com/rinchsan/gosimports/cmd/gosimports@%s", gosImportsVer),
		"-w",
		"-local",
		"github.com/coreruleset/ftw-tests-schema",
		".")
}

// Test runs all tests.
func Test() error {
	if err := sh.RunV("go", "test", "./..."); err != nil {
		return err
	}

	return nil
}

// Writes JSON schemas based on the current version.
// Make sure the update https://github.com/SchemaStore/schemastore when
// you create a new version.
func JsonSchemas() {
	specsDir := "spec"
	testJson, err := jsonschema.Reflect(&types.FTWTest{}).MarshalJSON()
	if err != nil {
		fmt.Print(err.Error())
	}

	overridesJson, err := jsonschema.Reflect(&overrides.FTWOverrides{}).MarshalJSON()
	if err != nil {
		fmt.Print(err.Error())
	}

	entries, err := os.ReadDir(specsDir)
	if err != nil {
		fmt.Print(err.Error())
	}
	slices.SortFunc(entries, func(a fs.DirEntry, b fs.DirEntry) int {
		// sort v2.0 before v1
		if a.Name() < b.Name() {
			return 1
		} else if a.Name() > b.Name() {
			return -1
		} else {
			return 0
		}
	})
	versionRegex := regexp.MustCompile(`v\d+(\.\d+)?`)
	outputDir := "."
	for _, entry := range entries {
		if entry.IsDir() && versionRegex.MatchString(entry.Name()) {
			outputDir = entry.Name()
			break
		}
	}
	err = os.WriteFile(path.Join(specsDir, outputDir, fmt.Sprintf("waf-tests-schema-%s.json", outputDir)), testJson, fs.ModePerm)
	if err != nil {
		fmt.Print(err.Error())
	}

	err = os.WriteFile(path.Join(specsDir, outputDir, fmt.Sprintf("waf-platform-overrides-schema-%s.json", outputDir)), overridesJson, fs.ModePerm)
	if err != nil {
		fmt.Print(err.Error())
	}
}
