//go:build !wasm

package icons_test

import (
	"strings"
	"testing"

	"github.com/tinywasm/icons/pencil"
	"github.com/tinywasm/icons/plus"
	"github.com/tinywasm/icons/trash"
	"github.com/tinywasm/icons/undo"
	"github.com/tinywasm/svg/sprite"
)

// The shape every consumer uses: pull the glyphs a view needs into ONE sprite
// from its IconSvg(), and render each Ref into markup. One <symbol> per glyph,
// each id matching its Ref, every body a currentColor path.
func TestConsumerBuildsSpriteAndMarkup(t *testing.T) {
	glyphs := []struct {
		id  string
		def sprite.Definition
		id2 string // Ref.ID()
		use string // Ref.Render(...) markup
	}{
		{"trash", trash.Def(), trash.Ref.ID(), trash.Ref.Render("m").String()},
		{"pencil", pencil.Def(), pencil.Ref.ID(), pencil.Ref.Render("m").String()},
		{"plus", plus.Def(), plus.Ref.ID(), plus.Ref.Render("m").String()},
		{"undo", undo.Def(), undo.Ref.ID(), undo.Ref.Render("m").String()},
	}

	defs := make([]sprite.Definition, len(glyphs))
	for i, g := range glyphs {
		defs[i] = g.def
	}
	sheet := sprite.NewSprite(defs...).String()

	if n := strings.Count(sheet, "<symbol"); n != len(glyphs) {
		t.Fatalf("expected %d symbols, got %d:\n%s", len(glyphs), n, sheet)
	}
	for _, g := range glyphs {
		if g.id2 != g.id {
			t.Errorf("Ref id = %q, want %q", g.id2, g.id)
		}
		if !strings.Contains(sheet, `<symbol id="`+g.id+`"`) {
			t.Errorf("sprite is missing the %q symbol:\n%s", g.id, sheet)
		}
		if !strings.Contains(g.use, `href='#`+g.id+`'`) {
			t.Errorf("%s Ref.Render must emit <use href='#%s'>, got %s", g.id, g.id, g.use)
		}
	}
	if !strings.Contains(sheet, `fill="currentColor"`) {
		t.Errorf("glyph bodies must be currentColor paths:\n%s", sheet)
	}
}
