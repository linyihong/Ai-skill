package app

import "testing"

func TestRetrievalLexicalCandidate(t *testing.T) {
	for _, tc := range []struct{ in, want string }{
		{"API SQL API", `"API" OR "SQL"`},
		{`a" OR (b*)`, `"a" OR "OR" OR "b"`},
		{"中文 說明", `"中文" OR "說明"`},
		{"", `""`},
	} {
		if got := runtimeFTSLexicalQuery(tc.in); got != tc.want {
			t.Errorf("%q: %q != %q", tc.in, got, tc.want)
		}
	}
}
