package a

import (
 _ "unsafe"
)

//go:linkname say a.say
//go:nosplit
func say(name string) string {
 return "hello, " + name
}

//go:linkname say2 fengzhu/golang-study/unsafe_test/b.Hi
//go:nosplit
func say2(name string) string {
 return "hi, " + name
}
