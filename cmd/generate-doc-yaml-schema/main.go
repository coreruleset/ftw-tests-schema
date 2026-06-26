// Copyright 2023 OWASP CRS
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"

	"github.com/coreruleset/ftw-tests-schema/v3/types"
	"github.com/coreruleset/ftw-tests-schema/v3/types/overrides"
)

func main() {
	data, err := types.GetFTWTestDoc().Encode()
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		panic(err)
	}

	data, err = overrides.GetFTWOverridesDoc().Encode()
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		panic(err)
	}
}
