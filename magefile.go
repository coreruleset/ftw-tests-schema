// Copyright 2023 CRS
// SPDX-License-Identifier: Apache-2.0

//go:build mage
// +build mage

package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"regexp"
	"slices"

	"github.com/invopop/jsonschema"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"

	"github.com/coreruleset/ftw-tests-schema/v2/types"
	"github.com/coreruleset/ftw-tests-schema/v2/types/overrides"
)

var addLicenseVersion = "v1.1.1" // https://github.com/google/addlicense
var golangCILintVer = "v1.63.4"  // https://github.com/golangci/golangci-lint/releases
var gosImportsVer = "v0.3.8"     // https://github.com/rinchsan/gosimports/releases/tag/v0.3.8

var errRunGoModTidy = errors.New("go.mod/sum not formatted, commit changes")
var errNoGitDir = errors.New("no .git directory found")
var errUpdateGeneratedFiles = errors.New("generated files need to be updated")

// Generate Go documentation files for YAML structures
func Generate() error {
	if err := sh.RunV("go", "generate", "./..."); err != nil {
		return err
	}
	return nil
}

// Format formats code in this repository.
func Format() error {
	mg.SerialDeps(Generate)

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

// Lint verifies code quality.
func Lint() error {
	mg.SerialDeps(Generate)

	if sh.Run("git", "diff", "--exit-code", "--", "'*_doc.go'") != nil {
		return errUpdateGeneratedFiles
	}

	if err := sh.RunV("go", "run", fmt.Sprintf("github.com/golangci/golangci-lint/cmd/golangci-lint@%s", golangCILintVer), "run"); err != nil {
		return err
	}

	if err := sh.RunV("go", "mod", "tidy"); err != nil {
		return err
	}

	if sh.Run("git", "diff", "--exit-code", "go.mod", "go.sum") != nil {
		return errRunGoModTidy
	}

	return nil
}

// Test runs all tests.
func Test() error {
	mg.SerialDeps(Generate)

	if err := sh.RunV("go", "test", "./..."); err != nil {
		return err
	}

	return nil
}

// Generate Markdown output (printed to terminal)
func Markdown() error {
	mg.SerialDeps(Generate)
	generatorBinary := "generate-doc-yaml-schema"

	if err := sh.RunV("go", "build", "./cmd/"+generatorBinary); err != nil {
		return err
	}

	defer os.Remove(generatorBinary)

	markdown, err := sh.Output("./" + generatorBinary)
	if err != nil {
		return err
	}
	// Write markdown to file
	fmt.Println(markdown)

	return nil

}

// Check runs lint and tests.
func Check() {
	mg.SerialDeps(Lint, Test)
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
