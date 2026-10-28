//go:build !wasm

package database

import "webtyp.com/svg/sprite"

// Def is the glyph geometry, for a consumer's IconSvg() sprite. IcoMoon Free
// (CC BY 4.0) "database", viewBox 0 0 16 16, three paths.
func Def() sprite.Definition {
	return sprite.Define(Ref, "0 0 16 16",
		sprite.Path("M8 0c-4.418 0-8 1.119-8 2.5v2c0 1.381 3.582 2.5 8 2.5s8-1.119 8-2.5v-2c0-1.381-3.582-2.5-8-2.5z"),
		sprite.Path("M8 8.5c-4.418 0-8-1.119-8-2.5v3c0 1.381 3.582 2.5 8 2.5s8-1.119 8-2.5v-3c0 1.381-3.582 2.5-8 2.5z"),
		sprite.Path("M8 13c-4.418 0-8-1.119-8-2.5v3c0 1.381 3.582 2.5 8 2.5s8-1.119 8-2.5v-3c0 1.381-3.582 2.5-8 2.5z"))
}
