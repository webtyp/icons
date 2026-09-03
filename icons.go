//go:build !wasm

// Package icons is the shared builder for TinyWasm's per-glyph icon packages.
//
// It is NOT the package a consumer imports. A consumer imports one glyph
// subpackage — github.com/tinywasm/icons/trash, .../pencil, .../plus,
// .../undo — and takes two things from it:
//
//   - Ref  (github.com/tinywasm/svg.Icon) for markup: Ref.Render(class), or a
//     bare href="#<id>". WASM-safe: only the id string ever reaches the browser.
//   - Def() (github.com/tinywasm/svg/sprite.Definition) for the backend sprite
//     a component ships from its own IconSvg(). Lives behind //go:build !wasm
//     in each glyph package, so importing a glyph for its Ref never drags the
//     sprite geometry (and tinywasm/json + tinywasm/model with it) into a WASM
//     bundle.
//
// This package exists only so a glyph package's svg.go is one line and the
// glyph-family conventions (single closed path, currentColor fill, the
// FontAwesome-style viewBox) live in exactly one place instead of being
// re-typed in every glyph.
package icons

import (
	"github.com/tinywasm/svg"
	"github.com/tinywasm/svg/sprite"
)

// Solid builds a single-path solid glyph definition: one closed <path> whose
// fill is driven by currentColor at the use-site (so a white-text box yields a
// white glyph, a red box a red glyph — the box owns the colour, never the
// path). ref is the symbol id, viewBox its coordinate space, d the path data.
//
// Every glyph package's Def() is exactly one call to this. A glyph that needs
// more than one path or a stroke is not this family — call sprite.Define
// directly in that package rather than widening this helper.
func Solid(ref svg.Icon, viewBox, d string) sprite.Definition {
	return sprite.Define(ref, viewBox, sprite.Path(d))
}
