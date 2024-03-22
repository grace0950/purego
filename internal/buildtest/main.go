// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 The grace0950 Authors

//go:build !windows

package main

import (
	_ "github.com/grace0950/purego"
)

import "C"

// This file tests that build Cgo and purego at the same time succeeds to build (#189).
func main() {
}
