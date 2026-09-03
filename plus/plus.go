// Package plus is the plus glyph — additive actions: new record, add row,
// create. Ref is WASM-safe (just the symbol id). Def() lives in svg.go behind
// //go:build !wasm.
package plus

import "github.com/tinywasm/svg"

// Ref is the symbol id, for markup: plus.Ref.Render(class), or href="#plus".
const Ref = svg.Icon("plus")
