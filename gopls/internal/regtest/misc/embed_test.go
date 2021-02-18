// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package misc

import (
	"testing"

	. "golang.org/x/tools/gopls/internal/regtest"
	"golang.org/x/tools/internal/testenv"
)

func TestMissingPatternDiagnostic(t *testing.T) {
	testenv.NeedsGo1Point(t, 16)
	const files = `
-- go.mod --
module example.com
-- x.go --
package x

import (
	_ "embed"
)

//go:embed NONEXISTENT
var foo string
`
	Run(t, files, func(t *testing.T, env *Env) {
		env.Await(env.DiagnosticAtRegexpWithMessage("x.go", `NONEXISTENT`, "no matching files found"))
	})
}