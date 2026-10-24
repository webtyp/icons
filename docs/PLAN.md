---
PLAN: "feat: laptop, location, network and terminal glyphs"
EXECUTOR: jules
REVIEWER: none
STATUS: running
SESSION: 11800991101015914685
---

> Read AGENTS.md at the repo root before any change.

# PLAN — four new shared glyphs

## Why

Two consumers draw a glyph this set does not have, and both are wrong today:

- `veltylabs/device_manager` ships an **apple** path under a comment that says "laptop", and reuses it
  for three different screens (devices, zones, network interfaces).
- `webtyp/layout/apiexplorer` references the symbol `terminal`, but nobody defines it: the nav item
  renders with no icon.

Both will import these packages in their own follow-up plans. This plan only adds the glyphs.

## Tools

```
go install webtyp.com/devflow/cmd/gotest@latest
```

`curl` (preinstalled) is used to read the source SVG files.

## Ecosystem search

Searched `webtyp/icons` (glyphs: trash, pencil, plus, undo, selectall), `webtyp/svg`, `webtyp/layout`
and every `veltylabs/modules/*/ui/svg.go` for laptop, terminal, network, location, map, pin.
No shared glyph exists; `device_manager/ui/svg.go` has a private (wrong) path. **Extend**
`webtyp/icons` with one package per glyph, the pattern of `trash/`.

## Design gate

1. **Prior art:** Font Awesome, Material Symbols and Lucide ship one file per glyph and let the
   consumer import only what it uses; this repo already does the same with one Go package per glyph.
2. **Novice-name test:** `laptop.Ref`, `location.Ref`, `network.Ref`, `terminal.Ref` read as what
   they draw.
3. **Complexity ledger:** concepts +0 (same shape as `trash`); ways to draw a laptop −1 once
   `device_manager` deletes its private path (its own plan).
4. **Where it belongs:** shared glyphs live here (AGENTS.md: one package per glyph).
5. **Deletes:** nothing in this repo; consumers delete their private paths in their own plans.

## Source of the geometry — copy, never draw

Every glyph is **Font Awesome Free 6.7.2, solid family**. Read each file and copy its `viewBox`
attribute and the `d` attribute of its single `<path>` **verbatim** (no rounding, no reformatting):

| Package | Symbol id | Source file |
|---|---|---|
| `laptop` | `laptop` | `https://unpkg.com/@fortawesome/fontawesome-free@6.7.2/svgs/solid/laptop.svg` |
| `location` | `location` | `https://unpkg.com/@fortawesome/fontawesome-free@6.7.2/svgs/solid/location-dot.svg` |
| `network` | `network` | `https://unpkg.com/@fortawesome/fontawesome-free@6.7.2/svgs/solid/network-wired.svg` |
| `terminal` | `terminal` | `https://unpkg.com/@fortawesome/fontawesome-free@6.7.2/svgs/solid/terminal.svg` |

Example: `curl -sL https://unpkg.com/@fortawesome/fontawesome-free@6.7.2/svgs/solid/laptop.svg`.
Each file has exactly one `<path>`; if one has more than one path, stop and say which in the PR.
If the URL cannot be fetched, stop and open the PR saying so — never write a path from memory.

## APIs used

| Symbol | Exact signature | Source |
|---|---|---|
| `icons.Solid` | `func Solid(ref svg.Icon, viewBox, d string) sprite.Definition` | this repo, `icons.go` |
| `svg.Icon` | `type Icon string`; `func (i Icon) ID() string`; `func (i Icon) Render(class string) …` | `webtyp.com/svg` (current `go.mod`) |
| `sprite.NewSprite` | `func NewSprite(defs ...sprite.Definition) *Sprite`; `(*Sprite).String() string` | `webtyp.com/svg/sprite` (current `go.mod`) |

If a signature in the published tag differs from this table, stop and name the symbol in the PR;
never adapt it by guessing.

## Stage 1 — `laptop/`

Copy `trash/` exactly (three files) and change only:

`laptop/laptop.go`:

```go
// Package laptop is the laptop glyph — a computer or device registered in the
// system: the devices screen and anything that lists them.
//
// Ref is WASM-safe (just the symbol id). Def() lives in svg.go behind
// //go:build !wasm.
package laptop

import "webtyp.com/svg"

// Ref is the symbol id, for markup: laptop.Ref.Render(class), or href="#laptop".
const Ref = svg.Icon("laptop")
```

`laptop/svg.go` (`//go:build !wasm`, package `laptop`):

```go
// Def is the glyph geometry, for a consumer's IconSvg() sprite. FontAwesome
// Free 6.7.2 "laptop" (solid), viewBox <the copied viewBox>.
func Def() sprite.Definition {
	return icons.Solid(Ref, "<the copied viewBox>", "<the copied d>")
}
```

`laptop/svg_test.go`: same two tests as `trash/svg_test.go`, with package `laptop_test`,
`laptop.Ref`, `laptop.Def()` and the id `"laptop"`.

## Stage 2 — `location/`, `network/`, `terminal/`

Same as Stage 1 for each, with these package comments (first paragraph):

- `location`: `// Package location is the map-pin glyph — a place or zone: where something is.`
- `network`: `// Package network is the wired-network glyph — network interfaces, cabling, connectivity.`
- `terminal`: `// Package terminal is the command-prompt glyph — developer tools: the API explorer, a console.`

The source comment in each `svg.go` names its Font Awesome file (`location-dot`, `network-wired`,
`terminal`).

## Stage 3 — `icons_test.go`

Add the four glyphs to the `glyphs` table of `TestConsumerBuildsSpriteAndMarkup` (same shape as
the existing rows), and their imports. Change nothing else in the file.

## Stage 4 — docs

- `README.md`, table under `## Glyphs`: add four rows, in this order after `selectall`:

  ```
  | `icons/laptop` | laptop | a device or computer: device lists |
  | `icons/location` | map pin | a place or zone |
  | `icons/network` | wired network | network interfaces, connectivity |
  | `icons/terminal` | command prompt | developer tools: API explorer, console |
  ```
- `docs/ARCHITECTURE.md` line 61: the dependency check lists packages by hand
  (`go list -deps ./trash ./pencil ./plus ./undo`, already missing `selectall`). Replace it with
  `GOOS=js GOARCH=wasm go list -deps ./... | grep webtyp/svg/sprite   # MUST be empty`, so a new glyph
  never has to be added there. Change nothing else in that file.

## Delivery check

```
GOPRIVATE=webtyp.com/*,github.com/veltylabs/* GOFLAGS=-mod=mod gotest
GOOS=js GOARCH=wasm go list -deps ./... | grep webtyp/svg/sprite    # must print nothing
```

- A PR that does not compile, or whose tests could not run, is a failed delivery.
- Never pause to ask permission or report progress: when the check passes, commit, push and open
  the PR. Stuck for 30 minutes: open the PR anyway listing what fails.
- Commit after every stage (`git add -A` then `git commit -m "stage N: …"`); never `git reset`,
  `git restore`, `git checkout -- <file>` or `git stash` on your own work.
- `git diff --stat main` must list every file below and nothing else.

## Stages

| # | Files | Done when |
|---|---|---|
| 1 | `laptop/laptop.go`, `laptop/svg.go`, `laptop/svg_test.go` | tests pass |
| 2 | `location/*`, `network/*`, `terminal/*` (3 files each) | tests pass |
| 3 | `icons_test.go` | table has 9 glyphs |
| 4 | `README.md`, `docs/ARCHITECTURE.md` | 4 rows added; line 61 uses `./...` |
