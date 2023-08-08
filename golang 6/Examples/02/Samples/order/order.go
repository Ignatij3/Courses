package order

type Ordered interface {
	Before(b Ordered) bool
	Show() string
}
