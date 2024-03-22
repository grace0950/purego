// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2023 The grace0950 Authors

//go:build darwin || freebsd || linux

package main

import "github.com/grace0950/purego"

func openLibrary(name string) (uintptr, error) {
	return purego.Dlopen(name, purego.RTLD_NOW|purego.RTLD_GLOBAL)
}
