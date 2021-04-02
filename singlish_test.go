package singlish

import "testing"

func TestTranslate(t *testing.T) {
	txt := []struct {
		in  string
		out string
	}{
		{"දැක්වෙන", "dhekwena"},
		{"රුක", "ruka"},
		{"රූක", "ruuka"},
		{"රෑක", "reaka"},
		{"අර්ශස්", "arshas"},
	}

	for _, s := range txt {
		if out := TranslateString(s.in); out != s.out {
			t.Fatalf("expected:%v , received:%v", s.out, out)
		}
	}
}

func BenchmarkTranslate1Char(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TranslateString("දැ")
	}
}

func BenchmarkTranslate4Char(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TranslateString("දැක්වෙන")
	}
}
