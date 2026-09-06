//go:build !wasm

package selectall

import (
	"github.com/tinywasm/icons"
	"github.com/tinywasm/svg/sprite"
)

// Def is the glyph geometry, for a consumer's IconSvg() sprite. FontAwesome
// Free 6 "check-double" (solid), viewBox 0 0 448 512.
func Def() sprite.Definition {
	return icons.Solid(Ref, "0 0 448 512",
		"M342.6 86.6c12.5-12.5 12.5-32.8 0-45.3s-32.8-12.5-45.3 0L160 178.7l-57.4-57.4c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3l80 80c12.5 12.5 32.8 12.5 45.3 0l160-160zm96 128c12.5-12.5 12.5-32.8 0-45.3s-32.8-12.5-45.3 0L160 402.7 54.6 297.4c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3l128 128c12.5 12.5 32.8 12.5 45.3 0l256-256z")
}
