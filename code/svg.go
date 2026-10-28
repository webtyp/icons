//go:build !wasm

package code

import "webtyp.com/svg/sprite"

// Def is the glyph geometry, for a consumer's IconSvg() sprite. IcoMoon Free
// (CC BY 4.0) "embed", viewBox 0 0 16 16, two paths.
func Def() sprite.Definition {
	return sprite.Define(Ref, "0 0 16 16",
		sprite.Path("M9 11.5l1.5 1.5 5-5-5-5-1.5 1.5 3.5 3.5z"),
		sprite.Path("M7 4.5l-1.5-1.5-5 5 5 5 1.5-1.5-3.5-3.5z"))
}
