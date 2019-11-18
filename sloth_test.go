// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"github.com/nbio/st"
	"testing"
)

// TestSlothRelease test cases
func TestSlothRelease(t *testing.T) {
	st.Expect(t, SlothRelease(), "Sloth - 0.0.1")
}
