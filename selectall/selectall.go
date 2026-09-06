// Package selectall is the double-check glyph — mark every row at once: a
// list's select-all / deselect-all toggle. It is deliberately NOT trash or
// pencil: those name a destructive/edit ACTION, and reusing either one on a
// select-all control had it silently repeating whichever action button
// already carries that glyph elsewhere on the same screen. selectall names
// the SELECTION itself, so it stays the same glyph no matter what the
// selected rows are about to be used for.
//
// Ref is WASM-safe (just the symbol id). Def() lives in svg.go behind
// //go:build !wasm.
package selectall

import "webtyp.com/svg"

// Ref is the symbol id, for markup: selectall.Ref.Render(class), or
// href="#selectall".
const Ref = svg.Icon("selectall")
