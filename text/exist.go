package textutil

import (
	"regexp"
	"unicode"
)

// IsHan 判断是否为中文...
func IsHan(text string) bool {
	for _, r := range text {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}
