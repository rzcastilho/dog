// Package dog provides functions to dealing with dogs
package dog

import (
	"strings"
)

func WhenGrowsUp(s string) string {
	return "When the puppy grows up it says: " + strings.ToUpper(s)
}

// Years returns dog years based on human years
func Years(i int) int {
	return i / 7
}
