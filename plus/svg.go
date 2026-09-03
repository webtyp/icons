//go:build !wasm

package plus

import (
	"github.com/tinywasm/icons"
	"github.com/tinywasm/svg/sprite"
)

// Def is the glyph geometry, for a consumer's IconSvg() sprite. FontAwesome
// Free 6 "plus" (solid), viewBox 0 0 448 512.
func Def() sprite.Definition {
	return icons.Solid(Ref, "0 0 448 512",
		"M256 80c0-17.7-14.3-32-32-32s-32 14.3-32 32V224H48c-17.7 0-32 14.3-32 32s14.3 32 32 32H192V432c0 17.7 14.3 32 32 32s32-14.3 32-32V288H400c17.7 0 32-14.3 32-32s-14.3-32-32-32H256V80z")
}
