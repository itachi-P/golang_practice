package ignore

import (
	"testing"
)

func TestParseMatchSingle(t *testing.T) {
	cases := []struct {
		pattern  string
		filename string
		want     bool
	}{
		{"*.c", "foo.c", true},
		{"*.c", "foo/bar.c", true},
		{"Documentation/*.html", "Documentation/git.html", true},
		{"Documentation/*.html", "Documentation/ppc/ppc.html", false},
		{"/*.c", "cat-file.c", true},
		{"/*.c", "mozilla-sha1/sha1.c", false},
		{"foo", "foo", true},
		{"**/foo", "./foo", true}, // <--- leading './' required
		{"**/foo", "junk/foo", true},
		{"**/foo/bar", "./foo/bar", true}, // <--- leading './' required
		{"**/foo/bar", "junk/foo/bar", true},
		{"abc/**", "abc/foo", true},
		{"abc/**", "abc/foo/bar", true},
		{"a/**/b", "a/b", true},
		{"a/**/b", "a/x/b", true},
		{"a/**/b", "a/x/y/b", true},

		{"*_test*", "foo_test.go", true},
		{"*_test*", "junk/foo_test.go", true},
		{"junk\n!junk", "foo", false},
		{"junk\n!junk", "junk", false},

		{"*.html\n!foo.html", "junk.html", true},
		{"*.html\n!foo.html", "foo.html", false},

		{"/*\n!/foo\n/foo/*\n!/foo/bar", "crap", true},
		{"/*\n!/foo\n/foo/*\n!/foo/bar", "foo/crap", true},
		{"/*\n!/foo\n/foo/*\n!/foo/bar", "foo/bar", false},
		{"/*\n!/foo\n/foo/*\n!/foo/bar", "foo/bar/other", false},
		{"/*\n!/foo\n/foo/*\n!/foo/bar", "foo", false},
	}
	for _, test := range cases {
		test := test
		t.Run(test.pattern, func(t *testing.T) {
			t.Parallel()
		})

		matcher, err := Parse([]byte(test.pattern))
		if err != nil {
			t.Errorf("error: %s", err)
		}
		got := matcher.Match(test.filename)
		if test.want != got {
			t.Errorf("%q.Match(%q) = %v, got %v", test.pattern, test.filename, test.want, got)
		}
	}
}
