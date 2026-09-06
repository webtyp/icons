//go:build !wasm

package pencil_test

import (
	"strings"
	"testing"

	"webtyp.com/icons/pencil"
	"webtyp.com/svg/sprite"
)

func TestRefIsPlainID(t *testing.T) {
	if pencil.Ref.ID() != "pencil" {
		t.Fatalf("Ref id = %q, want %q", pencil.Ref.ID(), "pencil")
	}
}

func TestDefRendersCurrentColorSymbol(t *testing.T) {
	out := sprite.NewSprite(pencil.Def()).String()
	if !strings.Contains(out, `<symbol id="pencil"`) {
		t.Errorf("sprite is missing the pencil symbol:\n%s", out)
	}
	if !strings.Contains(out, `fill="currentColor"`) {
		t.Errorf("glyph must be a currentColor path (box owns the colour):\n%s", out)
	}
}
