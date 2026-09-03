# Agent Guide — `tinywasm/icons`

Read before adding or changing a glyph.

## Language

Public library: code, comments, identifiers, error messages **in English**.
Only `PLAN.md` prose may be Spanish.

## One package per glyph — no exceptions

A glyph is a folder with its own package. Never add a second glyph to an
existing package "because they go together": a consumer that wants one would
then pull the other's geometry. `trash` and `pencil` are used side by side in
CRUD views and are still separate packages.

## The three files

```
<glyph>/<glyph>.go       (untagged)      const Ref = svg.Icon("<id>")
<glyph>/svg.go            //go:build !wasm  func Def() sprite.Definition { return icons.Solid(Ref, viewBox, d) }
<glyph>/svg_test.go       //go:build !wasm  Ref id + consumer-shaped sprite render
```

- The id in `svg.Icon("<id>")` and the folder/package name should match.
- `svg.go` MUST carry `//go:build !wasm`. `svg/sprite` compiles for WASM, so
  the compiler will not remind you — the dep check below is the only guard.
- `Def()` is one call to `icons.Solid`. If the source glyph is not a single
  solid path (multiple paths, a stroke, `<g>` transforms), it is not this
  family: call `sprite.Define` directly in that package. Do not add fields or
  variants to `icons.Solid`.

## No colour, no size in the geometry

The path is `fill="currentColor"` (enforced by `sprite.Path`). Size and colour
are the consumer's, set from CSS on the box around `Ref.Render(class)`. A glyph
that hard-codes a colour is a bug.

## Mandatory pre-publish check

```bash
go build ./... && gotest
GOOS=js GOARCH=wasm go list -deps ./... | grep tinywasm/svg/sprite   # MUST be empty
```

The second line failing means a `svg.go` is missing its `//go:build !wasm` tag
and path data is about to ship to every browser that loads a consumer.

## Testing

`gotest`, never `go test`. Stdlib `testing` only — no testify/gomega. Each
glyph's test proves (1) `Ref.ID()` is the plain id, (2) `sprite.NewSprite(Def())`
renders a `<symbol>` under that id with a `currentColor` path — the shape a
consumer's `IconSvg()` uses.

## Publish

`gopush 'message'` from the repo root. Update `README.md` (the glyph table) and
`docs/` before publishing a new glyph.
