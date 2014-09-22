package trygo

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	fmt.Println("Encode('3DE') is " + Encode("3DE"))

	if Encode("SOS") != "...---..." {
		t.Fail()
	}
}

func TestDecode(t *testing.T) {
	fmt.Println("Decode('SOS') is " + Decode("...---..."))

	if Encode("...---...") != "SOS" {
		t.Fail()
	}
}
