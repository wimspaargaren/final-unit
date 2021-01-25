package runtime

import (
	"fmt"
	"strings"
)

// Contains check if memory contains key
func Contains(mem []string, key string) bool {
	for _, x := range mem {
		if strings.HasPrefix(key, x) {
			return true
		}
	}
	return false
}

// StartName expected start tag for test case output
func StartName(funcName string, index int) string {
	return fmt.Sprintf("<START;%s%d>", funcName, index)
}

// EndName expected end tag for test case output
func EndName(funcName string, index int) string {
	return fmt.Sprintf("<END;%s%d>", funcName, index)
}
