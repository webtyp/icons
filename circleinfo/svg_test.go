//go:build !wasm

package circleinfo_test

import (
	"strings"
	"testing"

	"webtyp.com/icons/circleinfo"
	"webtyp.com/svg/sprite"
)

// The Ref is a plain id string — safe to carry into a WASM build.
func TestRefIsPlainID(t *testing.T) {
	if circleinfo.Ref.ID() != "circleinfo" {
		t.Fatalf("Ref id = %q, want %q", circleinfo.Ref.ID(), "circleinfo")
	}
}

// Consumer shape: a one-glyph sprite renders under the Ref's id, and every
// path is currentColor so the use-site box drives the colour.
func TestDefRendersCurrentColorSymbol(t *testing.T) {
	out := sprite.NewSprite(circleinfo.Def()).String()
	if !strings.Contains(out, `<symbol id="circleinfo"`) {
		t.Errorf("sprite is missing the circleinfo symbol:\n%s", out)
	}
	if strings.Contains(out, `fill="#`) {
		t.Errorf("glyph must not hardcode a colour:\n%s", out)
	}
	if !strings.Contains(out, `fill="currentColor"`) {
		t.Errorf("glyph must be a currentColor path (box owns the colour):\n%s", out)
	}
}
