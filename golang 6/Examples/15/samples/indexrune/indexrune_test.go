package indexrune
 
import 	(
	"testing"
	"strings"
	"unicode/utf8"
)
	
func TestIndexRune(t *testing.T) {
	tests := []struct {
		in   string
		rune rune
		want int
	}{
		{"", 'a', -1},
		{"", '☺', -1},
		{"foo", '☹', -1},
		{"foo", 'o', 1},
		{"foo☺bar", '☺', 3},
		{"foo☺☻☹bar", '☹', 9},
		{"a A x", 'A', 2},
		{"some_text=some_value", '=', 9},
		{"☺a", 'a', 3},
		{"a☻☺b", '☺', 4},

		// RuneError should match any invalid UTF-8 byte sequence.
		{"�", '�', 0},
		{"\xff", '�', 0},
		{"☻x�", '�', len("☻x")},
		{"☻x\xe2\x98", '�', len("☻x")},
		{"☻x\xe2\x98�", '�', len("☻x")},
		{"☻x\xe2\x98x", '�', len("☻x")},

		// Invalid rune values should never match.
		{"a☺b☻c☹d\xe2\x98�\xff�\xed\xa0\x80", -1, -1},
		{"a☺b☻c☹d\xe2\x98�\xff�\xed\xa0\x80", 0xD800, -1}, // Surrogate pair
		{"a☺b☻c☹d\xe2\x98�\xff�\xed\xa0\x80", utf8.MaxRune + 1, -1},
	}
	for _, tt := range tests {
		if got := strings.IndexRune(tt.in, tt.rune); got != tt.want {
			t.Errorf("IndexRune(%q, %d) = %v; want %v", tt.in, tt.rune, got, tt.want)
		}
	}
}
