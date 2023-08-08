package order

type Ordered interface {
	Before(b Ordered) bool
}

type Visual interface {
	Show() string
}

type Key interface {
	Ordered
	Visual
}

const (
	PreOrder         = iota // NLR
	InOrder                 // LNR = ascending order
	PostOrder               // LRN
	ReversePreOrder         // NRL
	ReverseInOrder          // RNL = descending order
	ReversePostOrder        // RLN
)

var ImageWidth int

func init() {
	ImageWidth = 5
}
