//Given a string s containing just the characters '(', ')', '{', '}', '[' and ']
//', determine if the input string is valid. 
//
// An input string is valid if: 
//
// 
// Open brackets must be closed by the same type of brackets. 
// Open brackets must be closed in the correct order. 
// 
//
// 
// Example 1: 
//
// 
//Input: s = "()"
//Output: true
// 
//
// Example 2: 
//
// 
//Input: s = "()[]{}"
//Output: true
// 
//
// Example 3: 
//
// 
//Input: s = "(]"
//Output: false
// 
//
// Example 4: 
//
// 
//Input: s = "([)]"
//Output: false
// 
//
// Example 5: 
//
// 
//Input: s = "{[]}"
//Output: true
// 
//
// 
// Constraints: 
//
// 
// 1 <= s.length <= 104 
// s consists of parentheses only '()[]{}'. 
// 
// Related Topics 栈 字符串 
// 👍 1846 👎 0


package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	n := len(s)
	if n % 2 == 1 {
		return false
	}

	pairs := map[byte]byte {
		')':'(',
		']':'[',
		'}':'{',
	}
	stack := []byte{}

	for i := 0; i < n; i++{
		if pair, ok := pairs[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != pair {
				return false
			}
			stack = stack[:len(stack)-1]
		}else{
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0

	return true
}
//leetcode submit region end(Prohibit modification and deletion)




func main() {
	fmt.Println(isValid("((()))"))
}