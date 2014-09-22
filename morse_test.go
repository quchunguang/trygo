package trygo

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	fmt.Println("TestEncode()")
	s := "Welcome to Wikipedia, the free encyclopedia that anyone can edit."
	fmt.Println("Encode('" + s + "') is " + Encode(s))

	if Encode("SOS") != "... --- ... " {
		t.Fail()
	}
}

func TestDecode(t *testing.T) {
	fmt.Println("TestDecode()")
	s := ".-- . .-.. -.-. --- -- . - --- .-- .. -.- .. .--. . -.. .. .- - .... . ..-. .-. . . . -. -.-. -.-- -.-. .-.. --- .--. . -.. .. .- - .... .- - .- -. -.-- --- -. . -.-. .- -. . -.. .. - "
	fmt.Println("Decode('" + s + "') is " + Decode(s))

	if Decode("... --- ... ") != "SOS" {
		t.Fail()
	}
}

func TestEncodeS(t *testing.T) {
	fmt.Println("TestEncodeS()")
	s := "1234567890"
	fmt.Println("Encode('" + s + "') is " + EncodeS(s))

	if EncodeS("1234567890") != ".- ..- ...- ....- . -.... -... -.. -. - " {
		t.Fail()
	}
}

func TestDecodeS(t *testing.T) {
	fmt.Println("TestDecodeS()")
	s := ".- ..- ...- ....- . -.... -... -.. -. - "
	fmt.Println("Decode('" + s + "') is " + DecodeS(s))

	if DecodeS(".- ..- ...- ....- . -.... -... -.. -. - ") != "1234567890" {
		t.Fail()
	}
}
