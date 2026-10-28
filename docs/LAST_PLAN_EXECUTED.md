---
PLAN: "feat: database and code glyphs"
TAG: v0.0.38
EXECUTOR: jules
REVIEWER: none
---

> Read AGENTS.md at the repo root before any change.

# Plan — `webtyp.com/icons`: two new glyphs

## Why

The app menu gets a «Developer tools» group (kept on `terminal`) with two children: the API explorer, which needs
its own glyph, and a new database-diagram screen. Two new packages, nothing else changes.

Source: IcoMoon Free (CC BY 4.0), `101-database.svg` and `384-embed.svg`, viewBox `0 0 16 16`. Both are several
paths, so per AGENTS.md («The three files») `Def()` calls `sprite.Define` with one `sprite.Path` per source path —
**not** `icons.Solid`, and `icons.Solid` is not touched.

## Stage 1 — `database/`

`database/database.go` (untagged):

```go
// Package database is the database glyph (a stack of disks) — a database, its tables, its schema diagram.
//
// Ref is WASM-safe (just the symbol id). Def() lives in svg.go behind
// //go:build !wasm.
package database

import "webtyp.com/svg"

// Ref is the symbol id, for markup: database.Ref.Render(class), or href="#database".
const Ref = svg.Icon("database")
```

`database/svg.go`:

```go
//go:build !wasm

package database

import "webtyp.com/svg/sprite"

// Def is the glyph geometry, for a consumer's IconSvg() sprite. IcoMoon Free
// (CC BY 4.0) "database", viewBox 0 0 16 16, three paths.
func Def() sprite.Definition {
	return sprite.Define(Ref, "0 0 16 16",
		sprite.Path("M8 0c-4.418 0-8 1.119-8 2.5v2c0 1.381 3.582 2.5 8 2.5s8-1.119 8-2.5v-2c0-1.381-3.582-2.5-8-2.5z"),
		sprite.Path("M8 8.5c-4.418 0-8-1.119-8-2.5v3c0 1.381 3.582 2.5 8 2.5s8-1.119 8-2.5v-3c0 1.381-3.582 2.5-8 2.5z"),
		sprite.Path("M8 13c-4.418 0-8-1.119-8-2.5v3c0 1.381 3.582 2.5 8 2.5s8-1.119 8-2.5v-3c0 1.381-3.582 2.5-8 2.5z"))
}
```

`database/svg_test.go`: copy `terminal/svg_test.go`, replacing `terminal` with `database` (package
`database_test`, both tests: `TestRefIsPlainID` and `TestDefRendersCurrentColorSymbol`).

## Stage 2 — `code/`

Same three files as stage 1, package `code`, id `"code"`. Package comment:
`// Package code is the code-brackets glyph (</>) — an API, source code, the API explorer.`

`code/svg.go` `Def()` comment: `IcoMoon Free (CC BY 4.0) "embed", viewBox 0 0 16 16, two paths.` Geometry:

```go
	return sprite.Define(Ref, "0 0 16 16",
		sprite.Path("M9 11.5l1.5 1.5 5-5-5-5-1.5 1.5 3.5 3.5z"),
		sprite.Path("M7 4.5l-1.5-1.5-5 5 5 5 1.5-1.5-3.5-3.5z"))
```

`code/svg_test.go`: same as stage 1 with `code`.

## Stage 3 — `README.md`

Add two rows after `icons/terminal` in the glyph table:

```markdown
| `icons/code` | code brackets `</>` | an API, source code: API explorer |
| `icons/database` | stack of disks | a database: tables, schema diagram |
```

Change the `icons/terminal` row's last cell to `developer tools: a console, the developer tools group`.

## Delivery

- `go build ./... && gotest` green.
- `GOOS=js GOARCH=wasm go list -deps ./... | grep webtyp/svg/sprite` prints nothing.
- `git diff --stat main` lists exactly: `code/code.go`, `code/svg.go`, `code/svg_test.go`, `database/database.go`,
  `database/svg.go`, `database/svg_test.go`, `README.md`.
- Never pause to ask permission or report progress: when the check passes, commit, push and open the PR. Stuck 30
  minutes → open the PR listing what fails.

| Stage | File | Action |
|---|---|---|
| 1 | `database/*.go` | new glyph |
| 2 | `code/*.go` | new glyph |
| 3 | `README.md` | two rows, `terminal` row reworded |
