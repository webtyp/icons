// Package network is the wired-network glyph — network interfaces, cabling, connectivity.
//
// Ref is WASM-safe (just the symbol id). Def() lives in svg.go behind
// //go:build !wasm.
package network

import "webtyp.com/svg"

// Ref is the symbol id, for markup: network.Ref.Render(class), or href="#network".
const Ref = svg.Icon("network")
