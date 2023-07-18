package main

import "testing"

func TestGoodStr(t *testing.T) {
	res, err := unPacking("a4bc2d5e")
	if res != "aaaabccddddde" || err != nil {
		t.Errorf("Test a4bc2d5e has errors %v, %v", res, err)
	}
}

func TestNotChangedStr(t *testing.T) {
	res, err := unPacking("abcd")
	if res != "abcd" || err != nil {
		t.Errorf("Test abcd has errors %v, %v", res, err)
	}
}
func TestBadStr(t *testing.T) {
	res, err := unPacking("45")
	if res != "" || err == nil {
		t.Errorf("Test 45 has errors %v, %v", res, err)
	}
}

func TestNullStr(t *testing.T) {
	res, err := unPacking("")
	if res != "" || err == nil {
		t.Errorf("Test str has errors %v, %v", res, err)
	}
}

func TestEscapeStr(t *testing.T) {
	res, err := unPacking(`qwe\4\5`)
	if res != "qwe45" || err != nil {
		t.Errorf(`Test qwe\4\5 has errors %v, %v`, res, err)
	}
}

func TestEscapeStr2(t *testing.T) {
	res, err := unPacking(`qwe\45`)
	if res != "qwe44444" || err != nil {
		t.Errorf(`Test qwe\45 has errors %v, %v`, res, err)
	}
}

func TestEscapeStrRepeatEscape(t *testing.T) {
	res, err := unPacking(`qwe\\5`)
	if res != `qwe\\\\\` || err != nil {
		t.Errorf(`Test qwe\\5 has errors %v, %v`, res, err)
	}
}
