// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package util

import (
	"testing"

	"github.com/clivern/sloth/pkg"
)

// TestCrypto test cases
func TestCrypto(t *testing.T) {
	encryptionKey, _ := GenerateRandomString(32)
	text := []byte(`{"Id":1,"hostname":"localhost","message":"Hello World","type":"agent"}`)
	ciphertext, err := Encrypt(text, []byte(encryptionKey))
	pkg.Expect(t, nil, err)
	plaintext, err := Decrypt(ciphertext, []byte(encryptionKey))
	pkg.Expect(t, nil, err)
	pkg.Expect(t, text, plaintext)
	pkg.Expect(t, string(text), string(plaintext))
}
