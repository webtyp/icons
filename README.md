# icons
<img src="docs/img/badges.svg">

Shared SVG icon set for the TinyWasm framework. **One package per glyph.** A
project imports only the glyphs it uses; adding a new glyph is a new folder, it
never touches the ones already here.

## Why a whole package per icon

An icon has two halves that must reach different places:

- the **reference** — the symbol id, a plain string. This is all that may reach
  the browser (the WASM binary). It is `const Ref = svg.Icon("trash")`.
- the **geometry** — the `<path>` data and viewBox. This is backend-only: it is
  pulled out at build time by `tinywasm/ssr`, injected once into the page, and
  referenced by `<use href="#trash">`. Shipping it to the browser would drag
  the SVG-serialization machinery (`tinywasm/json`, `tinywasm/model`) into the
  bundle for nothing.

Keeping each glyph in its own package is what lets the geometry sit behind
`//go:build !wasm` per-glyph, so importing one glyph for its `Ref` can never
leak another glyph's path data into a WASM build.

## Getting started

```go
import (
    "github.com/tinywasm/icons/trash"
    "github.com/tinywasm/svg/sprite"
)

// 1. In component code (compiles to WASM too) — render the reference:
node := trash.Ref.Render("myapp__delete-icon")   // <svg class=...><use href="#trash"/></svg>

// 2. In your svg.go (//go:build !wasm) — hand the geometry to your sprite so
//    ssr can extract it:
func (c *MyView) IconSvg() *sprite.Sprite {
    return sprite.NewSprite(trash.Def(), /* pencil.Def(), ... */)
}
```

The `class` you pass to `Ref.Render` is where size and colour come from — the
path itself is `fill="currentColor"`, so the box around it drives the colour
(a white-text button gives a white glyph; a red box, a red glyph).

## Glyphs

| Package | Glyph | Use |
|---|---|---|
| `icons/trash` | trash can | destructive: delete, remove, discard |
| `icons/pencil` | pencil | edit: modify, rename, correct |
| `icons/plus` | plus | additive: new, add, create |
| `icons/undo` | counter-clockwise arrow | reverse: cancel, revert, back out |

## Adding a glyph

Copy an existing folder (`trash/`), rename the package and the three files,
swap the `Ref` id, and paste the new viewBox + path into `svg.go` via
`icons.Solid(Ref, viewBox, d)`. Keep the source a solid single-path glyph
(FontAwesome "solid" family or equivalent); anything else is not this family —
call `sprite.Define` directly in that package instead of widening `icons.Solid`.

## Docs

- [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) — the reference/geometry split and how ssr consumes it
- [docs/AGENTS.md](docs/AGENTS.md) — constraints for agents adding or changing glyphs
