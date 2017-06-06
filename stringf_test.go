package stringf_test

import (
	"testing"
	stringf "github.com/thanhpk/stringf"
)

func TestFormat(t *testing.T) {
	if "hi thanh" != stringf.Format("hi #name", map[string]string{
		"name": "thanh",
	}) {
		t.Fatal("err")
	}

	if "hi thanh " != stringf.Format("hi #name ", map[string]string{
		"name": "thanh",
	}) {
		t.Fatal("err")
	}

	if "hi #namethanh" != stringf.Format("hi ##name#name", map[string]string{
		"name": "thanh",
	}) {
		t.Fatal("err")
	}
}
