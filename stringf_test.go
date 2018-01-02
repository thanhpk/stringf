package stringf_test

import (
	"testing"
	stringf "github.com/thanhpk/stringf"
)

func TestFormat(t *testing.T) {
	flagtests := []struct {
		in string
		m map[string]string
		out string
	}{
		{ "hi {name}", map[string]string{
			"name": "thanh",
		}, "hi thanh" },
		{ "hi {name}", map[string]string{
			"name": "{name}thanh",
		}, "hi {name}thanh" },
		{ "hi {{name{name}}", map[string]string{
			"name": "thanh",
		}, "hi {namethanh}" },
		{ "conversation_started.account.{account_id}.user.{user_id}", map[string]string{
			"account_id": "1",
			"user_id": "2",
			"client_id": "3",
		}, "conversation_started.account.1.user.2" },
		{ "hi {name2}", nil, "hi {name2}" },
		{ "hi {{name2} alo", nil, "hi {name2} alo" },
		{ "hi {{name2", nil, "hi {name2"},
		{ "hi {name2 {name}", map[string]string{
			"name": "thanh",
		}, "hi {name2 thanh" },
		{ "hi {name2 {name}$s", map[string]string{
			"name": "thanh",
		}, "hi {name2 thanh$s" },
		{ "hi {name2#{name$s", map[string]string{
			"name": "thanh",
		}, "hi {name2#{name$s" },
		{ "hi {name{name}$v", map[string]string{
			"name": "thanh",
		}, "hi {namethanh$v" },
	}

	for _, tt := range flagtests {
		s := stringf.Format(tt.in, tt.m)
		if s != tt.out {
			t.Errorf("should equal, input %s expect %s, got %s", tt.in, tt.out, s)
		}
	}
}
