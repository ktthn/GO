package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []string{"leet", "hw1", "leeb", "leer"}
	fmt.Println(longestCommonPrefix(a))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, str := range strs { //underscore is for ignoring the index of the current element
		for strings.HasPrefix(str, prefix) == false { // Check if the current string 'str' starts with the current 'prefix'
			// If not, reduce the prefix by one character
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
