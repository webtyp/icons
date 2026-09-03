// Package undo is the counter-clockwise arrow glyph — reverse the last move:
// cancel, revert, back out of a mode. Ref is WASM-safe (just the symbol id).
// Def() lives in svg.go behind //go:build !wasm.
package undo

import "github.com/tinywasm/svg"

// Ref is the symbol id, for markup: undo.Ref.Render(class), or href="#undo".
const Ref = svg.Icon("undo")
