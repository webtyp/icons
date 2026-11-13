// Package circleinfo is the information glyph — a circle with an "i" inside:
// help about a field or a module. The outline version (ring, dot and stem) is
// drawn as three paths, so it is not the single-path solid family: svg.go calls
// sprite.Define directly.
//
// Ref is WASM-safe (just the symbol id). Def() lives in svg.go behind
// //go:build !wasm.
package circleinfo

import "webtyp.com/svg"

// Ref is the symbol id, for markup: circleinfo.Ref.Render(class), or href="#circleinfo".
const Ref = svg.Icon("circleinfo")
