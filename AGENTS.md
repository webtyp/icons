# AGENTS.md — webtyp/icons

<!-- webtyp:base — copied from devskills skill agents-md; do not edit here -->
## Ecosystem rules — read before any change

`docs/PLAN.md` is the spec: no design judgment; anything unclear is a question in the PR, not a guess.

- Public API: typed (no `any`/`interface{}`), explicit, minimal; never export a symbol only a test uses.
- One intent, one path: the old path is deleted in the same change. No silent fallback, `TODO`, commented code or stub.
- Fail at compile time, else a loud diagnostic — never silently. Zero values deny.
- A defect in a dependency is reported, never wrapped or patched here.
- No string literals in logic: named constants. The same logic never lives in two repos: import the package the plan names; if you find an existing one the plan does not name that already does what you are about to write, say so in the PR instead of writing it.
- `cmd/*/main.go`: parse flags, inject, print/exit only. stdout = data, stderr = logs, exit codes are the contract.
- WASM code: `webtyp/fmt`, `webtyp/time`, `webtyp/json` instead of stdlib; no `map`; `dom.Element` embedded by value.
- `css.go`, `svg.go`, `js.go`, `html.go` carry `//go:build !wasm`; no `ssr.go`, no `front.go`.
- Build tags: `wasm`/`!wasm` only when the API does not exist on that platform; `browser`/`!browser` when the file belongs to another binary One main per folder, no destination tag: `web/browser/main.go` → `web/public/browser.wasm`, `web/server/main.go` → `web/server/server`, `web/edge/main.go` → `web/edge/edge.wasm`.
- Code and keys in English; UI text is `lang.Text`, translations live in `lang.json`. Comments and docs: English in `webtyp/*`; in `veltylabs/*` **every README, `docs/*` file, diagram and code comment is in Spanish**, written for a reader with little technical background, and each exported symbol's comment shows a short usage example.
- A translation key is reused, never duplicated as a near-identical phrase (`"not a valid integer"` + `"not a valid decimal number"` repeats `"not a valid"` three ways instead of composing it once) — every library ships to WASM/TinyGo, so each unreused phrase is dictionary rows AND binary bytes, in every consumer, forever. A token that doesn't need translating (a kind/type name like `int`) is passed as data (a variable), never as a string literal — a literal becomes a dictionary key whether or not it needed one.
- Tests cover behaviour and domain rules, never the look (CSS, colours, markup order, class strings). In `tests/` (external package), real stack with fakes only at the edge; run `gotest`.
- Docs (`README.md`, `docs/*`) updated with the API; permanent docs never link `docs/PLAN.md`.
- Before the PR: `git diff --stat main` lists every file of the plan's stages, and nothing the plan did not order is deleted.
<!-- /webtyp:base -->

## This repo

Read before adding or changing a glyph.


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
GOOS=js GOARCH=wasm go list -deps ./... | grep webtyp/svg/sprite   # MUST be empty
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
