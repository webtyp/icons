// Package database is the database glyph (a stack of disks) — a database, its tables, its schema diagram.
//
// Ref is WASM-safe (just the symbol id). Def() lives in svg.go behind
// //go:build !wasm.
package database

import "webtyp.com/svg"

// Ref is the symbol id, for markup: database.Ref.Render(class), or href="#database".
const Ref = svg.Icon("database")
