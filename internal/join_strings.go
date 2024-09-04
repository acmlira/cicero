package internal

import "strings"

func JoinStrings(slice []string, separator string) string {
	return strings.Join(slice, separator)
}
