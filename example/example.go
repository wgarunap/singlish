package main

import (
	"fmt"
	"github.com/wgarunap/singlish"
)

func main() {
	txt := singlish.TranslateString("මෙය url generate කිරීම උදෙසා ලියන ලද්දක් බැවින් නිරවද්‍යතාවය 100% තහවුරු කර නැත.")
	fmt.Println(txt)

	txt2 := singlish.TranslateSinhalaStr("මෙය සිංහල එකකි")
	fmt.Println(txt2)

	txt3 := singlish.TranslateString(`මෙය !"#$%&'()*+,-./:;<=>?@[\]^_{|}~ එකකි `)
	fmt.Println(txt3)
}
