//go:build !wasm

package laptop_test

import (
	"strings"
	"testing"

	"webtyp.com/icons/laptop"
	"webtyp.com/svg/sprite"
)

// The Ref is a plain id string — safe to carry into a WASM build. The geometry
// is behind !wasm (svg.go), so nothing about referencing the glyph pulls
// svg/sprite into the browser bundle.
func TestRefIsPlainID(t *testing.T) {
	if laptop.Ref.ID() != "laptop" {
		t.Fatalf("Ref id = %q, want %q", laptop.Ref.ID(), "laptop")
	}
}

// Consumer shape: build a one-glyph sprite the way a component's IconSvg()
// does, and render it. The symbol must land under the same id the Ref uses,
// and be a currentColor path so the use-site box drives its colour.
func TestDefRendersCurrentColorSymbol(t *testing.T) {
	out := sprite.NewSprite(laptop.Def()).String()
	if !strings.Contains(out, `<symbol id="laptop"`) {
		t.Errorf("sprite is missing the laptop symbol:\n%s", out)
	}
	if !strings.Contains(out, `fill="currentColor"`) {
		t.Errorf("glyph must be a currentColor path (box owns the colour):\n%s", out)
	}
}
