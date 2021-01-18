// Package utils contains utility functionality which is shared across mutiple packages
package utils

import "unicode"

// LowerCaseFirstLetter converts first character of a string to lower case
func LowerCaseFirstLetter(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
