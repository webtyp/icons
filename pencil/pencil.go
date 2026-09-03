// Package pencil is the pencil glyph — edit actions: modify, rename, correct.
// The commit button of a bulk edit and the mark a row carries while selected
// for editing draw this same glyph; importing this one package in both places
// keeps them in sync.
//
// Ref is WASM-safe (just the symbol id). Def() lives in svg.go behind
// //go:build !wasm.
package pencil

import "github.com/tinywasm/svg"

// Ref is the symbol id, for markup: pencil.Ref.Render(class), or href="#pencil".
const Ref = svg.Icon("pencil")
