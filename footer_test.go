package footer

import "testing"

var tests = []struct {
	currentPage, totalPages, bonduaries, around int
	expected                                    string
}{
	{1, 1, 0, 0, "1"},
	{4, 5, 1, 0, "1 ... 4 5"},
	{4, 10, 2, 2, "1 2 3 4 5 6 ... 9 10"},
	{1, 1000, 4, 0, "1 2 3 4 ... 997 998 999 1000"},
	{100, 1000, 4, 0, "1 2 3 4 ... 100 ... 997 998 999 1000"},
	{1, 1000, 4, 4, "1 2 3 4 5 ... 997 998 999 1000"},
	{6, 10, 1, 4, "1 2 3 4 5 6 7 8 9 10"},
}

func TestFooter(t *testing.T) {
	for _, test := range tests {
		if tested := footer(test.currentPage, test.totalPages, test.bonduaries, test.around); tested != test.expected {
			t.Fatalf("footer(%v, %v, %v, %v) = \"%v\", want \"%v\"", test.currentPage, test.totalPages,
				test.bonduaries, test.around, tested, test.expected)
		}
	}
}
