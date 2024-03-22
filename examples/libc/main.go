// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 The grace0950 Authors

//go:build darwin || freebsd || linux || windows

package main

import (
	"fmt"
	"runtime"

	"github.com/grace0950/purego"
)

func getSystemLibrary() string {
	switch runtime.GOOS {
	case "darwin":
		return "/usr/lib/libSystem.B.dylib"
	case "linux":
		return "libc.so.6"
	case "freebsd":
		return "libc.so.7"
	case "windows":
		return "ucrtbase.dll"
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}

func main() {
	libc, err := openLibrary(getSystemLibrary())
	if err != nil {
		panic(err)
	}
	var puts func(string)
	purego.RegisterLibFunc(&puts, libc, "puts")
	puts("Calling C from Go without Cgo!")
}
