// Package trash is the trash-can glyph — destructive actions: delete, remove,
// discard. The commit button of a bulk delete and the mark a row carries while
// selected for deletion draw this same glyph, so importing this one package in
// both places is what keeps them from drifting apart.
//
// Ref is WASM-safe (just the symbol id). Def() lives in svg.go behind
// //go:build !wasm.
package trash

import "github.com/tinywasm/svg"

// Ref is the symbol id, for markup: trash.Ref.Render(class), or href="#trash".
const Ref = svg.Icon("trash")
