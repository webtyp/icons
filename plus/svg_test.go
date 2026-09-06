//go:build !wasm

package plus_test

import (
	"strings"
	"testing"

	"webtyp.com/icons/plus"
	"webtyp.com/svg/sprite"
)

func TestRefIsPlainID(t *testing.T) {
	if plus.Ref.ID() != "plus" {
		t.Fatalf("Ref id = %q, want %q", plus.Ref.ID(), "plus")
	}
}

func TestDefRendersCurrentColorSymbol(t *testing.T) {
	out := sprite.NewSprite(plus.Def()).String()
	if !strings.Contains(out, `<symbol id="plus"`) {
		t.Errorf("sprite is missing the plus symbol:\n%s", out)
	}
	if !strings.Contains(out, `fill="currentColor"`) {
		t.Errorf("glyph must be a currentColor path (box owns the colour):\n%s", out)
	}
}
