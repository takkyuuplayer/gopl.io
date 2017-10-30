package anagram

import (
	"strings"
	"testing"
)

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for _, v := range s1 {
		if i2 := strings.IndexRune(s2, v); i2 == -1 {
			return false
		} else if i2 == len(s2) {
			s2 = s2[:i2]
		} else {
			s2 = s2[:i2] + s2[(i2+1):]
		}
	}

	return true
}

func TestIsAnagram(t *testing.T) {
	if isAnagram("abcde", "abcdef") != false {
		t.Errorf(`isAnagram("abcde", "abcdef") = %#v, want %#v`, isAnagram("abcde", "abcdef"), false)
	}

	if isAnagram("tea", "eat") != true {
		t.Errorf(`isAnagram("tea", "eat") = %#v, want %#v`, isAnagram("tea", "eat"), true)
	}

	if isAnagram("eat", "abc") != false {
		t.Errorf(`isAnagram("eat", "abc") = %#v, want %#v`, isAnagram("eat", "abc"), false)
	}

	if isAnagram("世界", "界世") != true {
		t.Errorf(`isAnagram("世界","界世") = %#v, want %#v`, isAnagram("世界", "界世"), true)
	}

}
