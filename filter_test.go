package sensitive

import (
	"testing"
)

func TestSensitiveFilter(t *testing.T) {
	filter := New()
	filter.LoadWordDict("../dict/dict.txt")
	filter.AddWord("有一个东西")
	filter.AddWord("一个东西")
	filter.AddWord("一个")
	filter.AddWord("东西")
	filter.AddWord("个东")

	testcases := []struct {
		Text   string
		Expect string
	}{
		{"我有一个东东西", "我有东"},
		{"我有一个东西", "我"},
		{"一个东西", ""},
		{"两个东西", "两西"},
		{"一个物体", "物体"},
	}

	for _, tc := range testcases {
		if got := filter.Filter(tc.Text); got != tc.Expect {
			t.Fatalf("filter %s, got %s, expect %s", tc.Text, got, tc.Expect)
		}
	}

}

func TestSensitiveValidate(t *testing.T) {
	filter := New()
	filter.LoadWordDict("../dict/dict.txt")
	filter.AddWord("有一个东西")
	filter.AddWord("一个东西")
	filter.AddWord("一个")
	filter.AddWord("东西")
	filter.AddWord("个东")

	testcases := []struct {
		Text        string
		ExpectPass  bool
		ExpectFirst string
	}{
		{"我有一@ |个东东西", false, "一个"},
		{"我有一个东东西", false, "一个"},
		{"我有一个东西", false, "有一个东西"},
		{"一个东西", false, "一个"},
		{"两个东西", false, "个东"},
		{"一样东西", false, "东西"},
	}

	for _, tc := range testcases {
		if pass, first := filter.Validate(tc.Text); pass != tc.ExpectPass || first != tc.ExpectFirst {
			t.Fatalf("validate %s, got %v, %s, expect %v, %s", tc.Text, pass, first, tc.ExpectPass, tc.ExpectFirst)
		}
	}

}

func TestSensitiveReplace(t *testing.T) {
	filter := New()
	filter.LoadWordDict("../dict/dict.txt")
	filter.AddWord("有一个东西")
	filter.AddWord("一个东西")
	filter.AddWord("一个")
	filter.AddWord("东西")
	filter.AddWord("个东")

	testcases := []struct {
		Text   string
		Expect string
	}{
		{"我有一个东东西", "我有**东**"},
		{"我有一个东西", "我*****"},
		{"一个东西", "****"},
		{"两个东西", "两**西"},
		{"一个物体", "**物体"},
	}

	for _, tc := range testcases {
		if got := filter.Replace(tc.Text, 42); got != tc.Expect {
			t.Fatalf("replace %s, got %s, expect %s", tc.Text, got, tc.Expect)
		}
	}

}
