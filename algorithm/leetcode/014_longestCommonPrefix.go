package main

import (
	"fmt"
	"sort"
)

func main() {
	var strs = []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	sort.Strings(strs)
	var ret string
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[0][i] != strs[j][i] {
				return ret
			}
		}
		ret = strs[0][:i+1]
	}
	return ret
}
