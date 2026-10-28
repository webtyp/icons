// Package code is the code-brackets glyph (</>) — an API, source code, the API explorer.
//
// Ref is WASM-safe (just the symbol id). Def() lives in svg.go behind
// //go:build !wasm.
package code

import "webtyp.com/svg"

// Ref is the symbol id, for markup: code.Ref.Render(class), or href="#code".
const Ref = svg.Icon("code")
