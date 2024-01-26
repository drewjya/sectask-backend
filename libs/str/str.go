package str

import (
	"sectask/libs/number"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Buffer struct {
	R         []byte
	RuneBytes [utf8.UTFMax]byte
}

func ContainsUpper(data string, contains string) bool {
	return strings.Contains(strings.ToUpper(data), contains)
}

func TrimSpaceToLower(data string) string {
	return strings.TrimSpace(strings.ToLower(data))
}

func TrimSpaceToUpper(data string) string {
	return strings.TrimSpace(strings.ToUpper(data))
}

func TrimSpaceUpperFirstLetter(data string) string {
	return strings.TrimSpace(cases.Title(language.Und).String(data))
}

func TrimDoubleQuote(data string) string {
	return strings.Trim(data, "\"")
}

func ParseEnglishDecimal(data string) string {
	data = strings.ReplaceAll(data, ",", "")
	return data
}

func FindFirstDataNumberFormated(data ...string) string {
	for _, d := range data {
		if d != "" {
			d = strings.ReplaceAll(d, ",", "")
			return number.FormattingNumber(d, "")
		}
	}
	return "0"
}

func FindFirstDataNumber(data ...string) string {
	for _, d := range data {
		if d != "" {
			d = strings.ReplaceAll(d, ",", "")
			return d
		}
	}
	return "0"
}

func (b *Buffer) Write(r rune) {
	if r < utf8.RuneSelf {
		b.R = append(b.R, byte(r))
		return
	}
	n := utf8.EncodeRune(b.RuneBytes[0:], r)
	b.R = append(b.R, b.RuneBytes[0:n]...)
}

func (b *Buffer) Indent() {
	if len(b.R) > 0 {
		b.R = append(b.R, '_')
	}
}

// Underscore ...
func Underscore(s string) string {
	b := Buffer{
		R: make([]byte, 0, len(s)),
	}
	var m rune
	var w bool
	for _, ch := range s {
		if unicode.IsUpper(ch) {
			if m != 0 {
				if !w {
					b.Indent()
					w = true
				}
				b.Write(m)
			}
			m = unicode.ToLower(ch)
		} else {
			if m != 0 {
				b.Indent()
				b.Write(m)
				m = 0
				w = false
			}
			b.Write(ch)
		}
	}
	if m != 0 {
		if !w {
			b.Indent()
		}
		b.Write(m)
	}

	return string(b.R)
}
