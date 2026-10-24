// Package laptop is the laptop glyph — a computer or device registered in the
// system: the devices screen and anything that lists them.
//
// Ref is WASM-safe (just the symbol id). Def() lives in svg.go behind
// //go:build !wasm.
package laptop

import "webtyp.com/svg"

// Ref is the symbol id, for markup: laptop.Ref.Render(class), or href="#laptop".
const Ref = svg.Icon("laptop")
