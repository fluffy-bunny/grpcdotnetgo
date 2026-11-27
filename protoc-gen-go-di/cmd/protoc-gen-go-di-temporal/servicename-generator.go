package main

import (
	"crypto/sha256"
	"encoding/hex"
)

// fileUniqueID returns an identifier derived from the proto file path, so it is
// unique per generated file and the same on every run.
func fileUniqueID(protoPath string) string {
	sum := sha256.Sum256([]byte(protoPath))
	return hex.EncodeToString(sum[:16])
}
