package testgo

import "testing"

func TestEven(t *testing.T) {
	if Even(2) != true {
		t.Log("2 should be even!")
		t.Fail()
	}
	if Even(3) != false {
		t.Log("3 should not be even!")
		t.Fail()
	}
}

// func (t *T) Fail()
// func (t *T) FailNow()
// func (t *T) Log(args ...interface{})
// func (t *T) Fatal(args ...interface{})
