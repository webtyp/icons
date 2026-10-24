// Package terminal is the command-prompt glyph — developer tools: the API explorer, a console.
//
// Ref is WASM-safe (just the symbol id). Def() lives in svg.go behind
// //go:build !wasm.
package terminal

import "webtyp.com/svg"

// Ref is the symbol id, for markup: terminal.Ref.Render(class), or href="#terminal".
const Ref = svg.Icon("terminal")
