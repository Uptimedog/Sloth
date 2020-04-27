// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package util

import (
	"testing"

	"github.com/clivern/sloth/pkg"
)

// TestInArray test cases
func TestInArray(t *testing.T) {
	pkg.Expect(t, true, InArray("A", []string{"A", "B", "C", "D"}))
	pkg.Expect(t, true, InArray("B", []string{"A", "B", "C", "D"}))
	pkg.Expect(t, false, InArray("H", []string{"A", "B", "C", "D"}))
	pkg.Expect(t, true, InArray(1, []int{2, 3, 1}))
	pkg.Expect(t, false, InArray(9, []int{2, 3, 1}))
}
