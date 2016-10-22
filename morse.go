package trygo

import "strings"

// MorseTable data
var MorseTable = map[byte]string{
	'A': ".- ",
	'B': "-... ",
	'C': "-.-. ",
	'D': "-.. ",
	'E': ". ",
	'F': "..-. ",
	'G': "--. ",
	'H': ".... ",
	'I': ".. ",
	'J': ".--- ",
	'K': "-.- ",
	'L': ".-.. ",
	'M': "-- ",
	'N': "-. ",
	'O': "--- ",
	'P': ".--. ",
	'Q': "--.- ",
	'R': ".-. ",
	'S': "... ",
	'T': "- ",
	'U': "..- ",
	'V': "...- ",
	'W': ".-- ",
	'X': "-..- ",
	'Y': "-.-- ",
	'Z': "--.. ",
	'0': "----- ",
	'1': ".---- ",
	'2': "..--- ",
	'3': "...-- ",
	'4': "....- ",
	'5': "..... ",
	'6': "-.... ",
	'7': "--... ",
	'8': "---.. ",
	'9': "----. ",
}

// MorseTableS data
var MorseTableS = map[byte]string{
	'1': ".- ",
	'2': "..- ",
	'3': "...- ",
	'4': "....- ",
	'5': ". ",
	'6': "-.... ",
	'7': "-... ",
	'8': "-.. ",
	'9': "-. ",
	'0': "- ",
}

var morseURL = "http://introcs.cs.princeton.edu/java/data/morse.csv"

// Encode func
func Encode(s string) (ret string) {
	for _, c := range []byte(strings.ToUpper(s)) {
		ret += MorseTable[c]
	}
	return
}

// Decode func
func Decode(s string) (ret string) {
	for s != "" {
		for k, v := range MorseTable {
			if strings.HasPrefix(s, v) {
				s = strings.TrimPrefix(s, v)
				ret += string(k)
			}
		}
	}
	return
}

// EncodeS func
func EncodeS(s string) (ret string) {
	for _, c := range []byte(strings.ToUpper(s)) {
		ret += MorseTableS[c]
	}
	return
}

// DecodeS func
func DecodeS(s string) (ret string) {
	for s != "" {
		for k, v := range MorseTableS {
			if strings.HasPrefix(s, v) {
				s = strings.TrimPrefix(s, v)
				ret += string(k)
			}
		}
	}
	return
}
