//go:build !wasm

package location

import (
	"webtyp.com/icons"
	"webtyp.com/svg/sprite"
)

// Def is the glyph geometry, for a consumer's IconSvg() sprite. FontAwesome
// Free 6.7.2 "location-dot" (solid), viewBox 0 0 384 512.
func Def() sprite.Definition {
	return icons.Solid(Ref, "0 0 384 512",
		"M215.7 499.2C267 435 384 279.4 384 192C384 86 298 0 192 0S0 86 0 192c0 87.4 117 243 168.3 307.2c12.3 15.3 35.1 15.3 47.4 0zM192 128a64 64 0 1 1 0 128 64 64 0 1 1 0-128z")
}
