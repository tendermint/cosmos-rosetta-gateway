package ensurehex

import "strings"

const zerox = "0x"

// String ensures that string representation of hex starts with 0x.
func String(hex string) string {
	if !strings.HasPrefix(hex, zerox) {
		return zerox + hex
	}
	return hex
}
