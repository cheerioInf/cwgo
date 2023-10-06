package utils

import (
	"net"
	"strings"
)

// IsTokenError Determine if it is a Token error issue
func IsTokenError(err error) bool {
	return strings.Contains(err.Error(), "401 Unauthorized")
}

// IsNetworkError Determine if it is a network timeout issue
func IsNetworkError(err error) bool {
	if netErr, ok := err.(net.Error); ok {
		if netErr.Timeout() {
			return true
		}
	}
	return false
}