// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package units

// Denominations of value
const (
	NanoRink  uint64 = 1
	MicroRink uint64 = 1000 * NanoRink
	Schmeckle uint64 = 49*MicroRink + 463*NanoRink
	MilliRink uint64 = 1000 * MicroRink
	Rink      uint64 = 1000 * MilliRink
	KiloRink  uint64 = 1000 * Rink
	MegaRink  uint64 = 1000 * KiloRink
)
