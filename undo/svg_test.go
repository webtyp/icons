//go:build !wasm

package undo_test

import (
	"strings"
	"testing"

	"webtyp.com/icons/undo"
	"webtyp.com/svg/sprite"
)

func TestRefIsPlainID(t *testing.T) {
	if undo.Ref.ID() != "undo" {
		t.Fatalf("Ref id = %q, want %q", undo.Ref.ID(), "undo")
	}
}

func TestDefRendersCurrentColorSymbol(t *testing.T) {
	out := sprite.NewSprite(undo.Def()).String()
	if !strings.Contains(out, `<symbol id="undo"`) {
		t.Errorf("sprite is missing the undo symbol:\n%s", out)
	}
	if !strings.Contains(out, `fill="currentColor"`) {
		t.Errorf("glyph must be a currentColor path (box owns the colour):\n%s", out)
	}
}
