package singlish

import (
	"strings"
	"unicode"
)

const escape = ` !"#$%&'()*+,-./:;<=>?@[\]^_{|}~` + "`"

// TranslateString will return the translated sinhala word and english words only.
// this function will remove any existing escape characters defined in this library.
func TranslateString(text string) (translated string) {
	l := []rune(text)
	for i, c := range l {
		if 0 < c && c < 128 {
			if strings.Contains(escape, string(c)) {
				if len(translated) > 0 && string(translated[len(translated)-1]) != " " {
					translated = translated + " "
				}
				continue
			}
			translated = translated + string(c)
		} else if unicode.Is(unicode.Sinhala, c) {
			if unicode.IsLetter(c) {
				if VowelsStart <= c && c <= VowelsEnd {
					let, ok := Vowels[string(c)]
					if !ok {
						continue
					}
					translated += let
					continue
				}
				if ConstStart <= c && c <= ConstEnd {
					let, ok := Constants[string(c)]
					if !ok {
						continue
					}
					if len(l)-1 >= i+1 && !unicode.IsMark(l[i+1]) || len(l) == i+1 {
						let += "a"
					}
					translated += let
					continue
				}
			} else if unicode.IsMark(c) {
				if i > 0 {
					if let, ok := R[string(l[i-1])]; ok {
						translated += let
						continue
					}
				}

				let, ok := LangModifiers[string(c)]
				if !ok {
					continue
				}
				translated += let
				continue
			}
		}
	}
	return translated
}

func TranslateSinhalaStr(text string) (translated string) {
	l := []rune(text)
	for i, c := range l {
		if unicode.Is(unicode.Sinhala, c) {
			if unicode.IsLetter(c) {
				if VowelsStart <= c && c <= VowelsEnd {
					let, ok := Vowels[string(c)]
					if !ok {
						continue
					}
					translated += let
					continue
				}
				if ConstStart <= c && c <= ConstEnd {
					let, ok := Constants[string(c)]
					if !ok {
						continue
					}
					if len(l)-1 >= i+1 && !unicode.IsMark(l[i+1]) || len(l) == i+1 {
						let += "a"
					}
					translated += let
					continue
				}
			} else if unicode.IsMark(c) {
				if i > 0 {
					if let, ok := R[string(l[i-1])]; ok {
						translated += let
						continue
					}
				}

				let, ok := LangModifiers[string(c)]
				if !ok {
					continue
				}
				translated += let
				continue
			}
		} else if unicode.IsSpace(c) {
			translated += " "
		}
	}
	return translated
}
