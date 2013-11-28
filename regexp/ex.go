// 扩展的正则方法

package regexp

import (
	"regexp/syntax"
)

func CompileEx(expr string, mode syntax.Flags) (*Regexp, error) {
	return compile(expr, syntax.Perl|mode, false)
}

func CompilePOSIXEx(expr string, mode syntax.Flags) (*Regexp, error) {
	return compile(expr, syntax.POSIX|mode, true)
}
