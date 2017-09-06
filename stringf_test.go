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


	if "hi thanh" != stringf.Format("hi #name ", map[string]string{
		"name": "thanh",
	}) {
		t.Fatal("err")
	}

	if "hi #namethanh" != stringf.Format("hi ##name#name", map[string]string{
		"name": "thanh",
	}) {
		t.Fatal("err")
	}

	if "conversation_started.account.1.user.2" != stringf.Format("conversation_started.account.#account_id .user.#user_id ", map[string]string{
		"account_id": "1",
		"user_id": "2",
		"client_id": "3",
	}) {
		t.Fatal("err")
	}
}
