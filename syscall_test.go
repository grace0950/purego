// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2023 The grace0950 Authors

package purego_test

import (
	"os"
	"testing"

	_ "github.com/grace0950/purego"
)

func TestOS(t *testing.T) {
	// set and unset an environment variable since this calls into fakecgo.
	err := os.Setenv("TESTING", "SOMETHING")
	if err != nil {
		t.Errorf("failed to Setenv: %s", err)
	}
	err = os.Unsetenv("TESTING")
	if err != nil {
		t.Errorf("failed to Unsetenv: %s", err)
	}
}
