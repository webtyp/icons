// Package location is the map-pin glyph — a place or zone: where something is.
//
// Ref is WASM-safe (just the symbol id). Def() lives in svg.go behind
// //go:build !wasm.
package location

import "webtyp.com/svg"

// Ref is the symbol id, for markup: location.Ref.Render(class), or href="#location".
const Ref = svg.Icon("location")
