package singlish

const (
	//Unicode values
	VowelsStart = 0x0D85
	VowelsEnd = 0x0D96

	ConstStart = 0x0D9A
	ConstEnd = 0x0DC6
)

var Vowels = map[string]string{
	"අ": "a", "ආ": "aa", "ඇ": "ae", "ඈ": "aa",
	"ඉ": "i", "ඊ": "ii", "උ": "u", "ඌ": "uu",
	"ඍ": "Ru", "ඎ": "Ruu", "ඏ": "au", "ඐ": "au",
	"එ": "e", "ඒ": "ee", "ඓ": "i", "ඔ": "o",
	"ඕ": "oo", "ඖ": "au",
}

var Constants = map[string]string{
	"ක": "k", "ඛ": "kh", "ග": "g", "ඝ": "gh", "ඞ": "N", "ඟ": "nng",
	"ච": "ch", "ඡ": "Ch", "ජ": "j", "ඣ": "q", "ඤ": "KN", "ඦ": "",
	"ට": "t", "ඨ": "T", "ඩ": "d", "ඪ": "D", "ණ": "n", "ඬ": "nnd",
	"ත": "th", "ථ": "Th", "ද": "dh", "ධ": "Dh", "න": "n", "ඳ": "nndh",
	"ප": "p", "ඵ": "ph", "බ": "b", "භ": "bh", "ම": "m", "ඹ": "B",
	"ය": "y", "ර": "r", "ල": "l", "ව": "w", "ශ": "sh", "ෂ": "Sh",
	"ස": "s", "හ": "h", "ළ": "L", "ෆ": "f",
}

var LangModifiers = map[string]string{
	"ං": "n",
	"ඃ": "h",
	"්": "",
	"ා": "a",
	"ැ": "e",
	"ෑ": "ea",
	"ි": "i",
	"ී": "ii",
	"ු": "u",
	"ූ": "uu",
	"ෘ": "ru",
	"ෙ": "e",
	"ේ": "ea",
	"ෛ": "ei",
	"ො": "o",
	"ෝ": "oo",
	"ෞ": "au",
	"ෟ": "uu",
	"ෲ": "ruu",
	"ෳ": "ru",
	"෴": "",
}

//special case ර
var R = map[string]string{
	"ැ": "u",
	"ෑ": "uu",
}
