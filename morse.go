package trygo

import "strings"

var Table = map[byte]string{
	'A': ".-",
	'B': "-...",
	'C': "-.-.",
	'D': "-..",
	'E': ".",
	'F': "..-.",
	'G': "--.",
	'H': "....",
	'I': "..",
	'J': ".---",
	'K': "-.-",
	'L': ".-..",
	'M': "--",
	'N': "-.",
	'O': "---",
	'P': ".--.",
	'Q': "--.-",
	'R': ".-.",
	'S': "...",
	'T': "-",
	'U': "..-",
	'V': "...-",
	'W': ".--",
	'X': "-..-",
	'Y': "-.--",
	'Z': "--..",
	'0': "------",
	'1': ".-----",
	'2': "..---",
	'3': "...--",
	'4': "....-",
	'5': ".....",
	'6': "-....",
	'7': "--...",
	'8': "---..",
	'9': "----.",
}

var morse_url string = "http://introcs.cs.princeton.edu/java/data/morse.csv"

func Encode(s string) (ret string) {
	for _, c := range []byte(strings.ToUpper(s)) {
		ret += Table[c]
	}
	return
}

func Decode(s string) (ret string) {
	for s != "" {
		for k, v := range Table {
			if len(s) < len(v) {
				continue
			}
			if s[:len(v)] == v {
				s = s[len(v):]
				ret += string(k)
			}
		}
	}
	return
}
