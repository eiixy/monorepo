package helpers

import "strings"

func SqlPlaceholder(length int) string {
	return strings.TrimSuffix(strings.Repeat("?,", length), ",")
}
