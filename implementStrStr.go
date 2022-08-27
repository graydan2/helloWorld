
package main
import "fmt"

func main () {
	found := strStr("mississippi", "pi")
    fmt.Printf("Pumpernickle")
    fmt.Printf("testing stuff")
	fmt.Printf("Result: %v", found)
}

func strStr(haystack string, needle string) int {
    if len(needle) == 0 {
        return 0
    }
    for i := 0; i <= len(haystack) - len(needle); i++ {
        if needle[0] == haystack[i] {
            found := check(needle, haystack, i)
            if found {
                return i
            }
        }
    }
    return -1
}

func check(needle string, haystack string, index int) bool {
    for i := index; i < index + len(needle); i++ {
        if needle[i] != haystack[index + i] {
            return false
        }
    }
    return true
}

